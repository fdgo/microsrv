package common

import "strconv"

// AppCfg common config
type SrvCfg struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

func (a *SrvCfg) Addr() string {
	return a.Address + ":" + strconv.Itoa(a.Port)
}
