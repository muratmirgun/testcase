1- First of all create rabbit and redis for backend apps

docker-compose up -d

2- Navigate to `api` folder and start 
```
$ cd api
$ go run main.go
```
3- Navigate to `consumer` and start 
```
$ cd consumer
$ go run main.go
```

4- Navigate to `list-api` and start
```
$ cd list-api
$ go run main.go
```

```
# Send
$ curl -X POST 'http://localhost:8080/message' -H 'Accept: application/json' -H 'Content-Type: application/json' -H 'Authorization;' -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2490.86 Safari/537.36' -H 'X-Forwarded-For: 23.235.46.133' --data-raw '{
    "sender": "Murat",
    "receiver":"Ali",
    "message":"Test Message"
}'

# List
$ curl -X GET 'http://localhost:8081/list/Murat/Ali'
```