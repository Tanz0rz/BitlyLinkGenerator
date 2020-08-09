package models

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

// v3

type BitlyLinkCreateResponseV3 struct {
	Data BitlyLinkCreationResponseV3Data `json:"data"`
}

type BitlyLinkCreationResponseV3Data struct {
	LinkSave BitlyLinkCreationResponseV3LinkSave `json:"link_save"`
}

type BitlyLinkCreationResponseV3LinkSave struct {
	Link string `json:"link"`
}