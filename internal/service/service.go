package service

type Service struct {
	Email *EmailService
}

func New(
	emailManager EmailManager,
) *Service {
	return &Service{
		Email: NewEmailService(emailManager),
	}
}
