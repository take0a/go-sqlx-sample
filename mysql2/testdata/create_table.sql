DROP TABLE 日ほンｺﾞ表;

CREATE TABLE 日ほンｺﾞ表 (
    列１ char(10) COMMENT '固定長文字列',
    れつ2 varchar(10) COMMENT '可変長文字列',
    レツ３ tinytext COMMENT 'tinytext',
    ﾚﾂ4 text COMMENT 'text',
    Ｒｅｔｓｕ５ mediumtext COMMENT 'mediumtext',
    列６ longtext COMMENT 'longtext',
    列７ enum('小', '中', '大') COMMENT '列挙',
    列８ set('あか', 'みどり', 'あお') COMMENT '集合',
    列９ json COMMENT 'JSON文字列',

    CONSTRAINT pk_日ほンｺﾞ表 PRIMARY KEY(列１)
) COMMENT='日本語テスト';
