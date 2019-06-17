package main

import (
	"flag"
	"github.com/futurehomeno/fimpgo"
	log "github.com/sirupsen/logrus"
	"time"
)

func onMsg(topic string, addr *fimpgo.Address, iotMsg *fimpgo.FimpMessage,rawMessage []byte){
	log.Infof("New message from topic %s",topic)
}

func main() {
	mqttHost := flag.String("host","cube.local:1883","MQTT broker URL , for instance cube.local:1883")
	flag.Parse()
	log.SetLevel(log.DebugLevel)
	log.Infof("Broker url %s",*mqttHost)
	mqtt := fimpgo.NewMqttTransport("tcp://"+*mqttHost,"","","",true,1,1)
	err := mqtt.Start()
	log.Infof("Connected to broker %s",*mqttHost)
	if err != nil {
		log.Error("Error connecting to broker ",err)
	}

	mqtt.SetMessageHandler(onMsg)
	time.Sleep(time.Second*1)
	mqtt.Subscribe("#")
	log.Info("Publishing message")

	msg := fimpgo.NewFloatMessage("evt.sensor.report", "temp_sensor", float64(35.5), nil, nil, nil)
	adr := fimpgo.Address{MsgType: fimpgo.MsgTypeEvt, ResourceType: fimpgo.ResourceTypeDevice, ResourceName: "test", ResourceAddress: "1", ServiceName: "temp_sensor", ServiceAddress: "300"}
	mqtt.Publish(&adr,msg)

	select {

	}
	
}
