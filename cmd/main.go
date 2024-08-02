package main

import (
	"flag"
	"log"

	"go.uber.org/zap"

	"dns_resolver/internal/cache"
	"dns_resolver/internal/dns"
	"dns_resolver/internal/logger"
)

// cmd tool impl
var (
	address  = flag.String("address", ":53", "Address to listen for DNS queries")
	logLevel = flag.String("loglevel", "info", "logging level")
)

func main() {
	flag.Parse()

	loggr := logger.InitLogger()
	defer loggr.Sync()

	cache := cache.NewCache()

	resolver, err := dns.NewResolver(loggr, cache)
	if err != nil {
		log.Fatalf("Failed to create DNS resolver: %v", zap.Error(err))
	}

	if err := resolver.Start(*address); err != nil {
		log.Fatalf("Failed to start DNS resolver: %v", zap.Error(err))
	}
}
