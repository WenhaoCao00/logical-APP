package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	initDB()
	insertDefaultAdmin()
	insertFixedQuestions()

	//user state
	http.HandleFunc("/signup/", handleSignup)
	http.HandleFunc("/login/", handleLogin)
	http.HandleFunc("/updateUserState/", handleUpdateUserState)
	http.HandleFunc("/getUserState/", handleGetUserState)
	

	//task generation
	http.HandleFunc("/generateConditionalTask/", handleGenerateConditionalTask)
	http.HandleFunc("/generateDeMorganTask/", handleGenerateDeMorganTask)
	http.HandleFunc("/generateTruthTableTask/", handleGenerateTruthTableTask)
	http.HandleFunc("/addTask/", handleAddTask)
	http.HandleFunc("/deleteTask/", handleDeleteTask)
	http.HandleFunc("/getTask/", handleGetTask)


	//admin page that can be accessed by admin
	http.HandleFunc("/admin/", handleAdminPage)
	http.HandleFunc("/admin/api/", handleApiPage)
	

	fmt.Println("Starting server on :7999")
	log.Fatal(http.ListenAndServe(":7999", nil))
}
