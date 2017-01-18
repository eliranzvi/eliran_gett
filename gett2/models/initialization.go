package models

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"io"
	"log"
	"github.com/astaxie/beego"
)

func Initialize() {
	CreateTables()
	DeleteAllMetrics()
	DeleteAllDrivers()

	drivers := readDrivers()
	InsertDrivers(drivers)
	metrics := ReadMetrics()
	InsertMetrics(metrics)
}

func CreateTables() {
	db := openConnection()
	fmt.Println("# Creating tables")
	db.QueryRow(DROP_DRIVERS_TABLE_STMT)
	db.QueryRow(CREATE_DRIVERS_TABLE_STMT)
	db.QueryRow(DROP_METRICS_TABLE_STMT)
	db.QueryRow(CREATE_METRICS_TABLE_STMT)
	fmt.Println("# Finished creating tables")

	closeConnection(db)
}

func readDrivers() []Driver {
	driversFilePath, _ := beego.GetConfig("String", "DRIVER_FILE_PATH", "C:/projects/go/src/gett2/conf/drivers.json")
	file, e := ioutil.ReadFile(driversFilePath.(string))
	if e != nil {
		panic("Could not load drivers file")
		os.Exit(1)
	}

	var drivers []Driver
	json.Unmarshal(file, &drivers)
	return drivers
}

func ReadMetrics() []Metric {
	metricsFilePath, _ := beego.GetConfig("String", "METRICS_FILE_PATH", "C:/projects/go/src/gett2/conf/metrics.json")
	configFile, err := os.Open(metricsFilePath.(string))

	if err != nil {
		panic("Could not load metrics file")
		os.Exit(1)
	}

	jsonParser := json.NewDecoder(configFile)
	var metrics []Metric

	for {
		var metric Metric
		if err := jsonParser.Decode(&metric); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// If there is no driver id, it will not be considered as a valid metric
		if len(metric.Driver_id) > 0 {
			metrics = append(metrics, metric)
		}
	}

	return metrics
}
