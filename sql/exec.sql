
select * from tenants;


select name, default_currency from tenants;



select * from products;
select name, email from users where consume_api = 1;


SELECT id, total_amount FROM sales ORDER BY total_amount DESC LIMIT 5;


select * from subscriptions where status  = 'active' and billing_time = 'monthly' limit 100;

select id, total_amount from sales where created_at BETWEEN  '2023-01-01 00:00:00' AND '2023-03-31 23:59:59';