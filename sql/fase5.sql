SELECT id, created_at, YEAR(created_at) AS year_sale, MONTH(created_at) AS month_sale
FROM sales
order by created_at DESC;

SELECT s.id, DATEDIFF(s.canceled_at, s.started_at) AS dias_ativo
FROM subscriptions AS s
WHERE s.canceled_at IS NOT NULL
ORDER BY dias_ativo DESC;