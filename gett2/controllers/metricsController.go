package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"gett2/models"
)

const (
	METRIC_NAME_PARAMETER = ":metricName"
)


type MetricsController struct {
	beego.Controller
}

// Returns the metrics of the specified type for the specified driver
func (c *MetricsController) GetDriverMetricsByType() {
	metrics, err := models.GetDriverMetricsByType(c.Ctx.Input.Param(ID_PARAMETER), c.Ctx.Input.Param(METRIC_NAME_PARAMETER))
	buildResponse(err, c.Ctx, metrics)
}

// Deletes all metrics of specified type of specified driver
func (c *MetricsController) DeleteDriverMetricsByType() {
	result := models.DeleteDriverMetricsByType(c.Ctx.Input.Param(ID_PARAMETER), c.Ctx.Input.Param(METRIC_NAME_PARAMETER))
	c.Ctx.Output.Body([]byte(result))
}

// Adds a driver metric
func (c *MetricsController) AddDriverMetric() {
	var metric models.Metric
	json.Unmarshal(c.Ctx.Input.RequestBody, &metric)
	result := models.AddDriverMetric(metric)
	c.Ctx.Output.Body([]byte(result))
}
