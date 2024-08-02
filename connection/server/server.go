package server

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "notification/api"
    queue "notification/connection/rabbitmq"
)


func StartHttpServer()  {
    myRouter := mux.NewRouter().StrictSlash(true)
    handleRequests(myRouter)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func handleRequests(myRouter *mux.Router) {
    myRouter.HandleFunc("/", api.HomePage)
    myRouter.HandleFunc("/ShowForm", api.MessengerPage)
    myRouter.HandleFunc("/SubmitForm", api.Submit).Methods("POST")
    myRouter.HandleFunc("/startConsumer", queue.ConsumerQueue)
    myRouter.HandleFunc("/CreateMessage", api.CreateNewMessage).Methods("POST")
}
