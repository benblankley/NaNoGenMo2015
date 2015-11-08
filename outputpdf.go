package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{Unit: "pt", PageSize: gopdf.Rect{W: 595.28, H: 841.89}}) //595.28, 841.89 = A4
	pdf.AddPage()
	err := pdf.AddTTFFont("DejaVuSerif", "./src/github.com/signintech/gopdf/gopdftest/ttf/DejaVuSerif.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetLineWidth(2)
	pdf.Line(10, 400, 585, 400)

	err = pdf.SetFont("DejaVuSerif", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "test")
	pdf.Br(20)
	pdf.Cell(nil, "test2")
	pdf.Br(20)
	pdf.WritePdf("line3.pdf")

}