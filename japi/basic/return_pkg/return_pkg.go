package return_pkg

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code   int32       `json:"code"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Detail string      `json:"detail"`
}

func NewResult() *Result {
	return &Result{
		Code:   -1,
		Data:   "NULL",
		Msg:    "NULL",
		Detail: "NULL",
	}
}

//---------------------------正常返回，有查询结果-----------------------------
func GinResponse(httpstatus int32, code int32, msg string, detail string, data interface{}, c *gin.Context) {
	json := NewResult()
	json.Code = code
	json.Msg = msg
	json.Detail = detail
	json.Data = data
	c.JSON(
		int(httpstatus),
		json,
	)
}

func NorMalResponse(code int32, msg string, data interface{}, r *http.Request, w http.ResponseWriter) {
	src := NewResult()
	src.Code = code
	src.Data = data
	src.Msg = msg
	buf, _ := json.Marshal(src)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(int(code))
	w.Write([]byte(string(buf)))
}
