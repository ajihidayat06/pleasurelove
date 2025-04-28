package utils

import "os"

func GetOSEnvGuestSecretKey() string {
	return os.Getenv("GUEST_SECRET_KEY")
}
