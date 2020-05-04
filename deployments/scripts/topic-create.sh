#!/usr/bin/env bash
export KAFKA_HOME=/Users/singaravelannandakumar/Desktop/kafka_2.11-2.4.0
export KAFKA=$KAFKA_HOME/bin
export KAFKA_CONFIG=$KAFKA_HOME/config
source ~/.bash_profile
$KAFKA/kafka-topics.sh --create --zookeeper localhost:2181 --topic crawler-data --partitions 3 --replication-factor 3