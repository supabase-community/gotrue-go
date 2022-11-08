package types

type SecurityEmbed struct {
	Security GoTrueMetaSecurity `json:"gotrue_meta_security"`
}

type GoTrueMetaSecurity struct {
	CaptchaToken string `json:"captcha_token"`
}
