package main

import (
	"html/template"
	"log"
	"net/http"
)

//var db *gorm.DB

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", Hello)

	log.Println("app running on http://localhost:8080")
	err := http.ListenAndServe(":8080", server)
	if err != nil {
		log.Println(err)
	}
}

func Hello(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		t := template.Must(template.ParseFiles("./templates/home.html"))
		err := t.Execute(writer, map[string]interface{}{})
		if err != nil {
			log.Println("Error: ", err.Error())
			return
		}
		return
	}
	return
}

func reverse(value string) string {
	result := ""
	for i := len(value); i > 0; i-- {
		result = result + string(value[i-1])
	}
	return result
}
