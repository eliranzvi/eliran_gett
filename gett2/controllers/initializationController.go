package controllers

import (
	"github.com/astaxie/beego"
	"gett2/models"
)

type InitializationController struct {
	beego.Controller
}


// Returns a driver by driver id
func (c *InitializationController) Initialization() {
	models.Initialize()
	c.Ctx.Output.Body([]byte("Initialized successfully"))
}