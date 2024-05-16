package main

import (
	"encoding/json"
	"github.com/AtaullinShamil/wbschool_exam_L2/tree/main/develop/dev11/internal/handler"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

type Config struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
}

func getConfig(fileName string) (string, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return "", err
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	var config Config

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return "", err
	}

	hostPort := config.Server.Host + ":" + config.Server.Port
	return hostPort, nil
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Incoming request:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	hostPort, err := getConfig("config.json")
	if err != nil {
		log.Fatalln(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", handler.CreateEventHandler)        //Post
	mux.HandleFunc("/update_event", handler.UpdateEventHandler)        //Post
	mux.HandleFunc("/delete_event", handler.DeleteEventHandler)        //Post
	mux.HandleFunc("/events_for_day", handler.EventsForDayHandler)     //Get
	mux.HandleFunc("/events_for_week", handler.EventsForWeekHandler)   //Get
	mux.HandleFunc("/events_for_month", handler.EventsForMonthHandler) //Get

	http.ListenAndServe(hostPort, loggingMiddleware(mux))
}

//{
//"title": "Party",
//"description": "for wb",
//"start": "2024.02.03",
//"end": "2024.02.04"
//}
