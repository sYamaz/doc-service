package web

import (
	"log"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	LoggerFactory interface {
		Middleware() echo.MiddlewareFunc
	}
	loggerFactory struct{}
)

func (f *loggerFactory) Middleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		HandleError:      true,
		LogLatency:       true,
		LogProtocol:      false,
		LogRemoteIP:      true,
		LogHost:          true,
		LogMethod:        true,
		LogURI:           true,
		LogURIPath:       true,
		LogRoutePath:     false,
		LogRequestID:     true,
		LogReferer:       false,
		LogUserAgent:     true,
		LogStatus:        true,
		LogError:         true,
		LogContentLength: true,
		LogResponseSize:  true,
		Skipper:          middleware.DefaultSkipper,

		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			errMsg := ""
			if v.Error != nil {
				errMsg = v.Error.Error()
			}
			clen := v.ContentLength
			if clen == "" {
				clen = "0"
			}

			log.Printf(`{"time":"%v","id":"%v","remote_ip":"%v","host":"%v","method":"%v","uri":"%v","uri_path":"%v","user_agent":"%v","status":%v,"error":"%v","latency":%v,"latency_human":"%v","bytes_in":%v,"bytes_out":%v}`,
				time.Now().Format(time.RFC3339Nano),
				v.RequestID,
				v.RemoteIP,
				v.Host,
				v.Method,
				v.URI,
				v.URIPath,
				v.UserAgent,
				v.Status,
				errMsg,
				strconv.FormatInt(int64(v.Latency), 10),
				v.Latency.String(),
				clen,
				v.ResponseSize,
			)
			return nil
		},
	})
}

func NewCustomLogger() LoggerFactory {
	return &loggerFactory{}
}
