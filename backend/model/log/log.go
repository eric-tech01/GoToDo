package logModel

type LogModel struct {
	Id      int    `json:"id"`
	Event   string `json:"event"`
	Content string `json:"content"` //发生内容

	CreateAt int64 `json:"createAt"`
	UpdateAt int64 `json:"updateAt"`
}

const (
	EVENT_TYPE_LOGIN        string = "login"
	EVENT_TYPE_LOGOUT       string = "logout"
	EVENT_TYPE_SHUTDOWN     string = "shutdown"
	EVENT_TYPE_UPDATE_MODEL string = "update_model"
	EVENT_TYPE_DETECT       string = "detect"
)
