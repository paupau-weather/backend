package main

import (
	"fmt"
	"io"
	"net"
	"database/sql"

	_ "github.com/lib/pq" // Импортируем драйвер PostgreSQL
)

const (
	host     = "localhost"
 	port     = 5432
 	user     = "postgres"
 	password = "1234"
 	dbname   = "HomeStation"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			fmt.Println("Error open: ", err)
			db.Close()
		}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error ping: ", err)
	}

	listener, err := net.Listen("tcp", "192.168.0.100:8080")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server listening...")
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn, db)
	}
}

func handleConnection(conn net.Conn, db *sql.DB) {
	defer conn.Close()

	buffer := make([]byte, 1024*2)
	var temperature float32
	var humidity float32
	var pressure float32
	var CO2 float32
	var id_station int

	for {
		n, err := conn.Read(buffer)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by client")
			} else {
				fmt.Println("Error: ", err)
			}
			return
		}

		if n > 0 {
			data := string(buffer[:n])
			_, err := fmt.Sscanf(data, "id_station=%d,temperature=%f,humidity=%f,pressure=%f,CO2=%f", &id_station, &temperature, &humidity, &pressure, &CO2)
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			fmt.Printf("id_station: %.06d, Temperature: %.2f, Humidity: %.2f, Pressure: %.2f, CO2: %.2f\n", id_station, temperature, humidity, pressure, CO2)
			query := "INSERT INTO measurement (station_id, temperature, humidity, pressure, \"CO2\") VALUES ($1, $2, $3, $4, $5)"
			_, err = db.Exec(query, id_station, temperature, humidity, pressure, CO2)
			if err != nil {
				fmt.Println("Error query: ", err)
				continue
			} else {
				fmt.Println("Database inserted successfully")
			}
		}
	}
}
