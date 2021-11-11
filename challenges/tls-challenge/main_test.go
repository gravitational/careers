package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegister(t *testing.T) {
	h := &handler{}

	//handler := func(w http.ResponseWriter, r *http.Request) {
	//	io.WriteString(w, "<html><body>Hello World!</body></html>")
	//}

	bytes := strings.NewReader(csr)
	req := httptest.NewRequest("POST", "http://example.com/register", bytes)

	w := httptest.NewRecorder()
	h.register(w, req)

	resp := w.Result()
	//body, err := io.ReadAll(resp.Body)

	//fmt.Println(resp.StatusCode)
	//fmt.Println(resp.Header.Get("Content-Type"))
	//fmt.Println(string(body))

}

const csr = `-----BEGIN CERTIFICATE REQUEST-----
MIIEnzCCAocCAQAwWjELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUtU3RhdGUx
ITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDETMBEGCSqGSIb3DQEJ
ARYEYWZhZTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAM+nvrB7PFnG
oDXdWoSmszyx7XG/4fIY9TrUpDEtnPgZyuWJMaISLJ/xVEH4HS7efMRC8zO5wEKZ
AEfkT6q85iLwuzVu0465HSvMhqR/+dCrr6d3MvRXK809XVT3irFlHkNUbNvmTcIH
OWt+ZriT96sXqS1INaBzAn6o3dDhNCwGfiPCromYVFmydxj+cJUIuuhQK55YHt9c
Y69SLyCtYowYhTZ+9c3ZuireMA/HNfxa8ptPJ1E4/6d4Y0s/pHrFT3dFTgEwGNBW
cRIFVmM8MwQZdSiAEhtl8tmyZGvO1d1QwKFawCSub+YUlF8DoWJ+5SZ0fsYIminM
pTGtbaBa2zFBN/IiehiG2TWbXX8l8F3JdxtYUaT1TUsR4oZLMvpEcOvzR4AmGJq1
sESV4s6mS/6ZP+cMk2zIOkxsV0UbNLHplsiu85l2uBIuTDYXPHct4iAI7YWQO/dv
T+JPfurZSRt+86WnberjCYgvJUwc/IK8VCXl9OUkuo3LNURdx0SRD/PF3HiUHnif
HTpvejj4YluE2dFqiPK11KEBXqr2VGBLbbB28o5wW9+nSK88C9ZgVCgGQ9l8rBfq
ZK45uHcUq//O/uSPBN3MUuWrwkJ1S0+E417g1Rpc8Wd6dYHJXiIEzI38MPmIlnQT
HVxnf3/ysPlk99PipSgvap017cdbgCztAgMBAAGgADANBgkqhkiG9w0BAQsFAAOC
AgEAVLbUmIto3qnwgBslZDnHyZKuvA0/6O64g6Qi5Um9gBwbvSbYO9XNMcjlPWGl
hqFwFA6m5lkADFZZkGA1HejZcL2PFUOFS12KEAZwo8dY5utAoa2J+BmmFRQuJhJm
PmpIHVnkSd+u2oNjPkjmEFZxns3ueVqJHbLHe52G8hQsj60mWMYQ7hea6zikCHEt
v57pzKdX7PwCohXvyPzcdW8BhvRV+bNKv2gRVaUMDZMtHjBd6GqLbI6fE+e+lLfx
mPzRvGQ8ilfjkipnKFqZaOzBYoCMjGGAqEPWA3PDN+CEkZg8oa377Hqwg7I5Ap+D
CA9znktwHCjgAPYkbmRESKSxf5YqVFbOWilvcb1Q5XyQ8DJr2YBTJrANPXfvFejI
pZia+SubtSNZeeE9xZW/RthW1XS9eCy5qC6OU610h4hpcqCOtdQGhr9mGwHwVm6h
Z5dDeyW9ycJ3mpOngkjehTrEFDLTwhUbH9r29q3JQMj2UjELbAY4XShvK7PjZ/NW
1Sa7uIWOZ+v0vmyCWq0g9XS8ar4cmxRmqNE595JypjYC2u2up+4ERvtP8RqAgdP+
Y67VuHfcbsZUeaYB77LBILnXMv9xYbW0/QZgGjhpebgBbDVXy/jOPGAxA07jpBpV
+EJ2Hcgf9r5JKgsSirnW9Eyp+UaTbt9cDxwoI5VQhIT8YJ0=
-----END CERTIFICATE REQUEST-----`
