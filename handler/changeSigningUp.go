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

// These hundle function the changeSigninUp page
func ChangeSigningUp(w http.ResponseWriter, r *http.Request) {
	// Get the id of the teacher from the context
	id := fmt.Sprintf("%s", r.Context().Value(structure.IdCtx))
	teacher, err := controller.SelectTeacher("teachers", []string{"id"}, id)
	if err != nil {
		// Write someone comming to the changeSigningUp route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/changeSigningUp route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}

	// If the method isn't post execute the template
	if r.Method != http.MethodPost {
		// Write someone comming to the changeSigningUp route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/changeSigningUp route - %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id))
		file.Close()

		utility.FullDecrypt(teacher)

		structure.Tpl.ExecuteTemplate(w, "changeSigningUp.html", structure.Result{ImagePath: teacher.SigningUpPath})

		go func() {
			time.Sleep(2 * time.Second)
			utility.FullEncrypt(teacher)
		}()

		return
	}

	// Write someone comming to the changeSigningUp route in the log file
	file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/changeSigningUp route - %s - datas: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, teacher.Id, "an image"))
	file.Close()

	errorCount := 0
	// Verify if there is no image
	in, header, err := r.FormFile("image")
	if err != nil {
		// Write someone comming to the changeSigningUp route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/changeSigningUp route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()

		utility.FullDecrypt(teacher)

		structure.Tpl.ExecuteTemplate(w, "changeSigningUp.html", structure.Result{Error: "L'image a été oubliée !", ImagePath: teacher.SigningUpPath})

		go func() {
			time.Sleep(2 * time.Second)
			utility.FullEncrypt(teacher)
		}()

		return
	}
	defer in.Close()

setImageName:
	// Set the path to save the image
	image_path := "/template/signinUp/" + utility.NewUUID() + "." + strings.Split(header.Filename, ".")[len(strings.Split(header.Filename, "."))-1]

	// Get the extension of he file and verify if the file is a JPG, JPEG, PNG
	image_extension := strings.Split(image_path, ".")[len(strings.Split(image_path, "."))-1]
	if image_extension != "jpg" && image_extension != "jpeg" && image_extension != "png" {
		// Write someone comming to the changeSigningUp route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/changeSigningUp route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, "forgotten image"))
		file.Close()

		utility.FullDecrypt(teacher)

		structure.Tpl.ExecuteTemplate(w, "changeSigningUp.html", structure.Result{Error: "Le format de l'image n'est pas supporter par le serveur !", ImagePath: teacher.SigningUpPath})

		go func() {
			time.Sleep(2 * time.Second)
			utility.FullEncrypt(teacher)
		}()
		return
	}

	// Verify to not have 2 image with the same name
	if _, err = os.Open(image_path[1:]); err == nil {
		if errorCount < 20 {
			errorCount++
			goto setImageName
		} else {
			// Write someone comming to the changeSigningUp route in the log file
			file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/changeSigningUp route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, "impossible generate unique image name"))
			file.Close()

			utility.FullDecrypt(teacher)

			structure.Tpl.ExecuteTemplate(w, "changeSigningUp.html", structure.Result{Error: "Une image avec le même nom est déjà présent dans le serveur veuillez changer le nom de l'image !", ImagePath: teacher.SigningUpPath})

			go func() {
				time.Sleep(2 * time.Second)
				utility.FullEncrypt(teacher)
			}()
			return
		}
	}

	// Crypt the image
	formatedCrypted := utility.Crypt(teacher, in)

	// Copy the image in the upload folder
	out, err := os.Create(image_path[1:])
	if err != nil {
		// Write someone comming to the changeSigningUp route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/changeSigningUp route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()

		utility.FullDecrypt(teacher)

		structure.Tpl.ExecuteTemplate(w, "changeSigningUp.html", structure.Result{Error: "Il y a un problème dans votre image !", ImagePath: teacher.SigningUpPath})

		go func() {
			time.Sleep(2 * time.Second)
			utility.FullEncrypt(teacher)
		}()
		return
	}
	defer out.Close()
	io.Copy(out, formatedCrypted)

	// Update the signinUp of a teacher in the BDD
	err = controller.Update("teachers", map[string]string{"signingPath": image_path}, []string{"id"}, teacher.Id)
	if err != nil {
		// Write someone comming to the changeSigningUp route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/changeSigningUp route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()

		utility.FullDecrypt(teacher)

		structure.Tpl.ExecuteTemplate(w, "changeSigningUp.html", structure.Result{Error: err.Error(), ImagePath: teacher.SigningUpPath})

		go func() {
			time.Sleep(2 * time.Second)
			utility.FullEncrypt(teacher)
		}()
		return
	}

	// Execute the template
	teacher.SigningUpPath = image_path

	utility.FullDecrypt(teacher)

	err = structure.Tpl.ExecuteTemplate(w, "changeSigningUp.html", structure.Result{ImagePath: teacher.SigningUpPath})
	if err != nil {
		// Write someone comming to the changeSigningUp route in the log file
		file, _ := os.OpenFile(structure.App.LogPath[1:], os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		file.WriteString(fmt.Sprintf("%s [%s] %s - /saisir/changeSigningUp route - %s - error: %s\n", time.Now().Format("2006-01-02 15:04:05"), r.RemoteAddr, r.Method, id, err))
		file.Close()
		return
	}
	
	go func() {
		time.Sleep(2 * time.Second)
		utility.FullEncrypt(teacher)
	}()
}
