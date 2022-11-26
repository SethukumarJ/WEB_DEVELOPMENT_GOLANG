package main

import (
	"log"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Auther		string		`json:"auther"`
	Title		string		`json:"title"`
	Publisher	string		`json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}
	err := context.BodyParser(&book)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message":"request body is not valid"},
		)
		return err
	}
	err = r.DB.Create(book).Error

	if err!= nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message":"could not create book"},
		)
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"book created"})
		return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Book{}
	err := r.DB.Find(bookModels).Error	
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message":"could not get books"},
		)
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"books fetched",
		"data":bookModels, 
	})

	return nil

}

func (r *Repository) SetupRoutes(app *fiber.App) {

	api := app.Group("/api")
	api.Post("/create_books",r.CreateBook)
	api.Delete("delete_book/:id",r.DeleteBook)
	api.Get("/get_books/:id",r.GetBookByID)
	api.Get("/books",r.GetBooks)

}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := storage.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("could not import database")
	}
	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
