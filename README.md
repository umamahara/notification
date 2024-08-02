#### Prerequisite for running code
1) Install erlang https://www.erlang.org/downloads
2) Install RabbitMQ https://www.rabbitmq.com/download.html



#### To execute the app, follow these steps in the /notification directory::
>> 1) First, compile the app with the following command:
>> ==> go build main.go
>> 3) Once the build is complete, run the app using this command:
>> ==> go run main.go
>> Executing these commands will successfully compile and run your Go application within the notification directory. Make sure you are in the correct directory when you run them to avoid any issues. 

>> Build the app using command inside the notification directory ->  **go build main.go**
>> Run the app using command inside the notification directory ->  **go run main.go**

#### To test the app:
>> http://localhost:10000/ 
>> http://localhost:10000/CreateMessage


Method: POST
{
    "Id": "1", 
    "Title": "Newly Created Post", 
    "desc": "The description for my new post", 
    "channel": "SLACK",
    "content": "My post content" 
}



#### RabbitMQ server details
>> **URL:** http://localhost:15672/#/ 
>> **Username/Password:** guest/guest 

#### Set up Incoming Webhook In Slack
>> https://www.youtube.com/watch?v=6NJuntZSJVA
