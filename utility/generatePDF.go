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
	pageHeight := 0
	categorieCount := 0
	currentI := 0
	currentPage := 0

	// Set a font and write the name of the student
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, translator(student.Name))

	// Set a font and write the school year
	pdf.SetFont("Arial", "", 16)
	pdf.CellFormat(0, 10, translator(fmt.Sprintf("Année scolaire %s", student.Year)), "", 0, "R", false, 0, "")

	// Make a break line
	pdf.Ln(10)
	pageHeight += 10

	// Write the name of the student class
	pdf.Cell(40, 10, translator(fmt.Sprintf("Classe de %s", student.Class)))

	// Make a break line
	pdf.Ln(10)
	pageHeight += 10

	// Draw a line
	pdf.SetDrawColor(255/2, 255/2, 255/2)
	pdf.SetLineWidth(0.5)
	pdf.Line(10, 30, width-10, 30)

	// Write the detail f the competence
	for i := range student.Competences {
		// Verify if there is the place to add a new line if not add a new page and reset some variable
		if height-float64(pageHeight) < 70 {
			pdf.AddPage()
			pageHeight = 0
			categorieCount = 0
			currentI = 0
			currentPage++
		}

		// If there is a new Categorie, write the name of the categorie in italic bold
		if i == 0 || student.Competences[i].Categorie.Name != student.Competences[i-1].Categorie.Name {
			if i != 0 {
				categorieCount++
			}

			pdf.Ln(5)
			pageHeight += 5

			pdf.SetFont("Arial", "BI", 16)
			pdf.Cell(40, 10, translator(student.Competences[i].Categorie.Name))

			pdf.Ln(20)
			pageHeight += 20
		}

		// If there is a new sub-Categorie, write the name of the categorie in bold
		if student.Competences[i].SubCategorie.Name != "" && (i == 0 || student.Competences[i].SubCategorie.Name != student.Competences[i-1].SubCategorie.Name) {
			if i != 0 {
				categorieCount++
			}

			pdf.Ln(5)
			pageHeight += 5

			pdf.SetFont("Arial", "B", 16)
			pdf.Cell(40, 10, translator(student.Competences[i].SubCategorie.Name))

			pdf.Ln(20)
			pageHeight += 10
		}

		// Set the font to normal and write the name of the competence
		pdf.SetFont("Arial", "", 16)
		pdf.Cell(40, 10, translator(student.Competences[i].Name))

		// Place the image of the competence
		if currentPage == 0 {
			pdf.Image(student.Competences[i].ImagePath[1:], width-65, float64(currentI*27+50+categorieCount*25), 55, 25, false, "", 0, "")
		} else {
			pdf.Image(student.Competences[i].ImagePath[1:], width-65, float64(currentI*27+3+categorieCount*25), 55, 25, false, "", 0, "")
		}
		pdf.Ln(27)
		pageHeight += 27

		currentI++
	}

	// Write a field to the parent Signin up
	pdf.Ln(10)
	pageHeight += 10
	pdf.Cell(0, 10, translator("Signature des parents"))

	// Write the signin up of the teacher field
	pdf.CellFormat(0, 10, translator("Signature de l'enseignant"), "", 0, "R", false, 0, "")
	pdf.Ln(10)
	pageHeight += 10

	if teacher.SigningUpPath != "" {
		FullDecrypt(teacher)

		pdf.Image(teacher.SigningUpPath[1:], width-55, float64(pageHeight+10), 25, 25, false, "", 0, "")

		go func() {
			time.Sleep(2 * time.Second)
			FullEncrypt(teacher)
		}()
	}

	return pdf
}
