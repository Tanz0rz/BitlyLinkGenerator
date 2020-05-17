package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type BitlyError struct {
	Field     string `json:"field"`
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
}

type BitlyErrorResponse struct {
	Message     string       `json:"message"`
	Errors      []BitlyError `json:"errors"`
	Resource    string       `json:"resource"`
	Description string       `json:"description"`
}

func ExecuteRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("unable to communicate with remote server: %v", err)
		return []byte{}, err
	}

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("unexpected status code from remote server: %v; unable to read response: %v", resp.StatusCode, err)
	}
	return respBytes, nil
}
