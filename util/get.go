package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)


func TestConnectivity(url string, timeout time.Duration) bool {
	client := http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:   true, 
			MaxIdleConnsPerHost: 512,  
		},
		Timeout: timeout, 
	}
	req, err := http.NewRequest("GET", url, bytes.NewReader([]byte("")))
	if err != nil {
		return false
	}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	return true
}
