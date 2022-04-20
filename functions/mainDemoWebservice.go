package functions

import (
	"net/http"

	"github.com/godarksystem/functions/controllers"
)

func MainDemoWebservice() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
