package main

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/emersion/go-imap/client"
)

func main() {
	// 连接到 IMAP 服务器
	// c, err := client.Dial("mail.thecompanydoesnotexist.store:143")
	// config :=
	c, err := client.DialTLS("mail.thecompanydoesnotexist.store:993", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Logout()

	// 登录
	if err := c.Login("u1@thecompanydoesnotexist.store", "12321"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("登录成功")

	// 列出邮箱
	// err = c.List("", "*", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, box := range boxes {
	// 	fmt.Println("邮箱:", box.Name)
	// }

	// SMTP 服务器地址
	// smtpServer := "mail.thecompanydoesnotexist.store:587" // 587 是常用的 SMTP 端口
	// auth := smtp.PlainAuth("", "u1@thecompanydoesnotexist.store", "12321", "mail.thecompanydoesnotexist.store")
	// // 邮件内容
	// from := "u1@thecompanydoesnotexist.store"
	// to := []string{"617178759@qq.com"}
	// msg := []byte("To: 617178759@qq.com\r\n" +
	// 	"Subject: Test Email\r\n" +
	// 	"\r\n" +
	// 	"This is a test email.\r\n")

	// // 发送邮件
	// err := smtp.SendMail(smtpServer, auth, from, to, msg)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("邮件发送成功")
}
