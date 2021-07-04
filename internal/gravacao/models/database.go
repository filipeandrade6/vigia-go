package models

import pb "github.com/filipeandrade6/vigia-go/internal/api/v1"

type Database struct {
	Host         string
	Port         int32
	User         string
	Password     string
	DBName       string
	PoolMaxConns int32
}

func (d *Database) FromProtobuf(db *pb.DatabaseConfigResp) {
	d.Host = db.GetHost()
	d.Port = db.GetPort()
	d.User = db.GetUser()
	d.Password = db.GetPassword()
	d.DBName = db.GetDbname()
	d.PoolMaxConns = db.GetPoolmaxconns()
}
