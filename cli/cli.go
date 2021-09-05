package cli

import (
	"homework-rakamin-go-sql/app"
	"homework-rakamin-go-sql/config"
	"homework-rakamin-go-sql/config/database"
	"homework-rakamin-go-sql/repository"
	"homework-rakamin-go-sql/services"
	"homework-rakamin-go-sql/handlers"
	route "homework-rakamin-go-sql/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os/signal"
	"os"
	log "github.com/sirupsen/logrus"

)

type Cli struct {
	Args []string
}

func NewCli(args []string) *Cli {
	return &Cli{
		Args: args,
	}
}

func (cli *Cli) Run(application *app.Application){
	fiberConfig := config.FiberConfig()
	app := fiber.New(fiberConfig)

	//set up connection
	connDB:= database.InitDb()

	//movies services
	moviesRepository := repository.NewMoviesRepository(connDB)
	moviesService := services.NewMoviesService(moviesRepository)
	moviesHandler := handlers.NewMoviesHandler(moviesService)

	//REGISTER HANDLER TO Routes
	routes := route.NewRoutes(moviesHandler)
	routes.InitializeRoutes(app)

	//not found routes
	route.NotFoundRoute(app)

	StartServerWithGracefulShutdown(app, application.Config.AppPort)

}

func StartServerWithGracefulShutdown(app *fiber.App,port string) {
	appPort:= fmt.Sprintf(`:%s`,port)
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal,1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := app.Shutdown(); err != nil {
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := app.Listen(appPort); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}