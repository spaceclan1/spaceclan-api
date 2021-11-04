package dao

import (
	log "github.com/sirupsen/logrus"
	"spaceclan1/spaceclan-api/datasource"
	"spaceclan1/spaceclan-api/models"
)

var (
	OptionsImpl = &options_impl{}
)

type options_impl struct {
}

func (oi options_impl) Get(k string) models.Option {
	o := models.Option{}
	err := datasource.MainDb.QueryRow("select `id`,`key`,`value` from wax_data.options o where o.key=?", k).Scan(&o.Id, &o.Key, &o.Value)
	if err != nil {
		log.Fatal(err)
	}
	return o
}

func (oi options_impl) Set(k string, v string) {
	stm, err := datasource.MainDb.Prepare("update wax_data.options o set o.value=? where o.key=?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stm.Exec(v, k)
	if err != nil {
		log.Fatal(err)
	}
}

func (oi options_impl) GetAll(k string) []models.Option {
	rows, err := datasource.MainDb.Query("select `id`,`key`,`value` from wax_data.options o where o.key=?", k)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	ops := make([]models.Option, 0)
	for rows.Next() {
		o := models.Option{}
		err := rows.Scan(&o.Id, &o.Key, &o.Value)
		if err != nil {
			log.Fatal(err)

		}
		ops = append(ops, o)
	}
	return ops
}
