-- https://github.com/ibmdb/go_ibm_db/blob/master/testdata/alldatatypes_test.go
-- create table arr (
--     c1 int, 
--     c2 SMALLINT, 
--     c3 BIGINT, 
--     c4 INTEGER, 
--     c5 DECIMAL(4,2), 
--     c6 NUMERIC, 
--     c7 float, 
--     c8 double, 
--     c9 decfloat, 
--     c10 char(10), 
--     c11 varchar(10), 
--     c12 char for bit data, 
--     c13 clob(10),
--     c14 dbclob(100), 
--     c15 date, 
--     c16 time, 
--     c17 timestamp, 
--     c18 blob(10), 
--     c19 boolean
-- ) ccsid unicode

-- db2cli.ini
-- [SAMPLE]
-- HOSTNAME=localhost
-- DATABASE=SAMPLE
-- PORT=50000
-- UID=db2inst1
-- PWD=password

-- db2cli execsql -dsn SAMPLE -inputsql create_table.sql -statementdelimiter ';'
drop table datatypes;
create table datatypes (
    col01 int NOT NULL, 
    col02 SMALLINT, 
    col03 BIGINT, 
    col04 INTEGER, 
    col05 DECIMAL(4,2), 
    col06 NUMERIC, 
    col07 float, 
    col08 double, 
    col09 decfloat, 
    col10 char(10), 
    col11 varchar(10), 
    col12 char for bit data, 
    col13 clob(10),
    col14 dbclob(100), 
    col15 date, 
    col16 time, 
    col17 timestamp, 
    col18 blob(10), 
    col19 boolean,
    col20 graphic(10),
    col21 vargraphic(10),
    col22 binary(10),
    col23 varbinary(10),
    col24 xml,
    CONSTRAINT pk_datatypes PRIMARY KEY(col01)
) ccsid unicode;
