package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HttpDo(url string, header map[string]string, httpMethod string, httpPostBody string) []byte {

	client := &http.Client{}
	req, err := http.NewRequest(strings.ToUpper(httpMethod), url, strings.NewReader(httpPostBody))
	if err != nil {

	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}

	fmt.Println(string(body))
	return body
}
