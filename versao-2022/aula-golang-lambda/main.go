package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
)

var fiberLambda *fiberadapter.FiberLambda

// Handlers
func healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK     APIGatewayV2HTTPRequest")
}

func returnUsers(c *fiber.Ctx) error {
	users := []string{"User1", "User2", "User3"}
	return c.JSON(users)
}

// Main function
func main() {
	app := fiber.New()

	// Define routes
	app.Get("/", healthCheck)
	app.Get("/users", returnUsers)

	// Check if running in Lambda
	if isLambda() {
		log.Println("Running in Lambda environment")
		fiberLambda = fiberadapter.New(app)
		lambda.Start(handler)
	} else {
		log.Println("Running locally")
		log.Fatal(app.Listen(":3000"))
	}
}

// Lambda handler
func handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	log.Printf("Path: %s", request.RawPath)
	return fiberLambda.ProxyWithContextV2(ctx, request)
}

// Check if running in Lambda environment
func isLambda() bool {
	value, exists := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME")
	if exists {
		log.Printf("AWS_LAMBDA_FUNCTION_NAME: %s", value)
	}
	return exists
}
