package gomail

import (
	"github.com/go-gomail/gomail"
	"strconv"
)

type MailConfForm struct {
	EmailName      string `json:"email_name" form:"email_name"`				// 发件人昵称
	EmailHost      string `json:"email_host" form:"email_host"`				// 主机地址
	EmailPort      string `json:"email_port" form:"email_port"`				// 端口
	EmailUser      string `json:"email_user" form:"email_user"`				// 发件邮箱地址
	EmailPwd       string `json:"email_pwd" form:"email_pwd"`				// 授权码
	EmailTest      string `json:"email_test" form:"email_test"`				// 测试邮箱地址
	EmailTestTitle string `json:"email_test_title" form:"email_test_title"`	// 测试邮件标题
	EmailTemplate  string `json:"email_template" form:"email_template"`		// 测试邮件正文
}

type Config struct {
	MailTo  []string
	Cc      string
	Subject string
	Config  MailConfForm
}

func SendMail(conf Config) error {
	port,err := strconv.Atoi(conf.Config.EmailPort)
	if err != nil{
		return err
	}
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(conf.Config.EmailUser, conf.Config.EmailName))
	m.SetHeader("To", conf.MailTo...)
	//m.SetAddressHeader("Cc", conf.Cc, conf.Config.EmailUser)
	m.SetHeader("Subject", conf.Subject)
	m.SetBody("text/html", conf.Config.EmailTemplate)
	d := gomail.NewDialer(conf.Config.EmailHost, port, conf.Config.EmailUser, conf.Config.EmailPwd)
	err = d.DialAndSend(m)
	return err
}
