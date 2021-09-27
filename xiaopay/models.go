package xiaopay

type RespBase struct {
	Success      bool        `json:"Success"`
	Code         string      `json:"Code"`
	ErrorMessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data"`
}
