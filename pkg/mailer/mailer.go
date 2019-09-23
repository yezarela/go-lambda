package mailer

var (
	mailer Mailer
)

// Mailer represents mailer
type Mailer struct{}

// SendEmail sends email
func (m *Mailer) SendEmail(toName string, toEmail string, content string) error {

	// some logic here
	return nil
}

// NewMailer returns mailer instance
func NewMailer() *Mailer {

	// Initialization
	mailer = Mailer{}

	return &mailer
}
