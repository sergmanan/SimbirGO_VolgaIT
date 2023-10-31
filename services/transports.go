package services

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Transport struct {
	CanBeRented   *bool    `json:"canBeRented"`
	TransportType *string  `json:"transportType"`
	Model         *string  `json:"model"`
	Color         *string  `json:"color"`
	Identifier    *string  `json:"identifier"`
	Description   *string  `json:"description"`
	Latitude      *float64 `json:"latitude"`
	Longitude     *float64 `json:"longitude"`
	MinutePrice   *float64 `json:"minutePrice"`
	DayPrice      *float64 `json:"dayPrice"`
	Account_id    *int     `json:"account_id"`
	Id            *int     `json:"id"`
}

type TransportService struct {
	db       *sql.DB
	c_params *ConnectionParams
}

func (t *TransportService) Insert(transport Transport) (int64, error) {
	var res sql.Result
	var err error

	res, err = t.db.Exec("INSERT INTO public.\"Transports\" (canBeRented, transportType, model, color, identifier, description, latitude, longitude, minutePrice, dayPrice, account_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
		transport.CanBeRented, transport.TransportType, transport.Model, transport.Color, transport.Identifier, transport.Description, transport.Latitude, transport.Longitude, transport.MinutePrice, transport.DayPrice, transport.Account_id)

	if err != nil {
		log.Println(err)
		return -1, err
	}
	var id int64
	id, err = res.LastInsertId()

	if err != nil {
		log.Println(err)
		return -1, err
	}

	return id, nil
}

func (t *TransportService) GetByID(id int) (Transport, error) {

	var transport Transport
	err := t.db.QueryRow("SELECT canBeRented, transportType, model, color, identifier, description, latitude, longitude, minutePrice, dayPrice, account_id FROM public.\"Transports\" WHERE id = $1", id).
		Scan(&transport.CanBeRented, &transport.TransportType, &transport.Model, &transport.Color, &transport.Identifier, &transport.Description, &transport.Latitude, &transport.Longitude, &transport.MinutePrice, &transport.DayPrice, &transport.Account_id)
	if err != nil {
		log.Println(err)
		return Transport{}, err
	}
	return transport, nil
}

func (t *TransportService) DeleteByID(id int) error {
	_, err := t.db.Exec("DELETE FROM public.\"Transports\" WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
