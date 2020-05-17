package requests

type BitlyLinkCreateRequest struct {
	Domain string `json:"domain"`
	LongUrl string `json:"long_url"`
}

type BitlyLinkCreateResponse struct {
	Id string `json:"id"`
}

type BitlyCustomLinkCreateRequest struct {
	BitlinkId string `json:"bitlink_id"`
}

type BitlyCustomLinkCreateResponse struct {
	CustomBitlink string `json:"custom_bitlink"`
}