package routes

import (
	"encoding/json"
	"net/http"

	"github.com/dapedrozo/go-gorm-resapi/db"
	"github.com/dapedrozo/go-gorm-resapi/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks) //busco la relacion de la tabla user con la tabla task
	json.NewEncoder(w).Encode(&user)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(&user)
}

//este metodo solo cambia el dato a eliminado en una fecha pero no lo elimina realmente
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	db.DB.Delete(&user)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user deleted"))
}

//este metodo elimina permanentemente de la bd
func DeleteUserPermanentlyHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}
	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user deleted permanently"))
}
