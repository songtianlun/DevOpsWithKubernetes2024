package main

import (
    "crypto/rand"
    "fmt"
    "time"
)

// Note - NOT RFC4122 compliant
func pseudo_uuid() (uuid string) {
    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }

    uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

    return
}

func print_time_random_uuid() {
    dt := time.Now() 
    uuid := pseudo_uuid()
    fmt.Printf("%v - %v\n",dt.Format(time.RFC3339), uuid)
}

func main() {
    for {
        print_time_random_uuid()
        time.Sleep(5 * time.Second)
    }

}
