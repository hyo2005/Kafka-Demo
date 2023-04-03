# Kafka-Demo
Demo kafka using Golang

#download kafka, copy folder onto disk C and change the folder name into kafka
#in kafka\config\zookeeper change the dataDir=c:/kafka/zookeeper 
#copy server.properties in kafka\config into 3 more properties file (server.1, server.2, server.3) 
#and change 3 properties

# server.1.properties
broker.id=1
listeners=PLAINTEXT://:9093
log.dirs=/tmp/kafka-logs1
# server.2.properties
broker.id=2
listeners=PLAINTEXT://:9094
log.dirs=/tmp/kafka-logs2
# server.3.properties
broker.id=3
listeners=PLAINTEXT://:9095
log.dirs=/tmp/kafka-logs3
#create the configured log directories and run cmd in kafka folder 
# start zookeeper
.\bin\windows\zookeeper-server-start.bat .\config\zookeeper.properties
# start server
.\bin\windows\kafka-server-start.bat .\config\server.1.properties
.\bin\windows\kafka-server-start.bat .\config\server.2.properties
.\bin\windows\kafka-server-start.bat .\config\server.3.properties
# create topic
.\bin\windows\kafka-topics.bat --create --topic message-log --bootstrap-server localhost:9093 --partitions 3 --replication-factor 2
# list topic
.\bin\windows\kafka-topic.bat --list --bootstrap-server localhost:9093
# create producer
.\bin\windows\kafka-console-producer.bat --broker-list localhost:9093,localhost:9094,localhost:9095 --topic message-log
# create consumer
.\bin\windows\kafka-console-consumer.bat --bootstrap-server localhost:9093 --topic message-log --from-beginning
