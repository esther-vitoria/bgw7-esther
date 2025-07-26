package product

var (
	GetQuery = `SELECT id, name, quantity, code_value, is_published, expiration, price FROM products`

	GetByIdQuery = `SELECT id, name, quantity, code_value, is_published, expiration, price FROM products WHERE id = ?`

	InsertQuery = `INSERT INTO products(name, quantity, code_value, is_published, expiration, price)
					VALUES (?, ?, ?, ?, ?, ?)`

	UpdateQuery = `UPDATE products SET
    		name = ?,
    		quantity = ?,
    		code_value = ?,
    		is_published = ?,
    		expiration = ?,
    		price = ?
		WHERE id = ?`

	DeleteQuery = `DELETE FROM products WHERE id = ?`
)
