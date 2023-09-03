package bifrost

import "net/http"



type BifrostBasicAuth struct {
	username string
	password string
}
type BifrostConfig struct {
	// The Headers to set on the request. This is a map
	Headers map[string]string
	// The Auth to set on the request if required. This is a struct of Basic authentication
	Auth *BifrostBasicAuth
}

type BifrostResponse struct {
	Status int // `Status` is the HTTP status code from the server response
	StatusText string // `statusText` is the HTTP status message from the server response
	Data interface{} // The unmarshalled response body
	Response *http.Response // The unfiltered *http.Respons object
	Bytes []byte // The response body as byte
}
