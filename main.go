package main

import (
	database "Recipe_App/database/mysql"
	"Recipe_App/handler"
	"Recipe_App/models"
	"Recipe_App/repository"
	"Recipe_App/usecase"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "Recipe_App/docs"

	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatal().Err(envErr).Msg("cannot load environment")
	}

	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8080"
	}

	var dbUsername = os.Getenv("DB_USERNAME")
	var dbPassword = os.Getenv("DB_PASSWORD")
	var dbName = os.Getenv("DB_DATABASE")
	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	db, dbErr := database.ConnectDB(dbUsername, dbPassword, dbHost, dbPort, dbName)

	if dbErr != nil {
		log.Fatal().Err(dbErr).Msg("cannot connect to database")
	}

	//auto migrate database
	db.AutoMigrate(&models.Bahan{})
	db.AutoMigrate(&models.Kategori{})
	db.AutoMigrate(&models.Resep{})
	db.AutoMigrate(&models.Komposisi{})

	//repository layer
	bahanRepository := repository.NewBahanRepositoryImpl(db)
	kategoriRepository := repository.NewKategoriRepositoryImpl(db)
	resepRepository := repository.NewResepRepositoryImpl(db)
	komposisiRepository := repository.NewKomposisiRepositoryImpl(db)

	//usecase layer
	bahanUsecase := usecase.NewBahanUsecaseImpl(bahanRepository)
	kategoriUsecase := usecase.NewKategoriUsecaseImpl(kategoriRepository)
	resepUsecase := usecase.NewResepUsecaseImpl(db, resepRepository, komposisiRepository)

	//handler layer
	bahanHandler := handler.NewBahanHandlerImpl(bahanUsecase)
	kategoriHandler := handler.NewKategoriHandlerImpl(kategoriUsecase)
	resepHandler := handler.NewResepHandlerImpl(resepUsecase)

	router := gin.Default()
	router.Use(cors.Default())
	// route swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	//route master bahan
	bahanRouter := api.Group("/bahan")
	bahanRouter.GET("", bahanHandler.GetAll)
	bahanRouter.POST("", bahanHandler.Create)
	bahanRouter.PUT("/:id", bahanHandler.Update)
	bahanRouter.GET("/:id", bahanHandler.GetById)
	bahanRouter.DELETE("/delete/:id", bahanHandler.Delete)

	//route master kategori
	kategoriRouter := api.Group("/kategori")
	kategoriRouter.GET("", kategoriHandler.GetAll)
	kategoriRouter.POST("", kategoriHandler.Create)
	kategoriRouter.PUT("/:id", kategoriHandler.Update)
	kategoriRouter.GET("/:id", kategoriHandler.GetById)
	kategoriRouter.DELETE("/delete/:id", kategoriHandler.Delete)

	//route master resep
	resepRouter := api.Group("/resep")
	resepRouter.POST("", resepHandler.Create)
	resepRouter.PUT("/:id", resepHandler.Update)
	resepRouter.GET("/:id", resepHandler.GetById)
	resepRouter.DELETE("/delete/:id", resepHandler.Delete)

	router.Run(":" + appPort)
}
