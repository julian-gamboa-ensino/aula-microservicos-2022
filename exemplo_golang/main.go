package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    _ "github.com/lib/pq"
)

var db *sql.DB

func main() {
    initDB()
    defer db.Close()

    http.HandleFunc("/", HomeHandler)           // Serve a página principal e exibe os itens
    http.HandleFunc("/create", CreateHandler)   // Manipula a criação de novos itens
    http.HandleFunc("/update", UpdateHandler)   // Manipula a atualização de itens existentes
    http.HandleFunc("/delete", DeleteHandler)   // Manipula a exclusão de itens

    log.Println("Listening on :8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}


func initDB() {
    var err error
    dbHost := os.Getenv("DB_HOST") // 
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")

    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }

    // Apenas pingue o banco de dados para testar a conexão quando tiver certeza que o banco de dados está pronto.
    // Isso pode exigir uma lógica de espera (retry logic) ou um script de espera.
    if err = db.Ping(); err != nil {
        log.Fatalf("Database connection test failed: %v", err)
    }
}

