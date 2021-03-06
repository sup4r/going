package logs

import (
	"fmt"
	"time"
	"github.com/smartwalle/going/email"
)

type MailWriter struct {
	level   int
	config	*email.MailConfig
	subject string
	from    string
	to		[]string
}

func NewMailWriter(level int) *MailWriter {
	var writer = &MailWriter{}
	writer.level = level
	return writer
}

func(this *MailWriter) SetLevel(level int) {
	this.level = level
}

func(this *MailWriter) GetLevel() int {
	return this.level
}

func (this *MailWriter) SetConfig(config *email.MailConfig) {
	this.config = config
}

func (this *MailWriter) GetConfig() *email.MailConfig {
	return this.config
}

func (this *MailWriter) SetSubject(subject string) {
	this.subject = subject
}

func (this *MailWriter) GetSubject() string {
	return this.subject
}

func (this *MailWriter) SetFrom(from string) {
	this.from = from
}

func (this *MailWriter) GetFrom() string {
	return this.from
}

func (this *MailWriter) SetToMailList(to []string) {
	this.to = to
}

func (this *MailWriter) GetToMailList() []string {
	return this.to
}

func(this *MailWriter) WriteMessage(level int, file string, line int, prefix string, msg string) {
	if level < this.level {
		return
	}

	if this.config == nil {
		return
	}

	if len(this.to) == 0 {
		return
	}

	var message = fmt.Sprintf("%s %s [%s:%d] %s", time.Now().String(), prefix, file, line, msg)

	var subject = this.GetSubject()
	if len(subject) == 0 {
		subject = file
	}

	var mail = email.NewTextMessage(subject, message)
	mail.To = this.to
	if len(this.from) > 0 {
		mail.From = this.from
	}

	email.SendMail(this.config, mail)
}

func(this *MailWriter) Close() {

}

func(this *MailWriter) Flush() {

}