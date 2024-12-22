package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	reviews "github.com/lavatee/shop_reviews"
	"github.com/lavatee/shop_reviews/internal/endpoint"
	"github.com/lavatee/shop_reviews/internal/repository"
	"github.com/lavatee/shop_reviews/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//	var zavozy = []string{
//		"Пересобрать кровать",
//		"Спиздить тарелку из столовой",
//		""
//	}
func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	if err := InitConfig(); err != nil {
		logrus.Fatalf("config open error: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("env open error: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(viper.GetString("db.host"), viper.GetString("db.port"), viper.GetString("db.user"), os.Getenv("DB_PASSWORD"), viper.GetString("db.dbname"), viper.GetString("db.sslmode"))
	if err != nil {
		logrus.Fatalf("db open error: %s", err.Error())
	}
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	endp := endpoint.NewEndpoint(services)
	server := &reviews.Server{}
	go func() {
		if err := server.Run(viper.GetString("port"), endp); err != nil {
			logrus.Fatalf("server run error: %s", err.Error())
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	server.Shutdown()
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
