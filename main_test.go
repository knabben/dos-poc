package main

import (
	"log"
	"testing"
)

func TestCertificate(t *testing.T) {
	certPath = "certs/server.crt"
	keyPath = "certs/server.key"

	config, err := fetchCertificate()
	if err != nil {
		log.Fatalf("error when parsing certificates: %v", err)
	}

	if config == nil {
		log.Fatalf("error acessing config when loading certificates: %v", err)
	}
}
