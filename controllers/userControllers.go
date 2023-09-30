package controllers

import (
	"encoding/json"
	"net/http"
	"sagara-try/handlers"
	"sagara-try/helpers"
	"strconv"

	"github.com/gorilla/mux"
)

func C_Login(w http.ResponseWriter, r *http.Request) {
	datum := &handlers.User{}

	err := json.NewDecoder(r.Body).Decode(datum)
	if err != nil {
		helpers.Logger("error", "In Server: Oopss server someting wrong"+err.Error())
		msg := helpers.MsgErr(http.StatusInternalServerError, "internal server error", err.Error())
		helpers.Response(w, http.StatusInternalServerError, msg)
		return
	}

	result, err := datum.H_Login()
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	Logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "Login success, response: "+string(Logger))
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_Register(w http.ResponseWriter, r *http.Request) {
	datum := &handlers.User{}

	err := json.NewDecoder(r.Body).Decode(datum)
	if err != nil {
		helpers.Logger("error", "In Server: Oopss server someting wrong"+err.Error())
		msg := helpers.MsgErr(http.StatusInternalServerError, "internal server error", err.Error())
		helpers.Response(w, http.StatusInternalServerError, msg)
		return
	}

	result, err := datum.H_Register()
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		format := helpers.FormatError(err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", format.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(201, "Successfully")
	rMsg["body"] = result
	logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "created user: "+string(logger))
	helpers.Response(w, http.StatusCreated, rMsg)
}

func C_GetOneUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["user_id"]

	result, err := handlers.H_GetOneUser(userId)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "view one data user: "+string(logger))
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_UpdateOneUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["user_id"]
	userIdInt, _ := strconv.Atoi(userId)

	datum := handlers.User{}
	err := json.NewDecoder(r.Body).Decode(&datum)
	if err != nil {
		helpers.Logger("error", "In Server: Oopss server someting wrong"+err.Error())
		msg := helpers.MsgErr(http.StatusInternalServerError, "internal server error", err.Error())
		helpers.Response(w, http.StatusInternalServerError, msg)
		return
	}

	result, err := datum.H_UpdateOneUser(userIdInt)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "updated one data user"+string(logger))
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_DeleteOneUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["user_id"]

	result, err := handlers.H_DeleteOneUser(userId)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Succesfully")
	rMsg["deleted"] = result
	helpers.Logger("info", "deleted one data user")
	helpers.Response(w, http.StatusOK, rMsg)
}
