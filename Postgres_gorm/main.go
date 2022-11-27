package main

import (
	"book/models"
	"book/storage"
	"log"
	"net/http"
	"os"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Auther    string `json:"auther"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}
	err := context.BodyParser(&book)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request body is not valid"},
		)
		return err
	}
	err = r.DB.Create(book).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"},
		)
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book created"})
	return nil
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	book := &models.Book{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id is required",
		})
		return nil

	}
	err := r.DB.Delete(book, id)
	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete book",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book deleted",
	})
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Book{}
	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"},
		)
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books fetched",
		"data":    bookModels,
	})

	return nil

}

func (r *Repository) GetBookByID(context *fiber.Ctx) error {
	id := context.Params("id")
	bookmodel := &models.Book{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id is required",
		})
		return nil
	}
	fmt.Println("the id is, ", id)

	err := r.DB.Where("id =?", id).Find(bookmodel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not get book",
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book fetched",
		"data":    bookmodel,
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {

	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)

}

func main() {
	fmt.Println("hi")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	fmt.Println(config)

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not import database")
	}

	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("could not migrate books")
	}
	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
