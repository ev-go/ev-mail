package main

import (
	"log"
	"net/smtp"
)

func main() {
	//auth := smtp.PlainAuth("", "evmailforgs@yandex.ru", "ttvopfpmphnwumui", "smtp.yandex.ru")
	//err := smtp.SendMail("smtp.yandex.ru:465", auth, "evmailforgs@yandex.ru", []string{"evmailforgs@yandex.ru"}, []byte("Текст письма."))
	//if err != nil {
	//	log.Fatal(err)
	//}
	////SendMail("smtp.mail.ru:995", "evmailforgs@mail.ru", "Subject text", "Body text", []string{"evmailforclient@mail.ru"})

	auth := smtp.PlainAuth("", "evmailforgs@yandex.ru", "ttvopfpmphnwumui", "smtp.yandex.ru")
	err := smtp.SendMail("smtp.yandex.ru:25", auth, "evmailforgs@yandex.ru", []string{"evmailforgs@gmail.com"}, []byte(" \r complete приложение тест \n "))
	if err != nil {
		log.Fatal(err)
	}
}

//func SendMail(addr, from, subject, body string, to []string) error {
//	r := strings.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")
//
//	c, err := smtp.Dial(addr)
//	if err != nil {
//		return err
//	}
//	defer c.Close()
//	if err = c.Mail(r.Replace(from)); err != nil {
//		return err
//	}
//	for i := range to {
//		to[i] = r.Replace(to[i])
//		if err = c.Rcpt(to[i]); err != nil {
//			return err
//		}
//	}
//
//	w, err := c.Data()
//	if err != nil {
//		return err
//	}
//
//	msg := "To: " + strings.Join(to, ",") + "\r\n" +
//		"From: " + from + "\r\n" +
//		"Subject: " + subject + "\r\n" +
//		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
//		"Content-Transfer-Encoding: base64\r\n" +
//		"\r\n" + base64.StdEncoding.EncodeToString([]byte(body))
//
//	_, err = w.Write([]byte(msg))
//	if err != nil {
//		return err
//	}
//	err = w.Close()
//	if err != nil {
//		return err
//	}
//	return c.Quit()
//}
