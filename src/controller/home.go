package controller

import (
	"net/http"
	"src/src/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Simple Bank")

}
