package main

import (
	"cinema/component/appctx"
	"cinema/docs"
	"cinema/middleware"
	ginauditorium "cinema/module/auditorium/transport/gin"
	gincinema "cinema/module/cinema/transport/gin"
	gincompany "cinema/module/company/transport/gin"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	// Read the content of the db_env file
	content, err := ioutil.ReadFile("local_db_env")
	if err != nil {
		log.Fatalln(err)
	}
	// Replace newline characters with spaces
	dsn := strings.ReplaceAll(string(content), "\n", " ")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	db.Debug()

	appCtx := appctx.NewAppContext(db)
	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	docs.SwaggerInfo.BasePath = "/v1"
	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		cinemas := v1.Group("/cinemas")
		//GET /v1/cinemas
		cinemas.GET("", gincinema.ListCinema(appCtx))

		//GET /v1/cinemas/name/:name
		cinemas.GET("/name/:name", gincinema.GetCinemaWithName(appCtx))

		//GET /v1/cinemas/:id
		cinemas.GET("/:id", gincinema.GetCinemaWithID(appCtx))

		//POST /v1/cinemas
		cinemas.POST("", gincinema.CreateCinema(appCtx))

		//GET /v1/cinemas/:id/auditoriums
		cinemas.GET("/:id/auditoriums", ginauditorium.ListAuditoriumWithCinemaID(appCtx))

		//GET /v1/cinemas/name/:name/auditoriums
		cinemas.GET("/name/:name/auditoriums", ginauditorium.ListAuditoriumWithCinemaName(appCtx))
	}

	{
		auditoriums := v1.Group("/auditoriums")
		//POST /v1/auditoriums
		auditoriums.POST("", ginauditorium.CreateAuditorium(appCtx))
	}

	{
		companies := v1.Group("/companies")
		//GET /v1/companies
		companies.GET("", gincompany.ListCompany(appCtx))
		//GET /v1/companies/:id
		companies.GET("/:id", gincompany.GetCompanyWithID(appCtx))
		//POST /v1/companies
		companies.POST("", gincompany.CreateCompany(appCtx))
	}
	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
