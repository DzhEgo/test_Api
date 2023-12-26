package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Book struct {
	gorm.Model
	Title string `json:"title"`
	Autor string `json:"autor"`
	Pages int    `json:"pages"`
}

type Delivery struct {
	gorm.Model
	Type    string `json:"type"`
	Address string `json:"address"`
}

type Vechicle struct {
	gorm.Model
	Brend       string `json:"brend"`
	Model_vech  string `json:"model_vech"`
	Dimensioins string `json:"dimensioins"`
}

type Task struct {
	gorm.Model
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var db *gorm.DB

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/user", UserPostHandler).Methods("POST")
	r.HandleFunc("/user", UserGetHandler).Methods("GET")
	r.HandleFunc("/user/{id}", UserFindHandler).Methods("GET")
	r.HandleFunc("/user/{id}", UserDeleteHandler).Methods("DELETE")
	r.HandleFunc("/book", BookPostHandler).Methods("POST")
	r.HandleFunc("/book", BookGetHandler).Methods("GET")
	r.HandleFunc("/book/{id}", BookFindHandler).Methods("GET")
	r.HandleFunc("/delivery", DeliveryPostHandler).Methods("POST")
	r.HandleFunc("/delivery", DeliveryGetHandler).Methods("GET")
	r.HandleFunc("/delivery/{id}", DeliveryFindHandler).Methods("GET")
	r.HandleFunc("/vechicle", VechiclePostHandler).Methods("POST")
	r.HandleFunc("/vechicle", VechicleGetHandler).Methods("GET")
	r.HandleFunc("/vechicle/{id}", VechicleFindHandler).Methods("GET")
	r.HandleFunc("/vechicle/{id}", VechiclePutHandler).Methods("PUT")
	r.HandleFunc("/task", TaskPostHandler).Methods("POST")
	r.HandleFunc("/task", TaskGetHandler).Methods("GET")
	r.HandleFunc("/task/{id}", TaskFindHandler).Methods("GET")
	r.HandleFunc("/task/{id}", TaskDeleteHandler).Methods("DELETE")

	err := http.ListenAndServe(":3030", r)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	var err error

	dsn := "user=postgres password=Lax212212 dbname=test_api sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{}, &Book{}, &Delivery{}, &Vechicle{}, Task{})

	fmt.Println("Запуск сервера...")
	StartServer()
}

func UserPostHandler(w http.ResponseWriter, r *http.Request) {
	var users User

	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	res := db.Create(&users)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Данные записаны!")
}

func UserGetHandler(w http.ResponseWriter, r *http.Request) {
	var users []User

	res := db.Find(&users)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func UserFindHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var users User

	if id == "" {
		http.Error(w, "Поле ID должно быть заполнено!", http.StatusBadRequest)
		return
	}

	res := db.First(&users, "id = ?", id)
	if res.Error != nil {
		http.Error(w, "Данные либо не найдены, либо ошибка сервера!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var users User

	err := db.First(&users, id)
	if err.Error != nil {
		http.Error(w, "Данные не найдены!", http.StatusNotFound)
		return
	}

	err = db.Where("id = ?", id).Delete(&users)
	if err.Error != nil {
		http.Error(w, "Ошибка при удалении!", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Данные удалены!")
}

func BookPostHandler(w http.ResponseWriter, r *http.Request) {
	var books Book

	err := json.NewDecoder(r.Body).Decode(&books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	res := db.Create(&books)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Данные записаны!")
}

func BookGetHandler(w http.ResponseWriter, r *http.Request) {
	var books []Book

	res := db.Find(&books)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func BookFindHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var books Book

	if id == "" {
		http.Error(w, "Поле ID должно быть заполнено!", http.StatusBadRequest)
		return
	}

	res := db.First(&books, "id = ?", id)
	if res.Error != nil {
		http.Error(w, "Данные либо не найдены, либо ошибка сервера!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func DeliveryPostHandler(w http.ResponseWriter, r *http.Request) {
	var deliver Delivery

	err := json.NewDecoder(r.Body).Decode(&deliver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	res := db.Create(&deliver)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Данные записаны!")
}

func DeliveryGetHandler(w http.ResponseWriter, r *http.Request) {
	var deliver []Delivery

	res := db.Find(&deliver)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deliver)
}

func DeliveryFindHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var deliver Delivery

	if id == "" {
		http.Error(w, "Поле ID должно быть заполнено!", http.StatusBadRequest)
		return
	}

	res := db.First(&deliver, "id = ?", id)
	if res.Error != nil {
		http.Error(w, "Данные либо не найдены, либо ошибка сервера!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deliver)
}

func VechiclePostHandler(w http.ResponseWriter, r *http.Request) {
	var vech Vechicle

	err := json.NewDecoder(r.Body).Decode(&vech)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	res := db.Create(&vech)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Данные записаны!")
}

func VechicleGetHandler(w http.ResponseWriter, r *http.Request) {
	var vech []Vechicle

	res := db.Find(&vech)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vech)
}

func VechicleFindHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var vech Vechicle

	if id == "" {
		http.Error(w, "Поле ID должно быть заполнено!", http.StatusBadRequest)
		return
	}

	res := db.First(&vech, "id = ?", id)
	if res.Error != nil {
		http.Error(w, "Данные либо не найдены, либо ошибка сервера!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vech)
}

func VechiclePutHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var vech Vechicle

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	search := db.First(&vech, id)
	if search.Error != nil {
		http.Error(w, "Данные не найдены!", http.StatusNotFound)
		return
	}

	res := json.Unmarshal(body, &vech)
	if res != nil {
		http.Error(w, "Ошибка при декодировании JSON!", http.StatusBadRequest)
		return
	}

	upd := db.Where("id = ?", id).Save(&vech)
	if upd.Error != nil {
		http.Error(w, "Ошибка при обновлении!", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Данные обновлены!")
}

func TaskPostHandler(w http.ResponseWriter, r *http.Request) {
	var tasks Task

	err := json.NewDecoder(r.Body).Decode(&tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	res := db.Create(&tasks)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Данные записаны!")
}

func TaskGetHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task

	res := db.Find(&tasks)
	if res.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func TaskFindHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var tasks Task

	if id == "" {
		http.Error(w, "Поле ID должно быть заполнено!", http.StatusBadRequest)
		return
	}

	res := db.First(&tasks, "id = ?", id)
	if res.Error != nil {
		http.Error(w, "Данные либо не найдены, либо ошибка сервера!", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

func TaskDeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var tasks Task

	err := db.First(&tasks, id)
	if err.Error != nil {
		http.Error(w, "Данные не найдены!", http.StatusNotFound)
		return
	}

	err = db.Where("id = ?", id).Delete(&tasks)
	if err.Error != nil {
		http.Error(w, "Ошибка при удалении!", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Данные удалены!")
}
