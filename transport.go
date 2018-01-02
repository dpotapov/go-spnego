package spnego

import "net/http"

// Transport extends the native http.Transport to provide SPNEGO communication
type Transport struct {
	http.Transport
	spnego Provider
}

// RoundTrip implements the RoundTripper interface.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.spnego == nil {
		t.spnego = New()
	}

	if err := t.spnego.SetSPNEGOHeader(req); err != nil {
		return nil, err
	}

	return t.Transport.RoundTrip(req)
	// ToDo: process negotiate token from response
}
