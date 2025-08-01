package invoice

var (
	GetQuery = `SELECT id, datetime, total, customer_id FROM invoices`

	UpdateInvoiceTotals = `UPDATE invoices i
    SET i.total = (
        SELECT COALESCE(SUM(p.price * s.quantity), 0)
        FROM sales s
            INNER JOIN products p ON p.id = s.product_id
        WHERE s.invoice_id = i.id
    )`

	InsertQuery = `INSERT INTO invoices (datetime, total, customer_id) VALUES (?, ?, ?)`
)
