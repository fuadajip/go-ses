package repository

import (
	"fmt"

	"github.com/mywishes/go-ses/models"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/mywishes/go-ses/domain/transactional"
	"github.com/mywishes/go-ses/shared/aws"
	"github.com/mywishes/go-ses/shared/config"
	Error "github.com/mywishes/go-ses/shared/error"
)

type repositoryHandler struct {
	sess *session.Session
}

var (
	conf  = config.NewImmutableConfig()
	myAWS = aws.NewAWS(conf)
)

// NewAWSRepository is a factory that return object method of its implementation
func NewAWSRepository(session *session.Session) transactional.Repository {
	return &repositoryHandler{
		sess: session,
	}
}

// SendDefaultTransactional return result of aws ses repository
func (r *repositoryHandler) SendDefaultTransactional(body *models.Transactional) (res *ses.SendEmailOutput, err error) {
	serviceSES := myAWS.GetSESService(r.sess)
	bodySES := myAWS.ConvertToSESFormat(body, conf.GetAWSSESMailFrom())

	res, err = myAWS.SESSendEmail(serviceSES, bodySES)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				return nil, Error.New(fmt.Sprintf("%s: %s", ses.ErrCodeMessageRejected, aerr.Error()))
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				return nil, Error.New(fmt.Sprintf("%s: %s", ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error()))
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				return nil, Error.New(fmt.Sprintf("%s: %s", ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error()))
			default:
				return nil, Error.New(fmt.Sprintf("Unsupported error: %s", aerr.Error()))
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			return nil, Error.New(fmt.Sprintf("Undefined error: %s", aerr.Error()))
		}
	}

	return res, nil
}
