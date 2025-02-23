package handler

import (
	"fmt"
	"net/http"
)

type Order struct {
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Create an order"))
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("List orders"))
}

func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Get By Id"))
}

func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Update By Id"))
}

func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Delete By Id"))
}
