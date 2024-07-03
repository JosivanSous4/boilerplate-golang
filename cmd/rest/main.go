package main

import (
	"context"
	"encoding/json"
	"log"

	"boilerplate-go/internal/config"
	"boilerplate-go/internal/delivery/http"
	"boilerplate-go/internal/domain/model"
	"boilerplate-go/internal/domain/repository"
	"boilerplate-go/internal/domain/service"
	"boilerplate-go/internal/infrastructure/database"
	"boilerplate-go/internal/infrastructure/messaging"

	"github.com/gofiber/fiber/v2"
)

func main() {
    cfg := config.LoadConfig()

    // Configurar o banco de dados MySQL
    db, err := database.NewMySQLConnection(cfg.MySQLDSN)
    if err != nil {
        log.Fatalf("failed to connect to MySQL: %v", err)
    }

    // Configurar o MongoDB
    // ctx := context.Background()
    // mongoClient, err := database.NewMongoDBConnection(ctx, cfg.MongoDBURI)
    // if err != nil {
    //     log.Fatalf("failed to connect to MongoDB: %v", err)
    // }

    // Configurar RabbitMQ
    producer, err := messaging.NewSQSProducer(cfg.MessagingURI)
    if err != nil {
        log.Fatalf("failed to connect to RabbitMQ: %v", err)
    }

    // Repositórios de Produto
    // mongoProductRepo := repository.NewMongoProductRepository(mongoClient, "yourdbname", "products")
    mysqlProductRepo := repository.NewMySQLProductRepository(db)

    // Serviços de Produto
    productService := service.NewProductService(mysqlProductRepo, producer)

    // Configurar Fiber
    app := fiber.New(fiber.Config{
        ErrorHandler: http.ErrorHandler,
    })

    // Handlers
    productHandler := http.NewProductHandler(productService)
    productHandler.RegisterRoutes(app)

    go func() {
        consumer, err := messaging.NewSQSConsumer(cfg.MessagingURI)
        if err != nil {
            log.Fatalf("failed to start SQS consumer: %v", err)
        }

        msgs, err := consumer.Consume("product_queue")
        if err != nil {
            log.Fatalf("failed to consume SQS messages: %v", err)
        }

        for body := range msgs {
            var product model.Product
            if err := json.Unmarshal(body, &product); err != nil {
                log.Printf("Error decoding JSON: %s", err)
                continue
            }
            if err := mysqlProductRepo.CreateProduct(context.Background(), &product); err != nil {
                log.Printf("Error saving product: %s", err)
            } else {
                log.Printf("Product saved: %v", product)
            }
        }
    }()

    log.Fatal(app.Listen(":8080"))
}
