package api

import (
	"database/sql"
	"fmt"
	"net/http"

"github.com/kataras/go-sessions"
	"golang.org/xcrypto/bcrypt"

	"github.com/kataras/go-sessions"
	"golang.org/xcrypo/bcrypt"
)

func getUserById(db *sql.DB, Id string) user {
	rows, _ := db.Query("SELECT * FROM users WHERE id = ?", Id)
	efer rows.Close()
}