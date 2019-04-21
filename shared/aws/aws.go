package aws

import (
	"fmt"

	"github.com/mywishes/go-ses/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/mywishes/go-ses/shared/config"
	Error "github.com/mywishes/go-ses/shared/error"
	"github.com/mywishes/go-ses/shared/logger"
)

type (
	// MywishesAWS return implementation of methods aws
	MywishesAWS interface {
		GetAWSSession() (*session.Session, error)
		GetSESService(sess *session.Session) *ses.SES
		ConvertToSESFormat(body *models.Transactional, sender string) *ses.SendEmailInput
		SESSendEmail(serviceSES *ses.SES, bodySES *ses.SendEmailInput) (result *ses.SendEmailOutput, err error)
	}

	mywishesAWS struct {
		config config.ImmutableConfig
	}
)

var (
	log = logger.NewMywishesLogger()
)

// GetAWSSession return created aws session
func (m *mywishesAWS) GetAWSSession() (sess *session.Session, err error) {

	sess, err = session.NewSession(&aws.Config{
		Region:      aws.String(m.config.GetAWSRegion()),
		Credentials: credentials.NewStaticCredentials(m.config.GetAWSAccessKey(), m.config.GetAWSSecretAccessKey(), m.config.GetAWSSessionToken()),
	})

	if err != nil {
		msgError := fmt.Sprintf("failed to create aws session %s", err.Error())
		return nil, Error.New(msgError)
	}

	return sess, nil

}

// GetSESService return aws ses as a service
func (m *mywishesAWS) GetSESService(sess *session.Session) *ses.SES {
	return ses.New(sess)
}

// ConvertToSESFormat will return SES mail format for given object body and sender
func (m *mywishesAWS) ConvertToSESFormat(body *models.Transactional, sender string) *ses.SendEmailInput {
	return &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: aws.StringSlice(body.CcAddresses),
			ToAddresses: aws.StringSlice(body.ToAddresses),
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(body.MessageBodyHTML.Charset),
					Data:    aws.String(body.MessageBodyHTML.Data),
				},
				Text: &ses.Content{
					Charset: aws.String(body.MessageBodyText.Charset),
					Data:    aws.String(body.MessageBodyText.Data),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(body.MessageSubject.Charset),
				Data:    aws.String(body.MessageSubject.Data),
			},
		},
		Source: aws.String(sender),
	}
}

// SESSendMail will send mail through SES service with given sesService session and SES body input
func (m *mywishesAWS) SESSendEmail(serviceSES *ses.SES, bodySES *ses.SendEmailInput) (result *ses.SendEmailOutput, err error) {
	result, err = serviceSES.SendEmail(bodySES)
	return result, err
}

// NewAWS is an instance that implement MywishesAWS
func NewAWS(config config.ImmutableConfig) MywishesAWS {
	if config == nil {
		panic("[APP CONFIG] immutable config required")
	}
	return &mywishesAWS{config}
}
