package model

import (
	"fmt"
	"github.com/burakkuru5534/src/helper"
	"strings"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type DocumentRequest struct {
	HtmlContent string `json:"HtmlContent"`
	PdfFileName string `json:"PdfFileName"`
}

type Document struct {
	ID          int    `json:"id" db:"id"`
	HtmlContent string `json:"HtmlContent" db:"html_content"`
	PdfFileName string `json:"PdfFileName" db:"pdf_file_name"`
}

func (d *Document) Create() error {

	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return err
	}
	//	htmlStr := `<html><body><h1 style="color:red;">This is an html
	// from pdf to test color<h1><img src="http://api.qrserver.com/v1/create-qr-
	//code/?data=HelloWorld" alt="img" height="42" width="42"></img></body></html>`

	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(d.HtmlContent)))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		return err
	}

	err = pdfg.WriteFile(d.PdfFileName)
	if err != nil {
		return err
	}

	sq := "INSERT INTO document (html_content, pdf_file_name) VALUES ($1, $2) RETURNING id"
	err = helper.App.DB.QueryRow(sq, d.HtmlContent, d.PdfFileName).Scan(&d.ID)
	if err != nil {
		return err
	}

	fmt.Println("insert document table Done! ")

	return nil
}

func (d *Document) Delete(id int64) error {

	sq := "DELETE FROM document WHERE id = $1"
	_, err := helper.App.DB.Exec(sq, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *Document) Get(id int64) error {

	sq := "SELECT id, html_content, pdf_file_name FROM document WHERE id = $1"
	err := helper.App.DB.QueryRow(sq, id).Scan(&d.ID, &d.HtmlContent, &d.PdfFileName)
	if err != nil {
		return err
	}
	return nil
}

func (d *Document) GetAll() ([]Document, error) {

	rows, err := helper.App.DB.Query("SELECT id,html_content,pdf_file_name FROM document")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// An album slice to hold data from returned rows.
	var Documents []Document

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var Document Document
		if err := rows.Scan(&Document.ID, &Document.HtmlContent, &Document.PdfFileName); err != nil {
			return Documents, err
		}
		Documents = append(Documents, Document)
	}
	if err = rows.Err(); err != nil {
		return Documents, err
	}
	return Documents, nil
}

