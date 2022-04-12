package handlers

import (
	"context"
	"log"

	"github.com/necutya-diploma/notification-service/pkg/grpc/gen"
)

type EmailService interface {
	SendEmail(from string, to []string, subject, body string) error
}

type Handler struct {
	gen.UnimplementedMailerServer
	service EmailService
}

func New(service EmailService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SendEmail(ctx context.Context, msg *gen.EmailMessage) (*gen.EmptyResponse, error) {
	err := h.service.SendEmail(msg.From, msg.To, msg.Subject, msg.Body)
	log.Println(err)
	return &gen.EmptyResponse{}, err
}
