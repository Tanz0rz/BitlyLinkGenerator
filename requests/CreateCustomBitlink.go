package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/TannerMoore/BitlyCodeChallenge/requests"
	"net/http"
	"time"
)

func SetCustomBackHalf(bitlink, domain, gameName, linkName, auth string) error {
	bitlyUserEndpoint := fmt.Sprintf("https://api-ssl.bitly.com/v4/custom_bitlinks/%v/%v%v", domain, gameName, linkName)

	requestBody := BitlyCustomLinkCreateRequest{
		BitlinkId: bitlink,
	}

	requestBodyMarshalled, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("PATCH", bitlyUserEndpoint, bytes.NewBuffer(requestBodyMarshalled))
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", auth)

	defaultGroupIdMarshaled, err := requests.ExecuteRequest(request)
	if err != nil {
		return err
	}

	var customLinkCreateResponse BitlyCustomLinkCreateResponse
	err = json.Unmarshal(defaultGroupIdMarshaled, &customLinkCreateResponse)
	if err != nil {
		return err
	}

	// Sleep after each request or we will get rate limited by Bitly
	time.Sleep(1 * time.Second)

	return nil
}
