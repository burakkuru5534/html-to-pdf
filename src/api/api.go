package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Shyp/go-dberror"
	"github.com/burakkuru5534/src/helper"
	"github.com/burakkuru5534/src/model"
	_ "github.com/letsencrypt/boulder/db"
	"io/ioutil"
	"log"
	"net/http"
)

func DocumentCreate(w http.ResponseWriter, r *http.Request) {

	var Document model.Document

	err := helper.BodyToJsonReq(r, &Document)
	if err != nil {
		http.Error(w, "{\"error\": \"Bad request\"}", http.StatusBadRequest)
		return
	}

	err = Document.Create()
	if err != nil {
		dberr := dberror.GetError(err)
		switch e := dberr.(type) {
		case *dberror.Error:
			if e.Code == "23505" {
				http.Error(w, "{\"error\": \"Document with that PdfFileName already exists\"}", http.StatusForbidden)
				return
			}
		}

		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	respBody := struct {
		ID          int    `json:"id"`
		HtmlContent string `json:"HtmlContent"`
		PdfFileName string `json:"PdfFileName"`
	}{
		ID:          Document.ID,
		HtmlContent: Document.HtmlContent,
		PdfFileName: Document.PdfFileName,
	}

	json.NewEncoder(w).Encode(respBody)

}

func DocumentDelete(w http.ResponseWriter, r *http.Request) {

	var Document model.Document

	//id := helper.StrToInt64(chi.URLParam(r, "id"))
	id := helper.StrToInt64(r.URL.Query().Get("id"))

	pdfFileName, err := helper.GetPdfFileName(id)
	if err != nil {
		log.Println("movie update get movie name error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}
	isExists, err := helper.CheckIfDocumentExists(pdfFileName)
	if err != nil {
		log.Println("movie update check if movie exists error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	if !isExists {
		log.Println("document update document does not exist error: ", err)
		http.Error(w, "{\"error\": \"Document with that name does not exist\"}", http.StatusNotFound)
		return
	}
	err = Document.Delete(id)
	if err != nil {
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("ok")

}

func DocumentGet(w http.ResponseWriter, r *http.Request) {

	var Document model.Document

	id := helper.StrToInt64(r.URL.Query().Get("id"))
	//id := helper.StrToInt64(chi.URLParam(r, "id"))
	pdfFileName, err := helper.GetPdfFileName(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "{\"error\": \"Document with that id does not exist\"}", http.StatusForbidden)
			return
		}
		log.Println("movie update get movie name error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}
	isExists, err := helper.CheckIfDocumentExists(pdfFileName)
	if err != nil {
		log.Println("movie update check if movie exists error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	if !isExists {
		log.Println("document update document does not exist error: ", err)
		http.Error(w, "{\"error\": \"Document with that name does not exist\"}", http.StatusNotFound)
		return
	}

	err = Document.Get(id)
	if err != nil {
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	respBody := struct {
		ID          int64  `json:"id"`
		HtmlContent string `json:"HtmlContent"`
		PdfFileName string `json:"PdfFileName"`
	}{
		ID:          id,
		HtmlContent: Document.HtmlContent,
		PdfFileName: Document.PdfFileName,
	}
	json.NewEncoder(w).Encode(respBody)

}

func DocumentList(w http.ResponseWriter, r *http.Request) {

	var Document model.Document

	DocumentList, err := Document.GetAll()
	if err != nil {
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	if len(DocumentList) == 0 {
		http.Error(w, "{\"error\": \"No document found\"}", http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(DocumentList)

}

func DocumentDownload(w http.ResponseWriter, r *http.Request) {

	var Document model.Document

	id := helper.StrToInt64(r.URL.Query().Get("id"))
	//id := helper.StrToInt64(chi.URLParam(r, "id"))
	pdfFileName, err := helper.GetPdfFileName(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "{\"error\": \"Document with that id does not exist\"}", http.StatusForbidden)
			return
		}
		log.Println("movie update get movie name error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}
	isExists, err := helper.CheckIfDocumentExists(pdfFileName)
	if err != nil {
		log.Println("movie update check if movie exists error: ", err)
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	if !isExists {
		log.Println("document update document does not exist error: ", err)
		http.Error(w, "{\"error\": \"Document with that name does not exist\"}", http.StatusNotFound)
		return
	}

	err = Document.Get(id)
	if err != nil {
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	contentType := "pdf"

	filename := fmt.Sprintf("./%s", Document.PdfFileName)

	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

	fileSize := len(fileBytes)
	//--> bu download için w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", data.FileName))
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", helper.Int64ToStr(int64(fileSize)))
	_, err = w.Write(fileBytes)
	if err != nil {
		http.Error(w, "{\"error\": \"server error\"}", http.StatusInternalServerError)
		return
	}

}
