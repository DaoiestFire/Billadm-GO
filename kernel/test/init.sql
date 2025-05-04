-- 创建账本表
CREATE TABLE IF NOT EXISTS tbl_billadm_ledger
(
    id         TEXT PRIMARY KEY,
    user_id    TEXT     NOT NULL,
    name       TEXT     NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL
);