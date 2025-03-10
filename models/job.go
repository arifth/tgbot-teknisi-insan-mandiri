package models

import "github.com/google/uuid"

type Job struct {
	ID                uuid.UUID
	MinuteHour        string
	Created_at        string
	Created_by        string
	Merek             string
	Type              string
	NoSeries          string
	ServiceType       string
	TindakanPerbaikan string
	ModfiedSparepart  string
	CounterMachine    string
	JobLocation       string
	Result            string
	TypeOfWork        string
	Points            uint8
}
