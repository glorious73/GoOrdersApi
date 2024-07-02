package handler

import (
	"fmt"
	"net/http"
)

type Order struct {
}

func (o *Order) Create(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Create an order")
}

func (o *Order) List(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("List orders")
}

func (o *Order) GetById(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Get an order by Id")
}
func (o *Order) UpdateById(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Update an order by Id")
}
func (o *Order) DeleteById(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Delete an order by Id")
}
