package util

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func PostJson(url string, bodyJson string, authorization string) string {

	contentType := "application/json"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(bodyJson)))

	if err != nil {
		log.Println(err)
		return ""
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authorization))
	req.Header.Set("Content-Type", contentType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}

	all, err := io.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()

	}(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(all)
}

func GetJson(url string, authorization string) string {

	contentType := "application/json"

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
		return ""
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authorization))
	req.Header.Set("Content-Type", contentType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}

	all, err := io.ReadAll(resp.Body)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if err != nil {
		log.Println(err)

		return ""
	}
	return string(all)
}
