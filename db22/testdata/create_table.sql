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

--  db2cli execsql -dsn SAMPLE -inputsql create_table.sql -statementdelimiter ';'
drop table 日ほンｺﾞ表;
create table 日ほンｺﾞ表 (
    列１ char(30) NOT NULL, 
    れつ2 varchar(30), 
    レツ３ clob(30),
    ﾚﾂ4 graphic(10),
    Ｒｅｔｓｕ５ vargraphic(20),
    CONSTRAINT pk_datatypes PRIMARY KEY(列１)
) ccsid unicode;
