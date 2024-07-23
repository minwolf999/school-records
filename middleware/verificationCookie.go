package middleware

import (
	"context"
	"dossier_scolaire/database/controller"
	"dossier_scolaire/structure"
	"net/http"
)

// These function look if the user have a correct cookie
func VerificationCookie(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify if there is a cookie named "cookie"
		data, err := r.Cookie("cookie")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Verify if there is a teacher with the id contain in the cookie
		user, err := controller.SelectTeacher("teachers", []string{"id"}, data.Value)
		if err != nil || user.Id == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Set the id in the context
		ctx := context.WithValue(r.Context(), structure.IdCtx, data.Value)
		newReq := r.WithContext(ctx)

		// Call the hundle function give in argument
		next(w, newReq)
	})
}
