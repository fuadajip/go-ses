package main

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mywishes/go-ses/shared/aws"
	"github.com/mywishes/go-ses/shared/config"
	"github.com/mywishes/go-ses/shared/container"
	"github.com/mywishes/go-ses/shared/logger"
	"github.com/mywishes/go-ses/shared/util"

	transactionalHandler "github.com/mywishes/go-ses/domain/transactional/delivery/http"
	transactionalUsecase "github.com/mywishes/go-ses/domain/transactional/usecase"
)

type (
	// CustomValidator return validation body of echo
	CustomValidator struct {
		validator *validator.Validate
	}
)

var (
	log = logger.NewMywishesLogger()
)

// Validate return result echo body validation
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	container := container.GoldiDefaultContainer()
	conf := container.MustGet("shared.config").(config.ImmutableConfig)
	aws := container.MustGet("shared.aws.session").(aws.MywishesAWS)

	awsSession, err := aws.GetAWSSession()
	if err != nil {
		log.Error(err.Error())
		panic(err.Error())
	}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ac := &util.CustomApplicationContext{
				Context:    c,
				Container:  container,
				AWSSession: awsSession,
			}
			return h(ac)
		}
	})

	transactionalUcase := transactionalUsecase.NewTransactionalUsecase()
	transactionalHandler.AddTransactionalHandler(e, transactionalUcase)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.GetPort())))

}
