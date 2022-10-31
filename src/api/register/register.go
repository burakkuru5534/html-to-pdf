package register

import (
	"encoding/json"
	"fmt"
	"github.com/burakkuru5534/src/helper"
	"github.com/burakkuru5534/src/model"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"log"
)

var (
	validate *validator.Validate
)

func NewRegister(w http.ResponseWriter, r *http.Request) {
	data := model.RegisterRequestData()

	err := helper.BodyToJsonReq(r, &data)
	if err != nil {
		log.Println("Register body to json error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	// clean trailing and ending spaces
	data.FirstName = strings.TrimSpace(data.FirstName)
	data.LastName = strings.TrimSpace(data.LastName)
	data.Email = strings.TrimSpace(data.Email)

	names := strings.Split(data.FirstName, " ")
	if len(names) > 1 {
		data.FirstName = names[0]
		data.MiddleName.SetValid(names[1])
	}

	validate = validator.New()
	err = validate.Struct(data)
	if err != nil {
		log.Println("Register body validate error: ", err)
		http.Error(w, "{\"error\": \"Bad request\"}", http.StatusBadRequest)
		return
	}

	data.Password, err = helper.HashPasswd(data.Password)
	if err != nil {
		log.Println("hash password error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	err = createSysUsr(data)
	if err != nil {
		log.Println("create sysusr error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	resp := fmt.Sprintf("User %s created:", data.FirstName+"."+data.LastName)
	json.NewEncoder(w).Encode(resp)
}

func createSysUsr(data *model.RegisterRequest) error {
	var err error
	qs := "insert into sysusr (code, full_name, email, upass) values ($1, $2, $3, $4) returning id"

	err = helper.App.DB.QueryRowx(qs, data.FirstName+"."+data.LastName, data.FirstName+" "+data.LastName, data.Email, data.Password).Scan(&data.ID)
	if err != nil {
		return err
	}

	return nil
}
