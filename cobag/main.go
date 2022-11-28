package main

import (
	// "fmt"
	"fmt"
	"log"
	"net/http"

	// "github.com/gorilla/mux"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm" // untuk migrate database
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shopspring/decimal"
)

var db *gorm.DB
var err error


//  membuat tabel product
type Product struct {
	ID		int							`json:"id"`
	Code	string					`json:"code"`
	Name	string					`json:"name"`
	Price	decimal.Decimal	`json:"price" sql:"type:decimal(16,2)"`
}

type Result struct {
	Code		int					`json:"code"`
	Data		interface{}	`json:"data"`
	Message	string			`json:"message"`
}

func main() {
    // Please define your username and password for MySQL.
    db, err = gorm.Open("mysql", "root:@/gogo?charset=utf8&parseTime=True")
    // NOTE: See we’re using = to assign the global var
    // instead of := which would assign it only in this function

    if err!=nil{
    log.Println("Koneksi Gagal! error: ", err)
    }else{ 
    log.Println("Koneksi Berhasil!")
    }

		db.AutoMigrate(&Product{})

		handleRequests()
}

// membuat routing
func handleRequests() {
	log.Println("Start the development server http:localhost:9999")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/api/products", createProduct).Methods("POST")

	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Selamat Datang!")
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Products")
}