# PUB-SUB Project

## Installing RabbitMQ with docker

- Download & Install docker `https://www.docker.com/get-started`

- Pull the docker image using below command

  ```
  docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
  ```

- Once done open `localhost:15672` for management console of rabbitMQ

## Setting up the Database

- Create a MySQL Database with name **pubsub**

- Create a new user for that database with username & password as **pubsub**

## Running the project

- Run the _Publisher_
  ```
  cd publisher
  go run pub.go
  ```
- Run the _Subscriber_
  ```
  cd publisher
  go run pub.go
  ```
