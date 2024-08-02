#!bin/bash

# clean prev build
rm -f dns_resoler

#build
go build -o dns_resolver "../cmd/main.go"

echo "Build successful"