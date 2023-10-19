package mail

import (
	"testing"

	"github.com/TheDP66/simple_bank_go/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	// ? skip test from make test
	// ? can only run test manually
	if testing.Short() {
		t.Skip()
	}

	config, err := util.LoadConfig("..")
	require.NoError(t, err)

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := `
		<h1>Hello world!</h1>
		<p>This is a test message from <a href="https://github.com/TheDP66">Dharma Putra</a></p>
	`
	to := []string{"test.pinguin66@gmail.com", "test.kukang66@gmail.com"}
	attachFiles := []string{"../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
