package main

import (
	"os"
	"sagara-try/controllers"

	"github.com/gorilla/mux"
)

func Factory(ctm *mux.Router) string {

	middleUrl := os.Getenv("MIDDLE_URL")

	//auth
	ctm.HandleFunc(middleUrl+"/login", controllers.C_Login).Methods("POST")
	ctm.HandleFunc(middleUrl+"/register", controllers.C_Register).Methods("POST")

	//user
	ctm.HandleFunc(middleUrl+"/user/{user_id}/single", controllers.C_GetOneUser).Methods("GET")
	ctm.HandleFunc(middleUrl+"/user/{user_id}/edit", controllers.C_UpdateOneUser).Methods("PUT")
	ctm.HandleFunc(middleUrl+"/user/{user_id}/delete", controllers.C_DeleteOneUser).Methods("DELETE")

	//blog
	ctm.HandleFunc(middleUrl+"/blog", controllers.C_SaveBlog).Methods("POST")
	ctm.HandleFunc(middleUrl+"/blog", controllers.C_GetBlogsbyUser).Methods("GET")
	ctm.HandleFunc(middleUrl+"/blog/paginated", controllers.C_GetPaginatedBlogs).Methods("GET")
	ctm.HandleFunc(middleUrl+"/blog/{blog_id}/single", controllers.C_GetOneBlog).Methods("GET")
	ctm.HandleFunc(middleUrl+"/blog/{blog_id}/edit", controllers.C_UpdateOneBlog).Methods("PUT")
	ctm.HandleFunc(middleUrl+"/blog/{blog_id}/delete", controllers.C_DeleteOneBlog).Methods("DELETE")

	return "presenter :"
}
