package main_

import (
   "encoding/csv"
   "fmt"
   "os"
)

func main() {
   // Open the CSV file
   file, err := os.Open("checkings_short.csv")
   if err != nil {
       panic(err)
   }
   defer file.Close()

   fmt.Println(file)