
# Provider Side case of Push
From the portal, the provider creates data stream (a call should go to the catalog service to create a new contract type and get the ID).

In the next page, the portal displays the topic and the endpoint to which the provider should publish his stream.

The topic is the Contract Type ID. The destination endpoint is: `https://<kafka_rest_proxy_url>/topics/`

This topic should be created by sending a query to this endpoint such as:
```
 curl -X POST -H "Content-Type: application/vnd.kafka.json.v2+json" -d '{"records":[{"key":"key","value":"value"},{"value":"value","partition":"0"},{"value":"value"}]}' https://<kafka_rest_proxy_url>/topics/<topic_ID>
```

# Provider Side case of Pull
From the portal, the provider creates data stream (a call should go to the catalog service to create a new contract type and get the ID).

The provider should specify the stream source URL, the protocol (KAFKA for now) and the topic.

Afterwards, the portal sends this query to the data-stream-delivery service as follow:
```
curl -X POST -k https://<stream_service_url>/v1/pullpushstream -H "Content-Type: text/plain" -d '{"DataContractTypeID":"<Contract_Type_ID>","DataStreamProtocol":"KAFKA","DataStreamSourceURL":"<source_of_the_stream>","DataStreamDestinationURL":"<destination>","DataStreamTopic":"<contract_Type_ID>"}'
```
Example:
```
curl -X POST -k https://<stream_service_url>/v1/pullpushstream -H "Content-Type: text/plain" -d '{"DataContractTypeID":"123","DataStreamProtocol":"KAFKA","DataStreamSourceURL":"http://stream.meetup.com/2/rsvps","DataStreamDestinationURL":"kafka-svc:9093","DataStreamTopic":"123"}'
```
As shown, the endpoint of the data-stream-delivery service is https://<stream_service_url>

This will make the data-stream-service start collecting the data from the source and publishing it to the destination endpoint.



# Customer Side
From the portal, the customer will buy the stream.
When the customer wants to visualize the stream, three queries should be sent to the to the kafka-rest-server:

# Create a consumer 

## Creating the Consumer:
In order to create a consumer in kafka rest server, the following query should be sent to the kafka-rest-server and it should have the name of the consumer (replace <consumer_ID>):
```
 curl -X POST -H "Content-Type: application/vnd.kafka.v2+json" -H "Accept: application/vnd.kafka.v2+json" --data '{"name":"<consumer_ID>","format":"json", "auto.offset.reset":"earliest"}' https://<kafka_rest_proxy_url>/consumers/<consumer_ID>
```
Example:
```
 curl -X POST -H "Content-Type: application/vnd.kafka.v2+json" -H "Accept: application/vnd.kafka.v2+json" --data '{"name":"portal_consumer","format":"json", "auto.offset.reset":"earliest"}' https://<kafka_rest_proxy_url>/consumers/portal_consumer
```

## Subscribing the Consumer:
The next step is to subscribe the consumer that we created. We can do that by sending the following query to the consumer that we created:

```
curl -X POST -H "Content-Type: application/vnd.kafka.v2+json" --data '{"topics":["<topic_ID>"]}' https://<kafka_rest_proxy_url>/consumers/<consumer_ID>/instances/<consumer_ID>/subscription
```
Example:
```
curl -X POST -H "Content-Type: application/vnd.kafka.v2+json" --data '{"topics":["123"]}' https://<kafka_rest_proxy_url>/consumers/portal_consumer/instances/portal_consumer/subscription
```

## Start collecting the data:
Next, the portal should be able to collect the events from kafka by sending rest queries as such:
```
 curl -X GET -H "Accept: application/vnd.kafka.json.v2+json" https://<kafka_rest_proxy_url>/consumers/<consumer_ID>/instances/<consumer_ID>/records
```
Example:
```
 curl -X GET -H "Accept: application/vnd.kafka.json.v2+json" https://<kafka_rest_proxy_url>/consumers/portal_consumer/instances/portal_consumer/records
```