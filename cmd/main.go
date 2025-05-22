package main

import (
	"context"
	"log"
	"my_project/configs"
	"my_project/db"
	"my_project/internal/auth"
	"my_project/internal/link"
	"my_project/internal/stat"
	"my_project/internal/user"
	"my_project/pkg/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func app() http.Handler {
	conf := configs.LoadConfig()

	router := http.NewServeMux()

	// DB
	Db := db.NewDb(conf)
	time.Sleep(5 * time.Second)
	//REPO
	LinkRepository := link.NewLinkRepository(Db)
	UserRepository := user.NewUserRepository(Db)
	StatRepository := stat.NewStatRepository(Db)

	AuthService := auth.NewAuthService(UserRepository)
	//APPENDING HANDLERS IN ROUTER
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf,
		Auth: AuthService})
	link.NewLinkHandler(router, link.LinkHandlerDeps{LinkRepository: LinkRepository,
		Config:         conf,
		StatRepository: StatRepository})
	stat.NewStatHandler(router, stat.StatHandlerDeps{StatRepository: StatRepository, Config: conf})

	stackMiddleware := middleware.Chain(middleware.CORS,
		middleware.Logging)

	return stackMiddleware(router)
}

func main() {
	application := app()

	server := http.Server{
		Addr:    ":8080",
		Handler: application,
	}
	log.Println("Server started in port 8080")
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Println("Server stops")
		}

	}()

	ch := make(chan (os.Signal), 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch
	log.Println("SPOPPING")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	server.Shutdown(ctx)

}
