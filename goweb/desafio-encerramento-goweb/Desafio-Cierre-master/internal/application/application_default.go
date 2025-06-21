package application

import (
	"app/desafio-goweb/internal/handler"
	"app/desafio-goweb/internal/loader"
	"app/desafio-goweb/internal/repository"
	"app/desafio-goweb/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ConfigAppDefault struct {
	ServerAddr string
	DbFilePath string
}

func NewServerChi(cfg *ConfigAppDefault) *ApplicationDefault {
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
	}

	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}

		if cfg.DbFilePath != "" {
			defaultConfig.DbFilePath = cfg.DbFilePath
		}
	}

	return &ApplicationDefault{
		serverAddr: defaultConfig.ServerAddr,
		dbFilePath: defaultConfig.DbFilePath,
	}
}

type ApplicationDefault struct {
	rt         *chi.Mux
	serverAddr string
	dbFilePath string
}

func (a *ApplicationDefault) SetUp() (err error) {
	ld := loader.NewLoaderTicketCSV(a.dbFilePath)
	db, err := ld.Load()
	lastId := len(db)
	if err != nil {
		return
	}

	rp := repository.NewRepositoryTicketMap(db, lastId)

	sv := service.NewServiceTicketDefault(rp)
	hd := handler.NewHandlerTicketDefault(sv)

	rt := chi.NewRouter()

	rt.Use(middleware.Logger)
	rt.Use(middleware.Recoverer)

	rt.Route("/tickets", func(rt chi.Router) {
		rt.Get("/get", hd.GetAllTickets())
		rt.Get("/getByCountry/{dest}", hd.GetTicketsAmountByDestinationCountry())
		rt.Get("/getAverage/{dest}", hd.GetPercentageTicketsByDestinationCountry())
	})

	err = http.ListenAndServe(a.serverAddr, rt)

	return
}
