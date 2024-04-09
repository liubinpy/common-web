package data

type CaptchaResponse struct {
	ID         string `json:"id"`
	Base64Data string `json:"base64_data"`
}
