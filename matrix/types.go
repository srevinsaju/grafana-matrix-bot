package matrix



type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Homeserver string `json:"homeserver"`
	ChannelId string `json:"channel_id"`
	SecretKey string `json:"secret_key"`
	Port int `json:"port"`
}