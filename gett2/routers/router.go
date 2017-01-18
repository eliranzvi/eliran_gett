package routers

import (
	"gett2/controllers"
	"github.com/astaxie/beego"
)

// rest endpoints constants
const (
	DRIVER_REST_ENDPOINT             = "/driver/"
	DRIVER_ID_REST_ENDPOINT          = "/driver/:id"
	DRIVER_METRIC_NAME_REST_ENDPOINT = "/driver/:id/metrics/:metricName"
	DRIVER_METRICS_REST_ENDPOINT     = "/driver/:id/metric"
	INITIALIZE_REST_ENDPOINT         = "/initialize"
)

func init() {
	beego.Router(DRIVER_REST_ENDPOINT, &controllers.DriversController{}, "post:AddDriver;put:UpdateDriver")
	beego.Router(DRIVER_ID_REST_ENDPOINT, &controllers.DriversController{}, "get:GetDriver;delete:DeleteDriver")
	beego.Router(DRIVER_METRIC_NAME_REST_ENDPOINT, &controllers.MetricsController{}, "get:GetDriverMetricsByType;delete:DeleteDriverMetricsByType")
	beego.Router(DRIVER_METRICS_REST_ENDPOINT, &controllers.MetricsController{}, "post:AddDriverMetric")
	beego.Router(INITIALIZE_REST_ENDPOINT, &controllers.InitializationController{}, "post:Initialization")
}
