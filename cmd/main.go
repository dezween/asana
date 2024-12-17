package main

import (
	"context"
	"log"
	"time"
	"os"

	"github.com/dezween/Calendar/internal/application"
	"github.com/dezween/Calendar/internal/config"
)

func main() {
	if err := os.MkdirAll(config.OutputFolder, os.ModePerm); err != nil {
		log.Fatalf("Ошибка создания папки: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	application.StartExtractors(ctx)

	time.Sleep(15 * time.Minute)
	cancel()
	log.Println("Завершение программы")
}
