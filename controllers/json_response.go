package controllers

import (
	"github.com/astaxie/beego"
)

type Json_Response struct {
	control beego.Controller
}

type JsonResponseSuccess struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type JsonResponseFailed struct {
	Status  string `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ResponseSuccess(controller *beego.Controller, data interface{}) {
	controller.Data["json"] = JsonResponseSuccess{"1", data}
	controller.ServeJSON()
}

func ResponseError(controller *beego.Controller, code string, message string) {
	controller.Data["json"] = JsonResponseFailed{"0", code, message}
	controller.ServeJSON()
}
