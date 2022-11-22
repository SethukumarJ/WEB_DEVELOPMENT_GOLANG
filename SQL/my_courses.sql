-- This is a demo for creating a table using an '.sql' file

DROP TABLE IF EXISTS my_courses;
 
CREATE TABLE my_courses (
    course_id text NOT NULL,
    created_at time(0) with time zone NOT NULL DEFAULT NOW()
);

INSERT INTO my_courses (course_id)
VALUES ('CMPS3162');


--Query to read the data from vs code is:

    
    --#\i 