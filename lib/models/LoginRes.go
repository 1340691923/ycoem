package model

type LoginRes struct {
	Code     int                    `json:"code"`
	Msg      string                 `json:"msg"`
	Uid      int64                  `json:"uid"`
	Token    string                 `json:"token"`
	Config   map[string]interface{} `json:"config"`
	X        bool                   `json:"x"`
	ChanUser string                 `json:"chan_user"`
	InSence  bool                   `json:"inSence"`
}
