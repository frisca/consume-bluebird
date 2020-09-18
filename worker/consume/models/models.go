package models

import (
	"github.com/wvanbergen/kafka/consumergroup"
)

type ConfigKafka struct {
	ConsumerGroup  string
	KafkaTopics    []string
	Zookeeper      []string
	ConsumerConfig *consumergroup.Config
}

type ReqOrder struct {
	ReffNo          string  	`json:"reff_no"`
	Products        []Products 	`json:"products"`
}

type Products struct {
	Product 	string  `json:"product"`
	Price       float64 `json:"price"`
	Qty         int     `json:"qty"`
}