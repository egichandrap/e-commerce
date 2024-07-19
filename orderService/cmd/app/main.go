package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"net/http"
	handler "orderService/internal/delivery/http"
	"orderService/internal/repository"
	"orderService/internal/repository/rabbitmq"
	"orderService/internal/usecase"
	"orderService/pkg/config"
	"os"
)

func main() {

	log.Infoln("app is running")

	if err := godotenv.Load(); err != nil {
		log.Warningln(err)
	}

	db, err := config.PSQL{
		DSN:                os.Getenv("DATABASE_URL"),
		MaxConnections:     config.DefaultValue("MAX_OPEN_CONNECTIONS", "10"),
		MaxIdleConnections: config.DefaultValue("MAX_IDLE_CONNECTIONS", "10"),
		ConnMaxLifetime:    config.DefaultValue("CONN_MAX_LIFETIME", "1m"),
	}.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	conn, channel, err := config.RabbitMQ{
		Url: os.Getenv("RABBITMQ_URL"),
	}.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
	defer channel.Close()

	orderRepo := repository.NewMysqlOrderRepository(db)
	client, err := rabbitmq.NewRabbitMQClient(channel, "orders_queue")
	if err != nil {
		log.Fatalf("Failed to Craete RabbitMQ: %v", err.Error())
	}
	useCase := usecase.NewOrderUseCase(orderRepo, client)
	orderHandler := handler.OrderHandler{UseCase: useCase}
	router := mux.NewRouter()
	router.HandleFunc("/orders", orderHandler.SendOrder).Methods("POST")
	err = http.ListenAndServe(fmt.Sprintf(":%s", config.DefaultValue("PORT", "8000")), router)
	if err != nil {
		panic(err)
	}

}
