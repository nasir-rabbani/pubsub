# PUB-SUB Project

## Installing RabbitMQ with docker

- Download & Install docker `https://www.docker.com/get-started`

- Pull the docker image using below command

  ```
  docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
  ```

- Once done open `localhost:15672` for management console of rabbitMQ

## Running the project

- Clone the project in your go workspace

  ```git
  git clone https://github.com/nasir-rabbani/pubsub.git
  ```

- Modify the .yaml files inside below locations as per your configurations before running

  ```
  subscriber/configs

  publisher/configs
  ```

- Run the _Subscriber_
  ```go
  cd subscriber
  go run main.go
  ```
- Run the _Publisher_
  ```go
  cd publisher
  go run main.go
  ```
