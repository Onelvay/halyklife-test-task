package main

import (
	"context"
	"fmt"
	"github.com/Onelvay/halyklife-test-task/pkg/domain"
	"github.com/Onelvay/halyklife-test-task/pkg/service"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var audit *service.AuditServer

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}
	Init()
	http.HandleFunc("/", handlerProxy)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}

func handlerProxy(w http.ResponseWriter, r *http.Request) {
	serverUrl, err := url.Parse("http://localhost:8000") //серверхост
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(serverUrl)

	id := uuid.New().String()
	log := domain.Log{id, r.Method, time.Now()}

	if err := audit.Log(context.Background(), log); err != nil {
		fmt.Println(err)
	}

	r.Header.Set("X-Request-Id", id)
	proxy.ServeHTTP(w, r)
}

// инициализация монгодб и структуры логирования
func Init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongoDB.host")))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("log").Collection("log")
	audit = service.NewAuditServer(collection)
}

// config
func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
