package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"calculator-grpc/pb"

	"github.com/gin-gonic/gin"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/version", getCalculatorVersion)

	listener, err := net.Listen("tcp", ":8111")
	if err != nil {
		log.Fatalf("Failed to listen to port 8111: %v", err)
	}

	m := cmux.New(listener)
	grpcListener := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpListener := m.Match(cmux.Any())

	g := new(errgroup.Group)
	g.Go(func() error { return grpcServe(grpcListener) })
	g.Go(func() error { return httpServe(httpListener, router) })
	g.Go(func() error { return m.Serve() })

	log.Println("Server started.")

	err = g.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
func grpcServe(l net.Listener) error {
	s := &pb.CalculatorServer{}
	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(grpcServer, s)
	reflection.Register(grpcServer)

	return grpcServer.Serve(l)
}

func httpServe(l net.Listener, router *gin.Engine) error {
	httpServer := &http.Server{
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return httpServer.Serve(l)
}

func getCalculatorVersion(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "1.0.0")
}
