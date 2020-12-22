package root

import (
	"fmt"
	"net/http"
)

func Root (res http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(res, "I'm in")
}
