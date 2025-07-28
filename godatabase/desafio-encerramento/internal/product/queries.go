package product

var (
	GetQuery = `SELECT id, description, price FROM products`

	GetTop5SoldProducts = `SELECT p.description, SUM(s.quantity) AS total_sold
    FROM products p
    INNER JOIN sales s ON p.id = s.product_id
    GROUP BY p.id, p.description
    ORDER BY total_sold DESC
    LIMIT 5`

	InsertQuery = `INSERT INTO products (description, price) VALUES (?, ?)`
)
