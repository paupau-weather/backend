package main

import (
	"database/sql"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/joho/godotenv" // Импортируем пакет для работы с .env файлами
	_ "github.com/lib/pq"      // Импортируем драйвер PostgreSQL
)

// init() вызывается до main()
func init() {
	// Загружаем значения из .env в систему
	if err := godotenv.Load(); err != nil {
		fmt.Print("No .env file found")
	}
}

func getPsqlInfo() string {
	dbAddress, dbAddressExists := os.LookupEnv("DB_ADDRESS")
	dbPort, dbPortExists := os.LookupEnv("DB_PORT")
	dbUser, dbUserExists := os.LookupEnv("DB_USER")
	dbPassword, dbPasswordExists := os.LookupEnv("DB_PASSWORD")
	dbName, dbNameExists := os.LookupEnv("DB_NAME")

	if !dbAddressExists || !dbPortExists || !dbUserExists || !dbPasswordExists || !dbNameExists {
		fmt.Println("Error: .env file is not correct")
		os.Exit(1)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbAddress, dbPort, dbUser, dbPassword, dbName)
	return psqlInfo
}

func getListenerInfo() string {
	listenerAddress, listenerAddressExists := os.LookupEnv("LISTENER_ADDRESS")
	listenerPort, listenerPortExists := os.LookupEnv("LISTENER_PORT")

	if !listenerAddressExists || !listenerPortExists {
		fmt.Println("Error: .env file is not correct")
		os.Exit(1)
	}

	listenerInfo := fmt.Sprintf("%s:%s", listenerAddress, listenerPort)
	return listenerInfo
}

func main() {
	psqlInfo := getPsqlInfo()

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

	listenerInfo := getListenerInfo()
	listener, err := net.Listen("tcp", listenerInfo)

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
