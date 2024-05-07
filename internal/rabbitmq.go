package internal

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func ConnectRabbitMQ(username string, password string, host string, vhost string) (*amqp.Connection, error) {
	return amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", username, password, host, vhost))
}

func NewRabbitMQClient(conn *amqp.Connection) (RabbitClient, error) {
	channel, err := conn.Channel()
	if err != nil {
		return RabbitClient{}, err
	}

	return RabbitClient{conn: conn, channel: channel}, nil
}

func (rc RabbitClient) Close() error {
	return rc.channel.Close()
}

func (rc RabbitClient) CreateQueue(queueName string, durable bool, autodelete bool) error {
	_, err := rc.channel.QueueDeclare(queueName, durable, autodelete, false, false, nil)
	return err
}

func (rc RabbitClient) CreateBinding(queueName string, binding string, exchange string) error {
	// NOTE: leaving nowait false, having nowait set to false will make the channel return an error if its failed to bind
	return rc.channel.QueueBind(queueName, binding, exchange, false, nil)
}

func (rc RabbitClient) Send(ctx context.Context, exchange string, routingKey string, options amqp.Publishing) error {
	return rc.channel.PublishWithContext(ctx, exchange, routingKey,
		// Mandartory: is used to determine if an error should be returned upon failed
		true,
		// Immediate: deprecated in rabbitmq 3
		false,
		options,
	)
}

func (rc RabbitClient) Consume(queue string, consumer string, autoAck bool) (<-chan amqp.Delivery, error) {
	return rc.channel.Consume(queue, consumer, autoAck, false, false, false, nil)
}
