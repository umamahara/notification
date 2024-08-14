# Notification Service

## Prerequisites for Running Code

1. **Erlang**: Make sure to install Erlang by following the instructions available at [Erlang Downloads](https://www.erlang.org/downloads).
2. **Docker Desktop**: Ensure that Docker Desktop is installed on your computer.

## Execution Instructions

To execute the application, follow these steps in the `/notification` directory:

1. Before running the application, ensure you have filled out the correct configuration details inside the files located in the `config` folder:
   - `dev_email_config.json`
   - `dev_slack_config.json`
   - `dev_sms_config.json`

2. Use the Docker Compose command:
   ```bash
   docker-compose up --build

### Note: 
Make sure you are in the correct directory when you run this command to avoid any issues

## Testing the Application
Once the application is running, you can test it using the following endpoints:

Base URL: http://localhost:10000/
Create Message Endpoint: http://localhost:10000/CreateMessage
HTTP Method: POST
You can test the CreateMessage endpoint with the following JSON body:

json

{
    "Id": "1", 
    "Title": "Newly Created Post", 
    "desc": "The description for my new post", 
    "channel": "SLACK",
    "content": "My post content" 
}

## RabbitMQ Server Details
You can access your RabbitMQ server using the following details:

Management URL: http://localhost:15672/#/
Username/Password:
Username: guest
Password: guest
## Setting Up Incoming Webhook in Slack
For instructions on setting up an incoming webhook in Slack, you can refer to the following video: [Setting Up Incoming Webhook in Slack](https://www.youtube.com/watch?v=6NJuntZSJVA)
