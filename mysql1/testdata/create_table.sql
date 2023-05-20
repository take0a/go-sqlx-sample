DROP TABLE datatypes;

CREATE TABLE datatypes (
    col01 tinyint COMMENT '1byte整数',
    col02 smallint COMMENT '2byte整数',
    col03 mediumint COMMENT '3byte整数',
    col04 int COMMENT '4byte整数',
    col05 bigint COMMENT '8byte整数',
    col06 decimal(10,4) COMMENT '固定小数点数',
    col07 float COMMENT '4byte浮動小数点数',
    col08 double COMMENT '8byte浮動小数点数',
    col09 bit(3) COMMENT 'ビット値',
    col10 date COMMENT '日付',
    col11 datetime COMMENT '日時',
    col12 timestamp COMMENT 'UNIXタイムスタンプ',
    col13 time COMMENT '時間値',
    col14 year COMMENT '年',
    col15 char(10) COMMENT '固定長文字列',
    col16 varchar(10) COMMENT '可変長文字列',
    col17 binary(10) COMMENT '固定長byte列',
    col18 varbinary(10) COMMENT '可変長byte列',
    col19 tinyblob COMMENT 'tinyblob',
    col20 blob COMMENT 'blob',
    col21 mediumblob COMMENT 'mediumblob',
    col22 longblob COMMENT 'longblob',
    col23 tinytext COMMENT 'tinytext',
    col24 text COMMENT 'text',
    col25 mediumtext COMMENT 'mediumtext',
    col26 longtext COMMENT 'longtext',
    col27 enum('S', 'M', 'L') COMMENT '列挙',
    col28 set('R', 'G', 'B')  COMMENT '集合',
    col29 geometry COMMENT 'POINTかLINESTRINGかPOLYGON',
    col30 point COMMENT '座標',
    col31 linestring COMMENT '直線で補間される一連の点',
    col32 polygon COMMENT '単純かつ閉じたLINESTRING',
    col33 multipoint COMMENT '座標の集まり',
    col34 multilinestring COMMENT 'LINESTRINGの集まり',
    col35 multipolygon COMMENT 'POLYGONの集まり',
    col36 geometrycollection COMMENT 'GEOMETRYの集まり',
    col37 json COMMENT 'JSON文字列',
    CONSTRAINT pk_datatypes PRIMARY KEY(col01)
) COMMENT='データ型テスト';
