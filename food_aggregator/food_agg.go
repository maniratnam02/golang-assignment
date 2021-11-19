package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Vegetables struct {
	ProductId   string
	ProductName string
	Quantity    int
	Price       string
}
type Fruits struct {
	Id       string
	Name     string
	Quantity int
	Price    string
}
type Grains struct {
	ItemId   string
	ItemName string
	Quantity int
	Price    string
}

const Fruit_Url = "https://run.mocky.io/v3/c51441de-5c1a-4dc2-a44e-aab4f619926b"
const Grain_Url = "https://run.mocky.io/v3/e6c77e5c-aec9-403f-821b-e14114220148"
const Vegetable_Url = "https://run.mocky.io/v3/4ec58fbc-e9e5-4ace-9ff0-4e893ef9663c"

var grains []Grains
var veg []Vegetables
var fruit []Fruits

var ListOfItems []interface{}

func Getbyname(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" Buy item")
	w.Header().Set("Content-Type", "application/json")

	value := mux.Vars(r)

	response, err1 := http.Get(Vegetable_Url)
	if err1 != nil {
		panic(err1)
	}

	content, _ := ioutil.ReadAll(response.Body)

	err2 := json.Unmarshal(content, &veg)
	if err2 != nil {
		panic(err2)
	}

	for _, m := range veg {

		if m.ProductName == value["itemname"] {
			json.NewEncoder(w).Encode(m)
			return
		}

	}

	response1, er := http.Get(Fruit_Url)
	if er != nil {
		panic(er)
	}

	content1, _ := ioutil.ReadAll(response1.Body)

	err3 := json.Unmarshal(content1, &fruit)
	if err3 != nil {
		panic(err3)
	}

	response2, er := http.Get(Grain_Url)
	if er != nil {
		panic(er)
	}

	for _, a := range fruit {
		if a.Name == value["itemname"] {
			json.NewEncoder(w).Encode(a)
			return
		}
	}

	content2, _ := ioutil.ReadAll(response2.Body)

	err4 := json.Unmarshal(content2, &grains)
	if err4 != nil {
		panic(err4)
	}

	for _, p := range grains {
		if p.ItemName == value["itemname"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	defer response.Body.Close()
	defer response1.Body.Close()
	defer response2.Body.Close()

	json.NewEncoder(w).Encode("!!---- Item Not Found ----!!")
	return
}

func Getbynamequantity(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" Buy item")
	w.Header().Set("Content-Type", "application/json")

	value := mux.Vars(r)
	k, _ := strconv.Atoi(value["quantity"])
	response, err1 := http.Get(Vegetable_Url)
	if err1 != nil {
		panic(err1)
	}

	content, _ := ioutil.ReadAll(response.Body)

	err2 := json.Unmarshal(content, &veg)
	if err2 != nil {
		panic(err2)
	}

	for _, m := range veg {

		if m.Quantity >= k && m.ProductName == value["itemname"] {
			json.NewEncoder(w).Encode(m)
			return
		}

	}

	response1, er := http.Get(Fruit_Url)
	if er != nil {
		panic(er)
	}

	content1, _ := ioutil.ReadAll(response1.Body)

	err3 := json.Unmarshal(content1, &fruit)
	if err3 != nil {
		panic(err3)
	}

	response2, er := http.Get(Grain_Url)
	if er != nil {
		panic(er)
	}

	for _, a := range fruit {
		if a.Quantity >= k && a.Name == value["itemname"] {
			json.NewEncoder(w).Encode(a)
			return
		}
	}

	content2, _ := ioutil.ReadAll(response2.Body)

	err4 := json.Unmarshal(content2, &grains)
	if err4 != nil {
		panic(err4)
	}

	for _, p := range grains {
		if p.Quantity >= k && p.ItemName == value["itemname"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	defer response.Body.Close()
	defer response1.Body.Close()
	defer response2.Body.Close()
	json.NewEncoder(w).Encode("!!---- Item Not Found ----!!")
	return
}
func Getbynamequantityprice(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" getBuy name qty price")
	w.Header().Set("Content-Type", "application/json")

	value := mux.Vars(r)
	str := "$" + value["price"]
	k, _ := strconv.Atoi(value["quantity"])
	response, err1 := http.Get(Vegetable_Url)
	if err1 != nil {
		panic(err1)
	}

	content, _ := ioutil.ReadAll(response.Body)

	err2 := json.Unmarshal(content, &veg)
	if err2 != nil {
		panic(err2)
	}

	for _, m := range veg {

		if m.Quantity >= k && m.ProductName == value["itemname"] && m.Price == str {
			ListOfItems = append(ListOfItems, m)
			json.NewEncoder(w).Encode(m)
			return
		}

	}

	response1, er := http.Get(Fruit_Url)
	if er != nil {
		panic(er)
	}

	content1, _ := ioutil.ReadAll(response1.Body)

	err3 := json.Unmarshal(content1, &fruit)
	if err3 != nil {
		panic(err3)
	}

	response2, er := http.Get(Grain_Url)
	if er != nil {
		panic(er)
	}

	for _, a := range fruit {
		if a.Quantity >= k && a.Name == value["itemname"] && a.Price == str {
			ListOfItems = append(ListOfItems, a)
			json.NewEncoder(w).Encode(a)
			return
		}
	}

	content2, _ := ioutil.ReadAll(response2.Body)

	err4 := json.Unmarshal(content2, &grains)
	if err4 != nil {
		panic(err4)
	}

	for _, val3 := range grains {
		if val3.Quantity >= k && val3.ItemName == value["itemname"] && val3.Price == str {
			ListOfItems = append(ListOfItems, val3)
			json.NewEncoder(w).Encode(val3)
			return
		}
	}

	defer response.Body.Close()
	defer response1.Body.Close()
	defer response2.Body.Close()

	json.NewEncoder(w).Encode("!!---- Item Not Found ----!!")

}
func Listitems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ListOfItems)

}
func main() {
	fmt.Println("application running")
	c := mux.NewRouter()
	c.HandleFunc("/buy-item/{itemname}", Getbyname).Methods("GET")
	c.HandleFunc("/buy-item-qty/{itemname}/{quantity}", Getbynamequantity).Methods("GET")
	c.HandleFunc("/buy-item-qty-price/{itemname}/{quantity}/{price}", Getbynamequantityprice).Methods("GET")
	c.HandleFunc("/show-summery/", Listitems).Methods("GET")
	log.Fatal(http.ListenAndServe(":8500", c))
}
