package config

type LinkConfigList []LinkConfig

type LinkConfig struct {
	Name string `json:"name"`
	JamPage string `json:"jam_page"`
	Ost string `json:"ost"`
	Github string `json:"github"`
	ItchIo string `json:"itch_io"`
}
