//Parameters

package functions

import (
	"fmt"
)

func main_Parameters() {
	// Duas maneiras diferentes de declarar um array
	port := 3000
	parametersStartWebServer(port, 2)
}

func parametersStartWebServer(port int, numberOfRetries int) {
	fmt.Println("Starting server...")
	//do important things
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries:", numberOfRetries)
}
