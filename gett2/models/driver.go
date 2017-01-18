package models

import (
	"fmt"
	"strconv"
	"database/sql"
)

// Error messages constants
const (
	INVALID_DRIVER_ID_ERROR = "Id of driver is invalid"
	INVALID_DRIVER_NAME = "Name of driver should not be empty"
	INVALID_DRIVER_LICENSE_NUMBER = "License number of the driver must not be empty"
	NON_EXISTING_DRIVER = "A driver with id[%s] does not exist"
	DRIVER_ALREADY_EXISTS = "A driver with id[%s] already exists"
	DRIVER_ADDED_SUCCESSFULLY = "Driver with id[%s] was successfully added"
	DRIVER_DELETED_SUCCESSFULLY = "Driver with id [%s] was successfully deleted"
	DRIVER_UPDATED_SUCCESSFULLY = "Driver with id [%s] was successfully updated"
	FAILED_ADDING_DRIVER = "Could not add driver with id [%s] due to internal error"
	FAILED_DELETING_DRIVER_METRICS = "Could not delete driver metrics for driver with id [%s] due to internal error"
	FAILED_DELETING_DRIVER = "Could not delete driver with id [%s] due to internal error"
	FAILED_UPDATING_DRIVER = "Could not update driver with id [%s] due to internal error"
)

// DB statements constants
const (
	DROP_DRIVERS_TABLE_STMT = "DROP TABLE drivers"
	CREATE_DRIVERS_TABLE_STMT = "CREATE TABLE drivers (Id varchar(100) UNIQUE NOT NULL, Name varchar(100) NOT NULL, License_number VARCHAR(100) NOT NULL)"
	INSERT_DRIVERS_STMT       = "INSERT INTO drivers(Id, Name, License_number) values ($1, $2, $3)"
	QUERY_DRIVER_BY_ID_STMT   = "SELECT * from drivers where Id=$1";
	DELETE_DRIVER_STMT = "DELETE from drivers where Id=$1"
	DELETE_ALL_DRIVERS_STMT = "DELETE from drivers"
	UPDATE_DRIVER_STMT = "UPDATE drivers set Name=$1, License_number=$2 where Id=$3"
)

// A struct representing a driver
type Driver struct {
	Id             int
	Name           string
	License_number string
}

func GetDriver(driverId string) (Driver, string) {
	db := openConnection()
	rows, err := db.Query(QUERY_DRIVER_BY_ID_STMT, driverId)
	closeConnection(db)
	checkErr(err)

	var driver Driver
	if rows.Next() {
		driver := new(Driver)
		rows.Scan(&driver.Id, &driver.Name, &driver.License_number)
		return *driver, ""
	} else {
		return driver, fmt.Sprintf(NON_EXISTING_DRIVER, driverId)
	}
}

func AddDriver(driver Driver) (string) {
	var result string

	if driver.Id == 0 {
		result = INVALID_DRIVER_ID_ERROR
	} else if len(driver.Name) == 0 {
		result = INVALID_DRIVER_NAME
	} else if len(driver.License_number) == 0 {
		result = INVALID_DRIVER_LICENSE_NUMBER
	} else {
		driverId := strconv.Itoa(driver.Id)
		existingDriver, _ := GetDriver(driverId)

		if existingDriver.Id != 0 {
			result = fmt.Sprintf(DRIVER_ALREADY_EXISTS, driverId)
		} else {
			db := openConnection()
			result = InsertDriver(db, driver)
			closeConnection(db)
			if len(result) == 0 {
				result = fmt.Sprintf(DRIVER_ADDED_SUCCESSFULLY, driverId)
			} else {
				result = fmt.Sprintf(FAILED_ADDING_DRIVER, driverId)
			}
		}
	}
	return result
}

func UpdateDriver(driver Driver) string {
	var result string

	if driver.Id == 0 {
		result = INVALID_DRIVER_ID_ERROR
	} else if len(driver.Name) == 0 {
		result = INVALID_DRIVER_NAME
	} else if len(driver.License_number) == 0 {
		result = INVALID_DRIVER_LICENSE_NUMBER
	}  else {
		driverId := strconv.Itoa(driver.Id)
		existingDriver, _ := GetDriver(driverId)

		if existingDriver.Id == 0 {
			result = fmt.Sprintf(NON_EXISTING_DRIVER, driverId)
		} else {
			db := openConnection()
			_, err := db.Exec(UPDATE_DRIVER_STMT, driver.Name, driver.License_number, driver.Id)

			if err != nil {
				result = fmt.Sprintf(FAILED_UPDATING_DRIVER, driverId)
			} else {
				result = fmt.Sprintf(DRIVER_UPDATED_SUCCESSFULLY, driverId)
			}
		}
	}
	return result
}

func DeleteDriver(driverId string) string {
	var result string

	if len(driverId) == 0 {
		result = INVALID_DRIVER_ID_ERROR
	} else {
		existingDriver, _ := GetDriver(driverId)

		if existingDriver.Id != 0 {
			db := openConnection()
			_, err := db.Exec(DELETE_DRIVER_METRICS_STMT,  driverId)

			if err != nil {
				result = fmt.Sprintf(FAILED_DELETING_DRIVER_METRICS, driverId)
			} else {
				_, err := db.Exec(DELETE_DRIVER_STMT, driverId)
				if err != nil {
					result = fmt.Sprintf(FAILED_DELETING_DRIVER, driverId)
				} else {
					result = fmt.Sprintf(DRIVER_DELETED_SUCCESSFULLY, driverId)
				}
			}

			closeConnection(db)
		} else {
			result = fmt.Sprintf(NON_EXISTING_DRIVER, driverId)
		}
	}
	return result
}

func DeleteAllDrivers() {
	db := openConnection()
	db.Exec(DELETE_ALL_DRIVERS_STMT)
	closeConnection(db)
}

func InsertDrivers(drivers []Driver) {
	if len(drivers) > 0 {
		db := openConnection()
		for _, driver := range drivers {
			_ = InsertDriver(db, driver)
		}

		closeConnection(db)
	}
}

func InsertDriver(db *sql.DB, driver Driver) string {
	result, error := db.Exec(INSERT_DRIVERS_STMT, driver.Id, driver.Name, driver.License_number)
	var err string
	driverId := strconv.Itoa(driver.Id)
	if error != nil {
		err = fmt.Sprintf(FAILED_ADDING_DRIVER, driverId)
	}
	rowsEffected, _ := result.RowsAffected()
	if rowsEffected == 0 {
		err = fmt.Sprintf(FAILED_ADDING_DRIVER, driverId)
	}
	return err
}