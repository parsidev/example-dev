package main

import (
   "database/sql"
   "fmt"
   "log"
   "net/http"

   _ "github.com/lib/pq"
)

func main() {
   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      db, err := sql.Open("postgres", "user=user password=password dbname=mydb sslmode=disable")
      if err != nil {
         log.Fatal(err)
      }
      defer db.Close()

      err = db.Ping()
      if err != nil {
         log.Fatal(err)
      }

      _, err = fmt.Fprintf(w, "Connected to PostgreSQL!")
      if err != nil {
         log.Fatal(err)
      }
   })

   log.Fatal(http.ListenAndServe(":8080", nil))
}
