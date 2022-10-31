package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/burakkuru5534/src/auth"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var App *app

type app struct {
	DB *DbHandle
}

type conf struct {
	Auth      *auth.Auth
	JwtSecret string `json:"jwt_secret"`
}

var Conf *conf

func InitConf() {
	Conf = &conf{}
}

func (c *conf) SetAuth(auth *auth.Auth) {
	c.Auth = auth
}

func InitApp(db *DbHandle) error {
	App = &app{
		DB: db,
	}

	return nil
}

func BodyToJsonReq(r *http.Request, data interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return errors.New(fmt.Sprintf("Body unmarshall error %s", string(body)))
	}

	defer r.Body.Close()

	return nil
}

func StrToInt64(aval string) int64 {
	aval = strings.Trim(strings.TrimSpace(aval), "\n")
	i, err := strconv.ParseInt(aval, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func Int64ToStr(aval int64) string {
	res := strconv.FormatInt(aval, 10)
	return res
}

func CheckIfDocumentExists(name string) (bool, error) {

	var isExists bool

	err := App.DB.Get(&isExists, "SELECT EXISTS (SELECT 1 FROM document WHERE pdf_file_name = $1)", name)
	if err != nil {
		return false, err
	}
	return isExists, nil

}

func GetPdfFileName(id int64) (string, error) {

	var name string

	err := App.DB.Get(&name, "SELECT pdf_file_name FROM document WHERE id = $1", id)
	if err != nil {
		return "", err
	}
	return name, nil

}

func CheckPass(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func CheckPassBasic(dbPass string, password string) bool {

	if dbPass == password {
		return true
	} else {
		return false

	}
}

func HashPasswd(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}
