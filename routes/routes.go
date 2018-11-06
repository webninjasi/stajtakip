package routes

import (
	"os"
	"strconv"
)

var MaxReqSize int64 = 32000000

func init() {
	env_mrs := os.Getenv("APP_MAX_REQ_SIZE")

	if env_mrs != "" {
		mrs, err := strconv.ParseInt(env_mrs, 10, 64)
		if err != nil && mrs > 0 {
			MaxReqSize = mrs
		}
	}
}
