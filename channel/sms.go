package channel

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
	"net/http"
	"net/url"
	"encoding/json"
	config "notification/config"
  )
  
  func SendSMS(msg string) {

	configuration := config.GetSMSConfig()
	// Set account keys & information
	accountSid := configuration.ACCOUNT_SID
	authToken := configuration.AUTH_TOKEN
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
  

	// Set up rand
	rand.Seed(time.Now().Unix())
  
	// Pack up the data for our message
	msgData := url.Values{}
	msgData.Set("To",configuration.TO_PHONE)
	msgData.Set("From",configuration.FROM_PHONE)
	msgData.Set("Body",msg)
	msgDataReader := *strings.NewReader(msgData.Encode())
  
	// Create HTTP request client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  
	// Make HTTP POST request and return message SID
	resp, _ := client.Do(req)
	if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
	  var data map[string]interface{}
	  decoder := json.NewDecoder(resp.Body)
	  err := decoder.Decode(&data)
	  if (err == nil) {
		fmt.Println(data["sid"])
	  }
	} else {
	  fmt.Println(resp.Status);
	}
  }