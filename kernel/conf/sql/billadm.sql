-- 创建账本表
CREATE TABLE
    IF NOT EXISTS tbl_billadm_ledger (
        id TEXT PRIMARY KEY,
        name TEXT NOT NULL,
        created_at INTEGER NOT NULL,
        updated_at INTEGER NOT NULL
    );

-- 创建交易记录表
CREATE TABLE
    IF NOT EXISTS tbl_billadm_transaction_record (
        transaction_id TEXT PRIMARY KEY,
        ledger_id TEXT NOT NULL,
        price REAL NOT NULL,
        transaction_type TEXT NOT NULL,
        category TEXT NOT NULL,
        description TEXT,
        transaction_at INTEGER NOT NULL,
        created_at INTEGER NOT NULL,
        updated_at INTEGER NOT NULL
    );

-- 创建交易记录标签表
CREATE TABLE
    IF NOT EXISTS tbl_billadm_transaction_record_tag (transaction_id TEXT NOT NULL, tag TEXT NOT NULL);

-- 创建消费分类表
CREATE TABLE
    IF NOT EXISTS tbl_billadm_category (name TEXT PRIMARY KEY, scope TEXT NOT NULL);

-- 更新内置消费分类
DELETE FROM tbl_billadm_category
WHERE
    scope = 'system';

INSERT INTO
    tbl_billadm_category (name, scope)
VALUES
    ('餐饮美食', 'system'),
    ('交通出行', 'system'),
    ('购物消费', 'system'),
    ('娱乐休闲', 'system'),
    ('生活缴费', 'system'),
    ('医疗健康', 'system'),
    ('人情往来', 'system'),
    ('教育学习', 'system');

-- 创建消费标签表
CREATE TABLE
    IF NOT EXISTS tbl_billadm_tag (name TEXT PRIMARY KEY, scope TEXT NOT NULL);

-- 更新内置的消费标签
DELETE FROM tbl_billadm_tag
WHERE
    scope = 'system';

INSERT INTO
    tbl_billadm_tag (name, scope)
VALUES
    ('三餐', 'system'),
    ('外卖', 'system'),
    ('奶茶', 'system'),
    ('零食', 'system'),
    ('水果', 'system'),
    ('房租', 'system'),
    ('物业', 'system'),
    ('水电', 'system'),
    ('燃气', 'system'),
    ('通讯', 'system'),
    ('社交', 'system'),
    ('打车', 'system'),
    ('地铁', 'system'),
    ('公交', 'system'),
    ('高铁', 'system'),
    ('油费', 'system'),
    ('停车', 'system'),
    ('ETC', 'system'),
    ('衣物', 'system'),
    ('数码', 'system'),
    ('家居', 'system'),
    ('书籍', 'system'),
    ('礼物', 'system'),
    ('宠物', 'system'),
    ('医药', 'system'),
    ('保险', 'system'),
    ('游戏', 'system'),
    ('快递', 'system'),
    ('维修', 'system'),
    ('清洁', 'system'),
    ('彩票', 'system'),
    ('红包', 'system'),
    ('旅游', 'system'),
    ('电影', 'system'),
    ('健身', 'system'),
    ('请客', 'system'),
    ('酒店', 'system'),
    ('烟酒', 'system'),
    ('培训', 'system'),
    ('充值', 'system'),
    ('还款', 'system'),
    ('汽车', 'system');