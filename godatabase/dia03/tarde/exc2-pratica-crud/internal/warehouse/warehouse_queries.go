package warehouse

var (
	GetQuery = `SELECT id, name, address, telephone, capacity FROM warehouses`

	GetByIdQuery = `SELECT id, name, address, telephone, capacity FROM warehouses WHERE id = ?`

	InsertQuery = `INSERT INTO warehouses(name, address, telephone, capacity)
					VALUES (?, ?, ?, ?)`

	UpdateQuery = `UPDATE warehouses SET
    		name = ?,
    		address = ?,
    		telephone = ?,
    		capacity = ?
		WHERE id = ?`

	DeleteQuery = `DELETE FROM warehouses WHERE id = ?`

	GetReport = `SELECT w.name, COUNT(p.id) AS product_count
				FROM warehouses w 
				LEFT JOIN products p ON w.id = p.warehouse_id
				WHERE w.id = ?
				GROUP BY w.id, w.name`
)
