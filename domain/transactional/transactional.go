package transactional

import (
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/labstack/echo"
	"github.com/mywishes/go-ses/models"
)

// Usecase is an interface of Transactional that return implementation of Transactional's methods
type Usecase interface {
	SendDefaultTransactional(c echo.Context, body *models.Transactional) (result *ses.SendEmailOutput, err error)
}
