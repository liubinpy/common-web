package param

type LoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CaptchaVerifyParam struct {
	ID   string `json:"id"`
	Code string `json:"code"`
}
