package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.LowQuality.Set(true)

	ref := "MOJ-x64t1wzrR"

	buf := new(bytes.Buffer)
	err = template.Must(template.ParseFiles("./template.html")).Execute(buf,
		struct{ Reference string }{
			Reference: ref,
		},
	)
	if err != nil {
		log.Fatalf("unable to execute files due: %v", err)
	}

	// Add to document
	pdfg.AddPage(wkhtmltopdf.NewPageReader(buf))

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile(fmt.Sprintf("./data/%v.pdf", ref))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done
}
