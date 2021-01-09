package main

import (
	"time"

	"github.com/clarketm/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel
var init_err error
var password string
var status StatusFH
var network Fault
var database Fault
var software Fault
var access Fault
var camera Fault
var day int
var email_changed bool
var first bool

func init() {
	log.Trace("Initialised rabbitmq package")
	email_changed = false
	first = true

	status = StatusFH{
		DailyFaults:  0,
		CommonFaults: "N/A"}

	_, _, day := time.Now().Date()
	log.Debug("Current day is set to: ", day)

	network = Fault{
		Count: 0,
		Name:  "Network Fault",
	}

	database = Fault{
		Count: 0,
		Name:  "Database Fault",
	}

	software = Fault{
		Count: 0,
		Name:  "Software Fault",
	}

	access = Fault{
		Count: 0,
		Name:  "Alarm Access Fault",
	}

	camera = Fault{
		Count: 0,
		Name:  "Camera Fault",
	}

}

func SetPassword(pass string) {
	password = pass
}

func failOnError(err error, msg string) {
	if err != nil {
		log.WithFields(log.Fields{
			"Message": msg, "Error": err,
		}).Trace("Rabbitmq error")
	}
}

func getTime() string {
	t := time.Now()
	log.Trace(t.Format(TIMEFORMAT))
	return t.Format(TIMEFORMAT)
}

func messages(routing_key string, value string) {
	log.Warn("Adding messages to map")
	if SubscribedMessagesMap == nil {
		log.Warn("Creation of messages map")
		SubscribedMessagesMap = make(map[uint32]*MapMessage)
		messages(routing_key, value)
	} else {
		if key_id >= 0 {
			_, valid := SubscribedMessagesMap[key_id]
			if valid {
				log.Debug("Key already exists, checking next field: ", key_id)
				if key_id == 100 {
					key_id = 0
				}
				key_id++
				messages(routing_key, value)
			} else {
				log.Debug("Key does not exists, adding new field: ", key_id)
				entry := MapMessage{value, routing_key, getTime(), true}
				SubscribedMessagesMap[key_id] = &entry
				key_id++
			}
		}
	}
}

func SetConnection() error {
	conn, init_err = amqp.Dial("amqp://guest:" + password + "@localhost:5672/")
	failOnError(init_err, "Failed to connect to RabbitMQ")

	ch, init_err = conn.Channel()
	failOnError(init_err, "Failed to open a channel")
	return init_err
}

func Subscribe() {
	init := SetConnection()
	log.Trace("Beginning rabbitmq initialisation")
	log.Warn("Rabbitmq error:", init)
	if init == nil {
		var topics = [6]string{
			FAILURE,
			MOTIONDETECTED,
			MONITORSTATE,
			DEVICEFOUND,
			GUIDUPDATE,
			EMAILRESPONSE,
		}

		err := ch.ExchangeDeclare(
			EXCHANGENAME, // name
			EXCHANGETYPE, // type
			true,         // durable
			false,        // auto-deleted
			false,        // internal
			false,        // no-wait
			nil,          // arguments
		)
		failOnError(err, "FH - Failed to declare an exchange")

		q, err := ch.QueueDeclare(
			"",    // name
			false, // durable
			false, // delete when usused
			true,  // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")

		for _, s := range topics {
			log.Printf("Binding queue %s to exchange %s with routing key %s",
				q.Name, EXCHANGENAME, s)
			err = ch.QueueBind(
				q.Name,       // queue name
				s,            // routing key
				EXCHANGENAME, // exchange
				false,
				nil)
			failOnError(err, "Failed to bind a queue")
		}

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto ack
			false,  // exclusive
			false,  // no local
			false,  // no wait
			nil,    // args
		)
		failOnError(err, "Failed to register a consumer")

		forever := make(chan bool)

		go func() {
			for d := range msgs {
				log.Trace("Sending message to callback")
				log.Trace(d.RoutingKey)
				s := string(d.Body[:])
				messages(d.RoutingKey, s)
				log.Debug("Checking states of received messages")
				checkState()
			}
			//This function is checked after to see if multiple errors occur then to
			//through an event message
		}()

		log.Trace(" [*] Waiting for logs. To exit press CTRL+C")
		<-forever
	}
}

func StatusCheck() {
	status.CommonFaults, status.DailyFaults = GetCommonFault()
	valid := PublishStatusFH()
	if valid != "" {
		log.Warn("Failed to publish")
	} else {
		log.Debug("Published Status FH")
	}
}

func Publish(message []byte, routingKey string) string {
	if init_err == nil {
		log.Debug(string(message))
		err := ch.Publish(
			EXCHANGENAME, // exchange
			routingKey,   // routing key
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(message),
			})
		if err != nil {
			log.Fatal(err)
			return FAILUREPUBLISH
		}
	}
	return ""
}

func publishAlarmEvent(user string, state string) string {
	alarm, err := json.Marshal(&AlarmEvent{
		User:  user,
		State: state})
	if err != nil {
		return "Failed to convert AlarmEvent"
	} else {
		return Publish(alarm, ALARMEVENT)
	}
}

func publishCameraStart() string {
	emailRequest, err := json.Marshal(&EmailRequest{
		Role: ""})
	if err != nil {
		return "Failed to convert CameraStart"
	} else {
		return Publish(emailRequest, CAMERASTART)
	}
}

func publishCameraStop() string {
	emailRequest, err := json.Marshal(&EmailRequest{
		Role: ""})
	if err != nil {
		return "Failed to convert CameraStop"
	} else {
		return Publish(emailRequest, CAMERASTOP)
	}
}

func publishEmailRequest(role string) string {
	emailRequest, err := json.Marshal(&EmailRequest{
		Role: role})
	if err != nil {
		return "Failed to convert EmailRequest"
	} else {
		return Publish(emailRequest, EMAILREQUEST)
	}
}

func PublishStatusFH() string {
	message, err := json.Marshal(&status)
	if err != nil {
		return "Failed to convert StatusFH"
	} else {
		return Publish(message, STATUSFH)
	}
}

func PublishEventFH(component string, message string, time string, event_type_id string) string {
	eventFH, err := json.Marshal(&EventFH{
		Component:   component,
		Time:        time,
		EventTypeId: event_type_id})
	if err != nil {
		return "Failed to convert EventFH"
	} else {
		return Publish(eventFH, EVENTFH)
	}
}
