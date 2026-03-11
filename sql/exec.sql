
select * from tenants;


select name, default_currency from tenants;



select * from products;
select name, email from users where consume_api = 1;


SELECT id, total_amount FROM sales ORDER BY total_amount DESC LIMIT 5;


select * from subscriptions where status  = 'active' and billing_time = 'monthly' limit 100;

select id, total_amount from sales where created_at BETWEEN  '2023-01-01 00:00:00' AND '2023-03-31 23:59:59';

select id, name, email from customers where email like '%@gmail.com';

select id, name from plans where interval_unit in ('daily', 'weekly', 'ten_days');

select id, name from customers where tax_identification_number is not null;
select id, name from plans where interval_unit not in ('daily', 'weekly', 'ten_days');


select id, code, value from events
                       where value between 10000 and 50000
                         and code like 'A%'
                         and deleted_at is null order by value desc limit 10;
