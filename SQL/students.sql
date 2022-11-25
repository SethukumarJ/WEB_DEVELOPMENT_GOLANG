DROP TABLE IF EXISTS students;
CREATE TABLE students (
    student_id integer PRIMARY KEY,
    first_name text,
    last_name text
);

INSERT INTO students(student_id,first_name,last_name)
VALUES
(20221,'PRASATH','pn'),
(20222,'FASIL','MUHAMMED'),
(20223,'SETHU','KUMAR');


DROP TABLE IF EXISTS grades;
CREATE TABLE grades (
    student_id integer PRIMARY KEY,
    grade text
);

INSERT INTO grades
VALUES
(20221,'A'),
(20222,'B+'),
(20223,'D+');