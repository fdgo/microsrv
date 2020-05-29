package common

type Redis struct {
	Enabled  bool     `json:"enabled"`
	Conn     string   `json:"conn"`
	DbNum    int      `json:"dbNum"`
	Password string   `json:"password"`
	Timeout  int      `json:"timeout"`
	Sentinel Sentinel `json:"sentinel"`
}
type Sentinel struct {
	Enabled bool   `json:"enabled"`
	Master  string `json:"master"`
	Nodes   string `json:"nodes"`
}
