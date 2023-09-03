<p align="center">
<h1 align="center">Bifrost Rest Client</h1>
<p align="center">Simple HTTP and REST client library for Go (inspired by Javascript's Got and Axios Rest client)</p>
<p align="center"><a href="#features">Features</a> section describes in detail about Bifrost capabilities</p>
</p>
<p align="center">
<h4 align="center">Bifrost Communication Channels</h4>
<p align="center"><a href="https://gitter.im/go_resty/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge"><img src="https://badges.gitter.im/go_resty/community.svg" alt="Chat on Gitter - Resty Community"></a> <a href="https://twitter.com/go_resty"><img src="https://img.shields.io/badge/twitter-@go__resty-55acee.svg" alt="Twitter @go_resty"></a></p>
</p>


## Table of Contents

  - [Features](#features)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Creating an instance](#creating-an-instance)
  - [Instance methods](#instance-methods)
  - [Initializing Bifrost](#initializing-bifrost)
  - [Response Struct](#response-struct)
  - [Interceptors](#interceptors)
    - [Multiple Interceptors](#multiple-interceptors)
  - [Contribution](#contribution)


## Features
- [x] Create and run REST HTTP requests
- [x] Basic configuration for REST HTTP requests
- [x] Create instance with common configuration for reusability
- [x] Validate Bifrost Struct before running request
- [ ] Interceptors for before Request and after Response

## Installation

```bash
go get github.com/unceentech/bifrost
```

## Usage

### Run a simple GET request

```go
package main

import (
    "fmt"
    "github.com/unceentech/bifrost"
)

func main() {
    bi := bifrost.NewClient()
    res, err := bi.Get("https://some-domain.com/api/v1/event/get", nil)
    if err != nil {
        log.Fatalf("err: %v", err)
    }

    fmt.Printf("Http Response: ", res)
}
```

### Creating an instance

You can create a new instance of axios with a custom config.

##### With BaseUrl: bifrost.NewClient(...config)

```go
  bi := bifrost.NewClient("https://some-domain.com/api/")
  res, err := bi.Get("v1/event/get", nil)
  if err != nil {
        log.Fatalf("err: %v", err)
  }
```
##### With Header: bifrost.NewClient(...config)
```go
  // Create a BifrostConfig instance
  config := bifrost.BifrostConfig{
      Headers: map[string]string{
          "Authorization": "Bearer YourAccessToken",
          "Content-Type": "application/json",
      },
  }

  bi := bifrost.NewClient("https://some-domain.com/api/", &config)
  res, err := bi.Get("v1/event/get", nil)
  if err != nil {
        log.Fatalf("err: %v", err)
  }
```

##### With BasicAuth: bifrost.NewClient(...config)
```go
  // Create a BifrostConfig instance
  config := bifrost.BifrostConfig{
      Auth: &bifrost.BifrostBasicAuth{
          Username: "yourUsername",
          Password: "yourPassword",
      },
  }

  bi := bifrost.NewClient("https://some-domain.com/api/", &config)
  res, err := bi.Get("v1/event/get", nil)
  if err != nil {
        log.Fatalf("err: %v", err)
  }
```

### Instance methods

The available instance methods are listed below. The specified config will be merged with the instance config.

##### bifrost#Get(url, [, config])


## Initializing Bifrost

These are the available config options for making requests in the NewClient function.

```go
{

  // `BaseURL` will be prepended to `url` unless `url` is absolute.
  // It can be convenient to set `baseURL` for an instance of axios to pass relative URLs
  // to methods of that instance.
  BaseURL: 'https://some-domain.com/api/',

  // `timeout` specifies the number of milliseconds before the request times out.
  // If the request takes longer than `timeout`, the request will be aborted.
  Timeout: 1000, // default is `0` (no timeout)


  config: BifrostConfig{
      // `Headers` to be used to make the request. This is optional
      Headers: map[string]string{
          "Authorization": "Bearer YourAccessToken",
          "Content-Type": "application/json",
      },

      // `Auth` indicates the HTTP Basic auth to be used to make the request. This is optional
      Auth: &BifrostBasicAuth{
          Username: "yourUsername",
          Password: "yourPassword",
      },
  }
}
```

## Response Struct

The response for a request contains the following information.

```go
{
  // `Data` is the unmarshalled response that was provided by the server
  Data: interface{},

  // `Status` is the HTTP status code from the server response
  Status: 200,

  // `statusText` is the HTTP status message from the server response
  StatusText: 'OK',

  // `Response` is the unfiltered *http.Response object
  Response *http.Response,

  // `Bytes` is the response body as byte
  Bytes []byte,
}
```

## Contribution

Contributions are welcome. If you find any improvement or issue you want to fix feel free to send a pull request [pull request](https://github.com/unceentech/bifrost/pulls) or open an [issue](https://github.com/unceentech/bifrost/issues) to discuss the change you wish to make.