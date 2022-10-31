package sys

import (
	"database/sql"
	"encoding/json"
	"github.com/burakkuru5534/src/auth"
	"github.com/burakkuru5534/src/helper"
	"github.com/burakkuru5534/src/model"
	"net/http"

	"log"
)

func Login(w http.ResponseWriter, r *http.Request) {

	var loginData model.LoginData
	var loginInfo model.LoginReqData

	err := helper.BodyToJsonReq(r, &loginInfo)
	if err != nil {
		log.Println("Login body to json error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	qs := "select id, coalesce(code, '') code, coalesce(full_name, '') full_name, coalesce(email, '') email, upass from sysusr where code = $1 or email = $1 and is_active"
	err = helper.App.DB.QueryRowx(qs, loginInfo.Username).Scan(&loginData.UserID, &loginData.UserName, &loginData.UserFullName, &loginData.UserEmail, &loginData.UpassFromDb)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Login user not found: ", err)
			http.Error(w, "{\"error\": \"Bad request\"}", http.StatusBadRequest)
		} else {
			log.Println("Login user query error: ", err)
			http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		}
		return
	}

	// compare password hashes
	loginResult := helper.CheckPass(*loginData.UpassFromDb, loginInfo.Password)
	if !loginResult {
		log.Println("Login password not match: ", err)
		http.Error(w, "{\"error\": \"Bad request\"}", http.StatusBadRequest)
		return
	}

	tc, err := auth.NewTokenClaimsForUser(loginData.UserID, loginData.UserName, loginData.UserEmail, loginData.UserFullName)
	if err != nil {
		log.Println("Login token claims error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	var respStruct struct {
		UserID       int64
		UserName     string
		UserEmail    string
		UserFullName string
		AccessToken  string
	}
	respStruct.UserID = loginData.UserID
	respStruct.UserName = loginData.UserName
	respStruct.UserEmail = loginData.UserEmail
	respStruct.UserFullName = loginData.UserFullName
	respStruct.AccessToken = tc.Encode(helper.Conf.Auth.JWTAuth)

	json.NewEncoder(w).Encode(respStruct)

}
