package main

import (
	"app/config"
	"fmt"
	"log"
	"net/http"

	"app/internal/handler"
	"app/internal/loader"
	"app/internal/repository"
	"app/internal/service"

	"github.com/go-chi/chi/v5"
)

func main() {

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Load error config: %v", err)
	}

	app := NewApplicationDefault(cfg)

	// Setup
	err = app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}

	// - Run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// NewApplicationDefault creates a new default application
func NewApplicationDefault(cfg *config.ConfigAppDefault) *ApplicationDefault {
	// default values
	defaultRouter := chi.NewRouter()
	defaultConfig := &config.ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}

	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		dbFile:     defaultConfig.DbFile,
	}
}

// ApplicationDefault represents the default application
type ApplicationDefault struct {
	// router represents the router of the application
	rt *chi.Mux
	// serverAddr represents the address of the server
	serverAddr string
	// dbFile represents the path to the database file
	dbFile string
}

// SetUp sets up the application
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	db := loader.NewLoaderTicketCSV(a.dbFile)

	data, err := db.Load()
	if err != nil {
		fmt.Println(err)
		return
	}

	rp := repository.NewRepositoryTicketMap(data, len(data))
	ticketService := service.NewServiceTicketDefault(rp)

	tickerHandler := handler.NewTicketHandler(ticketService)

	// routes
	(*a).rt.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("OK"))
	})

	(*a).rt.Get("/tickets/total", tickerHandler.GetTotalTickets)
	(*a).rt.Get("/tickets/total_amount", tickerHandler.GetTotalAmountTickets)
	(*a).rt.Get("/tickets/average/{destination}", tickerHandler.AverageDestination)

	return
}

// Run runs the application
func (a *ApplicationDefault) Run() (err error) {
	log.Printf("Server running on %s", a.serverAddr)
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}
