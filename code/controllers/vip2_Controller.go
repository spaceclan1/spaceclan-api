package controllers

import (
	"github.com/gin-gonic/gin"
	"spaceclan1/spaceclan-api/dao"
)

var (
	Vip2Controller = &vip_controller{}
)

type vip_controller struct {
}

func (v vip_controller) GetMonthlyRewards(c *gin.Context) {

}

func (v vip_controller) GetDeposits(c *gin.Context) {
	dao.Cacher.GetDailyRewardFromCache()
}
