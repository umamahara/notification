package rabbitmq

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	config "notification/config"
	channel "notification/channel"
	"github.com/streadway/amqp"
	msg "notification/dto"
)

var ch amqp.Channel
var q amqp.Queue
var err1 error

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
func StartPublisher(body string)  {
	log.Printf("Inside Publisher....")
	log.Printf(body)
	conn, err := amqp.Dial(config.Config.AMQPConnectionURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	
    q, err := ch.QueueDeclare(
		config.Queue.QueueName, // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	log.Printf("Received a message: %s", q.Name)

	err1 = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err1, "Failed to publish a message")
}

func ConsumerQueue(w http.ResponseWriter, r *http.Request)  {
	log.Printf("Inside consumer Receiver")
	conn, err := amqp.Dial(config.Config.AMQPConnectionURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		config.Queue.QueueName, // name
		false,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			var message msg.Message 
			json.Unmarshal(d.Body, &message)
			log.Printf("Inside Send Channel %s",message.Channel)

			switch { 
				case message.Channel == "EMAIL": 
					fmt.Println("Sending EMAIL .... ")
					channel.SendEmail(message) 
				case message.Channel == "SLACK": 
					fmt.Println("Sending message on SLACK ....") 
					err := channel.SendSlackNotification(message.Content)
					if err != nil {
						log.Fatal(err)
					}
				case message.Channel == "SMS": 
					fmt.Println("Sending SMS ....") 
					channel.SendSMS(message.Content)
				} 
			
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
