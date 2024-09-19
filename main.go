package main

import (
  "fmt"
  "github.com/gofiber/fiber/v2"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
  "log"
  "os"
  "time"
)

const (
  host     = "localhost"  // or the Docker service name if running in another container
  port     = 5432         // default PostgreSQL port
  user     = "myuser"     // as defined in docker-compose.yml
  password = "mypassword" // as defined in docker-compose.yml
  dbname   = "mydatabase" // as defined in docker-compose.yml
)

func main() {
  // Configure your PostgreSQL database details here
  dsn := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

  // New logger for detailed SQL logging
  newLogger := logger.New(
    log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
    logger.Config{
      SlowThreshold: time.Second, // Slow SQL threshold
      LogLevel:      logger.Info, // Log level
      Colorful:      true,        // Enable color
    },
  )

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: newLogger,
  })

  if err != nil {
    panic("failed to connect to database")
  }

  // Migrate the schema
  db.AutoMigrate(&Fruit{})

  // Setup Fiber
  app := fiber.New()

  // CRUD routes
  app.Get("/fruits", func(c *fiber.Ctx) error {
    return getFruits(db, c)
  })
  app.Get("/fruits/:id", func(c *fiber.Ctx) error {
    return getFruit(db, c)
  })
  app.Post("/fruits", func(c *fiber.Ctx) error {
    return createFruit(db, c)
  })
  app.Put("/fruits/:id", func(c *fiber.Ctx) error {
    return updateFruit(db, c)
  })
  app.Delete("/fruits/:id", func(c *fiber.Ctx) error {
    return deleteFruit(db, c)
  })

  // Start server
  log.Fatal(app.Listen(":8000"))
}