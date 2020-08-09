package requests

import (
	"encoding/json"
	"fmt"
	"github.com/TannerMoore/BitlyCodeChallenge/requests"
	"github.com/TannerMoore/BitlyGameLinkGenerator/models"
	"net/http"
	"time"
)

func CreateLink(domain, LinkDestination, auth string) (string, error) {
	bitlyUserEndpoint := "https://api-ssl.bitly.com/v3/user/link_save"
	parameters := fmt.Sprintf("access_token=%s&longUrl=%s&domain=%s", auth, LinkDestination, domain)
	requestPath := fmt.Sprintf("%s?%s", bitlyUserEndpoint, parameters)

	request, err := http.NewRequest("GET", requestPath, nil)
	if err != nil {
		return "", err
	}
	request.Header.Add("Authorization", auth)

	defaultGroupIdMarshaled, err := requests.ExecuteRequest(request)
	if err != nil {
		// This error comes back blank sometimes, so I am printing it with both levels of verbosity
		return "", fmt.Errorf("Http request error: %v:%+v\n", err, err)
	}

	var linkCreationResponse models.BitlyLinkCreateResponseV3
	err = json.Unmarshal(defaultGroupIdMarshaled, &linkCreationResponse)
	if err != nil {
		return "", err
	}

	// Sleep after each request or we will get rate limited by Bitly
	time.Sleep(1 * time.Second)

	linkIdRaw := linkCreationResponse.Data.LinkSave.Link
	return linkIdRaw[8:], nil
}