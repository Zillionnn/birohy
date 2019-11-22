package main

import(
	"fmt"
	"time"
	"net/smtp"
	"net/mail"
	"encoding/base64"
)

// var bodyState map[string]string
// bodyState = make(map[string]string)
// bodyState["HIGH"]= "HIGH PERIOD"
// bodyState["LOW"]= "LOW PERIOD"
// bodyState["CRITICAL"]= "WARNING, IN CRITICAL DAY!"


func getDaysFromBirth(birthday string, date string) int64{
	m_birthday,_ := time.Parse("01/02/2006", birthday);
	m_date, _ :=time.Parse("01/02/2006",date);
	duration := (m_date.Unix() - m_birthday.Unix()) / 60/60/24
	return	duration
}
func formatState(stateType string, days int64) string{
	day_s := fmt.Sprintf("%d", days)
	s:= stateType + " " + day_s + " "
	if stateType=="phiscal" {
		if days < 12 {
			s+= "HIGH PERIOD"
		} else if days ==12 {
			s+= "WARNING, IN CRITICAL DAY!"
		} else {
			s+="LOW PERIOD"
		}
	}
	if stateType=="emotion" {
		if days < 14 {
			s+= "HIGH PERIOD"
		} else if days ==14 {
			s+= "WARNING, IN CRITICAL DAY!"
		} else {
			s+="LOW PERIOD"
		}
	}
	if stateType=="inte" {
		if days < 17 {
			s+= "HIGH PERIOD"
		} else if days ==17 {
			s+= "WARNING, IN CRITICAL DAY!"
		} else {
			s+="LOW PERIOD"
		}
	}
	s+="\n"
	return s
}

func Up() {
	tm := time.Now()
	b64 := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
    host := "smtp.163.com"
    email := "skyestzhang@163.com"
    password := "4501122z"
    toEmail := "zhang_tianji@hoperun.com"
    from := mail.Address{">>BORIHY<<", email}
    to := mail.Address{"接收人", toEmail}
    header := make(map[string]string)
    header["From"] = from.String()
    header["To"] = to.String()
    header["Subject"] = fmt.Sprintf("=?UTF-8?B?%s?=", b64.EncodeToString([]byte("🚀节律"+tm.Format("01/02/2006"))))
    header["MIME-Version"] = "1.0"
    header["Content-Type"] = "text/html; charset=UTF-8"
	header["Content-Transfer-Encoding"] = "base64"

	body := "";
	
	duration := getDaysFromBirth("08/07/1993", tm.Format("01/02/2006"))
	phiscal := duration  % 23
	emotion := duration %28
	inte:=duration%33
	body += formatState("phiscal", phiscal)
	body += formatState("emotion", emotion)
	body += formatState("inte", inte)

 
    message := ""
    for k, v := range header {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + b64.EncodeToString([]byte(body))
    auth := smtp.PlainAuth(
        "",
        email,
        password,
        host,
    )
    err := smtp.SendMail(
        host+":25",
        auth,
        email,
        []string{to.Address},
        []byte(message),
    )
    if err != nil {
        panic(err)
    }
	}
	
func main()  {
	
	Up()
}