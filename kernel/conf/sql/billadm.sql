-- 创建账本表
CREATE TABLE IF NOT EXISTS tbl_billadm_ledger
(
    id         TEXT PRIMARY KEY,
    user_id    TEXT     NOT NULL,
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
    category_id      TEXT     NOT NULL,
    description      TEXT,
    tags             TEXT,
    transaction_at   DATETIME NOT NULL,
    created_at       DATETIME NOT NULL,
    updated_at       DATETIME NOT NULL
);

-- 创建索引
CREATE INDEX IF NOT EXISTS idx_transaction_records_ledger_id ON tbl_billadm_transaction_record (ledger_id);
CREATE INDEX IF NOT EXISTS idx_transaction_records_transaction_type ON tbl_billadm_transaction_record (transaction_type);
CREATE INDEX IF NOT EXISTS idx_transaction_records_category_id ON tbl_billadm_transaction_record (category_id);
CREATE INDEX IF NOT EXISTS idx_transaction_records_transaction_at ON tbl_billadm_transaction_record (transaction_at);