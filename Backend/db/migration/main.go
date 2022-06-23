package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Run This Script for migration db
func main() {
	db, err := sql.Open("sqlite3", "../usedbooks.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id integer not null primary key AUTOINCREMENT,
		username varchar(255) not null,
		email varchar(255) not null,
		nohp varchar(255) not null,
		password varchar(255) not null,
		role TEXT DEFAULT ('user'),
		created_at datetime DEFAULT (datetime('now')),
		updated_at datetime DEFAULT (datetime('now'))
);

	CREATE TABLE IF NOT EXISTS users_address (
		id integer not null primary key AUTOINCREMENT,
		user_id int not null,
		address_line_1 varchar(255) not null,
		address_line_2 varchar(255),
		city varchar(255) not null,
		postal_code varchar(255) not null,
		FOREIGN KEY (user_id) REFERENCES users(id)
);

	CREATE TABLE IF NOT EXISTS users_payment (
		id integer not null primary key AUTOINCREMENT,
		user_id int not null,
		payment_type varchar(255) not null,
		provider varchar(255),
		FOREIGN KEY (user_id) REFERENCES users(id)
);


	CREATE TABLE IF NOT EXISTS product (
		id integer not null primary key AUTOINCREMENT,
		name varchar(255) not null,
		description text not null,
		price int not null,
		image varchar(255) not null
);

	CREATE TABLE IF NOT EXISTS product_inventory (
		id integer not null primary key AUTOINCREMENT,
		product_id int not null,
		FOREIGN KEY (product_id) REFERENCES product(id)

);

	CREATE TABLE IF NOT EXISTS cart_item (
		id integer not null primary key AUTOINCREMENT,
		product_id int not null,
		quantity int not null,
		FOREIGN KEY (product_id) REFERENCES product(id)
		
);

	CREATE TABLE IF NOT EXISTS shoping_session (
		id integer not null primary key AUTOINCREMENT,
		user_id int not null,
		amount int not null,
		FOREIGN KEY (user_id) REFERENCES users(id)

);

	CREATE TABLE IF NOT EXISTS order_items (
		id integer not null primary key AUTOINCREMENT,
		quantity int not null
);

	CREATE TABLE IF NOT EXISTS order_details (
		id integer not null primary key AUTOINCREMENT,
		order_item_id int not null,
		user_id int not null,
		amount int not null,
		FOREIGN KEY (order_item_id) REFERENCES order_items(id),
		FOREIGN KEY (user_id) REFERENCES users(id)

);

	CREATE TABLE IF NOT EXISTS payment_details (
		id integer not null primary key AUTOINCREMENT,
		order_detail_id int not null,
		amount int not null,
		provider varchar(255) not null,
		FOREIGN KEY (order_detail_id) REFERENCES order_details(id)

);

INSERT INTO users (username, email, nohp, password, role) VALUES
    ('vini','vini@gmail.com', '082268', '1234', 'admin'),
    ('atika','atika@gmail.com', '082269', '1234', 'admin'),
    ('Afif','afif@gmail.com', '082270', '1234', 'admin'),
    ('riski','riski@gmail.com', '082271', '1234', 'admin'),
    ('jasmine','jasmine@gmail.com', '082272', '1234', 'admin');

	INSERT INTO product (name, description, price, image) VALUES
	('MATEMATIKA', 'Buku khusus untuk anak sekolah dasar', 10000, '../images'),
	('Bahasa Indonesia', 'Buku khusus untuk anak sekolah dasar', 15000, '../images');

	INSERT INTO  cart_item(product_id, quantity) VALUES
	(1, 3),
	(2, 3);

`)

	if err != nil {
		panic(err)
	}
	defer db.Close()
}
