package mail

import "gopkg.in/gomail.v2"

func Send(title string, content string, to string) (err error) {
	if to == "" {
		to = "zhou7419@foxmail.com"
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "zhou7419@qq.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", title)
	m.SetBody("text/plain", content)

	d := gomail.NewDialer("smtp.qq.com", 587, "zhou7419@qq.com", "yqgbptaimgdjbcjh")
	err = d.DialAndSend(m)

	return err
}
