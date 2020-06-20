package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/MrWormHole/todo-api-grpc/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewMathServiceClient(connection)

	g := gin.Default()
	g.GET("/add/:a/:b", func(context *gin.Context) {
		a, err := strconv.ParseUint(context.Param("a"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater A"})
		}

		b, err := strconv.ParseUint(context.Param("b"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater B"})
		}

		request := &proto.Request{A: int64(a), B: int64(b)}
		response, err := client.Add(context, request)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
		}
	})
	g.GET("/multiply/:a/:b", func(context *gin.Context) {
		a, err := strconv.ParseUint(context.Param("a"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater A"})
		}

		b, err := strconv.ParseUint(context.Param("b"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater B"})
		}

		request := &proto.Request{A: int64(a), B: int64(b)}
		response, err := client.Multiply(context, request)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
		}
	})
	g.GET("/divide/:a/:b", func(context *gin.Context) {
		a, err := strconv.ParseUint(context.Param("a"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater A"})
		}

		b, err := strconv.ParseUint(context.Param("b"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater B"})
		}

		request := &proto.Request{A: int64(a), B: int64(b)}
		response, err := client.Divide(context, request)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
		}
	})
	g.GET("/substract/:a/:b", func(context *gin.Context) {
		a, err := strconv.ParseUint(context.Param("a"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater A"})
		}

		b, err := strconv.ParseUint(context.Param("b"), 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater B"})
		}

		request := &proto.Request{A: int64(a), B: int64(b)}
		response, err := client.Substract(context, request)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
		}
	})

	err = g.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
