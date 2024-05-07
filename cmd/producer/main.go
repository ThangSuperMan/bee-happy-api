package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
	AppSetting "github.com/thangsuperman/bee-happy/config"
	"github.com/thangsuperman/bee-happy/internal"
)

// NOTE: create exchange(include exchange type) and set permission for how the user interact with the exchange

func main() {
	conn, err := internal.ConnectRabbitMQ(
		AppSetting.Envs.RabbitMQUsername,
		AppSetting.Envs.RabbitMQPassword,
		fmt.Sprintf("%s:%s", AppSetting.Envs.RabbitMQHost, AppSetting.Envs.RabbitMQPort),
		"customers")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client, err := internal.NewRabbitMQClient(conn)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	queue1 := "customers_created"
	queue2 := "customers_test"
	customerEventExchange := "customer_events"

	if err := client.CreateQueue(queue1, true, false); err != nil {
		panic(err)
	}

	if err := client.CreateQueue(queue2, false, false); err != nil {
		panic(err)
	}

	if err := client.CreateBinding(queue1, "customers.created.*", customerEventExchange); err != nil {
		panic(err)
	}

	if err := client.CreateBinding(queue2, "customers.*", customerEventExchange); err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Send(ctx, customerEventExchange, "customers.created.us", amqp091.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp091.Persistent,
		Body:         []byte(`Another cool message between services`),
	}); err != nil {
		panic(err)
	}

	if err := client.Send(ctx, customerEventExchange, "customers.test", amqp091.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp091.Persistent,
		Body:         []byte(`A uncool durable message`),
	}); err != nil {
		panic(err)
	}

	time.Sleep(20 * time.Second)

	log.Println(client)
}
