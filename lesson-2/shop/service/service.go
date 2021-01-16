package service

import (
	"errors"
	"log"

	"shop/models"
	"shop/pkg/smtpserv"
	tg "shop/pkg/tgbot"
	rep "shop/repository"
)

type Service interface {
	CreateItem(item *models.Item) (*models.Item, error)
	CreateOrder(order *models.Order) (*models.Order, error)
}

type service struct {
	tg   tg.TelegramAPI
	db   rep.Repository
	smtp smtpserv.SmtpServer
}

func (s *service) CreateOrder(order *models.Order) (*models.Order, error) {
	for _, itemID := range order.ItemIDs {
		_, err := s.db.GetItem(int32(itemID))
		if err != nil {
			return nil, errors.New("item not found")
		}
	}

	order, err := s.db.CreateOrder(order)
	if err != nil {
		return nil, err
	}
	if err := s.tg.SendOrderNotification(order); err != nil {
		log.Println(err)
	}
	if err := s.smtp.SendOrderConfirmationEmail(order); err != nil {
		log.Println(err)
	}
	return order, err
}

func (s *service) CreateItem(item *models.Item) (*models.Item, error) {
	if item.Name == "" {
		return nil, errors.New("item name is empty")
	}
	if item.Price <= 0 {
		return nil, errors.New("item price should be positive")
	}

	return s.db.CreateItem(item)
}

func NewService(tg tg.TelegramAPI, db rep.Repository, smtp smtpserv.SmtpServer) Service {
	return &service{
		db:   db,
		tg:   tg,
		smtp: smtp,
	}
}
