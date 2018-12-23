package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// ExitOnFailure prints a fatal error message and exits the process with status 1.
func ExitOnFailure(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "[CRIT] %s. ", err.Error())
	os.Exit(1)
}

// Дописать обработчик ошибок
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
