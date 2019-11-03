package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"secondAssignment"
)




func main() {

	secondAssignment.DBp.Init()
	secondAssignment.DBu.Init()
	secondAssignment.DBl.Init()
	secondAssignment.ST.Init()



	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

																						// Runs functions based on url typed
	http.HandleFunc("/repocheck/v1/commits/", secondAssignment.HandlerCommits)
	http.HandleFunc("/repocheck/v1/issues/", secondAssignment.HandlerIssues)
//	http.HandleFunc("/conservation/v1/diag/", firstAssignment.HandlerDiag)
	http.HandleFunc("/repocheck/v1/status/", secondAssignment.HandlerDiag)
	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
