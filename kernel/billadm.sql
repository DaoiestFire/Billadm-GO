-- 创建账本表
CREATE TABLE IF NOT EXISTS tbl_billadm_ledger
(
    id         TEXT PRIMARY KEY,
    name       TEXT    NOT NULL,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL
);

-- 创建交易记录表
CREATE TABLE IF NOT EXISTS tbl_billadm_transaction_record
(
    transaction_id   TEXT PRIMARY KEY,
    ledger_id        TEXT    NOT NULL,
    price            REAL    NOT NULL,
    transaction_type TEXT    NOT NULL,
    category         TEXT    NOT NULL,
    description      TEXT,
    transaction_at   INTEGER NOT NULL,
    created_at       INTEGER NOT NULL,
    updated_at       INTEGER NOT NULL
);

-- 创建交易记录标签表
CREATE TABLE IF NOT EXISTS tbl_billadm_transaction_record_tag
(
    ledger_id      TEXT NOT NULL,
    transaction_id TEXT NOT NULL,
    tag            TEXT NOT NULL
);

-- 创建消费分类表
CREATE TABLE IF NOT EXISTS tbl_billadm_category
(
    name             TEXT PRIMARY KEY,
    scope            TEXT NOT NULL,
    transaction_type TEXT NOT NULL DEFAULT ''
);

-- 更新内置消费分类
DELETE
FROM tbl_billadm_category
WHERE name in ('餐饮美食',
               '交通出行',
               '购物消费',
               '娱乐休闲',
               '生活缴费',
               '医疗健康',
               '人情往来',
               '教育学习',
               '工资奖金',
               '二手转卖',
               '彩票收入',
               '投资理财',
               '五险一金',
               '税费党费');

INSERT INTO tbl_billadm_category (name, scope, transaction_type)
VALUES ('餐饮美食', 'system', 'expense'),
       ('交通出行', 'system', 'expense'),
       ('购物消费', 'system', 'expense'),
       ('娱乐休闲', 'system', 'expense'),
       ('生活缴费', 'system', 'expense'),
       ('医疗健康', 'system', 'expense'),
       ('人情往来', 'system', 'expense'),
       ('教育学习', 'system', 'expense'),
       ('工资奖金', 'system', 'income'),
       ('二手转卖', 'system', 'income'),
       ('彩票收入', 'system', 'income'),
       ('投资理财', 'system', 'income'),
       ('五险一金', 'system', 'transfer'),
       ('税费党费', 'system', 'transfer');

-- 创建消费标签表
CREATE TABLE IF NOT EXISTS tbl_billadm_tag
(
    name     TEXT PRIMARY KEY,
    scope    TEXT NOT NULL,
    category TEXT NOT NULL DEFAULT ''
);

-- 更新内置的消费标签
DELETE
FROM tbl_billadm_tag
WHERE name in ('三餐',
               '商场',
               '外卖',
               '奶茶',
               '零食',
               '水果',
               '打车',
               '地铁',
               '公交',
               '高铁',
               '油费',
               '停车',
               'ETC',
               '车险',
               '衣物',
               '数码',
               '家居',
               '书籍',
               '礼物',
               '宠物',
               '游戏',
               '快递',
               '彩票',
               '电影',
               '健身',
               '酒店',
               '烟酒',
               '充值',
               '汽车',
               '房租',
               '物业',
               '水电',
               '燃气',
               '通讯',
               '人险',
               '还款',
               '医药',
               '医险',
               '红包',
               '请客',
               '养老',
               '医疗',
               '失业',
               '住房',
               '团费',
               '交税');

INSERT INTO tbl_billadm_tag (name, scope, category)
VALUES ('三餐', 'system', '餐饮美食'),
       ('商场', 'system', '餐饮美食'),
       ('外卖', 'system', '餐饮美食'),
       ('奶茶', 'system', '餐饮美食'),
       ('零食', 'system', '餐饮美食'),
       ('水果', 'system', '餐饮美食'),
       ('打车', 'system', '交通出行'),
       ('地铁', 'system', '交通出行'),
       ('公交', 'system', '交通出行'),
       ('高铁', 'system', '交通出行'),
       ('油费', 'system', '交通出行'),
       ('停车', 'system', '交通出行'),
       ('ETC', 'system', '交通出行'),
       ('车险', 'system', '交通出行'),
       ('衣物', 'system', '购物消费'),
       ('数码', 'system', '购物消费'),
       ('家居', 'system', '购物消费'),
       ('书籍', 'system', '购物消费'),
       ('礼物', 'system', '购物消费'),
       ('宠物', 'system', '购物消费'),
       ('游戏', 'system', '购物消费'),
       ('快递', 'system', '购物消费'),
       ('彩票', 'system', '购物消费'),
       ('电影', 'system', '购物消费'),
       ('健身', 'system', '购物消费'),
       ('酒店', 'system', '购物消费'),
       ('烟酒', 'system', '购物消费'),
       ('充值', 'system', '购物消费'),
       ('汽车', 'system', '购物消费'),
       ('还款', 'system', '购物消费'),
       ('房租', 'system', '生活缴费'),
       ('物业', 'system', '生活缴费'),
       ('水电', 'system', '生活缴费'),
       ('燃气', 'system', '生活缴费'),
       ('通讯', 'system', '生活缴费'),
       ('人险', 'system', '生活缴费'),
       ('还款', 'system', '生活缴费'),
       ('医药', 'system', '医疗健康'),
       ('医险', 'system', '医疗健康'),
       ('红包', 'system', '人情往来'),
       ('请客', 'system', '人情往来'),
       ('养老', 'system', '五险一金'),
       ('医疗', 'system', '五险一金'),
       ('失业', 'system', '五险一金'),
       ('住房', 'system', '五险一金'),
       ('团费', 'system', '税费党费'),
       ('交税', 'system', '税费党费');


