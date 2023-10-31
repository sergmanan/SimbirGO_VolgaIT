package services

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Rent struct {
	Transport_id *int    `json:"transport_id"`
	Tenant_id    *int    `json:"tenant_id"`
	RentType     *string `json:"rentType"`
	Owner_id     *int    `json:"owner_id"`
	Id           *int    `json:"id"`
}

type RentService struct {
	db       *sql.DB
	c_params *ConnectionParams
}

func (t *RentService) Insert(rent Rent) (int64, error) {
	var res sql.Result
	var err error

	res, err = t.db.Exec("INSERT INTO public.\"Rents\" (transport_id, tenant_id, rentType, owner_id) VALUES ($1, $2, $3, $4) RETURNING id", rent.Transport_id, rent.Tenant_id, rent.RentType, rent.Owner_id)

	if err != nil {
		return -1, err
	}
	var id int64
	id, err = res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (t *RentService) GetByID(id int) (Rent, error) {

	var rent Rent
	err := t.db.QueryRow("SELECT transport_id, tenant_id, rentType, owner_id FROM public.\"Rents\" WHERE id = $1", id).
		Scan(&rent.Transport_id, &rent.Tenant_id, &rent.RentType, &rent.Owner_id)
	if err != nil {
		return Rent{}, err
	}
	return rent, nil
}

func (t *RentService) DeleteByID(id int) error {
	_, err := t.db.Exec("DELETE FROM public.\"Rents\" WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
