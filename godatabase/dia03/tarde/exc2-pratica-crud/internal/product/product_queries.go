package product

var (
	GetQuery = `SELECT id, name, quantity, code_value, is_published, expiration, price, warehouse_id FROM products`

	GetByIdQuery = `SELECT id, name, quantity, code_value, is_published, expiration, price, warehouse_id FROM products WHERE id = ?`

	InsertQuery = `INSERT INTO products(name, quantity, code_value, is_published, expiration, price, warehouse_id)
					VALUES (?, ?, ?, ?, ?, ?, ?)`

	UpdateQuery = `UPDATE products SET
    		name = ?,
    		quantity = ?,
    		code_value = ?,
    		is_published = ?,
    		expiration = ?,
    		price = ?,
			warehouse_id = ?
		WHERE id = ?`

	DeleteQuery = `DELETE FROM products WHERE id = ?`
)
