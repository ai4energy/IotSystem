package model

import (
	"database/sql"
	"log"
	
	_"modernc.org/sqlite"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite", "./device.db")
	if err != nil {
		log.Fatal(err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS device_info (
		num INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL UNIQUE,
		value REAL NOT NULL
	);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

type SensorData struct {
	Num    int     `json:"num"`
	Name  string  `json:"name"`
	Value int `json:"value"`
}

func GetSensors() ([]SensorData, error) {
	rows, err := db.Query("SELECT num, name, value FROM device_info;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sensors []SensorData
	for rows.Next() {
		var sensor SensorData
		if err := rows.Scan(&sensor.Num, &sensor.Name, &sensor.Value); err != nil {
			return nil, err
		}
		sensors = append(sensors, sensor)
	}

	return sensors, nil
}

func AddSensor(sensor SensorData) (int64, error) {
	result, err := db.Exec("INSERT INTO device_info (name, value) VALUES (?, ?)", sensor.Name, sensor.Value)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func UpdateSensor(name string, value int) error {
	_, err := db.Exec("UPDATE device_info SET value=? WHERE name=?;", value, name)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSensor(name string) error {
	_, err := db.Exec("DELETE FROM device_info WHERE name=?", name)
	if err != nil {
		return err
	}
	return nil
}