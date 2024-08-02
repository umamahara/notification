package config

import (
	"fmt"
	"log"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	AMQPConnectionURL string
}

type QueueInfo struct {
	QueueName string
}

var Config = Configuration{
	AMQPConnectionURL: "amqp://guest:guest@localhost:5673/",
}

var Queue = QueueInfo{
	QueueName: "NotificationMessageQueue",
}

type EmailConfiguration struct {
	FROM_EMAIL     string
	EMAIL_PASSWORD string
	PORT           string
	HOST           string
	TO_EMAIL       string
	SUBJECT        string
}

type SMSConfiguration struct {
	ACCOUNT_SID string
	AUTH_TOKEN  string
	URL         string
	TO_PHONE    string
	FROM_PHONE  string
}

type SlackConfiguration struct {
	WEBHOOK_URL string
}

func GetEmailConfig(params ...string) EmailConfiguration {
	configuration := EmailConfiguration{}
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf("config/%s_email_config.json", env)
	log.Printf("FILE_NAME %s", fileName)
	gonfig.GetConf(fileName, &configuration)
	log.Printf("Config-FROM_EMAIL %s", configuration.FROM_EMAIL)
	return configuration
}

func GetSMSConfig(params ...string) SMSConfiguration {
	configuration := SMSConfiguration{}
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf("config/%s_sms_config.json", env)
	log.Printf("FILE_NAME %s", fileName)
	gonfig.GetConf(fileName, &configuration)
	log.Printf("Config-TO_PHONE %s", configuration.TO_PHONE)
	return configuration
}

func GetSlackConfig(params ...string) SlackConfiguration {
	configuration := SlackConfiguration{}
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf("config/%s_slack_config.json", env)
	log.Printf("FILE_NAME %s", fileName)
	gonfig.GetConf(fileName, &configuration)
	log.Printf("Config-WEBHOOK_URL %s", configuration.WEBHOOK_URL)
	return configuration
}
