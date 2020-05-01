package app

import (
	"context"
	"fantlab/server/internal/config"
	"fantlab/server/internal/helpers"
	"text/template"
)

const (
	pmMailTemplate = `
To: {{.Email}}
From: "{{.SiteName}}" <{{.SiteEmail}}>
Content-Type: text/plain; charset=utf-8
Content-Transfer-Encoding: 8bit
Subject: Сообщение от "{{.Login}}"

{{.Message}}

---
Это письмо отправлено с сайта {{.SiteUrl}} посетителем {{.Login}}. Не отвечайте на него почтой!
Для ответа зайдите на страницу {{.SiteUrl}}/user{{.UserId}}, авторизуйтесь, если нужно, и отправьте сообщение, используя форму внизу.
	`
)

var (
	parsedPmMailTemplate = template.Must(template.New("private_message").Parse(pmMailTemplate))
)

type privateMessageMailData struct {
	SiteUrl   string
	SiteName  string
	SiteEmail string
	Email     string
	UserId    uint64
	Login     string
	Message   string
}

func (s *Services) SendPrivateMessageMail(ctx context.Context, fromUserId uint64, fromLogin string, toEmails []string, message string, cfg *config.AppConfig) {
	for _, toEmail := range toEmails {
		data := privateMessageMailData{
			SiteUrl:   cfg.SiteURL,
			SiteName:  cfg.SiteName,
			SiteEmail: cfg.SiteEmail,
			Email:     toEmail,
			UserId:    fromUserId,
			Login:     fromLogin,
			Message:   message,
		}

		msg, err := helpers.InflateTextTemplate(parsedPmMailTemplate, data)

		if err == nil {
			_ = s.smtp.SendMail(ctx, cfg.SiteEmail, toEmail, "private message", msg)
		}
	}
}
