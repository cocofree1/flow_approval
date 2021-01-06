package lib

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

func ParseRequestBody(inputStruct interface{},c *beego.Controller)(r ResultInfo){
	var err error
	body := c.Ctx.Input.RequestBody
	strReqBody := strings.TrimSpace(string(body))
	if "" != strReqBody && strings.HasPrefix(strReqBody,"{"){
		if err = json.Unmarshal(body,inputStruct); nil != err {
			r = ErrorResulInfo(ERR_JSON_UNMARSHAL,err.Error(),nil,err)
		}
	}
	if r.IsOk(){
		r.Data = inputStruct
	}
	return
}

func ParseRequestParams(c *beego.Controller)(r ResultInfo){
	params := c.Ctx.Input.Params()
	id,err := strconv.Atoi(params[":id"])
	if err != nil{
		r = ErrorResulInfo(ERR_STRING_TO_INT,err.Error(),nil,err)
	}
	r.Data = id
	return
}