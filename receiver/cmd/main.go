package main

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sinyavcev/microservices/receiver/model"
	"github.com/sinyavcev/microservices/receiver/pkg/repository"
	"github.com/sinyavcev/microservices/receiver/server"
	"github.com/spf13/viper"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	client, err := repository.Init(context.Background(), viper.GetString("db.dbname"))

	db := repository.NewRepository(client.Collection(viper.GetString("db.dbname")))

	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
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
	var DataArr []model.Stock
	go func() {
		for d := range msgs {
			json.Unmarshal(d.Body, &DataArr)
			for _, item := range DataArr {
				log.Printf("stock: %+v\\", item)
				db.Create(context.Background(), item)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	server.Run(*client)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
