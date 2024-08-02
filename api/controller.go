package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	queue "notification/connection/rabbitmq"
	msg "notification/dto"

	"github.com/thedevsaddam/renderer"
)

var (
	rnd = renderer.New(renderer.Options{ParseGlobPattern: "html/*.html"})
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "home", nil)
}

func MessengerPage(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "messenger", nil)
}

func Submit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	fmt.Println("BODY:", r.Body)     //get request method
	reqBody, _ := ioutil.ReadAll(r.Body)

	var message msg.Message
	json.Unmarshal(reqBody, &message)

	body, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}

}
func CreateNewMessage(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var message msg.Message
	json.Unmarshal(reqBody, &message)

	body, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	queue.StartPublisher(string(body))
}
