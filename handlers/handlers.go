package handlers

import (
	"encoding/json"
	"gorm/db"
	"gorm/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {

	users := models.Users{}

	db.Database.Find(&users)
	sendData(rw, users, 200)

}

func GetUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, user, http.StatusOK)

	}

}

// Funcion para buscar por ID
func getUserById(r *http.Request) (models.User, *gorm.DB) {
	//Leo la url
	vars := mux.Vars(r)
	//Caputo el id de la url
	userId, _ := strconv.Atoi(vars["id"])
	user := models.User{}

	//Manejamos el error
	if err := db.Database.First(&user, userId); err.Error != nil {
		return user, err
	} else {
		return user, nil
	}

}

func CreateUser(rw http.ResponseWriter, r *http.Request) {

	//Obtener registro
	user := models.User{}

	//Obtengo todo el cuerpo de la peticion con json.NewDecoder
	decoder := json.NewDecoder(r.Body)

	//Devulve un error si no se cagaron los datos al objeto &user
	if err := decoder.Decode(&user); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database.Save(&user)
		sendData(rw, &user, http.StatusCreated)

	}

}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	//Obtener registro
	var userId int64

	if user_ant, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {

		userId = user_ant.Id

		//Obtengo todo el cuerpo de la peticion con json.NewDecoder
		user := models.User{}
		decoder := json.NewDecoder(r.Body)
		//Devulve un error si no se cagaron los datos al objeto &user
		if err := decoder.Decode(&user); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			user.Id = userId
			db.Database.Save(&user)
			sendData(rw, &user, http.StatusOK)

		}
	}

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserById(r); err != nil {
		// fmt.Println(rw)
		sendError(rw, http.StatusNotFound)

	} else {
		db.Database.Delete(&user)
		sendData(rw, user, http.StatusOK)
	}

}
