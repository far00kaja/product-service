package server

import (
	"net/http"
	"os"
	"sribuu/products/lib"
	"time"
)

var readTimeout, _ = time.ParseDuration(os.Getenv("READ_TIMEOUT"))
var writeTimeout, _ = time.ParseDuration(os.Getenv("WRITE_TIMEOUT"))
var iddleTimeout, _ = time.ParseDuration(os.Getenv("IDDLE_TIMEOUT"))

var HttpServer = &http.Server{
	Addr:              lib.GetEnv("PORT", "0.0.0.0:8129"),
	Handler:           nil,
	TLSConfig:         nil,
	ReadTimeout:       readTimeout * time.Second,
	ReadHeaderTimeout: 0,
	WriteTimeout:      writeTimeout * time.Second,
	IdleTimeout:       iddleTimeout * time.Second,
	MaxHeaderBytes:    0,
	TLSNextProto:      nil,
	ConnState:         nil,
	ErrorLog:          nil,
	BaseContext:       nil,
	ConnContext:       nil,
}

// func EchoServer(*echo.Echo, error) {
// 	e := echo.New()

// 	vc := validator.New()
// 	e.

// }
