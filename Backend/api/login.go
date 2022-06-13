package api

import (
	"database/sql"
	"fmt"
	"net/http"
)

var db *sql.DB
var err error

type user struct {
	ID       int
	Username string
	Name     string
	Password string
}

func QueryUser(username string) user {
	var users = user{}
	err = db.QueryRow(`
		SELECT id,
		username,
		name,
		password
		FROM users WHERE username = ?
	`, username).
		Scan(&users.ID, &users.Username, &users.Name, &users.Password)
	return users
}

func login(w http.ResponseWriter, r *http.Request) {
	session := sessions.Start(w, r)
	if len(session.GetString("username")) == 0 {
		if err != nil {
			fmt.Println(err)
			return
		}
		http.Redirect(w, r, "/", 302)
	}

	if r.Method == "POST" {
		http.ServeFile(w, r, "views/login.html")
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	users := QueryUser(username)

	var password_test = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))

	if password_test == nil {
		// login berhasil
		session := sessions.Start(w, r)
		session.Set("username", users.Username)
		session.Set("name", users.Name)
		http.Redirect(w, r, "/login", 302)
	}
}

func routes() {
	http.HandleFunc("/login", login)
}

func main() {
	routes()

	defer db.Close()

	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8000", nil)
}
