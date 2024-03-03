package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/http2"
)

const (
	PORT          = ":6443"
	READ_TIMEOUT  = 5 * time.Second
	WRITE_TIMEOUT = READ_TIMEOUT
)

var (
	certPath, keyPath string
)

func init() {
	flag.StringVar(&certPath, "cert", "certs/server.crt", "Server certificate path.")
	flag.StringVar(&keyPath, "key", "certs/server.key", "Server key path.")
}

func main() {
	server, err := configureServer()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Starting to listen HTTPS server on", PORT)
	// listen for HTTPS requests
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}

// configureServer returns the HTTPS server objects configured with TLS support.
func configureServer() (server *http.Server, err error) {
	var certificate *tls.Config
	if certificate, err = fetchCertificate(); err != nil {
		return server, err
	}
	server = &http.Server{
		Addr:         PORT,
		ReadTimeout:  READ_TIMEOUT,
		WriteTimeout: WRITE_TIMEOUT,
		TLSConfig:    certificate,
	}
	// configure server with HTTP2 and TLS
	if err := http2.ConfigureServer(server, nil); err != nil {
		return server, err
	}
	// add handlers in the server
	http.HandleFunc("/", handler)
	return
}

// handler returns the response for a particular HTTP path
func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("HTTP request from: %s\n", r.RemoteAddr))) // nolint
	w.Write([]byte(fmt.Sprintf("Protocol: %s", r.Proto)))                 // nolint
}

// fetchCertificate render the saved certificates
func fetchCertificate() (config *tls.Config, err error) {
	var (
		certificate tls.Certificate
		key, cert   []byte
	)
	if cert, err = os.ReadFile(certPath); err != nil {
		return nil, err
	}
	if key, err = os.ReadFile(keyPath); err != nil {
		return nil, err
	}
	if certificate, err = tls.X509KeyPair(cert, key); err != nil {
		return nil, err
	}
	config = &tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   "localhost",
	}
	return
}
