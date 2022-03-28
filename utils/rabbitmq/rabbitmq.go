package rabbitmq

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"psy-consult-backend/constant"
)

var conn *amqp.Connection
var channel *amqp.Channel

const (
	mq_address  = "124.221.197.218:5672"
	mq_user     = "ecnu"
	mq_password = "ecnu0006"
)

func GetConnection() *amqp.Connection {
	return conn
}

func GetChannel() *amqp.Channel {
	return channel
}

func Init() error {
	url := fmt.Sprintf("amqp://%s:%s@%s/", mq_user, mq_password, mq_address)
	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}
	channel, err = conn.Channel()
	if err != nil {
		return err
	}
	return nil
}

func PushMessage(queueName string, msg []byte) error {
	// 不存在则创建
	q, err := channel.QueueDeclare(queueName, true, true, false, false, nil)
	if err != nil {
		return err
	}
	err = channel.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         msg,
	})
	return err
}

func GetMessage(queueName string) ([]byte, error) {
	msg, ok, err := channel.Get(queueName, true)
	if !ok {
		return nil, err
	}
	if err != nil {
		logrus.Error(constant.DAO+"GetMessage Failed, err= %v", err)
		return nil, err
	}
	return msg.Body, nil
}
