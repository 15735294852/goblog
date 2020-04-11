package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data  interface{} `json:"data"`
}

type BaseController struct {
	Response Response
}

func (base *BaseController) Success(w http.ResponseWriter) {
	base.Response.Code = 0
	base.Response.Message = "Success"
	result,err := json.Marshal(base.Response)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	w.Write(result)
}

