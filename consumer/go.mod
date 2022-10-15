module consumer

go 1.18

require (
	api v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/rabbitmq/amqp091-go v1.5.0
)

require (
	github.com/google/uuid v1.3.0
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.21.1 // indirect
)

replace api => ../api
