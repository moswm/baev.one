
// by Baev, 2022

package main

import (
	"net/http"
)

func start_crypto(w http.ResponseWriter, r *http.Request) {
	
	display(w, "crypto", nil)
	
}
