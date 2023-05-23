use agra;
go

DROP TABLE IF EXISTS datatypes;
go

CREATE TABLE datatypes (
    col01 bigint,
    col02 int,
    col03 smallint,
    col04 tinyint,
    col05 bit,
    col06 decimal(10,4),
    col07 money,
    col08 smallmoney,
    col09 float(53),
    col10 real,
    col11 date,
    col12 datetime,
    col13 datetime2,
    col14 datetimeoffset,
    col15 smalldatetime,
    col16 time,
    col17 char(10),
    col18 varchar(10),
    col19 text,
    col20 nchar(10),
    col21 nvarchar(10),
    col22 ntext,
    col23 binary(10),
    col24 varbinary(10),
    col25 image,
    col26 uniqueidentifier,
    col27 sql_variant,
    col28 xml,
    CONSTRAINT pk_datatypes PRIMARY KEY(col01)
);
go
