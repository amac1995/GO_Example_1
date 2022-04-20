package programscrud

type HTTPRequest struct {
	Method string
}

func MainSwitch() {
	r := HTTPRequest{Method: "GET"}

	switch r.Method {
	case "GET":
		println("Get Request")
	case "DELETE":
		println("Delete Request")
	case "POST":
		println("Post Request")
	case "PUT":
		println("Put Request")
	default:
		println("Unhandled method")
	}

}
