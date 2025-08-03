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
    name in (
        '餐饮美食',
        '交通出行',
        '购物消费',
        '娱乐休闲',
        '生活缴费',
        '医疗健康',
        '人情往来',
        '教育学习'
    );

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
    name in (
        '三餐',
        '外卖',
        '奶茶',
        '零食',
        '水果',
        '房租',
        '物业',
        '水电',
        '燃气',
        '通讯',
        '社交',
        '打车',
        '地铁',
        '公交',
        '高铁',
        '油费',
        '停车',
        'ETC',
        '衣物',
        '数码',
        '家居',
        '书籍',
        '礼物',
        '宠物',
        '医药',
        '保险',
        '游戏',
        '快递',
        '维修',
        '清洁',
        '彩票',
        '红包',
        '旅游',
        '电影',
        '健身',
        '请客',
        '酒店',
        '烟酒',
        '培训',
        '充值',
        '还款',
        '汽车'
    );

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

-- 创建索引
CREATE INDEX IF NOT EXISTS idx_transaction_records_ledger_id ON tbl_billadm_transaction_record (ledger_id);

CREATE INDEX IF NOT EXISTS idx_transaction_records_transaction_type ON tbl_billadm_transaction_record (transaction_type);

CREATE INDEX IF NOT EXISTS idx_transaction_records_category ON tbl_billadm_transaction_record (category);

CREATE INDEX IF NOT EXISTS idx_transaction_records_transaction_at ON tbl_billadm_transaction_record (transaction_at);