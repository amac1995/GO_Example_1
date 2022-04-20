//Returning Data with a bool

package functions

import (
	"fmt"
)

func mainReturnData() {
	port := 3000
	isStarted := ReturnDataStartWebServer(port)
	fmt.Println(isStarted)

}

func ReturnDataStartWebServer(port int) bool {
	fmt.Println("Starting server...")
	//do important things
	fmt.Println("Server started on port", port)
	return true
}
