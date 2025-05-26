package coolify

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

type CoolifyRequestService struct {
	token        string
	base_api_url string
}

func NewCoolifyRequestService(api_url string, token string) *CoolifyRequestService {
	return &CoolifyRequestService{
		base_api_url: api_url,
		token:        token,
	}
}

func (svc *CoolifyRequestService) requestWithData(method string, endpoint string, data []byte) ([]byte, error) {
	url := svc.base_api_url + endpoint

	var body *bytes.Buffer
	if data != nil {
		body = bytes.NewBuffer(data)
	} else {
		body = bytes.NewBuffer([]byte{})
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", svc.token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resBody, _ := io.ReadAll(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Printf("Received unexpected response: [%d]: %s\n", resp.StatusCode, resBody)
		return nil, err
	}

	return resBody, nil
}

func (svc *CoolifyRequestService) Get(endpoint string) ([]byte, error) {
	return svc.requestWithData("GET", endpoint, nil)
}

func (svc *CoolifyRequestService) Post(endpoint string, data []byte) ([]byte, error) {
	return svc.requestWithData("POST", endpoint, data)
}

func (svc *CoolifyRequestService) Patch(endpoint string, data []byte) ([]byte, error) {
	return svc.requestWithData("PATCH", endpoint, data)
}

func (svc *CoolifyRequestService) Put(endpoint string, data []byte) ([]byte, error) {
	return svc.requestWithData("PUT", endpoint, data)
}

func (svc *CoolifyRequestService) Delete(endpoint string) ([]byte, error) {
	return svc.requestWithData("DELETE", endpoint, nil)
}
