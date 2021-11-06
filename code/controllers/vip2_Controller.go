package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"spaceclan1/spaceclan-api/dao"
	"spaceclan1/spaceclan-api/utils/errors"
)

var (
	Vip2Controller = &vip_controller{}
)

type vip_controller struct {
}

func (v *vip_controller) GetPoolDeposits(c *gin.Context) {
	d, err := dao.Cacher.GetToFrom("heroestaking", "spaceheroes1")
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, errors.OK(&d))
}

func (v *vip_controller) GetPoolRewards(c *gin.Context) {

	wallet := c.Param("wallet")
	log.WithField("wallet", c.Params).Info()
	var d map[string]string
	var err error
	if wallet == "" {
		d, err = dao.Cacher.GetRewards("TOTAL")
	} else {
		d, err = dao.Cacher.GetRewards(wallet)
	}
	if err != nil {
		c.JSON(500, errors.NewInternalRequestError(err.Error()))
		return
	}
	c.JSON(200, errors.OK(&d))
}
