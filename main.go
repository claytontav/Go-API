package main

import(
	"io/ioutil"
	"encoding/json"
	"net/http"
	"fmt"
)

type User struct{
	Id int `json:"identificador"` //Mudar nome da estrutura
	Name string
}

var dataUser []User = []User{
	User{
		Id: 1,
		Name: "Clayton",
	},
	User{
		Id: 2,
		Name: "Samanta",
	},
}


//GET
func routesGet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	encode := json.NewEncoder(w)

	encode.Encode(dataUser)

}


//POST
func routesPost(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")

	//Verifica erros
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		fmt.Println("Erro ao cadastrar")
	}

	//Registra um novo dado
	var newUser User
	json.Unmarshal(body, &newUser)
	newUser.Id = len(dataUser) + 1 //Soma mais um valor do id
	dataUser = append(dataUser, newUser)
	encode := json.NewEncoder(w)
	encode.Encode(newUser)
}

func routesUser(w http.ResponseWriter, r *http.Request){

	if r.Method == "GET"{

		routesGet(w, r)

	}else if r.Method == "POST"{
		routesPost(w, r)
	}
}

func routes(){

	http.HandleFunc("/users", routesUser)

}

func config(){	

	routes()

	fmt.Println("Servidor esta rodando na porta 8080")

	http.ListenAndServe(":8080", nil)

}

func main(){

	config()

}