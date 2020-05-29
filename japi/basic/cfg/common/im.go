package common

import "strconv"

// AppCfg common config
type ImCfg struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

func (a *ImCfg) Addr() string {
	return a.Address + ":" + strconv.Itoa(a.Port)
}
