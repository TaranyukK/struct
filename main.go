package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"struct/api"
	"struct/config"
)

func main() {
	cmdCreate := flag.Bool("create", false, "")
	cmdGet := flag.Bool("get", false, "")
	cmdUpdate := flag.Bool("update", false, "")
	cmdDelete := flag.Bool("delete", false, "")
	cmdList := flag.Bool("list", false, "")

	flagID := flag.String("id", "", "")
	flagFile := flag.String("file", "", "")

	flag.Parse()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	client := api.NewClient(cfg.Key)

	switch {
	case *cmdCreate:
		data := loadDataFromFile(*flagFile)
		err = client.Create(data)
	case *cmdGet:
		requireID(*cmdGet, *flagID)
		err = client.Get(*flagID)
	case *cmdUpdate:
		requireID(*cmdUpdate, *flagID)
		data := loadDataFromFile(*flagFile)
		err = client.Update(*flagID, data)
	case *cmdDelete:
		requireID(*cmdDelete, *flagID)
		err = client.Delete(*flagID)
	case *cmdList:
		err = client.List()
	default:
		fmt.Println("Не указано действие. Используйте -create, -get, -update, -delete или -list")
		flag.Usage()
		return
	}

	if err != nil {
		log.Fatal("Ошибка: ", err)
	}
}

func requireID(cmdName string, id string) {
	if id == "" {
		log.Fatalf("Для команды требуется указать флаг -id")
	}
}

func loadDataFromFile(filePath string) any {
	if filePath == "" {
		return map[string]string{"message": "default data"}
	}

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Не удалось прочитать файл %s: %v", filePath, err)
	}

	var data any
	if err := json.Unmarshal(fileData, &data); err != nil {
		log.Fatalf("Файл %s не является валидным JSON: %v", filePath, err)
	}
	return data
}