package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Структура ответа
type Response struct {
	Data       []*Data `json:"data"`
	Error      int     `json:"error"`
	Message    *string `json:"message"`
	StatusCode int     `json:"status_code"`
}
type Data struct {
	AdditionalCharact *int               `json:"additional_charact"`
	AdditionalSkills  []*AdditionalSkill `json:"additional_skills"`
	Experience        int                `json:"experience"`
	Id                int                `json:"id"`
	Isdeleted         bool               `json:"isdeleted"`
	Name              string             `json:"name"`
	PermanentAwards   []*PermanentAward  `json:"permanent_awards"`
	Protection        int                `json:"protection"`
	RandomAwards      []*RandomAward     `json:"random_awards"`
	Types             string             `json:"types"`
	UrlPhoto          string             `json:"url_photo"`
}
type AdditionalSkill struct {
	Name string `json:"name"`
}
type PermanentAward struct {
	Name string `json:"name"`
}
type RandomAward struct {
	Name        string `json:"name"`
	Probability int    `json:"probability"`
}

// Реализация интерфейса Stringer для вывода структуры на экран
func (d *Data) String() string {
	return fmt.Sprintf("\n%+v\n", *d)
}
func (a *AdditionalSkill) String() string {
	return fmt.Sprintf("%+v", *a)
}
func (p *PermanentAward) String() string {
	return fmt.Sprintf("%+v", *p)
}
func (r *RandomAward) String() string {
	return fmt.Sprintf("%+v", *r)
}

func main() {
	//Проверка url
	request, err := url.Parse("https://skytower-game.ru/api/enemy/all/get")
	if err != nil {
		log.Fatalf("%s", err)
	}
	//Выполнение запроса
	response, err := http.Get(request.String())
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer response.Body.Close()
	//Проверка статуса ответа и заголовка
	if response.StatusCode != 200 || response.Header["Content-Type"][0] != http.CanonicalHeaderKey("application/json") {
		log.Fatalf("invalid server response status or content type")
	}
	//Считывание body ответа
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("outsideApi: getInfo: %s", err)
	}
	//Сериализация body в структуру
	structResponse := new(Response)
	if err := json.Unmarshal(body, &structResponse); err != nil {
		log.Fatalf("%s", err)
	}
	//Печать в консоль
	fmt.Printf("%+v", structResponse)

}
