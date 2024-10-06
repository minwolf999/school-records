package utility

import (
	"dossier_scolaire/structure"
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
)

// These function take 2 arguments: a students and a teacher and return a pdf ready to be used
func GeneratePDF(student structure.Student, teacher structure.Teacher) *gofpdf.Fpdf {
	// Initialise the pdf
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a new page to the pdf
	pdf.AddPage()

	// Init variables to be used later
	translator := pdf.UnicodeTranslatorFromDescriptor("")
	width, height := pdf.GetPageSize()
	pageHeight := 0.0

	// Set a font and write the name of the student
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, translator(student.Name))

	// Set a font and write the school year
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, 10, translator(fmt.Sprintf("Année scolaire %s", student.Year)), "", 0, "R", false, 0, "")

	// Make a break line
	pdf.Ln(10)

	// Write the name of the student class
	pdf.Cell(40, 10, translator(fmt.Sprintf("Classe de %s", student.Class)))

	// Make a break line
	pdf.Ln(10)

	// Draw a line
	pdf.SetDrawColor(128, 128, 128)
	pdf.SetLineWidth(0.5)
	pdf.Line(10, 30, width-10, 30)
	pdf.Ln(10)

	// Write the detail of the competence
	for i, comp := range student.Competences {
		// Verify if there is the place to add a new line, if not, add a new page and reset some variables
		pageHeight = pdf.GetY()
		if height-pageHeight < 50 {
			pdf.AddPage()
			pageHeight = 0
		}

		// If there is a new Category, write the name of the category in italic bold
		if i == 0 || comp.Categorie.Name != student.Competences[i-1].Categorie.Name {
			pdf.SetFont("Arial", "BI", 15)
			pdf.Cell(0, 10, translator(comp.Categorie.Name))
			pdf.Ln(10)
		}

		// If there is a new sub-Category, write the name of the sub-category in bold
		if comp.SubCategorie.Name != "" && (i == 0 || comp.SubCategorie.Name != student.Competences[i-1].SubCategorie.Name) {
			pdf.SetFont("Arial", "B", 13)
			pdf.Cell(0, 10, translator(comp.SubCategorie.Name))
			pdf.Ln(10)
		}

		// Set the font to normal and write the name of the competence
		pdf.SetFont("Arial", "", 10)
		pdf.Cell(0, 10, translator(comp.Name))

		// Place the image of the competence
		if comp.ImagePath != "" {
			pageHeight = pdf.GetY() - 3
			pdf.Image(comp.ImagePath[1:], width-30, pageHeight, 20, 16, false, "", 0, "")
		}

		// Add space between competences
		pdf.Ln(20)
	}

	pageHeight = pdf.GetY()

	// Draw a line
	pdf.SetDrawColor(128, 128, 128)
	pdf.SetLineWidth(0.5)
	pdf.Line(10, pageHeight, width-10, pageHeight)
	pdf.Ln(10)

	// Write a field for the parent's signature
	pdf.Ln(10)
	pdf.Cell(0, 10, translator("Signature des parents"))

	// Write the teacher's signature field
	pdf.CellFormat(0, 10, translator("Signature de l'enseignant"), "", 0, "R", false, 0, "")
	pdf.Ln(10)

	if teacher.SigningUpPath != "" {
		FullDecrypt(teacher)
		pageHeight = pdf.GetY()

		pdf.Image(teacher.SigningUpPath[1:], width-55, pageHeight+10, 25, 25, false, "", 0, "")

		go func() {
			time.Sleep(2 * time.Second)
			FullEncrypt(teacher)
		}()
	}

	return pdf
}
