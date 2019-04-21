package http

import (
	"net/http"

	"github.com/mywishes/go-ses/models"
	"github.com/mywishes/go-ses/shared/config"
	"github.com/mywishes/go-ses/shared/logger"
	"github.com/mywishes/go-ses/shared/util"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mywishes/go-ses/domain/transactional"
)

type (
	handlerTransactional struct {
		usecase transactional.Usecase
	}
)

var (
	log  = logger.NewMywishesLogger()
	conf = config.NewImmutableConfig()
)

// AddTransactionalHandler return http handler for transactional
func AddTransactionalHandler(e *echo.Echo, usecase transactional.Usecase) {
	handler := handlerTransactional{
		usecase: usecase,
	}

	e.Use(middleware.JWT([]byte(conf.GetMywishesSecretKey())))
	e.POST("/api/transactional/v1/ses", handler.SendDefaultTransactional)
}

func (h *handlerTransactional) SendDefaultTransactional(c echo.Context) (err error) {
	ac := c.(*util.CustomApplicationContext)

	u := new(models.Transactional)
	if err = c.Bind(u); err != nil {
		return ac.CustomResponse("failed", nil, err.Error(), 422, nil)
	}

	// Validate request body with given SES input models
	if err = c.Validate(u); err != nil {
		return ac.CustomResponse("failed", nil, err.Error(), 422, nil)
	}

	result, err := h.usecase.SendDefaultTransactional(c, u)

	if err != nil {
		return ac.CustomResponse("failed", nil, err.Error(), http.StatusInternalServerError, nil)
	}

	return ac.CustomResponse("success", result, "success", 200, nil)

}
