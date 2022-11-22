-- The table stores my stocks 


DROP TABLE IF EXISTS stocks;

CREATE TABLE stocks (

    id serial PRIMARY KEY,
    symbol text NOT NULL,
    num_shares integer NOT NULL,
    date_acquired date NOT NULL
);

-- Query and out put of reading values from vs code and updating the table


-- assets=> \dS stocks
--                   Table "public.stocks"
--     Column     |  Type   | Collation | Nullable | Default 
-- ---------------+---------+-----------+----------+---------
--  symbol        | text    |           | not null | 
--  num_shares    | integer |           | not null | 
--  date_acquired | date    |           | not null | 

-- assets=> \i my_assets.sql
-- psql:my_assets.sql:4: ERROR:  syntax error at or near "IF"
-- LINE 1: DROP TABLE stocks IF EXISTS;
--                           ^
-- psql:my_assets.sql:12: ERROR:  relation "stocks" already exists
-- assets=> \i my_assets.sql
-- DROP TABLE
-- CREATE TABLE
-- assets=> \dS stocks
-- assets=> \COPY stocks(symbol,num_shares,date_acquired) FROM 'stocks.txt' WITH DELIMITER ',';
-- COPY 5
-- assets=> SELECT *
-- assets-> FROM stocks;
--  id | symbol | num_shares | date_acquired 
-- ----+--------+------------+---------------
--   1 | APPLE  |      10000 | 2020-11-17
--   2 | AMD    |        500 | 2000-10-03
--   3 | AMZN   |        120 | 2011-11-11
--   4 | LULU   |         50 | 2019-12-21
--   5 | TSLA   |       5000 | 2022-01-01
-- (5 rows)






-- Query and out put of reading values from vs code and updating the table


-- assets=> \dS stocks
--                   Table "public.stocks"
--     Column     |  Type   | Collation | Nullable | Default 
-- ---------------+---------+-----------+----------+---------
--  symbol        | text    |           | not null | 
--  num_shares    | integer |           | not null | 
--  date_acquired | date    |           | not null | 

-- assets=> \i my_assets.sql
-- psql:my_assets.sql:4: ERROR:  syntax error at or near "IF"
-- LINE 1: DROP TABLE stocks IF EXISTS;
--                           ^
-- psql:my_assets.sql:12: ERROR:  relation "stocks" already exists
-- assets=> \i my_assets.sql
-- DROP TABLE
-- CREATE TABLE
-- assets=> \dS stocks
-- assets=> \COPY stocks(symbol,num_shares,date_acquired) FROM 'stocks.txt' WITH DELIMITER ',';
-- COPY 5
-- assets=> SELECT *
-- assets-> FROM stocks;
--  id | symbol | num_shares | date_acquired 
-- ----+--------+------------+---------------
--   1 | APPLE  |      10000 | 2020-11-17
--   2 | AMD    |        500 | 2000-10-03
--   3 | AMZN   |        120 | 2011-11-11
--   4 | LULU   |         50 | 2019-12-21
--   5 | TSLA   |       5000 | 2022-01-01
-- (5 rows)
