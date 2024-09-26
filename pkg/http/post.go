package httphelpers

import (
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

func Request(client HTTPClient, reqData []byte, url string, method string) ([]byte, int, error) {
	log.Println("[LOG][Request] Making request ...")
	payload := strings.NewReader(string(reqData))

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Println("[ERROR][Request] Error creating request", err)
		return nil, http.StatusInternalServerError, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println("[ERROR][Request] Error making request", err)
		return nil, http.StatusInternalServerError, err
	}

	if res == nil {
		log.Println("[ERROR][Request] Response is nil")
		return nil, http.StatusInternalServerError, errors.New("response is nil")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("[ERROR][Request] Error reading response", err)
		return nil, http.StatusBadRequest, err
	}
	if res.StatusCode >= 300 {
		log.Println("[ERROR][Request] Error in response", res.StatusCode)
		log.Println("[ERROR][Request] Response body", string(body))
		return body, res.StatusCode, errors.New("error in response")
	}

	return body, res.StatusCode, nil
}
