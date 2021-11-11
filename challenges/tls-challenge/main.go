package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

type handler struct {
}

func (h *handler) root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	n, err := w.Write([]byte(instructions))
	if err != nil {
		log.Printf("Failed to write response: %v.", err)
	}
	if n != len(instructions) {
		log.Printf("Failed to write full response: %v.", n)
	}
}

func (h *handler) register(w http.ResponseWriter, r *http.Request) {
	// Read in the body, but don't read in more then 0.5 MB.
	limitedReader := &io.LimitedReader{
		R: r.Body,
		N: 10240,
	}
	bytes, err := ioutil.ReadAll(limitedReader)
	if err != nil {
		log.Printf("Failed to read CSR: %v.", err)
		http.Error(w, http.StatusText(StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Parse CSR.
	block, err := pem.Decode(bytes)
	if err != nil {
		log.Printf("Failed to decode PEM: %v.", err)
		http.Error(w, http.StatusText(StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	csr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		log.Printf("Failed to parse CSR: %v.", err)
		http.Error(w, http.StatusText(StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Printf("--> %v.\n", csr.EmailAddresses)
}

func (h *handler) apply(w http.ResponseWriter, r *http.Request) {
	// Seems r *http.Request has TLS *tls.ConnectionState field which in turn has PeerCertificates []*x509.Certificate field, so
}

func main() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

	s = httptest.NewUnstartedServer(http.HandlerFunc(ok))
	s.TLS = &tls.Config{
		Certificates: []tls.Certificate{servTLSCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	s.StartTLS()
	_, err = client.Get(s.URL)
	s.Close()
	fmt.Println(err)

}

const instructions = `Welcome to the Teleport interview challenge!

To apply, you'll generate and use a x509 client certificate. This sounds
complex, but don't worry, we'll walk you through it.

1. Generate a private key. You can pick any algorithm and key size you would
   like as long as it's supported by Go.
   
2. Generate a CSR (Certificate Signing Request) with the email address you
   would like to be contacted at in the SAN (Subject Alternative Name).

3. Submit your CSR for us to sign by submitting an HTTP POST to
      
     http://challenge.gravitational.io/register

   No need for JSON or form encoding, just send the CSR as text/plain.

4. With your signed certificate in hand, use it to submit an HTTP POST
   with an empty body to

     http://challenge.gravitational.io/apply
`
