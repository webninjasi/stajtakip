package routes

import "net/http"

// Web klasörünü sunar
var Index http.Handler = http.StripPrefix("/", http.FileServer(http.Dir("web")))
