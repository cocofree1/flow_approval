package request

type UserLogin struct{
	Regex     string     `json:"regex"`
	Password  string     `json:"password"`
}

type UserRegister struct{
	UserName   string  	 `json:"user_name" description:"用户名"`
	Telephone  string  	 `json:"telephone" description:"电话号码"`
	Email      string  	 `json:"email" description:"邮箱"`
	Password   string    `json:"password" description:"密码"`
}
