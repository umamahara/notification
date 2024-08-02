package channel

import (
	"log"
    config "notification/config"
	"fmt"
    "bytes"
    "net/smtp"
    "mime/quotedprintable"
	msg "notification/dto"
)

// smtpServer data to smtp server
type smtpServer struct {
    host string
    port string
}

   // Address URI to smtp server
func (s *smtpServer) Address() string {
    return s.host + ":" + s.port
}


func SendEmail(message msg.Message) {
	log.Printf("Inside Send Email")
    configuration := config.GetEmailConfig()
	log.Printf("FROM_EMAIL %s",configuration.FROM_EMAIL)
	log.Printf("TO_EMAIL %s",configuration.TO_EMAIL)
	
	
	from_email:= configuration.FROM_EMAIL
    password  := configuration.EMAIL_PASSWORD
    host      := configuration.HOST+configuration.PORT
    auth      := smtp.PlainAuth("", from_email, password, configuration.HOST)

     header := make(map[string]string)
     to_email        := configuration.TO_EMAIL
     header["From"]   = from_email
     header["To"]     = to_email
     header["Subject"]= configuration.SUBJECT

    header["MIME-Version"]              = "1.0"
    header["Content-Type"]              = fmt.Sprintf("%s; charset=\"utf-8\"", "text/html")
    header["Content-Disposition"]       = "inline"
    header["Content-Transfer-Encoding"] = "quoted-printable"

    header_message := ""
    for key, value := range header {
       header_message += fmt.Sprintf("%s: %s\r\n", key, value)
    }

    body := "<h1>"+message.Content+"</h1>"
    var body_message bytes.Buffer
    temp := quotedprintable.NewWriter(&body_message)
    temp.Write([]byte(body))
    temp.Close()

    final_message := header_message + "\r\n" + body_message.String()

    status  := smtp.SendMail(host, auth, from_email, []string{to_email}, []byte(final_message))
    if status != nil {
        log.Printf("Error from SMTP Server: %s", status)
    }
    log.Print("Email Sent Successfully")
}