package controllers

import "github.com/astaxie/beego"

//定义响应json数据格式
type ResponseInfo struct {
	Code string
	Message string
	Data interface{}
}

type BaseController struct {
	beego.Controller
}

func init() {

}