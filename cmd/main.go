package cmd

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func Runner() {
	server := http.NewServeMux()
	routes(server)
	fs := http.FileServer(http.Dir("templates/assets"))
	server.Handle("/assets/", http.StripPrefix("/assets", fs))
	fmt.Println("\n\033[34m[http://127.0.0.1"+port+"]\033[32m \033[4mServer run on port", port, ".\033[0m")

	err := http.ListenAndServe(port, server)

	if err != nil {
		fmt.Println("error :", err)
		return
	}
}

func routes(server *http.ServeMux) {
	//server.HandleFunc("/", authGuestSecurity(rootHandler))
	server.HandleFunc("/homepage", mainHandlers)
	server.HandleFunc("/login", loginHandlers)
	server.HandleFunc("/login/google", loginGoogle)
	server.HandleFunc("/login/github", loginGithub)
	server.HandleFunc("/register", registerHandlers)
	server.HandleFunc("/admin", adminHandlers)
	server.HandleFunc("/post", postHandlers)
	server.HandleFunc("/api/login", loginPost)
	//server.HandleFunc("/api/test", CreateAccountGoogle)
	server.HandleFunc("/api/loginGoogle", loginGoogle)
	server.HandleFunc("/api/loginGithub", loginGithub)
	server.HandleFunc("/api/callbacklogingoogle", callbackLoginGoogle)
	server.HandleFunc("/api/callbacklogingithub", callbackLoginGithub)
	server.HandleFunc("/api/take-ban", adminPanel)
	server.HandleFunc("/api/register", CreateUser)
	server.HandleFunc("/api/adminpanel", adminPanel)
	server.HandleFunc("/api/catch-info-admin", sendInfoAdmin)
	server.HandleFunc("/api/create-post", createPostHandler)
	server.HandleFunc("/api/display-post", displayPostVisible)
	server.HandleFunc("/api/createcomment", createComment)
	server.HandleFunc("/api/editPost", authUserSecurity(editPost))
	server.HandleFunc("/api/getComments", getComments)
	server.HandleFunc("/api/takepostid", sendDataPostWithId)
	server.HandleFunc("/api/likeordislike", postLikeOrDislike)
}
