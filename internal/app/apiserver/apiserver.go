package apiserver

import (
	"github.com/sirupsen/logrus"
)

// APISERVER ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
}

// New ..
func New(config *Config) *APIServer { // возвращает указатель а APIServer,здесь мы инициализируем наш APIServer
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

// Start ...
func (s *APIServer) Start() error { // для запуска http сервера и подключения к БД и тд

	if err := s.configureLogger(); err != nil {
		return err
	}

	s.logger.Info("starting API server")

	return nil
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
