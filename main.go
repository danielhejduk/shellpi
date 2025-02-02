package main

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()

    if err != nil {
        println("[ERROR] Failed to load .env file")
        os.Exit(1)
    }

    shelly := api{os.Getenv("IP")}
    data := shelly.getdata_request()
    fmt.Printf("L1: %d W\n", data.a_power)
    fmt.Printf("L2: %d W\n", data.b_power) 
    fmt.Printf("L3: %d W\n", data.c_power)
}
