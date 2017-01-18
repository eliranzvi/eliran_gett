package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"gett2/models"
)

type DriversController struct {
	beego.Controller
}

// Returns a driver by driver id
func (c *DriversController) GetDriver() {
	driver, err := models.GetDriver(c.Ctx.Input.Param(ID_PARAMETER))
	buildResponse(err, c.Ctx, driver)
}

// Adds the given driver
func (c *DriversController) AddDriver() {
	var driver models.Driver
	json.Unmarshal(c.Ctx.Input.RequestBody, &driver)
	result := models.AddDriver(driver)
	buildResponse(result, c.Ctx, driver)
}

// Updates the given driver
func (c *DriversController) UpdateDriver() {
	var driver models.Driver
	json.Unmarshal(c.Ctx.Input.RequestBody, &driver)
	result := models.UpdateDriver(driver)
	c.Ctx.Output.Body([]byte(result))
}

// Deletes the given driver
func (c *DriversController) DeleteDriver() {
	result := models.DeleteDriver(c.Ctx.Input.Param(ID_PARAMETER))
	c.Ctx.Output.Body([]byte(result))
}