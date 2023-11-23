package mail

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

var (
	hostname = "mailhog"
	port     = 1025
	username = "user@example.com"
	password = "password"
	from     = "kintai@example.com"
	subject  = "ログインしてください"
	loginUrl = "http://localhost:4200/login"
)

func SendEmail(email string, initPassword string) {
	receiver := []string{email}
	body := fmt.Sprintf("下記のリンクからログインしてください。\n\n%s\n\nメールアドレス： %s\n初期パスワード： %s", loginUrl, email, initPassword)
	smtpServer := fmt.Sprintf("%s:%d", hostname, port)
	auth := smtp.CRAMMD5Auth(username, password)
	msg := []byte(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(receiver, ","), subject, body))

	if err := smtp.SendMail(smtpServer, auth, from, receiver, msg); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
