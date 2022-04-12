package service

type EmailManager interface {
	SendEmail(from string, to []string, subject, body string) error
}

type EmailService struct {
	emailManager EmailManager
}

func NewEmailService(emailManager EmailManager) *EmailService {
	return &EmailService{
		emailManager: emailManager,
	}
}

func (es *EmailService) SendEmail(from string, to []string, subject, body string) error {
	return es.emailManager.SendEmail(from, to, subject, body)
}
