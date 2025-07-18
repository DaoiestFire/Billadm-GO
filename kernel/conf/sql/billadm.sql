-- 创建账本表
CREATE TABLE IF NOT EXISTS tbl_billadm_ledger
(
    id         TEXT PRIMARY KEY,
    name       TEXT     NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);

-- 创建交易记录表
CREATE TABLE IF NOT EXISTS tbl_billadm_transaction_record
(
    transaction_id   TEXT PRIMARY KEY,
    ledger_id        TEXT     NOT NULL,
    price            REAL     NOT NULL,
    transaction_type TEXT     NOT NULL,
    category         TEXT     NOT NULL,
    description      TEXT,
    tags             TEXT,
    transaction_at   DATETIME NOT NULL,
    created_at       DATETIME NOT NULL,
    updated_at       DATETIME NOT NULL
);

-- 创建消费分类表
CREATE TABLE IF NOT EXISTS tbl_billadm_category
(
    name TEXT PRIMARY KEY
);

-- 更新内置消费分类
DELETE FROM tbl_billadm_category WHERE name in ('餐饮美食', '交通出行', '购物消费', '娱乐休闲','生活缴费', '医疗健康', '人情往来', '教育学习');
INSERT INTO tbl_billadm_category (name) VALUES ('餐饮美食'),('交通出行'),('购物消费'),('娱乐休闲'),('生活缴费'),('医疗健康'),('人情往来'),('教育学习');

-- 创建消费标签表
CREATE TABLE IF NOT EXISTS tbl_billadm_tag
(
    name TEXT PRIMARY KEY
);

-- 更新内置的消费标签
DELETE FROM tbl_billadm_tag WHERE name in ('三餐','外卖','奶茶','零食','水果',
                                           '房租','物业','水电','燃气','通讯',
                                           '社交','打车','地铁','公交','高铁',
                                           '油费','停车','ETC','衣物','数码',
                                           '家居','书籍','礼物','宠物','医药',
                                           '保险','游戏','快递','维修','清洁',
                                           '彩票','红包','旅游','电影','健身',
                                           '请客','酒店','烟酒','培训','充值',
                                           '还款','汽车');
INSERT INTO tbl_billadm_tag (name) VALUES ('三餐'),('外卖'),('奶茶'),('零食'),('水果'),
                                          ('房租'),('物业'),('水电'),('燃气'),('通讯'),
                                          ('社交'),('打车'),('地铁'),('公交'),('高铁'),
                                          ('油费'),('停车'),('ETC'),('衣物'),('数码'),
                                          ('家居'),('书籍'),('礼物'),('宠物'),('医药'),
                                          ('保险'),('游戏'),('快递'),('维修'),('清洁'),
                                          ('彩票'),('红包'),('旅游'),('电影'),('健身'),
                                          ('请客'),('酒店'),('烟酒'),('培训'),('充值'),
                                          ('还款'),('汽车');

-- 创建索引
CREATE INDEX IF NOT EXISTS idx_transaction_records_ledger_id ON tbl_billadm_transaction_record (ledger_id);
CREATE INDEX IF NOT EXISTS idx_transaction_records_transaction_type ON tbl_billadm_transaction_record (transaction_type);
CREATE INDEX IF NOT EXISTS idx_transaction_records_category ON tbl_billadm_transaction_record (category);
CREATE INDEX IF NOT EXISTS idx_transaction_records_transaction_at ON tbl_billadm_transaction_record (transaction_at);