package main

import (
	"flag"
	"http-rest-api/internal/app/apiserver"
	"log"

	"github.com/BurntSushi/toml"
)

// чтобы путь к конфиг-файлу(configs/apiserver.toml) можно было задавать в качестве флага при запуске бинарника
var configPath string

func init() {
	// мы хотим парсить пер.config-path, по умолчанию = configs/apiserver.toml, и описание для help`a
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	// вызываем флаг парс, чтобы распарсить наши флаги и записать их в нужные переменные
	flag.Parse()

	// воспользуемся библиотекой, чтобы прочитать файл .toml, распарсить его и записать
	// все его значения в нашу переменную конфиг
	config := apiserver.NewConfiпg()
	_, err := toml.DecodeFile(configPath, config) // 1-путь до конфиг-файлаб 2-переменная, в которую мы хотим записать
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

}
