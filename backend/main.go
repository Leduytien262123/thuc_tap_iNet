package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	Label    string
	Key      string
	Link     string
	Metadata datatypes.JSON
}

func main() {

	dsn := "letien:letien@tcp(127.0.0.1:3306)/backend?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
	}

	db.AutoMigrate(&Page{})

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/page", func(c *fiber.Ctx) error {

		var pages []Page
		if err := db.Select("Label", "Metadata").Find(&pages).Error; err != nil {
			return err
		}

		return c.JSON(pages)
	})

	app.Get("/page/get", func(c *fiber.Ctx) error {
		var page []Page
		db.Find(&page)
		return c.JSON(page)
	})

	app.Post("/page/create", func(c *fiber.Ctx) error {

		page := new(Page)

		if err := c.BodyParser(&page); err != nil {
			return err
		}

		err := db.Create(&page).Error
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"page":     page,
			"thongbao": "Đã thêm thành công",
		})
	})

	app.Patch("/page", func(c *fiber.Ctx) error {

		requestData := new(Page)
		if err := c.BodyParser(requestData); err != nil {
			return err
		}
		existingPage := &Page{}
		if err := db.Where("label = ?", requestData.Label).First(&existingPage).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Không tìm thấy bản ghi"})
		}
		existingPage.Key = requestData.Key
		existingPage.Link = requestData.Link
		existingPage.Metadata = requestData.Metadata
		if err := db.Save(&existingPage).Error; err != nil {
			return err
		}
		return c.JSON(fiber.Map{"page": existingPage})
	})
	app.Listen(":8000")
}
