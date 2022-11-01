package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/burakkuru5534/src/api"
	"github.com/burakkuru5534/src/helper"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "inspakt",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	req, err := http.NewRequest("GET", "/api/documents", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.DocumentList)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		if status == http.StatusBadRequest {
			expected := `{"error": "Bad request"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusForbidden {
			expected := `{"error": "Document with that email already exists"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusInternalServerError {
			expected := `{"error": "Internal server error"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusNotFound {
			expected := `{"error": "Document with that id does not exist"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	} else {
		expected := `[{"id":1,"HtmlContent":"\u003chtml\u003e\u003cbody\u003e\u003ch1 style=\"color:red;\"\u003eThis is an html from pdf to test color\u003ch1\u003e\u003cimg src=\"http://api.qrserver.com/v1/create-qr-code/?data=HelloWorld \" alt=\"img\" height=\"42\" width=\"42\"\u003e\u003c/img\u003e\u003c/body\u003e\u003c/html\u003e","PdfFileName":"mytestpdfFile","DirectoryPath":"","FileSize":0}]
`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
	// Check the response body is what we expect.

}

func TestCreate(t *testing.T) {

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "inspakt",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	var jsonStr = []byte(`{
    "HtmlContent":"<html><body><h1 style=\"color:red;\">This is an html from pdf to test color<h1><img src=\"http://api.qrserver.com/v1/create-qr-code/?data=HelloWorld \" alt=\"img\" height=\"42\" width=\"42\"></img></body></html>",
    "PdfFileName":"mytestpdfFile"
}`)

	req, err := http.NewRequest("POST", "/api/document", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.DocumentCreate)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		if status == http.StatusBadRequest {
			expected := `{"error": "Bad request"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusForbidden {
			expected := `{"error": "Document with that email already exists"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusInternalServerError {
			expected := `{"error": "Internal server error"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusNotFound {
			expected := `{"error": "Document with that id does not exist"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	} else {
		var id int64

		err = db.Get(&id, "SELECT id from document order by id desc limit 1")
		if err != nil {
			errors.New("get id error.")
		}

		expected := fmt.Sprintf(`{"id":%d,"HtmlContent":""<html><body><h1 style=\"color:red;\">This is an html from pdf to test color<h1><img src=\"http://api.qrserver.com/v1/create-qr-code/?data=HelloWorld \" alt=\"img\" height=\"42\" width=\"42\"></img></body></html>"","PdfFileName":"mytestpdfFile"}
`, id)
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}

}

func TestGet(t *testing.T) {

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "inspakt",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	req, err := http.NewRequest("GET", "/api/document", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "22")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.DocumentGet)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		if status == http.StatusBadRequest {
			expected := `{"error": "Bad request"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusForbidden {
			expected := `{"error": "Document with that email already exists"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusInternalServerError {
			expected := `{"error": "Internal server error"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusNotFound {
			expected := `{"error": "Document with that id does not exist"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	} else {
		expected := `{"id":1,"html_content":""<html><body><h1 style=\"color:red;\">This is an html from pdf to test color<h1><img src=\"http://api.qrserver.com/v1/create-qr-code/?data=HelloWorld \" alt=\"img\" height=\"42\" width=\"42\"></img></body></html>"","pdf_file_name":"mytestpdfFile"}
`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
	// Check the response body is what we expect.

}

func TestDelete(t *testing.T) {

	conInfo := helper.PgConnectionInfo{
		Host:     "127.0.0.1",
		Port:     5432,
		Database: "inspakt",
		Username: "postgres",
		Password: "tayitkan",
		SSLMode:  "disable",
	}

	db, err := helper.NewPgSqlxDbHandle(conInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
	}

	// Create Appplication Service
	err = helper.InitApp(db)
	if err != nil {
		errors.New("init app error.")
	}

	req, err := http.NewRequest("DELETE", "/api/document", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.DocumentDelete)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		if status == http.StatusBadRequest {
			expected := `{"error": "Bad request"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusForbidden {
			expected := `{"error": "Document with that email already exists"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusInternalServerError {
			expected := `{"error": "Internal server error"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else if status == http.StatusNotFound {
			expected := `{"error": "Document with that id does not exist"}
`
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		} else {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

	} else {
		// Check the response body is what we expect.
		expected := `"ok"
`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}

}
