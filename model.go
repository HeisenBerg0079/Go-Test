package main

import (
  "github.com/gofiber/fiber/v2"
  "gorm.io/gorm"
)

type Fruit struct {
  gorm.Model
  Name        string `json:"name"`
}

// getFruits retrieves all fruits
func getFruits(db *gorm.DB, c *fiber.Ctx) error {
  var fruits []Fruit
  db.Find(&fruits)
  return c.JSON(fruits)
}

// getFruit retrieves a fruit by id
func getFruit(db *gorm.DB, c *fiber.Ctx) error {
  id := c.Params("id")
  var fruit Fruit
  db.First(&fruit, id)
  return c.JSON(fruit)
}

// createFruit creates a new fruit
func createFruit(db *gorm.DB, c *fiber.Ctx) error {
  fruit := new(Fruit)
  if err := c.BodyParser(fruit); err != nil {
    return err
  }
  db.Create(&fruit)
  return c.JSON(fruit)
}

// updateFruit updates a fruit by id
func updateFruit(db *gorm.DB, c *fiber.Ctx) error {
  id := c.Params("id")
  fruit := new(Fruit)
  db.First(&fruit, id)
  if err := c.BodyParser(fruit); err != nil {
    return err
  }
  db.Save(&fruit)
  return c.JSON(fruit)
}

// deleteFruit deletes a fruit by id
func deleteFruit(db *gorm.DB, c *fiber.Ctx) error {
  id := c.Params("id")
  db.Delete(&Fruit{}, id)
  return c.SendString("Fruit successfully deleted")
}