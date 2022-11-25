--A demonstration table
DROP TABLE IF EXISTS lecturer CASCADE;
CREATE TABLE lecturer (
    id serial PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    gender text NOT NULL
);

--A demonstration table
DROP TABLE IF EXISTS course;
CREATE TABLE course (
    id serial PRIMARY KEY,
    name text NOT NULL,
    lecturer_id integer REFERENCES lecturer(id) ON DELETE CASCADE 

);

INSERT INTO lecturer (first_name,last_name,gender)
VALUES
('Jane','Smith','female'),
('Jill','Sosa','female'),
('Jashe','sullivan','man'),
('Joseph','Suggs','male');

INSERT INTO course (name, lecturer_id)
VALUES
('Introduction to Databases',1),
('Principles of Programming 1',1),
('Principles of Programming 2',1),
('Web Development',2),
('Principles of Programming 1',2),
('Principles of Programming 2',2),
('Introduction to Databases',3),
('Principles of Programming 1',3),
('Algorithms',4),
('Fundaments of Computing',1),
('Fundaments of Computing',2),
('Fundaments of Computing',3),
('Fundaments of Computing',4);



--lectures and their courses 
-- SELECT L.first_name, L.last_name, C.name
-- FROM lecturer AS L
-- INNER JOIN course AS C
-- ON L.id = C.lecturer_id
-- ORDER BY C.name;

-- SELECT L.first_name, L.last_name, C.name
-- FROM lecturer AS L
-- INNER JOIN course AS C
-- ON L.id = C.lecturer_id;


-- -- lecturer that teach Introduction to Databases
-- SELECT L.first_name, L.last_name, C.name
-- FROM lecturer AS L
-- INNER JOIN course AS C
-- ON L.id = C.lecturer_id
-- WHERE C.name = 'Introduction to Databases';

-- -- lecturer that teach Introduction to Databases and funcdaametals of computing

-- SELECT L.first_name, L.last_name
-- FROM lecturer AS L
-- INNER JOIN course AS C
-- ON L.id = C.lecturer_id
-- WHERE C.name = 'Introduction to Databases'
-- AND C.name = 'Fundaments of Computing';

-- SELECT L1.first_name , L2.first_name
-- FROM lecturer AS L1, lecturer AS L2;

-- SELECT L.first_name, L.last_name, C1.name, C2.name
-- FROM lecturer AS L
-- INNER JOIN course AS C1
-- ON L.id = C1.lecturer_id
-- INNER JOIN course AS C2
-- ON L.id = C2.lecturer_id
-- WHERE C1.name = 'Introduction to Databases'
-- AND C2.name = 'Fundaments of Computing';



-- week5=> SELECT *
-- week5-> FROM lecturer;




--  id | first_name | last_name | gender 
-- ----+------------+-----------+--------
--   1 | Jane       | Smith     | female
--   2 | Jill       | Sosa      | female
--   3 | Jashe      | sullivan  | man
--   4 | Joseph     | Suggs     | male
-- (4 rows)

-- week5=> SELECT *
-- week5-> FROM course;
--  id |            name             | lecturer_id 
-- ----+-----------------------------+-------------
--   1 | Introduction to Databases   |           1
--   2 | Principles of Programming 1 |           1
--   3 | Principles of Programming 2 |           1
--   4 | Web Development             |           2
--   5 | Principles of Programming 1 |           2
--   6 | Principles of Programming 2 |           2
--   7 | Introduction to Databases   |           3
--   8 | Principles of Programming 1 |           3
--   9 | Algorithms                  |           4
--  10 | Fundaments of Computing     |           1
--  11 | Fundaments of Computing     |           2
--  12 | Fundaments of Computing     |           3
--  13 | Fundaments of Computing     |           4
-- (13 rows)

-- week5=> SELECT name
-- week5-> FROM course;
--             name             
-- -----------------------------
--  Introduction to Databases
--  Principles of Programming 1
--  Principles of Programming 2
--  Web Development
--  Principles of Programming 1
--  Principles of Programming 2
--  Introduction to Databases
--  Principles of Programming 1
--  Algorithms
--  Fundaments of Computing
--  Fundaments of Computing
--  Fundaments of Computing
--  Fundaments of Computing
-- (13 rows)

-- week5=> SELECT DISTINCT name
-- week5-> FROM course;
--             name             
-- -----------------------------
--  Fundaments of Computing
--  Introduction to Databases
--  Web Development
--  Algorithms
--  Principles of Programming 2
--  Principles of Programming 1
-- (6 rows)

-- week5=> SELECT COUNT(DISTINC)T name
-- FROM course;
-- ERROR:  syntax error at or near "name"
-- LINE 1: SELECT COUNT(DISTINC)T name
--                                ^
-- week5=> SELECT COUNT(DISTINCT) name
-- FROM course;
-- ERROR:  syntax error at or near ")"
-- LINE 1: SELECT COUNT(DISTINCT) name
--                              ^
-- week5=> SELECT COUNT(DISTINCT name)
-- FROM course;
--  count 
-- -------
--      6
-- (1 row)

