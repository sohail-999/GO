package mocksstubs

import "testing"

type StubEmailSender struct{}

func (s *StubEmailSender) Send(to, subject, body string) error {
	// Always succeed.
	return nil
}

func TestRegister(t *testing.T) {
	stub := &StubEmailSender{}

	service := NewUserService(stub)

	err := service.Register("john@example.com")

	if err != nil {
		t.Fatal(err)
	}
}
