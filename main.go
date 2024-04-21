package main 

import (
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func main() {
	abc := gomail.NewMessage()

	abc.SetHeader("From", "luumanhquan.91@gmail.com")
	abc.SetHeader("To", "abitato.jiro@gmail.com")
	abc.SetHeader("Subject", "Test email Subject abc")
	abc.SetBody("text/plain", "this is test body")

	a := gomail.NewDialer("smtp.gmail.com", 587, "luumanhquan.91@gmail.com", "######")
	if err := a.DialAndSend(abc); err != nil {
		fmt.Println(err)
		panic(err)
	}
}