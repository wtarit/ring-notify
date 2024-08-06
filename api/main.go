package main

import (
	"api/handler"
)

func main() {
	// http.HandleFunc("/*", Handler)
	// log.Fatal(http.ListenAndServe(":1323", nil))
	e := handler.InitEcho()
	e.Logger.Fatal(e.Start(":1323"))
}
