package main

import (
	"api/handler"
)

func main() {
	// http.HandleFunc("/*", Handler)
	// log.Fatal(http.ListenAndServe(":1323", nil))
	e := handler.Init()
	e.Logger.Fatal(e.Start(":1323"))
}