-- week5=> SELECT name 
-- week5-> FROM course
-- week5-> WHERE name = 
-- week5-> 
-- week5-> 'Algorithms
-- week5'> 
-- week5'> ';
--  name 
-- ------
-- (0 rows)

-- week5=> SELECT *
-- week5-> FROM course 
-- week5-> WHERE name = 'Algorithms';
--  id |    name    | lecturer_id 
-- ----+------------+-------------
--   9 | Algorithms |           4
-- (1 row)

-- week5=> SELECT name
-- week5-> FROM course
-- week5-> GROUP BY name;
--             name             
-- -----------------------------
--  Fundaments of Computing
--  Introduction to Databases
--  Web Development
--  Algorithms
--  Principles of Programming 2
--  Principles of Programming 1
-- (6 rows)

-- week5=> SELECT name, COUNT(name)
-- FROM course
-- GROUP BY name;
--             name             | count 
-- -----------------------------+-------
--  Fundaments of Computing     |     4
--  Introduction to Databases   |     2
--  Web Development             |     1
--  Algorithms                  |     1
--  Principles of Programming 2 |     2
--  Principles of Programming 1 |     3
-- (6 rows)

-- week5=> SELECT lecturer_id,name,COUT(name)
-- FROM course 
-- GROUP BY name;
-- ERROR:  function cout(text) does not exist
-- LINE 1: SELECT lecturer_id,name,COUT(name)
--                                 ^
-- HINT:  No function matches the given name and argument types. You might need to add explicit type casts.
-- week5=> SELECT lecturer_id,name,COUNT(name)
-- FROM course
-- GROUP BY name;
-- ERROR:  column "course.lecturer_id" must appear in the GROUP BY clause or be used in an aggregate function
-- LINE 1: SELECT lecturer_id,name,COUNT(name)
--                ^
-- week5=> SELECT lecturer_id,name,COUNT(name)
-- FROM course
-- GROUP BY name;
-- ERROR:  column "course.lecturer_id" must appear in the GROUP BY clause or be used in an aggregate function
-- LINE 1: SELECT lecturer_id,name,COUNT(name)
--                ^
-- week5=> SELECT id,name,COUNT(name)
-- FROM course
-- GROUP BY name;
-- ERROR:  column "course.id" must appear in the GROUP BY clause or be used in an aggregate function
-- LINE 1: SELECT id,name,COUNT(name)
--                ^
-- week5=> SELECT name, COUNT(name)
-- FROM course
-- GROUP BY name;
--             name             | count 
-- -----------------------------+-------
--  Fundaments of Computing     |     4
--  Introduction to Databases   |     2
--  Web Development             |     1
--  Algorithms                  |     1
--  Principles of Programming 2 |     2
--  Principles of Programming 1 |     3
-- (6 rows)

-- week5=> SELECT name, COUNT(name)
-- FROM course
-- GROUP BY name
-- week5-> HAVING COUNT(name) > 2;
--             name             | count 
-- -----------------------------+-------
--  Fundaments of Computing     |     4
--  Principles of Programming 1 |     3
-- (2 rows)


-- week5=> SELECT name, COUNT(name)
-- FROM couEXIT
--         SELECT name, COUNT(name)
-- FROM course
-- GROUP BY name 
-- HAVING COUNT(name) > 2
-- week5-> ORDER BY name ASC;
--             name             | count 
-- -----------------------------+-------
--  Fundaments of Computing     |     4
--  Principles of Programming 1 |     3
-- (2 rows)

-- week5=> SELECT name, COUNT(name)
-- FROM course
-- GROUP BY name
-- HAVING COUNT(name) > 2
-- ORDER BY name DESC;
--             name             | count 
-- -----------------------------+-------
--  Principles of Programming 1 |     3
--  Fundaments of Computing     |     4
-- (2 rows)

-- week5=> 








-- lecturer that teach Introduction to Databases and funcdaametals of computing






-- week5=> \i code.sql
-- psql:code.sql:2: NOTICE:  drop cascades to constraint course_lecturer_id_fkey on table course
-- DROP TABLE
-- CREATE TABLE
-- DROP TABLE
-- CREATE TABLE
-- INSERT 0 4
-- INSERT 0 13
-- psql:code.sql:49: ERROR:  relation "courses" does not exist
-- LINE 3: INNER JOIN courses AS C



SELECT L.first_name, L.last_name, C.name
FROM lecturer AS L
INNER JOIN course AS C
ON L.id = C.lecturer_id
ORDER BY C.name;

-- week5=> \i code.sql
-- psql:code.sql:2: NOTICE:  drop cascades to constraint course_lecturer_id_fkey on table course
-- DROP TABLE
-- CREATE TABLE
-- DROP TABLE
-- CREATE TABLE
-- INSERT 0 4
-- INSERT 0 13
--  first_name | last_name |            name             
-- ------------+-----------+-----------------------------
--  Joseph     | Suggs     | Algorithms
--  Joseph     | Suggs     | Fundaments of Computing
--  Jashe      | sullivan  | Fundaments of Computing
--  Jill       | Sosa      | Fundaments of Computing
--  Jane       | Smith     | Fundaments of Computing
--  Jashe      | sullivan  | Introduction to Databases
--  Jane       | Smith     | Introduction to Databases
--  Jashe      | sullivan  | Principles of Programming 1
--  Jill       | Sosa      | Principles of Programming 1
--  Jane       | Smith     | Principles of Programming 1
--  Jill       | Sosa      | Principles of Programming 2
--  Jane       | Smith     | Principles of Programming 2
--  Jill       | Sosa      | Web Development
-- (13 rows)


