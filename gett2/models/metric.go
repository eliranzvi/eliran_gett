package models

import (
	"fmt"
)

// Error messages constants
const (
	INVALID_METRIC_NAME = "Metric name is invalid"
	FAILED_GETTING_DRIVER_METRICS = "Could not get driver [%s] metrics due to internal error"
	DRIVER_METRICS_DELETED_SUCCESSFULLY = "Driver [%s] metrics of type [%s] were deleted successfully"
	FAILED_ADDING_METRIC_FOR_DRIVER = "Driver [%s] metric of type [%s] was not added due to internal error"
	INSERT_DRIVER_METRIC_SUCCESSFULLY = "Driver [%s] metric of type [%s] was added successfully"
)

// DB statements constants
const (
	DROP_METRICS_TABLE_STMT = "DROP TABLE metrics"
	CREATE_METRICS_TABLE_STMT = "CREATE TABLE metrics(Metric_name varchar(100) NOT NULL, Value varchar(100) NOT NULL, Lat Numeric NOT NULL, Lon Numeric NOT NULL, Timestamp Integer NOT NULL, Driver_id varchar(100))"
	INSERT_METRICS_STMT       = "INSERT INTO metrics(Metric_name, Value, Lat, Lon, Timestamp, Driver_id) values ($1,$2,$3,$4,$5,$6);"
	DELETE_DRIVER_METRICS_STMT = "DELETE FROM metrics where Driver_id=$1"
	DELETE_ALL_METRICS = "Delete from metrics"
	QUERY_DRIVER_METRICS_BY_TYPE = "select * from metrics where Driver_id=$1 and Metric_name=$2 order by Timestamp DESC"
	DELETE_DRIVER_METRICS_BY_TYPE = "DELETE from metrics where Driver_id=$1 and Metric_name=$2"
)

// A struct representing a metric of a driver
type Metric struct {
	Metric_name string
	Value       string
	Lat         float32
	Lon         float32
	Timestamp   int
	Driver_id   string
}


func AddDriverMetric(metric Metric) (string) {
	var result string

	if len(metric.Driver_id) == 0 {
		result = INVALID_DRIVER_ID_ERROR
	} else if len(metric.Metric_name) == 0 {
		result = INVALID_METRIC_NAME
	} else {
		existingDriver, _ := GetDriver(metric.Driver_id)
		if existingDriver.Id == 0 {
			result = fmt.Sprintf(NON_EXISTING_DRIVER, metric.Driver_id)
		} else {
			db := openConnection()
			_, error := db.Exec(INSERT_METRICS_STMT, metric.Metric_name, metric.Value, metric.Lat, metric.Lon, metric.Timestamp, metric.Driver_id)
			closeConnection(db)

			if error != nil {
				result = fmt.Sprintf(FAILED_ADDING_METRIC_FOR_DRIVER, metric.Driver_id, metric.Metric_name)
			} else {
				result = fmt.Sprintf(INSERT_DRIVER_METRIC_SUCCESSFULLY, metric.Driver_id, metric.Metric_name)
			}
		}
	}

	return result
}

func GetDriverMetricsByType(driverId string, metricName string) ([]Metric, string) {
	var result string
	var metrics []Metric
	if len(driverId) == 0 {
		result = INVALID_DRIVER_ID_ERROR
	} else if len(metricName) == 0 {
		result = INVALID_METRIC_NAME
	} else {
		existingDriver, _ := GetDriver(driverId)
		if existingDriver.Id == 0 {
			result = fmt.Sprintf(NON_EXISTING_DRIVER, driverId)
		} else {
			db := openConnection()
			rows, error := db.Query(QUERY_DRIVER_METRICS_BY_TYPE, driverId, metricName)
			closeConnection(db)

			if error != nil {
				result = fmt.Sprintf(FAILED_GETTING_DRIVER_METRICS, driverId)
			} else {
				for rows.Next() {
					metric := new(Metric)
					rows.Scan(&metric.Metric_name, &metric.Value, &metric.Lat, &metric.Lon, &metric.Timestamp, &metric.Driver_id)
					metrics = append(metrics, *metric)
				}
			}
		}
	}

	return metrics, result
}

func DeleteDriverMetricsByType(driverId string, metricName string) (string) {
	var err string

	if len(driverId) == 0 {
		err = INVALID_DRIVER_ID_ERROR
	} else if len(metricName) == 0 {
		err = INVALID_METRIC_NAME
	} else {
		existingDriver, _ := GetDriver(driverId)
		if existingDriver.Id == 0 {
			err = fmt.Sprintf(NON_EXISTING_DRIVER, driverId)
		} else {
			db := openConnection()
			_, error := db.Exec(DELETE_DRIVER_METRICS_BY_TYPE, driverId, metricName)
			closeConnection(db)

			if error != nil {
				err = fmt.Sprintf(FAILED_DELETING_DRIVER_METRICS, driverId)
			} else {
				err = fmt.Sprintf(DRIVER_METRICS_DELETED_SUCCESSFULLY, driverId, metricName)
			}
		}
	}

	return err
}

func DeleteAllMetrics() {
	db := openConnection()
	db.Exec(DELETE_ALL_METRICS)
	closeConnection(db)
}

func InsertMetrics(metrics []Metric) {
	if len(metrics) > 0 {
		db := openConnection()

		for _, metric := range metrics {
			_, error := db.Exec(INSERT_METRICS_STMT, metric.Metric_name, metric.Value, metric.Lat, metric.Lon, metric.Timestamp, metric.Driver_id)

			checkErr(error)
		}

		closeConnection(db)
	}
}