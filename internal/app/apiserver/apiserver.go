package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APISERVER ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New ..
func New(config *Config) *APIServer { // возвращает указатель а APIServer,здесь мы инициализируем наш APIServer
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error { // для запуска http сервера и подключения к БД и тд

	if err := s.configureLogger(); err != nil {
		return err
	}

	// т.к. этот метода не возвращает ошибок, то мы его никак не обрабатываем
	s.configureRouter()

	s.logger.Info("starting API server")

	// в качествк адреса мы возьмем BindAddr из config`a, 2-возьмем поле router
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error { // конфигурируем логгер
	// она может вернуть ошибку, потому что возможен неправильный уровень логгирования
	// или неправильная передача строки, поэтому можем ожидать ошибку от этого метода

	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	// если ошибки нет, то ставим нашему логгеру соответствующий уровень
	s.logger.SetLevel(level)

	//выходим из функции
	return nil
}

func (s *APIServer) configureRouter() {
	// никакую ошибку вернуть не может, потому что он просто описывает роутинг

	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	// возвращает HandlerFunc
	// идея у хэндлера возвращать не привычную функцию (как в большгинстве туториалов)
	// а именно вот такой интерфейс (http.HandlerFunc)-функцию
	// прикол в том что мы можем *тут* определить какие-то переменные, которые
	// будут использоваться только в этом хэндлере, и код *тут*
	// выполнится всего 1 раз
	// вся логика обработки запроса будет описываться в функции ниже :)

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}

}
