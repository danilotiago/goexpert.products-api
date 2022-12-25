package main

import (
	"log"
	"net/http"

	"github.com/backendengineerark/products-api/configs"
	_ "github.com/backendengineerark/products-api/docs"
	"github.com/backendengineerark/products-api/internal/entity"
	"github.com/backendengineerark/products-api/internal/infra/database"
	"github.com/backendengineerark/products-api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Danilo Tiago
// @contact.url    https://medium.com/@danilotiago
// @contact.email  danilotiago1996@gmail.com

// @license.name   Full Cycle License
// @license.url    http://www.fullcycle.com.br

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs := configs.LoadConfig(".")
	println(configs.DBDriver)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandlers(userDB)

	r := chi.NewRouter()
	// r.Use(middleware.Logger) // logger middleware (log all requests)
	r.Use(LogRequest)           // custom logger middleware (log all requests)
	r.Use(middleware.Recoverer) // if any errors has throws, this middleware not permit the server will be down

	r.Use(middleware.WithValue("jwt", configs.TokenAuth))             // inject key value in request context
	r.Use(middleware.WithValue("jwtExpiresIn", configs.JWTExpiresIn)) // inject key value in request context

	r.Route("/products", func(r chi.Router) { // group products routes
		r.Use(jwtauth.Verifier(configs.TokenAuth)) // middleware to find token in request
		r.Use(jwtauth.Authenticator)               // middleware to validate token

		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Receive request to path %s - %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
