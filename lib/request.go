package lib

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"strings"
)

func ParseRequestParams(inputStruct interface{},c *beego.Controller)(r ResultInfo){
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