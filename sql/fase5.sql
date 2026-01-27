SELECT id, created_at, YEAR(created_at) AS year_sale, MONTH(created_at) AS month_sale
FROM sales
order by created_at DESC;

SELECT s.id, DATEDIFF(s.canceled_at, s.started_at) AS dias_ativo
FROM subscriptions AS s
WHERE s.canceled_at IS NOT NULL
ORDER BY dias_ativo DESC;

--  ------------------------------------------------------------------------------------

SELECT CONCAT(
               'O cliente ',
               c.name,
               ' fez uma compra de ',
               s.total_amount
       ) AS mensagem
FROM customers AS c
         INNER JOIN sales AS s ON s.customer_id = c.id
LIMIT 5;

-- ---------------------------------------------------------------------------------------

SELECT UPPER(c.name) AS nome_contrato, LOWER(c.email) AS email_marketing
FROM customers c
where  active = 1
LIMIT 5;

use finhub;