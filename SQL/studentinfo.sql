--This is an example of a one to many relationship in database 
-- the one table (student) can have many rows in the many table(address)
-- A student can have many addresses.

DROP TABLE IF EXISTS student CASCADE;
CREATE TABLE student (
    id serial PRIMARY KEY,
    first_name text,
    last_name text
);

DROP TABLE IF EXISTS address;
CREATE TABLE address (
    id serial PRIMARY KEY,
    name text,
    student_id integer REFERENCES student(id) ON DELETE CASCADE
);

--Step #1: add data to the student table
INSERT INTO student(first_name, last_name)
VALUES
('prasanth','pn'),
('fasil','muhammed');

--Step #2: add data to the address table
INSERT INTO ADDRESS(name, student_id)
VALUES
('#3 ATTAPPADI', 1),
('#8 PALAKKAD', 1),
('#434 NILAMBOOR', 2);

--Get sudent information
SELECT S.first_name, S.last_name, A.name
FROM student AS S
INNER JOIN address AS A
ON S.id = A.student_id;

--  first_name | last_name |     name     
-- ------------+-----------+--------------
--  prasanth   | pn        | #3 ATTAPPADI
--  prasanth   | pn        | #3 PALAKKAD
--  fasil      | muhammed  | #3 NILAMBOOR
-- (3 rows)

-- week4=> SELECT *
-- FROM student;
--  id | first_name | last_name 
-- ----+------------+-----------
--   1 | prasanth   | pn
--   2 | fasil      | muhammed
-- (2 rows)

-- week4=> SELECT *
-- FROM 
-- address;
--  id |      name      | student_id 
-- ----+----------------+------------
--   1 | #3 ATTAPPADI   |          1
--   2 | #8 PALAKKAD    |          1
--   3 | #434 NILAMBOOR |          2
-- (3 rows)


-- week4=> DELETE FROM student
-- week4-> WHERE id=1;
-- DELETE 1
-- week4=> SELECT *
-- week4-> FROM student;
--  id | first_name | last_name 
-- ----+------------+-----------
--   2 | fasil      | muhammed
-- (1 row)

-- week4=> SELECT *
-- FROM address;
--  id |      name      | student_id 
-- ----+----------------+------------
--   3 | #434 NILAMBOOR |          2
-- (1 row)





