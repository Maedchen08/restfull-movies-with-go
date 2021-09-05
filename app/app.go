package app

import ("homework-rakamin-go-sql/config")

type Application struct {
	Config *config.Config
}

func Init() *Application {
	application := &Application{
		Config: config.Init(),
	}

	return application
}