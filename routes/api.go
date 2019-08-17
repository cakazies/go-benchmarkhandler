package routes

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	ctr "github.com/local/go-benchmark/controllers"
)

func RunMux() {
	r := mux.NewRouter()
	r.HandleFunc("/pong", ctr.MuxGetProfile)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func RunGin() {
	r := gin.Default()
	r.GET("/ping", ctr.GinGetProfile)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func RunEcho() {
	e := echo.New()
	e.GET("/peng", ctr.EchoGetProfile)
	e.Logger.Fatal(e.Start(":8080"))
}
