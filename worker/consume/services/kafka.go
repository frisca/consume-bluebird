package services

import (
	"encoding/json"
	"fmt"
	"consume-bluebird/worker/consume/kafka"
	"consume-bluebird/worker/consume/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	cluster "github.com/bsm/sarama-cluster"
	"consume-bluebird/dbmodels"
	repository "consume-bluebird/database"
)

var (
	consumerGroup         string
	kafkaTopics           string
	zookeeper             string
	brokers               string
)


func StartServicesKafka() {
	fmt.Println("START SERVICE KAFKA..")

	consumerGroup = beego.AppConfig.DefaultString("kafka.group", "bluebird-group")
	brokers = beego.AppConfig.DefaultString("kafka.brokers", "localhost:9092")
	kafkaTopics = beego.AppConfig.DefaultString("kafka.topics", "order-topic")
	zookeeper = beego.AppConfig.DefaultString("kafka.zookeeper", "localhost:2181")

	configKafka := models.ConfigKafka{}
	configKafka.ConsumerGroup = consumerGroup
	configKafka.KafkaTopics = strings.Split(kafkaTopics, ",")
	configKafka.Zookeeper = strings.Split(zookeeper, ",")
	configKafka.ConsumerConfig = kafka.GetConsumerConfig()

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	fmt.Println("List Brokers :", brokers)
	fmt.Println("Kafka Topics :", kafkaTopics)

	listBroker := strings.Split(brokers, ",")
	topics := strings.Split(kafkaTopics, ",")

	consumer, err := cluster.NewConsumer(listBroker, consumerGroup, topics, config)
	if err != nil {
		fmt.Println("Cluster Connect problem ", err)
	}
	// consume errors
	go func() {
		for err := range consumer.Errors() {
			fmt.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			logs.Info("Rebalanced: %+v\n", ntf)
			//fmt.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Printf("%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				SaveDB(msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed\
			}
		}
	}
}


// SaveDB ...
func SaveDB(data []byte) {
	var req models.ReqOrder
	var dataorder dbmodels.Order

	json.Unmarshal(data, &req)

	newOrderNo := repository.GenerateOrderNo()
	totals := int64(0)
	if len(req.Products) > 0 {
		for _, product := range req.Products {
			totals = totals + (int64(product.Qty) * int64(product.Price))
		}
	}

	dataorder.ReffNo = req.ReffNo
	dataorder.Total = totals
	dataorder.OrderNo = newOrderNo

	repository.SaveOrder(dataorder, req.Products)
}