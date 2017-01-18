package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
)

const (
	ID_PARAMETER = ":id"
)

func buildResponse(err string, ctx *context.Context, response interface{}) {
	if len(err) > 0 {
		ctx.Output.Body([]byte(err))
	} else {
		responseBytes, err := json.Marshal(response)
		checkErr(err)
		ctx.Output.Body([]byte(responseBytes))
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


