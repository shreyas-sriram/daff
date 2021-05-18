package daff

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// createRequest parses the challenge details and create a HTTP request
func createRequest(c Challenge) (*http.Request, error) {
	headers, err := parseKeyValue(c.Request.Headers)
	if err != nil {
		log.Printf("Error parsing request headers: %v\n", err)
		return nil, err
	}

	log.Printf("Request headers: %+v\n", headers)

	cookies, err := parseKeyValue(c.Request.Cookies)
	if err != nil {
		log.Printf("Error parsing request headers: %v\n", err)
		return nil, err
	}

	log.Printf("Request cookies: %+v\n", cookies)

	req, err := http.NewRequest(c.Request.Method, c.URL, nil)
	if err != nil {
		log.Printf("Error creating new HTTP request: %v\n", err)
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	for key, value := range cookies {
		req.AddCookie(&http.Cookie{Name: key, Value: value})
	}

	if c.Request.Method == http.MethodPost {
		req.Body = ioutil.NopCloser(strings.NewReader(c.Request.Body))
	}

	return req, nil
}

// sendRequest sends the HTTP request and returns the response
func sendRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request: %v\n", err)
		return nil, err
	}

	return res, nil
}
