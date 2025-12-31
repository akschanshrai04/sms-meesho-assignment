
# Polyglot Distributed SMS Service

A Polyglot, event-driven SMS system split into two services: a Java service that sends SMS events and a Go service that receives and stores them, showcasing how different languages can work together in a distributed setup.


## API Reference

#### Send SMS

```http
  POST /v0/sms/send
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `phoneNumber` | `string` | **Required**. The PhoneNumber you want to send the message to |
| `message` | `string` | **Required**. the message you want to send |

#### Get SMS History

```http
  GET /v0/user/{User_id}/messages 
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `User_id`      | `string` | **Required**. The phone number(used as User_id) for which the SMS History is requested. |

## Running the Project Locally

### 1. SMS Sender Service (Java – Spring Boot)
The Java service is responsible for exposing REST APIs and producing SMS events to Kafka.

Setup instructions, configuration, and run commands are as follows:

- Clone the javaMeesho folder: 
```bash
git clone https://github.com/akschanshrai04/sms-meesho-assignment.git
cd javaMeesho
```
- Configure the `application.properties` file as mentioned in the `README.md`

- Build and Run using maven: 
```bash
mvn clean install
mvn spring-boot:run
```

- once the application starts succesfully, it will be available on: 
```bash
http://localhost:8080
```

---

### 2. SMS Store Service (Go)
The Go service consumes SMS events from Kafka and persists them to MongoDB. It also exposes APIs to fetch SMS history.

Setup instructions, configuration, and run commands are as follows:

- clone the goMeesho folder: 
```bash
git clone <go-service-repository-url>
cd <go-project-root>
```

- create a `.env` in the root of the project and setup the environment variables as mentioned in the `README.md`

- install the dependencies: 
```bash
go mod tidy
```

- run the services: 
start the http api
```bash
go run cmd/api/main.go 
```

start the kafka consumer
```bash
go run cmd/consumer/main.go 
```


---

### 3. Running Kafka Locally (Docker)
Kafka is used for asynchronous communication between the Java and Go services.

Run the following command: 

```bash
docker-compose up -d
```



kafka will be available on:
```bash
localhost:9092
```

To configure Kafka for both the Java and Go services, the Kafka broker address and topic details must be provided in their respective configuration files. Kafka runs locally and is used for asynchronous communication between the services.

Go Service(.env File):
```bash
KAFKA_SERVER : <kafka_broker_address>
KAFKA_TOPIC_NAME : <kafka_topic_name>
KAFKA_GROUP_ID : <consumer_group_id>
```

Java Service(application.properties File): 
```bash
spring.kafka.bootstrap-servers : <kafka_broker_address>
kafka.topicname : <kafka_topic_name>
```

### 4. MongoDB (Cloud)
A cloud-hosted MongoDB instance is used for persisting SMS data.

- The Go service connects to MongoDB using a connection URI provided via environment variables.
- No local MongoDB setup is required.
- Any MongoDB-compatible cloud provider (e.g., MongoDB Atlas) can be used.
- To configure MongoDB for the Go service, the following environment variables must be set in the .env file using the credentials provided by the cloud-hosted MongoDB instance.
```bash
MONGO_URI : <mongodb_connection_uri>
MONGO_DB_NAME : <database_name>
```

---

### 5. Redis (Cloud)
A cloud-hosted Redis instance is used by the Java service for rate limiting and temporary state management.

- The Java service connects to Redis using host, port, and credentials provided via environment variables.
- No local Redis setup is required.
- Any managed Redis provider (e.g., Redis Cloud) can be used.
- To configure Redis for the Java service, the following properties must be set in the application.properties file using the credentials provided by the cloud-managed Redis instance.
```bash
spring.data.redis.host : <redis_host>
spring.data.redis.port : <redis_port>
spring.data.redis.username : <redis_username>
spring.data.redis.password : <redis_password>
spring.data.redis.ssl.enabled : false
```

