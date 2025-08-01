package customer

var (
	GetQuery = "SELECT id, first_name, last_name, `condition` FROM customers"

	GetTop5ActiveCustomersByTotalSpent = `
    SELECT
    c.first_name ,
    c.last_name,
    ROUND(SUM(i.total), 2) AS Amount
    FROM
    customers c
    INNER JOIN invoices i ON c.id = i.customer_id
    WHERE
    c.condition = '1'
    GROUP BY
    c.id, c.first_name, c.last_name
    ORDER BY
    Amount DESC
    LIMIT 5;
    `
	InsertQuery = "INSERT INTO customers (first_name, last_name, `condition`) VALUES (?, ?, ?)"

	GetTotalByCondition = `SELECT CASE 
        WHEN c.condition = '0' THEN 'Inactivo ( 0 )'
        WHEN c.condition = '1' THEN 'Activo ( 1 )'
        ELSE c.condition
    END AS ConditionNew,
    ROUND(SUM(i.total), 2) AS TotalAmount
    FROM
    customers c
    INNER JOIN invoices i ON c.id = i.id
    GROUP BY    
    c.condition
    ORDER BY
    c.condition;
`
)
