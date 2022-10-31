package model

import (
	"fmt"
	"github.com/burakkuru5534/src/helper"
	"log"
	"strings"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type DocumentRequest struct {
	HtmlContent string `json:"HtmlContent"`
	PdfFileName string `json:"PdfFileName"`
}

type Document struct {
	ID            int    `json:"id" db:"id"`
	HtmlContent   string `json:"HtmlContent" db:"html_content"`
	PdfFileName   string `json:"PdfFileName" db:"pdf_file_name"`
	DirectoryPath string `json:"DirectoryPath" db:"directory_path"`
	FileSize      int64  `json:"FileSize" db:"file_size"`
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
	size := len(pdfg.Bytes())
	fmt.Println("create pdf file Done! ")

	sq := "INSERT INTO document (html_content, pdf_file_name, directory_path, file_size) VALUES ($1, $2, $3, $4) RETURNING id"
	err = helper.App.DB.QueryRow(sq, d.HtmlContent, d.PdfFileName, d.DirectoryPath, size).Scan(&d.ID)
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

	sq := "SELECT id, html_content, pdf_file_name, directory_path FROM document WHERE id = $1"
	err := helper.App.DB.QueryRow(sq, id).Scan(&d.ID, &d.HtmlContent, &d.PdfFileName, &d.DirectoryPath)
	if err != nil {
		return err
	}
	return nil
}

func (d *Document) GetAll() ([]Document, error) {

	rows, err := helper.App.DB.Query("SELECT id,html_content,pdf_file_name,directory_path FROM document")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// An album slice to hold data from returned rows.
	var Documents []Document

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var Document Document
		if err := rows.Scan(&Document.ID, &Document.HtmlContent, &Document.PdfFileName, &Document.DirectoryPath); err != nil {
			return Documents, err
		}
		Documents = append(Documents, Document)
	}
	if err = rows.Err(); err != nil {
		return Documents, err
	}
	return Documents, nil
}

//func htmlToPdf(html string) ([]byte, error) {
//	// Create a new PDF document.
//	pdf := gofpdf.New("P", "mm", "A4", "")
//	pdf.AddPage()
//	pdf.SetFont("Arial", "", 14)
//	pdf.Write(5, html)
//	// Get content as io.Reader.
//	buf := new(bytes.Buffer)
//	if err := pdf.Output(buf); err != nil {
//		return nil, err
//	}
//	return buf.Bytes(), nil
//}

//func (d *Document) GeneratePdf() ([]byte, error) {
//	// Get content as io.Reader.
//	buf := new(bytes.Buffer)
//	if err := pdf.Output(buf); err != nil {
//		return nil, err
//	}
//	return buf.Bytes(), nil
//}

func myPdfFunc(htmlContent string, pdfFileName string, directoryPath string) error {
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return err
	}
	//	htmlStr := `<html><body><h1 style="color:red;">This is an html
	// from pdf to test color<h1><img src="http://api.qrserver.com/v1/create-qr-
	//code/?data=HelloWorld" alt="img" height="42" width="42"></img></body></html>`

	pdfg.AddPage(wkhtml.NewPageReader(strings.NewReader(htmlContent)))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	//Your Pdf Name
	//directoryPath = "./"
	pdfFileName = fmt.Sprintf("%s%s.pdf", directoryPath, pdfFileName)
	err = pdfg.WriteFile(pdfFileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
	return nil

}

//func mySecHtmlToPdfFunc() {
//	r := u.NewRequestPdf("")
//
//	//html template path
//	templatePath := "templates/sample.html"
//
//	//path for download pdf
//	outputPath := "storage/example.pdf"
//
//	//html template data
//	templateData := struct {
//		Title       string
//		Description string
//		Company     string
//		Contact     string
//		Country     string
//	}{
//		Title:       "HTML to PDF generator",
//		Description: "This is the simple HTML to PDF file.",
//		Company:     "Jhon Lewis",
//		Contact:     "Maria Anders",
//		Country:     "Germany",
//	}
//
//	if err := r.ParseTemplate(templatePath, templateData); err == nil {
//		ok, _ := r.GeneratePDF(outputPath)
//		fmt.Println(ok, "pdf generated successfully")
//	} else {
//		fmt.Println(err)
//	}
//}
