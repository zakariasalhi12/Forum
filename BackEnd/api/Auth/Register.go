package auth

import (
	"net/http"

	models "forum/BackEnd/Models"
	"forum/BackEnd/helpers"
)

// RegisterAPI handles the registration of a new user.
func RegisterAPI(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		helpers.Writer(w, map[string]string{"Error": helpers.ErrMethod.Error()}, http.StatusMethodNotAllowed)
		return
	}
	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Create a new instance of NewUser model
	NewUser := models.NewUser()

	// Get the Body Request And Parse It into my newuser Model
	Status, err := helpers.ParseRequestBody(r, &NewUser)
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, Status)
	}

	// Check if any of the required fields (Email, Password, UserName) are empty
	if helpers.CheckEmpty(NewUser.Email, NewUser.Password, NewUser.UserName) {
		helpers.Writer(w, map[string]string{"Error": helpers.ErrInvalidRequest.Error()}, 400)
		return
	}

	// Attempt to add the new user to the database
	err = NewUser.AddUserTodb(w)
	if err == models.EmailAlreadyUsed || err == models.InvalidEmail || err == models.InvalidPassword {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 400)
		return
	}
	if err != nil {
		helpers.Writer(w, map[string]string{"Error": err.Error()}, 500)
		return
	}

	// If the registration is successful, return a success message with HTTP Status 200 OK
	helpers.Writer(w, map[string]string{"Message": "Registration successful!"}, 200)
}
