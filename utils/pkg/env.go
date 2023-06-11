package pkg

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/vanneeza/e-commerce-lite/utils/helper"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	helper.PanicError(err)

	return os.Getenv(key)
}
