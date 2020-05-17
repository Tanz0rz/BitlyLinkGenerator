package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/TannerMoore/BitlyCodeChallenge/requests"
	"github.com/TannerMoore/BitlyGameLinkGenerator/models"
	"net/http"
	"time"
)

func CreateLink(domain, LinkDestination, auth string) (string, error) {
	bitlyUserEndpoint := "https://api-ssl.bitly.com/v4/bitlinks"

	requestBody := models.BitlyLinkCreateRequest{
		Domain: domain,
		LongUrl: LinkDestination,
	}

	requestBodyMarshalled, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", bitlyUserEndpoint, bytes.NewBuffer(requestBodyMarshalled))
	if err != nil {
		return "", err
	}
	request.Header.Add("Authorization", auth)

	defaultGroupIdMarshaled, err := requests.ExecuteRequest(request)
	if err != nil {
		return "", fmt.Errorf("Http request error: %+v\n", err)
	}

	var linkCreationResponse models.BitlyLinkCreateResponse
	err = json.Unmarshal(defaultGroupIdMarshaled, &linkCreationResponse)
	if err != nil {
		return "", err
	}

	// Sleep after each request or we will get rate limited by Bitly
	time.Sleep(1 * time.Second)

	return linkCreationResponse.Id, nil
}