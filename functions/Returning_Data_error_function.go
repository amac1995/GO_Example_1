//A função startWebServer nao pode ser booleana pq nao fornece nenhuma informação sobre o que deu erro

//Returning Data with a error function

package functions

import (
	"fmt"
)

func mainDataError() {
	port := 3000
	//Capta uma mensagem de erro e imprime-a
	//Variaveis de gravação -> Permite despejar qualquer tipo de dados (_)
	_, err := dataErrorStartWebServer(port)
	fmt.Println(err)

}

func dataErrorStartWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	//do important things
	fmt.Println("Server started on port", port)
	return port, nil
}
