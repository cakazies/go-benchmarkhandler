package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"
	ctr "github.com/local/go-benchmark/controllers"
)

func BenchmarkHanlderGin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, r := gin.CreateTestContext(resp)
		c.Request, _ = http.NewRequest("GET", "/pong", nil)
		r.GET("/pong", ctr.GinGetProfile)
		r.ServeHTTP(resp, c.Request)
	}
}

func BenchmarkHanlderMux(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("GET", "/ping", nil)
		r := mux.NewRouter()
		r.HandleFunc("/ping", ctr.MuxGetProfile)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)
	}
}
func BenchmarkHanlderEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/peng", nil)
		rec := httptest.NewRecorder()
		e.GET("/peng", ctr.EchoGetProfile)
		e.ServeHTTP(rec, req)
	}
}
