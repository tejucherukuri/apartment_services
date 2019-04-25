//package queue

//import (
//	"fmt"

//	"github.com/astaxie/beego"
//	"github.com/streadway/amqp"
//)

//var QueueConnection *amqp.Connection
//var Channel *amqp.Channel
//var GcmQueue amqp.Queue
//var SmsQueue amqp.Queue
//var LogQueue amqp.Queue
//var EmailQueue amqp.Queue
//var GradebookQueue amqp.Queue

//func InitializeQueues() {
//	QueueConnection, err := amqp.Dial(beego.AppConfig.String("rabbitmq_uri"))
//	if err != nil {
//		beego.Error(" Error Initializing Queue : ", err)
//		panic(fmt.Sprintf("Error Initializing Queue: %s", err))
//	}
//	Channel, err = QueueConnection.Channel()
//	if err != nil {
//		beego.Error(" Error Initializing Channel : ", err)
//		panic(fmt.Sprintf("Error Initializing Channel: %s", err))
//	}
//	GcmQueue, err = Channel.QueueDeclare(
//		beego.AppConfig.String("gcm_queue_name"),
//		true,
//		false,
//		false,
//		false,
//		nil,
//	)
//	if err != nil {
//		beego.Error(" Error Initializing GCM Queue : ", err)
//		panic(fmt.Sprintf("Error Initializing GCM Queue: %s", err))
//	}

//	SmsQueue, err = Channel.QueueDeclare(
//		beego.AppConfig.String("sms_queue_name"),
//		true,
//		false,
//		false,
//		false,
//		nil,
//	)
//	if err != nil {
//		beego.Error(" Error Initializing SMS Queue : ", err)
//		panic(fmt.Sprintf("Error Initializing SMS Queue: %s", err))
//	}

//	LogQueue, err = Channel.QueueDeclare(
//		beego.AppConfig.String("log_queue_name"),
//		true,
//		false,
//		false,
//		false,
//		nil,
//	)
//	if err != nil {
//		beego.Error(" Error Initializing Log Queue : ", err)
//		panic(fmt.Sprintf("Error Initializing Log Queue: %s", err))
//	}

//	EmailQueue, err = Channel.QueueDeclare(
//		beego.AppConfig.String("email_queue_name"),
//		true,
//		false,
//		false,
//		false,
//		nil,
//	)
//	if err != nil {
//		beego.Error(" Error Initializing Email Queue : ", err)
//		panic(fmt.Sprintf("Error Initializing Email Queue: %s", err))
//	}
//	GradebookQueue, err = Channel.QueueDeclare(
//		beego.AppConfig.String("grade_queue_name"),
//		true,
//		false,
//		false,
//		false,
//		nil,
//	)
//	if err != nil {
//		beego.Error(" Error Initializing Gradebook Queue : ", err)
//		panic(fmt.Sprintf("Error Initializing Gradebook Queue: %s", err))
//	}
//}
