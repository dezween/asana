package application

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"encoding/json"
	"os"

	"github.com/dezween/Calendar/internal/config"
	"github.com/dezween/Calendar/internal/controllers"
)

var wg sync.WaitGroup

func StartExtractors(ctx context.Context) {
	go periodicExtractor(ctx, config.FetchIntervalFastTest, "fast")
	go periodicExtractor(ctx, config.FetchIntervalSlowTest, "slow")
}

func periodicExtractor(ctx context.Context, interval time.Duration, mode string) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Printf("[%s] Извлечение отменено", mode)
			return
		case <-ticker.C:
			log.Printf("[%s] Начинаю извлечение данных...", mode)
			extractData(ctx)
		}
	}
}

func extractData(ctx context.Context) {
	dataCh := make(chan string)
	wg.Add(1)

	go workerPool(ctx, dataCh)

	go func() {
		defer close(dataCh)
		dataCh <- "users"
		dataCh <- "projects"
	}()
	wg.Done()
}

func workerPool(ctx context.Context, jobs <-chan string) {
	for i := 0; i < config.WorkerCount; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs {
				select {
				case <-ctx.Done():
					log.Printf("Worker %d завершает работу", workerID)
					return
				default:
					if job == "users" {
						users, err := controllers.FetchUsers()

						if err != nil {
							log.Printf("Ошибка получения пользователей: %v", err)
							continue
						}
						for _, user := range users {
							saveToFile("user", user.GID, user)
						}
					} else if job == "projects" {
						projects, err := controllers.FetchProjects()

						if err != nil {
							log.Printf("Ошибка получения проектов: %v", err)
							continue
						}
						for _, project := range projects {
							saveToFile("project", project.GID, project)
						}
					}
				}
			}
		}(i)
	}
}

func saveToFile(baseFilename string, id string, data interface{}) {
	timestamp := time.Now().Format("20060102_150405")

	filename := fmt.Sprintf("%s/%s_%s_%s.json", config.OutputFolder, baseFilename, id, timestamp)
	
	file, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		log.Printf("Ошибка сериализации данных: %v", err)
		return
	}

	if err := os.WriteFile(filename, file, 0644); err != nil {
		log.Printf("Ошибка сохранения файла %s: %v", filename, err)
		return
	}

	log.Printf("Файл сохранен: %s", filename)
}

