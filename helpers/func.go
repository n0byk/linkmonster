package helpers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

//GetMD5Hash хеширует строку
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

//ShowError dsdjl b htlbhtrn yf cnhfybwe jib,rb
func ShowError(w http.ResponseWriter, number string) {
	errortpl, err := ioutil.ReadFile("public/templates/errors/" + number + ".html")
	if err != nil {
		fmt.Print(err)
	}

	switch number {
	case "404":
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errortpl))
	case "500":
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errortpl))
	case "405":
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(errortpl))
	}
}

//Metric удобный вид больших цифр
func Metric(n int64) string {

	switch {
	case n >= 1000:
		return fmt.Sprint(n/1000, "k")
	case n >= 1000000:
		return fmt.Sprint(n/1000000, "m")
	default:
		return fmt.Sprint(n)
	}
}

// ExitOnFailure prints a fatal error message and exits the process with status 1.
func ExitOnFailure(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "[CRIT] %s. ", err.Error())
	os.Exit(1)
}

//CheckErr Дописать обработчик ошибок
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

//CheckPassword  проверка пароля
func CheckPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString(`^[\@A-Za-z0-9\!\#\$\%\^\&\*\~\{\}\[\]\.\,\<\>\(\)\_\+\=]{4,30}$`, password); !ok {
		return false
	}
	return true
}

//CheckUsername проверка пользователя
func CheckUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}A-Z0-9a-z_-]{2,30}$", username); !ok {
		return false
	}
	return true
}

//CheckEmail проверка EMAIL
func CheckEmail(email string) (b bool) {
	if ok, _ := regexp.MatchString(`^([a-zA-Z0-9._-])+@([a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$`, email); !ok {
		return false
	}
	return true
}
