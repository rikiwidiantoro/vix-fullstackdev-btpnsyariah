package main

import (
	// "fmt"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	myRouter.HandleFunc("/api/products", getProducts).Methods("GET")
	myRouter.HandleFunc("/api/products/{id}", getProduct).Methods("GET")
	myRouter.HandleFunc("/api/products/{id}", updateProduct).Methods("PUT")
	myRouter.HandleFunc("/api/products/{id}", deleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Selamat Datang!")
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Create Products")

	// bikin variabel untuk menangkap payload
	payloads, _ := ioutil.ReadAll(r.Body)

	// bikin variabel struct product
	var product Product

	// casting ke stuck product
	json.Unmarshal(payloads, &product)

	// data yg sudah ditangkap lalu di inputkan ke tabel products
	db.Create(&product)

	// result
	res:= Result{Code: 200, Data: product, Message: "Berhasil menambahkan produk"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// membuat respon
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}


func getProducts(w http.ResponseWriter, r *http.Request) {
	// bikin variabel
	products := []Product{}

	// mengambil data
	db.Find(&products)

	// struct result
	res := Result{Code: 200, Data: products, Message: "Berhasil menampilkan produk"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// membuat respon
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}


// single product
func getProduct(w http.ResponseWriter, r *http.Request) {
	// mengambil single product
	vars := mux.Vars(r)
	productID := vars["id"]

	// bikin variabel
	var product Product
	db.First(&product, productID)

	// struct result
	res := Result{Code: 200, Data: product, Message: "Berhasil menampilkan produk"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// membuat respon
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}


// update
func updateProduct(w http.ResponseWriter, r *http.Request) {
	// mengambil single product
	vars := mux.Vars(r)
	productID := vars["id"]

	// bikin variabel untuk menangkap payload
	payloads, _ := ioutil.ReadAll(r.Body)

	// bikin variabel struct product
	var productUpdates Product

	// casting ke stuck product
	json.Unmarshal(payloads, &productUpdates)

	// mengambil eksisting product
	var product Product
	db.First(&product, productID)

	// update ke database
	db.Model(&product).Update(productUpdates)

	// result
	res:= Result{Code: 200, Data: product, Message: "Berhasil mengubah produk"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// membuat respon
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}


// delete
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// mengambil single product
	vars := mux.Vars(r)
	productID := vars["id"]

	var product Product

	db.First(&product, productID)
	db.Delete(&product)

	// struct result
	res := Result{Code: 200, Message: "Berhasil menghapus produk"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// membuat respon
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}