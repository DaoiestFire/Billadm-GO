/**
 * 表示一个前端使用的消费记录
 */
export interface TrForm {
    /**
     * 交易ID. 对应 transaction_id.
     */
    id: string;
    /**
     * 金额. 对应 price.
     */
    price: number;
    /**
     * 交易类型 (e.g., 'income', 'expense', 'transfer'). 对应 transaction_type.
     */
    type: string;
    /**
     * 消费类型. 对应 category.
     */
    category: string;
    /**
     * 描述. 对应 description.
     */
    description: string;
    /**
     * 标签. 对应 tags.
     */
    tags: string[];
    /**
     * 交易时间. 对应 transaction_at. 时间戳转 Date 对象.
     */
    time: Date;
}

export interface ApiClient {
    get<T = any>(url: string): Promise<Result<T>>;

    post<T = any>(url: string, data?: any): Promise<Result<T>>;

    isRespSuccess(result: Result, prefix?: string): void;
}

/**
 * 后端返回的响应的规范结构
 */
export interface Result<T = any> {
    code: number;
    msg: string;
    data: T;
}

/**
 * 账本
 */
export interface Ledger {
    id: string;           // 账本UUID
    name: string;         // 账本名称
    createdAt: number;   // 创建时间（Unix 时间戳，单位秒）
    updatedAt: number;   // 更新时间（Unix 时间戳，单位秒）
}

/**
 * 消费记录
 */
export interface TransactionRecord {
    ledgerId: string;
    transactionId: string;
    price: number;
    transactionType: string;
    category: string;
    description: string;
    tags: string[];
    transactionAt: number;
}

/**
 * 消费类型
 */
export interface Category {
    name: string;
    transactionType: string;
}

/**
 * 消费标签
 */
export interface Tag {
    name: string;      // 标签名称
    scope: string;     // 作用域
    category: string;  // 分类ID
}

/**
 * 消费记录统计数据
 */
export interface TrStatistics {
    income: number;    // 收入金额
    expense: number;   // 支出金额
    transfer: number;  // 转账金额
}

/**
 * 消费记录条件查询
 */
export interface TrQueryCondition {
    ledgerId: string;
    offset: number;
    limit: number;
    tsRange: number[];
    transactionType: string[];
}

/**
 * 工作空间状态
 */
export interface WorkspaceStatus {
    isOpened: boolean;
    workspaceDir: string;
}

/**
 * 时间选择器类型
 */
type TimeRangeType = 'year' | 'month' | 'date'