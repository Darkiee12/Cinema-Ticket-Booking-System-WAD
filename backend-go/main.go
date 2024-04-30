package main

import (
	"cinema/component/appctx"
	"cinema/docs"
	"cinema/middleware"
	ginauditorium "cinema/module/auditorium/transport/gin"
	gincinema "cinema/module/cinema/transport/gin"
	gincompany "cinema/module/company/transport/gin"
	ginmovie "cinema/module/movie/transport/gin"
	ginshow "cinema/module/show/transport/gin"
	ginticket "cinema/module/ticket/transport/gin"
	"cinema/module/user/transport/ginuser"
	"cinema/module/user/userstore"
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

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
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
	db = db.Debug()

	key := "my_secret"
	appCtx := appctx.NewAppContext(db, key)

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))

	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	userStore := userstore.NewSQLStore(appCtx.GetMainDBConnection())
	{
		cinemas := v1.Group("/cinemas")
		//GET /v1/cinemas
		cinemas.GET("", gincinema.ListCinema(appCtx))

		//GET /v1/cinemas/name/:name
		cinemas.GET("/name/:name", gincinema.GetCinemaWithName(appCtx))

		//GET /v1/cinemas/:id
		cinemas.GET("/:id", gincinema.GetCinemaWithID(appCtx))

		//POST /v1/cinemas
		cinemas.POST("",
			middleware.RequireAuth(appCtx, userStore),
			middleware.CheckRole(appCtx, "admin", "cinema_owner"),
			gincinema.CreateCinema(appCtx))

		//GET /v1/cinemas/:id/auditoriums
		cinemas.GET("/:id/auditoriums", ginauditorium.ListAuditoriumWithCinemaID(appCtx))

		//GET /v1/cinemas/name/:name/auditoriums
		cinemas.GET("/name/:name/auditoriums", ginauditorium.ListAuditoriumWithCinemaName(appCtx))
	}

	{
		auditoriums := v1.Group("/auditoriums")
		//POST /v1/auditoriums
		auditoriums.POST("",
			middleware.RequireAuth(appCtx, userStore),
			middleware.CheckRole(appCtx, "admin", "cinema_owner"),
			ginauditorium.CreateAuditorium(appCtx))
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

	{
		movies := v1.Group("/movies")
		//GET /v1/movies
		movies.GET("", ginmovie.ListMovie(appCtx))
		//GET /v1/movies/:imdb_id
		movies.GET("/:imdb_id", ginmovie.GetMovieWithID(appCtx))
		//POST /v1/movies
		movies.POST("",
			middleware.RequireAuth(appCtx, userStore),
			middleware.CheckRole(appCtx, "admin"),
			ginmovie.CreateMovie(appCtx))
	}

	{
		shows := v1.Group("/shows")
		//GET /v1/shows
		shows.GET("", ginshow.ListShow(appCtx))
		//GET /v1/shows/:id
		shows.GET("/:id", ginshow.GetShowWithID(appCtx))
		//POST /v1/shows
		shows.POST("", middleware.RequireAuth(appCtx, userStore), ginshow.CreateShow(appCtx))
	}

	{
		tickets := v1.Group("/tickets")
		//GET /v1/tickets
		tickets.GET("", ginticket.ListTickets(appCtx))
		//UPDATE /v1/tickets/
		tickets.PUT("", middleware.RequireAuth(appCtx, userStore), ginticket.UpdateTicket(appCtx))
	}

	{
		users := v1
		//GET /v1/profile
		users.GET("/profile", middleware.RequireAuth(appCtx, userStore), ginuser.GetProfile(appCtx))
		//POST /v1/register
		users.POST("/register", ginuser.Register(appCtx))
		//POST /v1/login
		users.POST("/login", ginuser.Login(appCtx))
	}
	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
