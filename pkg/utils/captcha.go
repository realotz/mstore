package utils

import (
	captcha "github.com/mojocn/base64Captcha"
	"image/color"
)

func NewCaptchaDriver() *captcha.DriverString {
	return captcha.NewDriverString(
		44,
		122,
		0,
		0,
		4,
		"qwertyuipkjhgfdsazxcvbnm",
		&color.RGBA{R: 255, G: 255, B: 255, A: 255},
		nil,
		[]string{"wqy-microhei.ttc"},
	)
}
