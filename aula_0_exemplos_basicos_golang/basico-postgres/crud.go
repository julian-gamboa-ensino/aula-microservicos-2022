package main

import (
    "html/template"
    "net/http"
    "strconv"

    _ "github.com/lib/pq"
)


func HomeHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT item FROM items")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var items []string
    for rows.Next() {
        var item string
        if err := rows.Scan(&item); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        items = append(items, item)
    }

    tmpl := template.Must(template.ParseFiles("templates/home.html"))
    tmpl.Execute(w, items)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseForm()
        item := r.FormValue("item")

        _, err := db.Exec("INSERT INTO items (item) VALUES ($1)", item)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseForm()
        indexStr, newItem := r.FormValue("index"), r.FormValue("newItem")
        index, err := strconv.Atoi(indexStr)
        if err != nil {
            http.Error(w, "Index conversion error", http.StatusBadRequest)
            return
        }

        _, err = db.Exec("UPDATE items SET item = $1 WHERE id = $2", newItem, index)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseForm()
        indexStr := r.FormValue("index")
        index, err := strconv.Atoi(indexStr)
        if err != nil {
            http.Error(w, "Index conversion error", http.StatusBadRequest)
            return
        }

        _, err = db.Exec("DELETE FROM items WHERE id = $1", index)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/", http.StatusSeeOther) // Correção aplicada aqui
    }
}

