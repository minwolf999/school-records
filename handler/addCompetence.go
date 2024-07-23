package handler

import (
	"dossier_scolaire/database/controller"
	"dossier_scolaire/structure"
	"dossier_scolaire/utility"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// These hundle function the addCompetence page
func AddCompetenceHandler(w http.ResponseWriter, r *http.Request) {
	// Get the id of the teacher from the context
	id := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, id)
	if err != nil {
		// Write someone comming to the addCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}

	// Get all the categories and remove the categorie not accessible to the teacher
	categories, err := controller.SelectCategories("categories", []string{})
	if err != nil {
		// Write someone comming to the addCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}

	categories = utility.FilterCategoriesForTeacher(categories, teacher)

	// If the method isn't post execute the template
	if r.Method != http.MethodPost {
		// Write someone comming to the addCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addCompetence route - %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "addCompetence.html", structure.Result{Categories: categories})
		return
	}

	errorCount := 0

	// Get the value from the form
	categorieId := r.FormValue("categorie")
	subCategorieId := r.FormValue("subCategorie")
	name := r.FormValue("name")

	// Write someone comming to the addCompetence route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addCompetence route - %s - datas: [%s, %s, %s]\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, name, subCategorieId, categorieId))
	file.Close()

	// Verify if there isn't an empty field
	for _, v := range categories {
		if v.Id == categorieId && len(v.SubCategories) != 0 {
			if subCategorieId == "" {
				// Write someone comming to the addCompetence route in the log file
				file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, "empty field"))
				file.Close()

				structure.Tpl.ExecuteTemplate(w, "addCompetence.html", structure.Result{Categories: categories, Error: "Aucune sous-catégorie n'a été sélectionné !"})
				return
			}
		}
	}

	in, header, err := r.FormFile("image")
	if categorieId == "" || name == "" || err != nil || header.Filename == "" {
		if err != nil {
			// Write someone comming to the addCompetence route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
			file.Close()
		} else {
			// Write someone comming to the addCompetence route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, "empty field"))
			file.Close()
		}
		structure.Tpl.ExecuteTemplate(w, "addCompetence.html", structure.Result{Categories: categories, Error: "Il y a un champ vide !"})
		return
	}
	defer in.Close()

setImageName:
	// Set the path to save the image
	image_path := "/template/upload/" + utility.NewUUID() + "." + strings.Split(header.Filename, ".")[len(strings.Split(header.Filename, "."))-1]

	// Get the extension of he file and verify if the file is a JPG, JPEG, PNG
	image_extension := strings.Split(image_path, ".")[len(strings.Split(image_path, "."))-1]
	if image_extension != "jpg" && image_extension != "jpeg" && image_extension != "png" {
		structure.Tpl.ExecuteTemplate(w, "addCompetence.html", structure.Result{Categories: categories, Error: "Le format de l'image n'est pas supporter par le serveur !"})
		return
	}

	// Verify to not have 2 image with the same name
	if _, err = os.Open(image_path[1:]); err == nil {
		if errorCount > 20 {
			// Write someone comming to the addCompetence route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, "error image name"))
			file.Close()

			structure.Tpl.ExecuteTemplate(w, "addCompetence.html", structure.Result{Categories: categories, Error: "Une image avec le même nom est déjà présent dans le serveur veuillez changer le nom de l'image !"})
		} else {
			errorCount++
			goto setImageName
		}
		return
	}

	// Copy the image in the upload folder
	out, err := os.Create(image_path[1:])
	if err != nil {
		// Write someone comming to the addCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "addCompetence.html", structure.Result{Categories: categories, Error: "Il y a un problème dans votre image !"})
		return
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	if err != nil {
		// Write someone comming to the addCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}

	// Add the new competence in the BDD
	err = controller.AddNew("competences", utility.NewUUID(), name, image_path, categorieId, subCategorieId, teacher.Id)
	if err != nil {
		// Write someone comming to the addCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "addCompetence.html", structure.Result{Categories: categories, Error: err.Error()})
		return
	}

	// Execute the template with a success message
	err = structure.Tpl.ExecuteTemplate(w, "addCompetence.html", structure.Result{Categories: categories, Success: fmt.Sprintf("La compétence `%s` a été créé avec succès!", name)})
	if err != nil {
		// Write someone comming to the addCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/addCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, err))
		file.Close()
		return
	}
}
