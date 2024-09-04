package main

import (
	"fmt"

	"github.com/Askanin/personalidade-go.git/models"
	"github.com/Askanin/personalidade-go.git/routes"
)





func main() {
	models.Personalidades = []models.Personalidade {
		{Id: 1, Nome: "nome1", Historia: "historia1"},
		{Id: 2, Nome: "nome2", Historia: "historia2"},
	}
	
	fmt.Println("Iniciando o servidro Rest com Go")
	routes.HandleRequest()
}
