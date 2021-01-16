package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"shop/pkg/conf"
	"shop/pkg/smtpserv"
	"shop/pkg/tgbot"
	"shop/repository"
	"shop/service"
)

func main() {
	senderFrom := flag.String("from", "jakoloborodun@gmail.com", "smtp sender email")
	senderPass := flag.String("pass", "", "smtp sender password")

	cfgPath, err := conf.ParseConfigFlag()
	if err != nil {
		panic(fmt.Sprintf("can't find config file: %s", err))
	}
	cfg, err := conf.NewConfig(cfgPath)
	if err != nil {
		panic(fmt.Sprintf("can't read config file: %s", err))
	}

	if len(cfg.Smtp.From) == 0 {
		cfg.Smtp.From = *senderFrom
	}
	if len(cfg.Smtp.Pass) == 0 {
		cfg.Smtp.Pass = *senderPass
	}

	tg, err := tgbot.NewTelegramAPI(cfg.Telegram.Token, cfg.Telegram.ChatId)
	if err != nil {
		log.Fatal("Unable to init telegram bot")
	}

	db := repository.NewMapDB()

	smtp, err := smtpserv.NewSmtpServer(cfg.Smtp.Host, cfg.Smtp.Port, cfg.Smtp.From, cfg.Smtp.Pass)
	if err != nil {
		log.Fatal("Unable to load SMTP server")
	}

	newService := service.NewService(tg, db, smtp)
	handler := &shopHandler{
		service: newService,
		db:      db,
	}

	router := mux.NewRouter()

	router.HandleFunc("/item", handler.createItemHandler).Methods("POST")
	router.HandleFunc("/item/{id}", handler.getItemHandler).Methods("GET")
	router.HandleFunc("/item/{id}", handler.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/item/{id}", handler.updateItemHandler).Methods("PUT")

	router.HandleFunc("/order", handler.createOrderHandler).Methods("POST")
	router.HandleFunc("/order/{id}", handler.getOrderHandler).Methods("GET")

	srv := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
