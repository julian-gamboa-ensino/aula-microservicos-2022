package main

import (
    "context"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/gin-gonic/gin"
    "net/http"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Olá do AWS Lambda com Gin!"})
    })

    return events.APIGatewayProxyResponse{
        StatusCode: http.StatusOK,
        Body:       "Olá do AWS Lambda com Gin!",
    }, nil
}

func main() {
    lambda.Start(handleRequest)
}
