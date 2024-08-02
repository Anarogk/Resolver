package dns

import (
	"testing"

	"github.com/miekg/dns"

	"dns_resolver/internal/cache"
	"dns_resolver/internal/logger"
)

func TestDNSResolver(t *testing.T) {
	logger := logger.InitLogger()
	cache := cache.NewCache()

	resolver, err := NewResolver(logger, cache)
	if err != nil {
		t.Fatalf("Failed to create DNS resolver: %v", err)
	}

	question := dns.Question{
		Name:   "example.com",
		Qtype:  dns.TypeA,
		Qclass: dns.ClassINET,
	}

	_, err = resolver.queryUpstream(question)
	if err != nil {
		t.Errorf("Failed to query upstream DNS serer: %v", err)
	}
}
