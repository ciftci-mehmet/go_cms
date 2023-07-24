package main

import (
	"github.com/go-chi/chi/v5"

	"github.com/ciftci-mehmet/go_cms/cmd/web/handlers"
	db "github.com/ciftci-mehmet/go_cms/db/sqlc"
)

type Application struct {
	//AppName       string
	//Debug         bool
	//Version       string
	RootPath string
	//config        config
	//EncryptionKey string
	Server   Server
	Routes   *chi.Mux
	Handlers *handlers.Handlers
	//ErrorLog      *log.Logger
	//InfoLog       *log.Logger
	Query *db.Queries
	//Render        *render.Render
	//Session       *scs.SessionManager
	//Cache         cache.Cache
	//Mail          mailer.Mail
	//Scheduler     *cron.Cron
}

type Server struct {
	ServerName string
	Port       string
	Secure     bool
	URL        string
}

//type config struct {
//	port string
//	//renderer    string
//	//cookie      cookieConfig
//	//sessionType string
//	//database    databaseConfig
//	//redis       redisConfig
//}
