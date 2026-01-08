SELECT p.name, COUNT(s.id) AS total_assinaturas
FROM plans AS p
INNER JOIN subscriptions AS s ON s.plan_id = p.id
WHERE s.deleted_at IS NULL
GROUP BY p.name
ORDER BY total_assinaturas DESC;

-- -------------------------------------------------
SELECT c.name, SUM(s.total_amount) AS total_amount
FROM customers AS c
INNER JOIN sales AS s on s.customer_id = c.id
WHERE s.deleted_at IS NULL
GROUP BY c.name
ORDER BY total_amount DESC
limit 3;

-- --------------------------------------------------
SELECT p.name, COUNT(s.id) AS sumplans
FROM plans AS p
INNER JOIN subscriptions AS s ON s.plan_id = p.id
WHERE s.deleted_at IS NULL
GROUP BY p.name
HAVING sumplans > 1;

-- ---------------------------------------------------------------------
SELECT c.name, SUM(s.total_amount) AS total_gasto
FROM customers AS c
INNER JOIN sales AS s ON s.customer_id = c.id
WHERE s.deleted_at IS NULL
GROUP BY c.name
HAVING total_gasto > 100000;
