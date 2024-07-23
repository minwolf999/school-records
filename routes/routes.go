package routes

import (
	"dossier_scolaire/XML"
	"dossier_scolaire/fetch"
	"dossier_scolaire/handler"
	"dossier_scolaire/middleware"
	"net/http"
)

// This function create all the routes of the web site for the different page
func Routes() {
	// The user page
	http.HandleFunc("/", handler.RedirectHandler)

	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/register", handler.RegisterHandler)

	http.HandleFunc("/home", middleware.VerificationCookie(handler.HomeHandler))
	http.HandleFunc("/saisir", middleware.VerificationCookie(handler.SaisirHandler))

	http.HandleFunc("/saisir/listCompetence", middleware.VerificationCookie(handler.ListCompetenceHandler))
	http.HandleFunc("/saisir/addCompetence", middleware.VerificationCookie(handler.AddCompetenceHandler))
	http.HandleFunc("/saisir/modifyCompetence", middleware.VerificationCookie(handler.ModifyCompetence))

	http.HandleFunc("/saisir/listEleve", middleware.VerificationCookie(handler.ListEleve))
	http.HandleFunc("/saisir/addEleve", middleware.VerificationCookie(handler.AddEleve))
	http.HandleFunc("/saisir/modifyEleve", middleware.VerificationCookie(handler.ModifyEleve))
	
	http.HandleFunc("/saisir/addLinkCompetenceEleve", middleware.VerificationCookie(handler.AddLinkCompetenceEleve))

	http.HandleFunc("/saisir/changeSigningUp", middleware.VerificationCookie(handler.ChangeSigningUp))
	http.HandleFunc("/saisir/removeLinkCompetenceEleve", middleware.VerificationCookie(handler.RemoveLinkCompetenceEleve))

	http.HandleFunc("/observer", middleware.VerificationCookie(handler.ObserverHandler))
	http.HandleFunc("/createPDF", middleware.VerificationCookie(handler.CreatePDF))

	// The fetch page
	http.HandleFunc("/getStudents", middleware.VerificationCookie(fetch.GetStudents))
	http.HandleFunc("/getCompetencesWithStudents", middleware.VerificationCookie(fetch.GetStudentsByCompetence))
	http.HandleFunc("/getCompetences", middleware.VerificationCookie(fetch.GetCompetences))
	http.HandleFunc("/removeLink", middleware.VerificationCookie(fetch.RemoveLink))

	http.HandleFunc("/getSubCategories", middleware.VerificationCookie(fetch.GetSubCategorie))

	// The XML page
	http.HandleFunc("/deleteCompetence", middleware.VerificationCookie(XML.DeleteCompetence))
	http.HandleFunc("/deleteEleve", middleware.VerificationCookie(XML.DeleteEleve))
}
