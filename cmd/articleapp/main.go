package main

import (
	"fmt"
	article "github.com/cymon1997/go-backend/module/article/model"
	"log"

	"encoding/json"
	"net/http"

	"github.com/cymon1997/go-backend/internal/entity"
	"github.com/gorilla/mux"
)

var (
	articleFactory article.Factory
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/article")
	router.HandleFunc("/", GetArticle).Methods(http.MethodGet)
	fmt.Println("Listen at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func init() {
	articleFactory = article.NewFactory()
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	var response entity.BaseResponse
	response = entity.BaseResponse{
		Status:  http.StatusOK,
		Message: "ok",
	}
	result, err := articleFactory.NewGetByIDModel().Call()
	response.Payload = result
	if err != nil {
		response.Status = http.StatusInternalServerError
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Printf("Error while marshall %s\n", err.Error())
	}
}
