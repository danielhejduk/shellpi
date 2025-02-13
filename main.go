package main

import (
    "fmt"
    "os"
    "time"

    "github.com/joho/godotenv"
)

const (
    CREATE_TABLE = "CREATE TABLE IF NOT EXISTS hourly_readings (date TEXT NOT NULL, time TEXT NOT NULL, l1 INTEGER NOT NULL, l2 INTEGER NOT NULL, l3 INTEGER NOT NULL, total INTEGER NOT NULL, l1_hour INTEGER NOT NULL, l2_hour INTEGER NOT NULL, l3_hour INTEGER NOT NULL, total_hour NOT NULL);"
)

func main() {
    err := godotenv.Load()

    if err != nil {
        println("[ERROR] Failed to load .env file")
        os.Exit(1)
    }

    shelly := api{os.Getenv("IP")}

    ticker := time.NewTicker(1 * time.Hour)
    defer ticker.Stop()

    go func() {
        for {
            select {
                case <-ticker.C:
                    // TODO
            }
        }
    }()

    data := shelly.getdata_request()
    fmt.Printf("L1: %d W\n", data.a_power)
    fmt.Printf("L2: %d W\n", data.b_power) 
    fmt.Printf("L3: %d W\n", data.c_power)
}
