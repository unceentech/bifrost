package bifrost

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Bifrost struct {
	// @@ Optional
	// @@ BaseURL will be prepended to Url unless Url is absolute.
	// @@ It can be convenient to set BaseURL for an instance of Bifrost to pass relative URLs to methods of that instance.
	BaseUrl string
	// @@ Optional
	// @@ Timeout is the timeout to use for the request.
	Timeout time.Duration
	// @@ Optional
	// @@ Configuration for setting the Headers and Basic Authentication credentials.
	Config *BifrostConfig
}

/*
Create a new Bifrost instance
This instance gives access to helper methods. Acceptable parameters are:
  - @param {string or ""} The BaseUrl if provided, it will be prepended to the Url provided in the Get method.
  - @param { *BifrostConfig or nil } Config for setting the Headers and Basic Authentication credentials.
  - @return { *Bifrost } Bifrost instance is returned
*/
func NewClient(params ...interface{}) *Bifrost {
	var BaseUrl string
	var Config *BifrostConfig // Default configuration

	for _, param := range params {
		switch v := param.(type) {
		case string:
			BaseUrl = v
		case *BifrostConfig:
			Config = v
		default:
			BaseUrl = ""
			Config = nil
		}
	}

	return &Bifrost{
		BaseUrl: BaseUrl,
		Config:  Config,
	}
}

/*
Get receives a string Url
If a BaseUrl is provided, it will be prepended to the Url provided in the Get method.
This method also sets the content header to application/json.
  - @param {string} The request Url
  - @param { interface{} or nil } ResponseStruct for Unmarshalling the response body.
  - @return { interface{}, error } The response
*/
func (bi *Bifrost) Get(Url string, ResponseStruct interface{}) (BifrostResponse, error) {

	// Validate the Url, check if baseUrl is provided
	url, err := bi.validateUrl(Url)
	if err != nil {
		return BifrostResponse{}, err
	}

	request, requestErr := http.NewRequest("GET", url, nil)
	if requestErr != nil {
		return BifrostResponse{}, requestErr
	}

	// Check if a timeout was specified else use the default timeout
	var timeoutValue time.Duration
	if bi.Timeout != 0 {
		timeoutValue = bi.Timeout
	} else {
		timeoutValue = time.Second * 10
	}

	httpclient := &http.Client{
		Timeout: timeoutValue,
	}

	// Check if the Bifrost Config was set
	bi.checkBifrostConfig(request)

	// Send the request and get the response from the request
	response, err := httpclient.Do(request)
	if err != nil {
		return BifrostResponse{}, err
	}

	// Call the close function at the end of this function
	defer response.Body.Close()

	// Get the data in the response body
	responseByte, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		return BifrostResponse{}, readErr
	}

	// Declare an empty response, make it an empty interface since we don't konw the type of the response.
	var EmptyResponse interface{}

	// Check if a struct was passed, and use that
	if ResponseStruct != nil {
		EmptyResponse = ResponseStruct
	}

	// Parse the JSON response into the empty interface created.
	json.Unmarshal(responseByte, &EmptyResponse)

	return BifrostResponse{
		Status: response.StatusCode,
		StatusText: response.Status,
		Data: EmptyResponse,
		Response: response,
		Bytes: responseByte,
	}, nil
}

// func (bi *Bifrost) Post(Url string, Body interface{}, Config *BifrostConfig) (*http.Response, error) {
// 	url, err := bi.validateUrl(Url)
// 	if err != nil {
// 		return nil, err
// 	}

// }
