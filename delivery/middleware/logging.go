package middleware

import (
	"github.com/EkoEdyPurwanto/goforit/config"
	"github.com/EkoEdyPurwanto/goforit/model/dto/req"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func Logging(log *logrus.Logger) fiber.Handler {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	log.Out = file

	return func(ctx *fiber.Ctx) error {
		startTime := time.Now()
		err := ctx.Next()
		endTime := time.Since(startTime)

		requestLog := req.LoggingRequest{
			StartTime:  startTime,
			EndTime:    endTime,
			StatusCode: ctx.Response().StatusCode(),
			ClientIP:   ctx.IP(),
			Method:     ctx.Method(),
			Path:       ctx.Path(),
			UserAgent:  ctx.Get("User-Agent"),
		}
		switch {
		case ctx.Response().StatusCode() >= 500:
			log.Error(requestLog)
		case ctx.Response().StatusCode() >= 400:
			log.Warn(requestLog)
		case ctx.Response().StatusCode() >= 300:
			log.Info(requestLog)
		case ctx.Response().StatusCode() >= 200:
			log.Debug(requestLog)
		case ctx.Response().StatusCode() >= 100:
			log.Trace(requestLog)
		default:
			log.Warnf("Unexpected status code: %d", ctx.Response().StatusCode())
		}

		return err
	}
}
