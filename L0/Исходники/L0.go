package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	//"os"

	//"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nats-io/stan.go"
)

func main() {
	//cache - слайс для хранения JSON из NATS streaming
	cache:= make(map[string]data)

	//dbURL:=fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",os.Getenv("PSQL_NAME"),os.Getenv("PSQL_PASS"),os.Getenv("PSQL_HOST"),os.Getenv("PSQL_PORT"),os.Getenv("PSQL_DB"),os.Getenv("PSQL_SLL"))
	dbURL:="postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable"
	conn, err:=pgxpool.Connect(context.Background(),dbURL)
	if err != nil {
		log.Fatalf("Не получилось подключиться к базе данных:%v\n", err)
	}
	defer conn.Close()

	rows, err:=conn.Query(context.Background(),"select data from taskl0")
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
	}

	sc, _ := stan.Connect("test-cluster","artap", stan.NatsURL("nats://localhost:4222"))

	_, err = sc.Subscribe("foo1", func(m *stan.Msg) {
		var d data

		err:=json.Unmarshal(m.Data,&d)
		if err != nil {
			log.Println(err)
		}

		if _, ok := cache[d.OrderUid]; !ok {
			cache[d.OrderUid]=d

			_,err=conn.Exec(context.Background(),"insert into taskl0 values ($1, $2)",d.OrderUid,m.Data)
			if err != nil {
				log.Println(err)
			}
		}
	}, stan.StartWithLastReceived())
	if err != nil {
		log.Println(err)
	}

	defer sc.Close()

	http.HandleFunc("/",func (w http.ResponseWriter,req *http.Request)  {
		switch req.Method {
		case "GET":
			tmpl, err:= template.ParseFiles("start.html")
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}

			err = tmpl.Execute(w,nil)
			if err !=nil{
				http.Error(w, err.Error(), 400)
				return
			}

		case "POST":	
			if val, ok := cache[req.PostFormValue("order_uid")]; ok {
				
				b, err:=json.MarshalIndent(val,"","\t")
				if err != nil {
					log.Println(err)
				}
				fmt.Fprint(w,string(b))
			} else {
				fmt.Fprint(w,"Структура не найдена.😿")
			}
		}
	})

	log.Fatal(http.ListenAndServe(":4969",nil))
	
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