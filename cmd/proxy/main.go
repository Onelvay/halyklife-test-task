package main

import (
	"context"
	"fmt"
	"github.com/Onelvay/halyklife-test-task/internal/service"
	"github.com/Onelvay/halyklife-test-task/pkg/domain"
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

var proxy *httputil.ReverseProxy
var audit *service.AuditServer

func main() {
	if err := initConfig(); err != nil { //config
		log.Fatal(err)
	}

	Init() //инициализация монгодб и структуры логирования

	serverUrl, err := url.Parse("http://localhost:8000") //серверхост
	if err != nil {
		panic(err)
	}
	proxy = httputil.NewSingleHostReverseProxy(serverUrl)

	http.HandleFunc("/", handlerProxy)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}

func handlerProxy(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
	log := domain.Log{id.String(), r.Method, time.Now()}
	fmt.Println(log)
	if err := audit.Log(context.Background(), log); err != nil {
		fmt.Println(err)
	}
	proxy.ServeHTTP(w, r)
}

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

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
