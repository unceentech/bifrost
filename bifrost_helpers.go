package bifrost

import (
	"errors"
	"net/http"
)

// Validate the request Url
// Check if the BaseUrl is provided and construct the proper url string.
func (bi *Bifrost) validateUrl(url string) (string, error) {
	// check if both BaseUrl and Url are empty
	if bi.BaseUrl == "" && url == "" {
		return "", errors.New("BaseUrl or Url cannot be empty")
	}

	// check if both BaseUrl and Url are not empty
	if bi.BaseUrl != "" && url != "" {
		requestUrl := bi.BaseUrl + url
		return requestUrl, nil
	}

	// check if BaseUrl is empty and Url is not empty
	if bi.BaseUrl == "" && url != "" {
		return url, nil
	}

	// check if BaseUrl is not empty and Url is empty
	if bi.BaseUrl != "" && url == "" {
		return bi.BaseUrl, nil
	}

	return "", errors.New("BaseUrl or Url cannot be empty")
}

// Check if BifrostConfig was provided, and construct the configuration properly
// If config was not provided, set some default values
func (bi *Bifrost) checkBifrostConfig(req *http.Request) {
	// SET HEADERS if provided, else set the default headers to application/json
	if bi.Config != nil {
		if bi.Config.Headers != nil {
			// Add headers
			// READ AS: for every key (k) and value (v) in Config.Headers
			for k, v := range bi.Config.Headers {
				req.Header.Add(k, v)
			}
		} else {
			req.Header.Add("Content-Type", "application/json")
		}

		// SET AUTH if provided
		if bi.Config.Auth != nil {
			// Add auth
			req.SetBasicAuth(bi.Config.Auth.username, bi.Config.Auth.password)
		}
	}
}
