package usecase

import (
	"github.com/mywishes/go-ses/shared/aws"

	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/labstack/echo"
	"github.com/mywishes/go-ses/domain/transactional"
	"github.com/mywishes/go-ses/models"
	"github.com/mywishes/go-ses/shared/config"
	"github.com/mywishes/go-ses/shared/logger"
	"github.com/mywishes/go-ses/shared/util"
)

var (
	log   = logger.NewMywishesLogger
	conf  = config.NewImmutableConfig()
	myAWS = aws.NewAWS(conf)
)

type usecase struct {
}

// NewTransactionalUsecase return implementation of transactional methods
func NewTransactionalUsecase() transactional.Usecase {
	return &usecase{}
}

// SendDefaultTransactional will process send transactional mail as default config and given body
func (u *usecase) SendDefaultTransactional(c echo.Context, body *models.Transactional) (result *ses.SendEmailOutput, err error) {
	ac := c.(*util.CustomApplicationContext)
	sess := ac.AWSSession

	serviceSES := myAWS.GetSESService(sess)
	bodySES := myAWS.ConvertToSESFormat(body, conf.GetAWSSESMailFrom())
	res, err := myAWS.SESSendEmail(serviceSES, bodySES)

	return res, err
}
