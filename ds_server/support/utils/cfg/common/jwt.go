package common

type Jwt struct {
	SecretKey string `json:"secretKey"`
	Exptime   int64  `json:"exptime"`
}
