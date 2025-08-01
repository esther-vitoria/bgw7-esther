package sale

var (
	GetQuery = `SELECT id, quantity, product_id, invoice_id FROM sales`

	InsertQuery = `INSERT INTO sales (quantity, product_id, invoice_id) VALUES (?, ?, ?)`
)
