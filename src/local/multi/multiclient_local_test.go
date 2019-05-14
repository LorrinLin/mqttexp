package main

import (
	"testing"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
	"log"
	"strconv"
)

func BenchmarkMqttLocalMulti(b *testing.B){
	uri := "localhost:1883"
	topic := "testTimeTopic"
	publisher := connect("pub",uri)
	for i:=0;i<b.N;i++{
		listen(uri,topic,i)
		publisher.Publish(topic, 0, false, "hello")
	}
}

func listen(uri string, topic string, i int){
	consumer := connect("sub"+strconv.Itoa(i), uri)
	consumer.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message){
//		log.Print("message:", string(msg.Payload()))
	})
}

func connect(clientId string, uri string) mqtt.Client{
	opts := mqtt.NewClientOptions()
	opts.AddBroker(uri)
	opts.SetClientID(clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	
	for !token.WaitTimeout(1 * time.Second){
	
	}
	
	if err := token.Error();err != nil{
		log.Fatal(err)
	}
	return client
}