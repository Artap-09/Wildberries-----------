package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var Users map[int][5]*Year

func main() {

	http.HandleFunc("/create_event", CreateEvent)
	http.HandleFunc("/update_event", UpdateEvent)
	http.HandleFunc("/delete_event", DeleteEvent)
	http.HandleFunc("/events_for_day", EventsForDay)
	http.HandleFunc("/events_for_week", EventsForWeek)
	http.HandleFunc("/events_for_month", EventsForMonth)
	err:=http.ListenAndServe(os.Args[1], nil)
	log.Println(err)
}

func CreateEvent(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "ожидается метод POST", 503)
		return
	}

	event, err := Parser(req)
	if err != nil {
		http.Error(w, "невалидные данные", 400)
		return
	}

	eventTime, err := time.Parse("2006-01-02",event.Time)
	if err != nil {
		http.Error(w, "невалидные данные", 400)
		return
	}
	
	year, week := eventTime.ISOWeek()

	switch year - time.Now().Year() {
	case 0:
		yearN:=Users[event.IdUser][0]
		if yearN == nil {
			yearN = &Year{}
		}
		
		day:=yearN.Months[eventTime.Month()][eventTime.Day()]
		if day == nil {
			day = make([]*Event,0,0)
		}
		day = append(day, event)
		yearN.Weeks[week][eventTime.Weekday()] = day
		event.IdEvent = len(day) - 1
		fmt.Println(yearN)
	case 1:
		yearN:=Users[event.IdUser][1]
		if yearN == nil {
			yearN = &Year{}
		}
		day:=yearN.Months[eventTime.Month()][eventTime.Day()]
		if day == nil {
			day = make([]*Event,0,0)
		}
		day = append(day, event)
		yearN.Weeks[week][eventTime.Weekday()] = day
		event.IdEvent = len(day) - 1

	case 2:
		yearN:=Users[event.IdUser][2]
		if yearN == nil {
			yearN = &Year{}
		}
		day:=yearN.Months[eventTime.Month()][eventTime.Day()]
		if day == nil {
			day = make([]*Event,0,0)
		}
		day = append(day, event)
		yearN.Weeks[week][eventTime.Weekday()] = day
		event.IdEvent = len(day) - 1

	case 3:
		yearN:=Users[event.IdUser][3]
		if yearN == nil {
			yearN = &Year{}
		}
		day:=yearN.Months[eventTime.Month()][eventTime.Day()]
		if day == nil {
			day = make([]*Event,0,0)
		}
		day = append(day, event)
		yearN.Weeks[week][eventTime.Weekday()] = day
		event.IdEvent = len(day) - 1

	case 4:
		yearN:=Users[event.IdUser][4]
		if yearN == nil {
			yearN = &Year{}
		}
		day:=yearN.Months[eventTime.Month()][eventTime.Day()]
		if day == nil {
			day = make([]*Event,0,0)
		}
		day = append(day, event)
		yearN.Weeks[week][eventTime.Weekday()] = day
		event.IdEvent = len(day) - 1

	default:
		http.Error(w, fmt.Sprintf("можно создать задачу только в диапазоне 5 лет от %d до %d", time.Now().Year(), time.Now().Year()+5), 500)
		return
	}
	
}

func UpdateEvent(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "ожидается метод POST", 503)
		return
	}

	event, err := Parser(req)
	if err != nil {
		http.Error(w, "невалидные данные", 400)
		return
	}

	eventTime, err := time.Parse("2006-01-02",event.Time)
	if err != nil {
		http.Error(w, "невалидные данные", 400)
		return
	}

	year, week := eventTime.ISOWeek()

	switch year - time.Now().Year() {
	case 0:
		year:=Users[event.IdUser][0]
		if year == nil {
			http.Error(w,"невалидные данные", 400)
			return
		}
		day:=year.Months[eventTime.Month()][eventTime.Day()]
		if day == nil {
			http.Error(w,"невалидные данные", 400)
			return
		}
		day = append(day, event)
		year.Weeks[week][eventTime.Weekday()] = day

	case 1:
		year:=Users[event.IdUser][1]
		if year == nil {
			http.Error(w,"невалидные данные", 400)
			return
		}
		day:=year.Months[eventTime.Month()][eventTime.Day()]
		if day == nil {
			http.Error(w,"невалидные данные", 400)
			return
		}
		day = append(day, event)
		year.Weeks[week][eventTime.Weekday()] = day

	case 2:
		year:=Users[event.IdUser][2]
		if year == nil {
			http.Error(w,"невалидные данные", 400)
			return
		}
		day:=year.Months[eventTime.Month()][eventTime.Day()]
		if day == nil {
			http.Error(w,"невалидные данные", 400)
			return
		}
		day = append(day, event)
		year.Weeks[week][eventTime.Weekday()] = day

	case 3:
		year:=Users[event.IdUser][3]
		if year == nil {
			http.Error(w,"невалидные данные", 400)
			return
		}
		day:=year.Months[eventTime.Month()][eventTime.Day()]
		if day == nil {
			http.Error(w,"невалидные данные", 400)
			return
		}
		day = append(day, event)
		year.Weeks[week][eventTime.Weekday()] = day

	case 4:
		year:=Users[event.IdUser][4]
		if year == nil {
			http.Error(w,"невалидные данные", 400)
			return
		}
		day:=year.Months[eventTime.Month()][eventTime.Day()]
		if day == nil {
			http.Error(w,"невалидные данные", 400)
			return
		}
		day = append(day, event)
		year.Weeks[week][eventTime.Weekday()] = day

	default:
		http.Error(w, fmt.Sprintf("можно создать задачу только в диапазоне 5 лет от %d до %d", time.Now().Year(), time.Now().Year()+5), 500)
		return
	}
}

