package v2

import (
	"net/http"
	. "github.com/cool2645/kotori-ng/httputils"
)

func Pong(w http.ResponseWriter, req *http.Request) {
	res := map[string]interface{}{
		"code":   http.StatusOK,
		"result": true,
		"msg":    "OK",
	}
	RespondJson(w, res, http.StatusOK)
	return
}
