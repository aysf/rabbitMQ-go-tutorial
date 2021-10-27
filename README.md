# rabbitMQ-go-tutorial

## how to use this tutorial

1. see and try to understand the writting article in https://tutorialedge.net/golang/go-rabbitmq-tutorial/
2. clone this repo `git clone git@github.com:aysf/rabbitMQ-go-tutorial.git`
3. change the directory `cd rabbitMQ-go-tutorial`
4. install and/or run RabbitMQ via docker

```sh
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.9-management
```

5. run the publisher `go run main.go`
6. run the consumer app service `go run consumer.go`

## Screenshot
