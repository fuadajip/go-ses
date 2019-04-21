package usecase

import (
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/labstack/echo"
	"github.com/mywishes/go-ses/domain/transactional"
	"github.com/mywishes/go-ses/models"
	"github.com/mywishes/go-ses/shared/logger"
)

var (
	log = logger.NewMywishesLogger
)

type usecase struct {
	repo transactional.Repository
}

// NewTransactionalUsecase return implementation of transactional methods
func NewTransactionalUsecase(repository transactional.Repository) transactional.Usecase {
	return &usecase{
		repo: repository,
	}
}

// SendDefaultTransactional will process send transactional mail as default config and given body
func (u *usecase) SendDefaultTransactional(c echo.Context, body *models.Transactional) (result *ses.SendEmailOutput, err error) {

	res, err := u.repo.SendDefaultTransactional(body)
	return res, err
}
