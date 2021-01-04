package lib

// 结果
type ResultInfo struct {
	ErrNo    int          `json:"errno"`   //错误码
	ErrMsg   string       `json:"errmsg"`  //提示信息
	Data     interface{}  `json:"data"`    //数据
	RawErr   error        `json:"-"`       //原始错误
}

// 正确的错误码
const ERRNO_TRUE = 0

// IsOk 判断结果是否正确
func (r *ResultInfo)IsOk()bool{
	if r.ErrNo == ERRNO_TRUE{
		return true
	}
	return false
}

// ErrorResultInfo 错误的结果信息
func ErrorResulInfo(errNo int,errMsg string,data interface{},err error)(r ResultInfo){
	r.ErrNo = errNo
	msg,ok := ErrnoMap[errNo]
	if ok == false {
		msg = "unkown"
	}
	r.ErrMsg = msg
	if errMsg != ""{
		r.ErrMsg += "(" + errMsg + ")"
	}
	if r.ErrMsg != ""{
		Log.Error(r.ErrMsg)
	}
	r.Data = data
	r.RawErr = err
	return
}