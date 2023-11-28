package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/renatocantarino/go/APIS/configs"
	_ "github.com/renatocantarino/go/APIS/docs"
	"github.com/renatocantarino/go/APIS/internal/entity"
	"github.com/renatocantarino/go/APIS/internal/infra/database"
	"github.com/renatocantarino/go/APIS/internal/infra/server/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   RenatoCantarino
// @contact.url    http://www.xxxxx.com.br
// @contact.email  xxxxx@xxxx.com.br

// @license.name   xxx xxx License
// @license.url    http://www.xxxxx.com.br

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	confs, err := configs.Load(".")
	if err != nil {
		panic(err)
	}

	db, err := openAndApllyAutoMigrateDB()
	runnerServer(db, confs)
}

func openAndApllyAutoMigrateDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.Product{}, &entity.User{})
	if err != nil {
		panic(err)
	}
	return db, err
}

func runnerServer(db *gorm.DB, cnfs *configs.Config) {

	produtoDb := database.NewProduct(db)
	userDb := database.NewUser(db)

	productSender := handlers.NewProductHandler(produtoDb)
	userSender := handlers.NewUserHandler(userDb, cnfs.JwtTokenAuth, cnfs.JwtExpiresIn)

	routers := chi.NewRouter()
	routers.Use(middleware.Logger)
	routers.Use(middleware.Recoverer)
	//routers.Use(LogRequestMiddleware)

	routers.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cnfs.JwtTokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productSender.CreateProduct)
		r.Get("/", productSender.GetProducts)
		r.Get("/{id}", productSender.GetProduct)
		r.Put("/{id}", productSender.UpdateProduct)
		r.Delete("/{id}", productSender.DeleteProduct)
	})

	routers.Post("/user", userSender.CreateUser)
	routers.Post("/user/token", userSender.GetJwt)

	routers.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8011/docs/doc.json")))
	http.ListenAndServe(":8011", routers)
}

func LogRequestMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})

}

/*

routers: atua principalmente na camada de web
echo:  https://echo.labstack.com/
fiber: https://gofiber.io/ like node express
gin:   https://github.com/gin-gonic/gin


framework: é um laravel, rails da vida
tem todas as libs ja existentes
precisa seguir o pattern do frameworks

poucos pessoas usam esses tipos de frameworks



bufalo: https://gobuffalo.io/pt/
IRIS: https://www.iris-go.com/docs/#/


*/
