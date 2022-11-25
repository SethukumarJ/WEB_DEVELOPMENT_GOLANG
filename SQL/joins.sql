 -- An INNER JOIN returns a new if the row matches in all the tables

--  SELECT S.student_id, S.first_name,S.last_name, G.student_id,G.grade 
--  FROM students AS S 
--  INNER JOIN grades AS G
--  ON S.student_id = G.student_id;

-- school=> \i joins.sql
--  student_id | first_name | last_name | student_id | grade 
-- ------------+------------+-----------+------------+-------
--       20221 | PRASATH    | pn        |      20221 | A
--       20222 | FASIL      | MUHAMMED  |      20222 | B+
--       20223 | SETHU      | KUMAR     |      20223 | D+
-- (3 rows)


  -- An LEFT OUTER JOIN determines what is in the first table but not in the second table

--  SELECT S.student_id, S.first_name,S.last_name, G.student_id,G.grade 
--  FROM students AS S 
--  LEFT OUTER JOIN grades AS G
--  ON S.student_id = G.student_id; 
 
--  school=> \i joins.sql
--  student_id | first_name | last_name | student_id | grade 
-- ------------+------------+-----------+------------+-------
--       20221 | PRASATH    | pn        |      20221 | A
--       20222 | FASIL      | MUHAMMED  |      20222 | B+
--       20223 | SETHU      | KUMAR     |      20223 | D+
--       20224 | AJAL       | CK        |            | 
-- (4 rows)

-- SELECT S.student_id, S.first_name,S.last_name, G.student_id,G.grade 
--  FROM students AS S 
--  LEFT OUTER JOIN grades AS G
--  ON S.student_id = G.student_id 
-- WHERE S.stuent_id IS NULL;

--  school=> \i joins.sql
--  student_id | first_name | last_name | student_id | grade 
-- ------------+------------+-----------+------------+-------
--       20224 | AJAL       | CK        |            | 
-- (1 row)

--An RIGHT OUTER JOIN determines what is in the SECOND table but not in the FIRST table

--  SELECT S.student_id, S.first_name,S.last_name, G.student_id,G.grade 
--  FROM students AS S 
--  RIGHT OUTER JOIN grades AS G
--  ON S.student_id = G.student_id; 

--  school=> \i joins.sql
--  student_id | first_name | last_name | student_id | grade 
-- ------------+------------+-----------+------------+-------
--       20221 | PRASATH    | pn        |      20221 | A
--       20222 | FASIL      | MUHAMMED  |      20222 | B+
--       20223 | SETHU      | KUMAR     |      20223 | D+
--             |            |           |      20225 | f
-- (4 rows)


--  SELECT S.student_id, S.first_name,S.last_name, G.student_id,G.grade 
--  FROM students AS S 
--  RIGHT OUTER JOIN grades AS G
--  ON S.student_id = G.student_id
--  WHERE S.student_id IS NULL;

--   student_id | first_name | last_name | student_id | grade 
-- ------------+------------+-----------+------------+-------
--             |            |           |      20225 | f
-- (1 row)


--An FULL OUTER JOIN determines what is in the SECOND table but not in the FIRST table and vice versa

--  SELECT S.student_id, S.first_name,S.last_name, G.student_id,G.grade 
--  FROM students AS S 
--  FULL OUTER JOIN grades AS G
--  ON S.student_id = G.student_id; 
--  school=> \i joins.sql
--  student_id | first_name | last_name | student_id | grade 
-- ------------+------------+-----------+------------+-------
--       20221 | PRASATH    | pn        |      20221 | A
--       20222 | FASIL      | MUHAMMED  |      20222 | B+
--       20223 | SETHU      | KUMAR     |      20223 | D+
--             |            |           |      20225 | f 
--       20224 | AJAL       | CK        |            | 
-- (5 rows)


--  SELECT S.student_id, S.first_name,S.last_name, G.student_id,G.grade 
--  FROM students AS S 
--  FULL OUTER JOIN grades AS G
--  ON S.student_id = G.student_id
--  WHERE S.student_id IS NULL OR G.student_id IS NULL;

-- school=> \i joins.sql
--  student_id | first_name | last_name | student_id | grade 
-- ------------+------------+-----------+------------+-------
--             |            |           |      20225 | f
--       20224 | AJAL       | CK        |            | 
-- (2 rows)
