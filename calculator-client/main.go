package main

import (
	"log"
	"net/http"
	"time"

	"calculator-grpc/pb"

	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
)

var client pb.CalculatorServiceClient

func main() {
	conn, err := grpc.Dial("localhost:8111", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client = pb.NewCalculatorServiceClient(conn)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/add", Add)
	router.GET("/substract", Substract)
	router.GET("/multiply", Multiply)

	port := "8112"
	serviceServer := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("Grpc client is running on port: " + port)
	log.Fatal(serviceServer.ListenAndServe())

}

func Add(ctx *gin.Context) {
	in := &pb.CalculatorPayload{X: 20, Y: 3}
	msg, err := client.Add(ctx, in)
	if err != nil {
		log.Println("Error: ", err)
		ctx.IndentedJSON(http.StatusBadRequest, err)
	}

	ctx.IndentedJSON(http.StatusOK, msg)
}
func Substract(ctx *gin.Context) {
	in := &pb.CalculatorPayload{X: 20, Y: 3}
	msg, err := client.Substract(ctx, in)
	if err != nil {
		log.Println("Error: ", err)
		ctx.IndentedJSON(http.StatusBadRequest, err)
	}

	ctx.IndentedJSON(http.StatusOK, msg)
}
func Multiply(ctx *gin.Context) {
	in := &pb.CalculatorPayload{X: 20, Y: 3}
	msg, err := client.Multiply(ctx, in)
	if err != nil {
		log.Println("Error: ", err)
		ctx.IndentedJSON(http.StatusBadRequest, err)
	}

	ctx.IndentedJSON(http.StatusOK, msg)
}