SELECT L.first_name, L.last_name, C.name
FROM lecturer AS L
INNER JOIN course AS C
ON L.id = C.lecturer_id;

-- week5=> \i code.sql
-- psql:code.sql:2: NOTICE:  drop cascades to constraint course_lecturer_id_fkey on table course
-- DROP TABLE
-- CREATE TABLE
-- DROP TABLE
-- CREATE TABLE
-- INSERT 0 4
-- INSERT 0 13
--  first_name | last_name |            name             
-- ------------+-----------+-----------------------------
--  Joseph     | Suggs     | Algorithms
--  Joseph     | Suggs     | Fundaments of Computing
--  Jashe      | sullivan  | Fundaments of Computing
--  Jill       | Sosa      | Fundaments of Computing
--  Jane       | Smith     | Fundaments of Computing
--  Jashe      | sullivan  | Introduction to Databases
--  Jane       | Smith     | Introduction to Databases
--  Jashe      | sullivan  | Principles of Programming 1
--  Jill       | Sosa      | Principles of Programming 1
--  Jane       | Smith     | Principles of Programming 1
--  Jill       | Sosa      | Principles of Programming 2
--  Jane       | Smith     | Principles of Programming 2
--  Jill       | Sosa      | Web Development
-- (13 rows)

--  first_name | last_name |            name             
-- ------------+-----------+-----------------------------
--  Jane       | Smith     | Introduction to Databases
--  Jane       | Smith     | Principles of Programming 1
--  Jane       | Smith     | Principles of Programming 2
--  Jill       | Sosa      | Web Development
--  Jill       | Sosa      | Principles of Programming 1
--  Jill       | Sosa      | Principles of Programming 2
--  Jashe      | sullivan  | Introduction to Databases
--  Jashe      | sullivan  | Principles of Programming 1
--  Joseph     | Suggs     | Algorithms
--  Jane       | Smith     | Fundaments of Computing
--  Jill       | Sosa      | Fundaments of Computing
--  Jashe      | sullivan  | Fundaments of Computing
--  Joseph     | Suggs     | Fundaments of Computing
-- (13 rows)

-- lecturer that teach Introduction to Databases
SELECT L.first_name, L.last_name, C.name
FROM lecturer AS L
INNER JOIN course AS C
ON L.id = C.lecturer_id
WHERE C.name = 'Introduction to Databases';

--  first_name | last_name |           name            
-- ------------+-----------+---------------------------
--  Jane       | Smith     | Introduction to Databases
--  Jashe      | sullivan  | Introduction to Databases
-- (2 rows)



SELECT L.first_name, L.last_name
FROM lecturer AS L
INNER JOIN course AS C
ON L.id = C.lecturer_id
WHERE C.name = 'Introduction to Databases'
AND C.name = 'Fundaments of Computing';


--  first_name | last_name 
-- ------------+-----------
-- (0 rows)


SELECT L1.first_name , L2.first_name
FROM lecturer AS L1, lecturer AS L2;

-- week5=> \i code.sql
-- psql:code.sql:2: NOTICE:  drop cascades to constraint course_lecturer_id_fkey on table course
-- DROP TABLE
-- CREATE TABLE
-- DROP TABLE
-- CREATE TABLE
-- INSERT 0 4
-- INSERT 0 13
--  first_name | first_name 
-- ------------+------------
--  Jane       | Jane
--  Jane       | Jill
--  Jane       | Jashe
--  Jane       | Joseph
--  Jill       | Jane
--  Jill       | Jill
--  Jill       | Jashe
--  Jill       | Joseph
--  Jashe      | Jane
--  Jashe      | Jill
--  Jashe      | Jashe
--  Jashe      | Joseph
--  Joseph     | Jane
--  Joseph     | Jill
--  Joseph     | Jashe
--  Joseph     | Joseph
-- (16 rows)

SELECT L.first_name, L.last_name, C1.name, C2.name
FROM lecturer AS L
INNER JOIN course AS C1
ON L.id = C1.lecturer_id
INNER JOIN course AS C2
ON L.id = C2.lecturer_id
WHERE C1.name = 'Introduction to Databases'
AND C2.name = 'Fundaments of Computing';


week5=> \i code.sql
psql:code.sql:2: NOTICE:  drop cascades to constraint course_lecturer_id_fkey on table course
DROP TABLE
CREATE TABLE
DROP TABLE
CREATE TABLE
INSERT 0 4
INSERT 0 13
 first_name | last_name |           name            |          name           
------------+-----------+---------------------------+-------------------------
 Jane       | Smith     | Introduction to Databases | Fundaments of Computing
 Jashe      | sullivan  | Introduction to Databases | Fundaments of Computing
