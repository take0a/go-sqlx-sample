DROP TABLE datatypes;

CREATE TABLE datatypes (
    col01 bigint,
    col02 bigserial,
    col03 bit(1),
    col04 varbit(1),
    col05 boolean,
    col06 box,
    col07 bytea,
    col08 varchar(10),
    col09 character(10),
    col10 cidr,
    col11 circle,
    col12 date,
    col13 float8,
    col14 inet,
    col15 integer,
    col16 interval,
    col17 json,
    col18 line,
    col19 lseg,
    col20 macaddr,
    col21 money,
    col22 numeric(10,4),
    col23 path,
    col24 point,
    col25 polygon,
    col26 real,
    col27 smallint,
    col28 smallserial,
    col29 serial,
    col30 text,
    col31 time,
    col32 timetz,
    col33 timestamp,
    col34 timestamptz,
    col35 tsquery,
    col36 tsvector,
    col37 txid_snapshot,
    col38 uuid,
    col39 xml,
    CONSTRAINT pk_datatypes PRIMARY KEY(col01)
);

COMMENT ON TABLE datatypes IS 'データ型テスト';
COMMENT ON COLUMN datatypes.col01 IS '8バイト符号付き整数';
COMMENT ON COLUMN datatypes.col02 IS '自動増分8バイト整数';
COMMENT ON COLUMN datatypes.col03 IS '固定長ビット列';
COMMENT ON COLUMN datatypes.col04 IS '可変長ビット列';
COMMENT ON COLUMN datatypes.col05 IS '論理値（真/偽）';
COMMENT ON COLUMN datatypes.col06 IS '平面上の矩形';
COMMENT ON COLUMN datatypes.col07 IS 'バイナリデータ（バイトの配列（bytearray））';
COMMENT ON COLUMN datatypes.col08 IS '可変長文字列';
COMMENT ON COLUMN datatypes.col09 IS '固定長文字列';
COMMENT ON COLUMN datatypes.col10 IS 'IPv4もしくはIPv6ネットワークアドレス';
COMMENT ON COLUMN datatypes.col11 IS '平面上の円';
COMMENT ON COLUMN datatypes.col12 IS '暦の日付（年月日）';
COMMENT ON COLUMN datatypes.col13 IS '倍精度浮動小数点（8バイト）';
COMMENT ON COLUMN datatypes.col14 IS 'IPv4もしくはIPv6ホストアドレス';
COMMENT ON COLUMN datatypes.col15 IS '4バイト符号付き整数';
COMMENT ON COLUMN datatypes.col16 IS '時間間隔';
COMMENT ON COLUMN datatypes.col17 IS 'JSONデータ';
COMMENT ON COLUMN datatypes.col18 IS '平面上の無限直線';
COMMENT ON COLUMN datatypes.col19 IS '平面上の線分';
COMMENT ON COLUMN datatypes.col20 IS 'MAC（メディアアクセスコントロール）アドレス';
COMMENT ON COLUMN datatypes.col21 IS '貨幣金額';
COMMENT ON COLUMN datatypes.col22 IS '精度の選択可能な高精度数値';
COMMENT ON COLUMN datatypes.col23 IS '平面上の幾何学的経路';
COMMENT ON COLUMN datatypes.col24 IS '平面上の幾何学的点';
COMMENT ON COLUMN datatypes.col25 IS '平面上の閉じた幾何学的経路';
COMMENT ON COLUMN datatypes.col26 IS '単精度浮動小数点（4バイト）';
COMMENT ON COLUMN datatypes.col27 IS '2バイト符号付き整数';
COMMENT ON COLUMN datatypes.col28 IS '自動増分2バイト整数';
COMMENT ON COLUMN datatypes.col29 IS '自動増分4バイト整数';
COMMENT ON COLUMN datatypes.col30 IS '可変長文字列';
COMMENT ON COLUMN datatypes.col31 IS '時刻（時間帯なし）';
COMMENT ON COLUMN datatypes.col32 IS '時間帯付き時刻';
COMMENT ON COLUMN datatypes.col33 IS '日付と時刻（時間帯なし）';
COMMENT ON COLUMN datatypes.col34 IS '時間帯付き日付と時刻';
COMMENT ON COLUMN datatypes.col35 IS 'テキスト検索問い合わせ';
COMMENT ON COLUMN datatypes.col36 IS 'テキスト検索文書';
COMMENT ON COLUMN datatypes.col37 IS 'ユーザレベルのトランザクションIDスナップショット';
COMMENT ON COLUMN datatypes.col38 IS '汎用一意識別子';
COMMENT ON COLUMN datatypes.col39 IS 'XMLデータ';
