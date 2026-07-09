package mocksstubs

type Emailsender interface {
	Send(to, subject, body string) error
}
