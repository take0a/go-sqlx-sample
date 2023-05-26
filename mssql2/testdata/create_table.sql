use agra;
go

DROP TABLE IF EXISTS 日ほンｺﾞ表;
go

CREATE TABLE 日ほンｺﾞ表 (
    列１ char(30) collate Japanese_CI_AS,
    れつ2 varchar(30) collate Japanese_CI_AS,
    レツ３ text collate Japanese_CI_AS,
    ﾚﾂ4 nchar(10),
    Ｒｅｔｓｕ５ nvarchar(10),
    列６ ntext,
    列７ xml,
    CONSTRAINT pk_日ほンｺﾞ表 PRIMARY KEY(列１)
);
go
