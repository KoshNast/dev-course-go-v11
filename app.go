package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Обработчик корневого маршрута (главная страница)
func handler(w http.ResponseWriter, r *http.Request) {
	// Реализация новой фичи: динамическое определение и форматирование времени
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Установка заголовка ответа для корректного отображения кодировки UTF-8
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Вывод исходного приветствия и новой фичи
	fmt.Fprintln(w, "Hello, 世界")
	fmt.Fprintf(w, "Текущее системное время сервера: %s\n", currentTime)
	fmt.Fprintf(w, "Контрольная работа: Вариант 11 (GitHub/GitFlow/WSL)\n")
}

// Новый эндпоинт для автоматической проверки работоспособности (Health Check)
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Возвращаем статус-код 200 OK и текстовое подтверждение
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	// Регистрация обработчиков для соответствующих URI-маршрутов
	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)

	// Логирование запуска сервера в консоль Linux
	fmt.Println("Демонстрационное Go-приложение успешно запущено на порту :8888.")
	fmt.Println("Для принудительной остановки нажмите комбинацию Ctrl+C...")

	// Запуск HTTP-сервера на прослушивание порта 8888
	log.Fatal(http.ListenAndServe(":8888", nil))
}
