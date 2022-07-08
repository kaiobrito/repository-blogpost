package todoapi

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func sendRequest(path string, method string, body io.Reader, headers map[string]string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 5,
	}
	fmt.Println(path, method)

	req, err := http.NewRequest(method, BASE_URL+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	return ioutil.ReadAll(res.Body)
}

func requestAndMarshall[Response any](path string, method string, body io.Reader, headers map[string]string) (*Response, error) {
	res, err := sendRequest(path, method, body, headers)
	if err != nil {
		return nil, err
	}

	var r Response
	err = json.Unmarshal(res, &r)

	return &r, err
}
