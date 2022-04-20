package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

//Método Handler sem criar uma interface
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the user controller!"))
}

//Nao existem construtores em GO, mas existem Funcoes Construtoras
//Funcao Construtora que retorna um apontador para um objeto userController
func newUserController() *userController {
	return &userController{
		//vou procurar no caminho /users/(seguido de um numero)
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
