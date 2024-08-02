package dns

import (
	"github.com/miekg/dns"
	"go.uber.org/zap"

	"dns_resolver/internal/cache"
)

// struct for resolver
type Resolver struct {
	logger *zap.Logger
	cache  *cache.Cache
}

// create resolver instance with logger and cache set
func NewResolver(logger *zap.Logger, cache *cache.Cache) (*Resolver, error) {
	return &Resolver{logger: logger, cache: cache}, nil
}

// start resolver server with an address addr, udp prot,
func (r *Resolver) Start(addr string) error {
	r.logger.Info("Starting DNS server", zap.String("address", addr))
	server := &dns.Server{Addr: addr, Net: "udp"}
	dns.HandleFunc(".", r.handleDNSRequest)

	err := server.ListenAndServe()
	if err != nil {
		r.logger.Error("Failed to start DNS server", zap.Error(err))
		return err
	}
	return nil
}

// handle DNS req : main logic
func (r *Resolver) handleDNSRequest(w dns.ResponseWriter, req *dns.Msg) {

	defer w.Close()

	// create m with msg and req created
	m := new(dns.Msg)
	m.SetReply(req)
	m.Authoritative = true

	for _, q := range req.Question {
		r.logger.Info("Received DNS query", zap.String("question", q.Name))

		if cachedResponse, found := r.cache.Get(q.Name); found {
			r.logger.Info("Cache hit", zap.String("question", q.Name))
			m.Answer = append(m.Answer, cachedResponse.([]dns.RR)...)
		} else {
			r.logger.Info("Cache miss", zap.String("question", q.Name))
			answers, err := r.queryUpstream(q)
			if err != nil {
				r.logger.Error("Failed to query upstream DNS serer", zap.Error(err))
				continue
			}
			r.cache.Set(q.Name, answers)
			m.Answer = append(m.Answer, answers...)
		}
	}

	if err := w.WriteMsg(m); err != nil {
		r.logger.Error("Failed to write DNS repsonse", zap.Error(err))
	}
}

func (r *Resolver) queryUpstream(q dns.Question) ([]dns.RR, error) {
	client := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(q.Name, q.Qtype)
	m.RecursionDesired = true

	upstream := "8.8.8.8:53"
	response, _, err := client.Exchange(m, upstream)
	if err != nil {
		return nil, err
	}

	r.logger.Info("Queried upstream DNS server", zap.String("question", q.Name))
	return response.Answer, nil
}
