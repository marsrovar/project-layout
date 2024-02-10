package env

import "os"

var (
	Env string
)

func Get(key string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return ""
}
