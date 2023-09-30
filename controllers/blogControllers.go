package controllers

import (
	"encoding/json"
	"net/http"
	"sagara-try/handlers"
	"sagara-try/helpers"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func C_SaveBlog(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(jwt.MapClaims)
	userId := user["user_id"].(string)
	userIdInt, _ := strconv.Atoi(userId)

	datum := handlers.Blog{}
	err := json.NewDecoder(r.Body).Decode(&datum)
	if err != nil {
		helpers.Logger("error", "In Server: Oopss server someting wrong"+err.Error())
		msg := helpers.MsgErr(http.StatusInternalServerError, "internal server error", err.Error())
		helpers.Response(w, http.StatusInternalServerError, msg)
		return
	}

	result, err := datum.H_SaveBlog(userIdInt)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(201, "Successfully")
	rMsg["body"] = result
	logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "saved blog: "+string(logger))
	helpers.Response(w, http.StatusCreated, rMsg)
}

func C_GetBlogsbyUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(jwt.MapClaims)
	userId := user["user_id"].(string)
	userIdInt, _ := strconv.Atoi(userId)

	result, err := handlers.H_GetBlogByUser(userIdInt)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	helpers.Logger("info", "view all blog by user")
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_GetPaginatedBlogs(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(jwt.MapClaims)
	userId := user["user_id"].(string)
	userIdInt, _ := strconv.Atoi(userId)

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	result, totalRows, err := handlers.H_GetPaginatedBlog(userIdInt, page, limit)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	totalPages := (totalRows + limit - 1) / limit

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	rMsg["total_pages"] = totalPages
	rMsg["current_page"] = page
	helpers.Logger("info", "view blog by paginated")
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_GetOneBlog(w http.ResponseWriter, r *http.Request) {
	blogId := mux.Vars(r)["blog_id"]
	blogIdInt, _ := strconv.Atoi(blogId)

	result, err := handlers.H_GetOneBlog(blogIdInt)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "view one blog: "+string(logger))
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_UpdateOneBlog(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(jwt.MapClaims)
	userId := user["user_id"].(string)
	userIdInt, _ := strconv.Atoi(userId)

	blogId := mux.Vars(r)["blog_id"]
	blogIdInt, _ := strconv.Atoi(blogId)

	datum := handlers.Blog{}
	err := json.NewDecoder(r.Body).Decode(&datum)
	if err != nil {
		helpers.Logger("error", "In Server: Oopss server someting wrong"+err.Error())
		msg := helpers.MsgErr(http.StatusInternalServerError, "internal server error", err.Error())
		helpers.Response(w, http.StatusInternalServerError, msg)
		return
	}

	result, err := datum.H_UpdateOneBlog(blogIdInt, userIdInt)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	logger, _ := json.Marshal(rMsg)
	helpers.Logger("info", "updated one data blog: "+string(logger))
	helpers.Response(w, http.StatusOK, rMsg)
}

func C_DeleteOneBlog(w http.ResponseWriter, r *http.Request) {
	blogId := mux.Vars(r)["blog_id"]
	blogIdInt, _ := strconv.Atoi(blogId)

	result, err := handlers.H_DeleteOneBlog(blogIdInt)
	if err != nil {
		helpers.Logger("error", "In Server: "+err.Error())
		msg := helpers.MsgErr(http.StatusBadRequest, "bad request", err.Error())
		helpers.Response(w, http.StatusBadRequest, msg)
		return
	}

	rMsg := helpers.MsgOk(200, "Successfully")
	rMsg["body"] = result
	helpers.Logger("info", "deleted one blog ")
	helpers.Response(w, http.StatusOK, rMsg)
}
