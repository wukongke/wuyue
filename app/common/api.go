package common

type Result struct {
	Code int64       `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func ResponseJSON(code int64, msg string, data map[string]interface{}) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
