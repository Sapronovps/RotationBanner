package http

import (
	"context"
	"github.com/Sapronovps/RotationBanner/internal/app"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

type Server struct {
	address string
	logg    *zap.Logger
	app     *app.App
	server  *http.Server
}

func NewServer(address string, app *app.App, logger *zap.Logger) *Server {
	return &Server{
		address: address,
		logg:    logger,
		app:     app,
	}
}

func (s *Server) Start(ctx context.Context) error {
	// Создаем новый роутер
	r := mux.NewRouter()
	_ = r

	// Регистрируем обработчики
	r.HandleFunc("/", home).Methods("GET")

	// Создание слота
	r.HandleFunc("/slots", func(w http.ResponseWriter, r *http.Request) {
		addSlot(w, r, s.app)
	}).Methods("POST")
	// Получение слота
	r.HandleFunc("/slots/{id}", func(w http.ResponseWriter, r *http.Request) {
		getSlot(w, r, s.app)
	}).Methods("GET")

	// Создание баннера
	r.HandleFunc("/banners", func(w http.ResponseWriter, r *http.Request) {
		addBanner(w, r, s.app)
	}).Methods("POST")
	// Получение баннера
	r.HandleFunc("/banners/{id}", func(w http.ResponseWriter, r *http.Request) {
		getBanner(w, r, s.app)
	}).Methods("GET")

	// Создание группы
	r.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		addGroup(w, r, s.app)
	}).Methods("POST")
	// Получение группы
	r.HandleFunc("/groups/{id}", func(w http.ResponseWriter, r *http.Request) {
		getGroup(w, r, s.app)
	}).Methods("GET")

	// Создание статистики по баннеру в разрезе слота и группы
	r.HandleFunc("/bannerGroupStats", func(w http.ResponseWriter, r *http.Request) {
		addBannerGroupStats(w, r, s.app)
	}).Methods("POST")

	// Добавляем middleware для логирования
	r.Use(s.loggingMiddleware)

	// Настраиваем HTTP сервер
	s.server = &http.Server{
		Addr:         s.address,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Запускам сервер
	s.logg.Info("Server starting", zap.String("address", s.address))
	err := s.server.ListenAndServe()
	if err != nil {
		s.logg.Fatal(err.Error())
	}

	<-ctx.Done()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	os.Exit(1)
	return nil
}