func DeleteEvent(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "ожидается метод POST", 503)
		return
	}

	event, err := Parser(req)
	if err != nil {
		http.Error(w, "невалидные данные", 400)
		return
	}

	eventTime, err := time.Parse("2006-01-02",event.Time)
	if err != nil {
		http.Error(w, "невалидные данные", 400)
		return
	}

	year, week := eventTime.ISOWeek()

	switch year - time.Now().Year() {
	case 0:
		Users[event.IdUser][0].Months[eventTime.Month()][eventTime.Day()] = Delete(Users[event.IdUser][0].Months[eventTime.Month()][eventTime.Day()], event.IdEvent)
		Users[event.IdUser][0].Weeks[week][eventTime.Weekday()] = Delete(Users[event.IdUser][0].Weeks[week][eventTime.Weekday()], event.IdEvent)
	case 1:
		Users[event.IdUser][1].Months[eventTime.Month()][eventTime.Day()] = Delete(Users[event.IdUser][1].Months[eventTime.Month()][eventTime.Day()], event.IdEvent)
		Users[event.IdUser][1].Weeks[week][eventTime.Weekday()] = Delete(Users[event.IdUser][1].Weeks[week][eventTime.Weekday()], event.IdEvent)
	case 2:
		Users[event.IdUser][2].Months[eventTime.Month()][eventTime.Day()] = Delete(Users[event.IdUser][2].Months[eventTime.Month()][eventTime.Day()], event.IdEvent)
		Users[event.IdUser][2].Weeks[week][eventTime.Weekday()] = Delete(Users[event.IdUser][2].Weeks[week][eventTime.Weekday()], event.IdEvent)
	case 3:
		Users[event.IdUser][3].Months[eventTime.Month()][eventTime.Day()] = Delete(Users[event.IdUser][3].Months[eventTime.Month()][eventTime.Day()], event.IdEvent)
		Users[event.IdUser][3].Weeks[week][eventTime.Weekday()] = Delete(Users[event.IdUser][3].Weeks[week][eventTime.Weekday()], event.IdEvent)
	case 4:
		Users[event.IdUser][4].Months[eventTime.Month()][eventTime.Day()] = Delete(Users[event.IdUser][4].Months[eventTime.Month()][eventTime.Day()], event.IdEvent)
		Users[event.IdUser][4].Weeks[week][eventTime.Weekday()] = Delete(Users[event.IdUser][4].Weeks[week][eventTime.Weekday()], event.IdEvent)
	default:
		http.Error(w, fmt.Sprintf("можно создать задачу только в диапазоне 5 лет от %d до %d", time.Now().Year(), time.Now().Year()+5), 500)
		return
	}
}

func EventsForDay(w http.ResponseWriter, req *http.Request) {
	idUser, err := strconv.Atoi(req.URL.Query().Get("id_user"))
	if err != nil {
		http.NotFound(w, req)
		return
	}

	data, err:=time.Parse("2006-01-02", req.URL.Query().Get("data"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "невалидные данные", 400)
		return
	}

	fmt.Println(data.Format(time.RFC3339))
	fmt.Println(Users[idUser])
	var b []byte
	for _, v := range Users[idUser][data.Year()-time.Now().Year()].Months[data.Month()][data.Day()] {
		b1, err := json.Marshal(v)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		b = append(b, b1...)
	}

	w.Write(b)
}

func EventsForWeek(w http.ResponseWriter, req *http.Request) {
	idUser, err := strconv.Atoi(req.URL.Query().Get("id_user"))
	if err != nil {
		http.NotFound(w, req)
		return
	}

	data, err := time.Parse("2006-01-02", req.URL.Query().Get("data"))
	if err != nil {
		http.Error(w, "невалидные данные", 400)
		return
	}

	var b []byte
	for _, v := range Users[idUser][data.Year()-time.Now().Year()].Weeks[data.Month()][data.Day()] {
		b1, err := json.Marshal(v)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		b = append(b, b1...)
	}

	w.Write(b)
}

func EventsForMonth(w http.ResponseWriter, req *http.Request) {
	idUser, err := strconv.Atoi(req.URL.Query().Get("id_user"))
	if err != nil {
		http.NotFound(w, req)
		return
	}

	data, err := time.Parse("2006-01-02", req.URL.Query().Get("data"))
	if err != nil {
		http.Error(w, "невалидные данные", 400)
		return
	}

	var b []byte
	for _, val := range Users[idUser][data.Year()-time.Now().Year()].Months[data.Month()] {
		for _, v := range val {
			b1, err := json.Marshal(v)
			if err != nil {
				http.Error(w, err.Error(), 500)
			}

			b = append(b, b1...)
		}
	}
	w.Write(b)
}

func Delete(slice []*Event, idx int) []*Event {
	slice[len(slice)-1], slice[idx] = slice[idx], slice[len(slice)-1]
	slice[idx].IdEvent = idx
	return slice[:len(slice)-1]
}

func Parser(req *http.Request) (*Event, error) {
	var event *Event

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

type Year struct {
	Months Months
	Weeks  Weeks
}

type Months [12][31]Day

type Weeks [53][7]Day

type Day []*Event

type Event struct {
	IdUser  int       `json:"id_user"`
	IdEvent int       `json:"id_event"`
	Header  string    `json:"header"`
	Event   string    `json:"event"`
	Time    string `json:"time"`
}
