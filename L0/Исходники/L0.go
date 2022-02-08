package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"github.com/jackc/pgx/v4"
	"github.com/nats-io/stan.go"
)

func main() {
	//cache - слайс для хранения JSON из NATS streaming
	cache:= make(map[string]data)

	var (
		name string
		password string
		conn *pgx.Conn
	)

	for {
		fmt.Println("Введите имя пользователя postgres:")
		_,err:=fmt.Scan(&name)
		if err != nil{
			log.Println(err)
			continue
		}
		fmt.Printf("Введите пароль для пользователя %s:\n",name)
		_,err=fmt.Scan(&password)
		if err != nil{
			log.Println(err)
			continue
		}

		// dbURL - DSN для подключения к БД (postgres)
		dbURL:=fmt.Sprintf("postgres://%s:%s@localhost:5432/l0?sslmode=disable",name,password)
		conn, err=pgx.Connect(context.Background(),dbURL) // Подключение к БД
		if err != nil {
			log.Printf("Не получилось подключиться к базе данных:%v\n", err)
		} else {
			log.Println("Подключились к БД")
			break
		}
		defer conn.Close(context.Background()) 

	}	

	rows, err:=conn.Query(context.Background(),"select data from taskl0") // Получаем данные JSON из БД
	if err != nil {
		log.Println(err)
	}

	
	for rows.Next() {
		var ( 
			b []byte
			da data
		)

		err=rows.Scan(&b)
		if err != nil {
			log.Println(err)
		}

		err=json.Unmarshal(b,&da)
		if err != nil{
			log.Println(err)
		}
		cache[da.OrderUid]=da
		log.Printf("Получили данные из бд Order UID: %s\n",da.OrderUid)
	}
	rows.Close()

	sc, err := stan.Connect("test-cluster","artap", stan.NatsURL("nats://localhost:4222")) // Подключаемся к NATS-streaming
	if err != nil && err!=io.EOF {
		log.Fatalln(err)
	} else {
		log.Println("Подключились к Nats-streaming")
	}
	defer sc.Close()

	_, err = sc.Subscribe("foo1", func(m *stan.Msg) { //Оформляем подписку
		var d data

		err:=json.Unmarshal(m.Data,&d) // Получаем данные из NATS-streaming и проверяем на соответствие
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Получили данные из NATS-streaming")
		}

		if _, ok := cache[d.OrderUid]; !ok { // Проверяем есть ли данные уже в кэше 
			cache[d.OrderUid]=d // Добавляем данные в кэш

			_,err=conn.Exec(context.Background(),"insert into taskl0 values ($1, $2)",d.OrderUid,m.Data) // Добавляем данные в БД
			if err != nil {
				log.Println(err)
			} else {
				log.Printf("Добавили в БД Order UID: %s\n",d.OrderUid)
			}
		}
	}, stan.StartWithLastReceived())
	if err != nil {
		log.Println(err)
	}

	http.HandleFunc("/",func (w http.ResponseWriter,req *http.Request)  { // Устанавливаем обработчик запросов
		switch req.Method {
		case "GET": // Получение страницы
			tmpl, err:= template.ParseFiles("start.html")
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			err = tmpl.Execute(w,nil) // Отправляем страницу
			if err !=nil{
				http.Error(w, err.Error(), 400)
				return
			}

			log.Printf("Зашли в интерфейс")

		case "POST": // Запрос на получение данных
			if val, ok := cache[req.PostFormValue("order_uid")]; ok { // Ищем данные
				
				b, err:=json.MarshalIndent(val,"","\t")
				if err != nil {
					log.Println(err)
				}
				log.Printf("Отправили данные с Order UID: %s\n",req.PostFormValue("order_uid"))
				fmt.Fprint(w,string(b))
			} else {
				log.Println("Не верный запрос")
				fmt.Fprint(w,"Структура не найдена.😿")
			}
		}
	})

	log.Fatal(http.ListenAndServe(":4969",nil)) // Запускаем сервер
	
}

type data struct {
	OrderUid    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Entry       string `json:"entry"`
	Delivery    struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction"`
		RequestId    string `json:"request_id"`
		Currency     string `json:"currency"`
		Provider     string `json:"provider"`
		Amount       int    `json:"amount"`
		PaymentDt    int    `json:"payment_dt"`
		Bank         string `json:"bank"`
		DeliveryCost int    `json:"delivery_cost"`
		GoodsTotal   int    `json:"goods_total"`
		CustomFee    int    `json:"custom_fee"`
	} `json:"payment"`
	Items []struct {
		ChrtId      int    `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       int    `json:"price"`
		Rid         string `json:"rid"`
		Name        string `json:"name"`
		Sale        int    `json:"sale"`
		Size        string `json:"size"`
		Total_price int    `json:"total_price"`
		Nm_id       int    `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      int    `json:"status"`
	} `json:"items"`
	Locale string `json:"locale"`
	InternalSignature string `json:"internal_signature"`
	CustomerId string `json:"customer_id"`
	DeliveryService string `json:"delivery_service"`
	Shardkey string `json:"shardkey"`
	SmId int `json:"sm_id"`
	DateCreated string `json:"date_created"`
	OofShard string `json:"oof_shard"`
}