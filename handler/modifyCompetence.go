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

// These hundle function the modifyCompetence page
func ModifyCompetence(w http.ResponseWriter, r *http.Request) {
	// Get the id of the teacher from the context
	id := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, id)
	if err != nil {
		// Write someone comming to the modifyCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}

	// Get the id of the competence selected to modify these informations
	ck, err := r.Cookie("competenceId")
	if err != nil {
		// Write someone comming to the modifyCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()

		http.Redirect(w, r, "/saisir/listCompetence", http.StatusSeeOther)
		return
	}
	competenceId := ck.Value

	// Get the detail of the competence selected
	competences, err := controller.SelectCompetences("competences", []string{"id", "teacherId"}, competenceId, teacher.Id)
	if err != nil {
		// Write someone comming to the modifyCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}

	// Get all the categories and remove the categorie not accessible to the teacher
	categories, err := controller.SelectCategories("categories", []string{})
	if err != nil {
		// Write someone comming to the modifyCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}
	categories = utility.FilterCategoriesForTeacher(categories, teacher)

	// If the method isn't post execute the template
	if r.Method != http.MethodPost {
		// Write someone comming to the modifyCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "modifyCompetence.html", structure.Result{Categories: categories, Competence: competences[0]})
		return
	}

	// Get the value from the form
	categorieId := r.FormValue("categorie")
	subCategorieId := r.FormValue("subCategorie")
	name := r.FormValue("name")

	// Write someone comming to the modifyCompetence route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s - datas: [%s, %s, %s]\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, name, categorieId, subCategorieId))
	file.Close()

	// Verify if there isn't an empty field
	for _, v := range categories {
		if v.Id == categorieId && len(v.SubCategories) != 0 {
			if subCategorieId == "" {
				// Write someone comming to the modifyCompetence route in the log file
				file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, "empty field"))
				file.Close()

				structure.Tpl.ExecuteTemplate(w, "modifyCompetence.html", structure.Result{Categories: categories, Competence: competences[0], Error: "Aucune sous-catégorie n'a été sélectionné !"})
				return
			}
		}
	}

	if categorieId == "" || name == "" {
		// Write someone comming to the modifyCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, "empty field"))
		file.Close()

		structure.Tpl.ExecuteTemplate(w, "modifyCompetence.html", structure.Result{Categories: categories, Competence: competences[0], Error: "Il y a un champ vide !"})
		return
	}

	in, header, err := r.FormFile("image")
	if err != nil {
		// If there is no image update the other information
		err = controller.Update("competences", map[string]string{"name": name, "categorieId": categorieId, "subCategorieId": subCategorieId}, []string{"id", "teacherId"}, competenceId, teacher.Id)
		if err != nil {
			// Write someone comming to the modifyCompetence route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
			file.Close()

			structure.Tpl.ExecuteTemplate(w, "modifyCompetence.html", structure.Result{Categories: categories, Error: err.Error()})
			return
		}
	} else {
		defer in.Close()
		errorCount := 0

	setImageName:

		// Set the path to save the image
		image_path := "/template/upload/" + utility.NewUUID() + "." + strings.Split(header.Filename, ".")[len(strings.Split(header.Filename, "."))-1]

		// Get the extension of he file and verify if the file is a JPG, JPEG, PNG
		image_extension := strings.Split(image_path, ".")[len(strings.Split(image_path, "."))-1]
		if image_extension != "jpg" && image_extension != "jpeg" && image_extension != "png" {
			// Write someone comming to the modifyCompetence route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, "incompatible image"))
			file.Close()

			structure.Tpl.ExecuteTemplate(w, "modifyCompetence.html", structure.Result{Categories: categories, Competence: competences[0], Error: "Le format de l'image n'est pas supporter par le serveur !"})
			return
		}

		// Remove the previous image
		if _, err = os.Open(image_path[1:]); err == nil {
			err = os.Remove(competences[0].ImagePath[1:])
			if err != nil {
				// Write someone comming to the modifyCompetence route in the log file
				file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
				file.Close()
				return
			}
		}

		// Verify to not have 2 image with the same name
		if _, err = os.Open(image_path[1:]); err == nil {
			if errorCount > 20 {
				// Write someone comming to the modifyCompetence route in the log file
				file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
				file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, "impossible to create a unique iage name"))
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
			// Write someone comming to the modifyCompetence route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
			file.Close()

			structure.Tpl.ExecuteTemplate(w, "modifyCompetence.html", structure.Result{Categories: categories, Competence: competences[0], Error: "Il y a un problème dans votre image !"})
			return
		}
		defer out.Close()
		io.Copy(out, in)

		// Update the information of the teacher in the BDD
		err = controller.Update("competences", map[string]string{"name": name, "imagePath": image_path, "categorieId": categorieId, "subCategorieId": subCategorieId}, []string{"id", "teacherId"}, competenceId, teacher.Id)
		if err != nil {
			// Write someone comming to the modifyCompetence route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
			file.Close()

			structure.Tpl.ExecuteTemplate(w, "modifyCompetence.html", structure.Result{Categories: categories, Competence: competences[0], Error: err.Error()})
			return
		}
	}

	// Execute the template with a success message
	competences, err = controller.SelectCompetences("competences", []string{"id", "teacherId"}, competenceId, teacher.Id)
	if err != nil {
		// Write someone comming to the modifyCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}

	err = structure.Tpl.ExecuteTemplate(w, "modifyCompetence.html", structure.Result{Categories: categories, Competence: competences[0], Success: fmt.Sprintf("La compétence `%s` a été mis à jour avec succès!", name)})
	if err != nil {
		// Write someone comming to the modifyCompetence route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/modifyCompetence route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
	}
}
