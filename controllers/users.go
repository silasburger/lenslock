package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
	"use-go/lenslocked.com/views"
)

// NewUsers is used to create a new Users controller.
// This function will panic if the templates are not
// parsed correctly, and should only be used during
// initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Email    string `json: "email"`
	Password string `json: "password"`
}

// New is used to render the form where a user can
// create a new user account.
//
//GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// Create is used to process the signup form when a user
// submitw it. This is used to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	var form SignupForm
	decoder := schema.NewDecoder()
	err = decoder.Decode(&form, r.PostForm)
}
