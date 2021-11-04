package controllers

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"spaceclan1/spaceclan-api/config"
	"spaceclan1/spaceclan-api/dao"
	models "spaceclan1/spaceclan-api/models/actions"
	"time"
)

var (
	HeroestakingController = &heroestaking_controller{}
)

type heroestaking_controller struct {
	main_controller
}

func (c heroestaking_controller) FetchPoolIncreaseTransactions() {
	k := "heroestaking_date"
	od := dao.OptionsImpl.Get(k)
	ol := dao.OptionsImpl.Get("heroestaking_limit")
	url := fmt.Sprintf(config.BuildUrl("get_actions", ol.Value, od.Value))
	r := models.ActionRes{}
	err := c.fetchUrl(url, &r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(url)
	dao.HeroeStakingTransactionsImpl.CreateBulk(r.Actions)
	a := r.Actions[len(r.Actions)-1]
	t, err2 := time.Parse("2006-01-02T15:04:05.000", a.Timestamp)
	if err2 != nil {
		log.Fatal(err2)

	}
	then := t.Add(time.Duration(-1) * time.Second)
	dao.OptionsImpl.Set(k, then.Format("2006-01-02T15:04:05"))
}
