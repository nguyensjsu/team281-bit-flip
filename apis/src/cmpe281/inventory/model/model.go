package model

import (
	"github.com/gocql/gocql"
	"time"
)

type StoreDetails struct {
	Id   gocql.UUID `json:"id" cql:"id"`
	Name string     `json:"name" cql:"name"`
}

type ProductDetails struct {
	StoreId  gocql.UUID `json:"-" cql:"store_id"`
	Id       gocql.UUID `json:"id" cql:"id"`
	Name     string     `json:"name" cql:"name"`
	Quantity string     `json:"quantity" cql:"quantity"`
	Size     string     `json:"size" cql:"size"`
}

type AllocationDetails struct {
	Id      gocql.UUID `json:"id" cql:"id"`
	Status  string     `json:"status" cql:"status"`
	Expires time.Time  `json:"expires" cql:"expires"`
}