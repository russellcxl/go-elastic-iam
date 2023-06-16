package middlewares

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s][%s] -- %s %s (status: %d latency: %s)\n ",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC822),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency,
		)
	})
}

func SetLogOutput(path string) {
	f, _ := os.Create(path)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
