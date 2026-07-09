package mocksstubs

type userservice struct {
	sender Emailsender
}

func NewUserService(sender Emailsender) *userservice {
	return &userservice{
		sender: sender,
	}
}

func (s *userservice) Register(email string) error {

	return s.sender.Send(email, "welcome", "Thanks for joining!")

}
