package main

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
)

// GetRequest general HTTP GET method
func GetRequest(headers map[string]string, url string) ([]byte, int, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 5 * time.Second}
	req, _ := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, 500, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, res.StatusCode, err
	}
	defer res.Body.Close()

	return body, res.StatusCode, nil
}
