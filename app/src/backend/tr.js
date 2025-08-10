import api_client from "@/backend/api_client.js";
import {getUnixTimestampInSeconds} from "@/backend/functions.js";

export async function getAllTrsFromLedgerById(id) {
    return await api_client.post('/v1/tr/get', {'ledger_id': id});
}

export async function createTrForLedger(data) {
    return await api_client.post('/v1/tr/post', data);
}

export async function deleteTrById(id) {
    return await api_client.post('/v1/tr/delete', {'tr_id': id});
}

/**
 * 构造符合后端 TransactionRecordDto 的请求对象
 * @param {Object} data - 输入数据，字段可选
 * @returns {Object} 符合后端结构的对象
 */
export function buildTransactionRecordDto(data = {}) {

    const transactionRecord = {
        ledger_id: data.ledger_id || '',           // string，默认为空字符串
        transaction_id: data.transaction_id || '', // string，默认为空字符串

        // 交易核心信息
        price: typeof data.price === 'number' ? data.price : 0.0,        // float64，默认 0.0
        transaction_type: data.transaction_type || '',                   // string，默认空字符串

        // 分类与描述
        category: data.category || '',                                   // string，默认空字符串
        description: data.description || '',                           // string，默认空字符串
        tags: Array.isArray(data.tags) ? data.tags : [],                // []string，默认空数组
        transaction_at: getUnixTimestampInSeconds(),
    };

    if (data.transaction_at !== undefined && !isNaN(data.transaction_at)) {
        transactionRecord.transaction_at = data.transaction_at // 转换为秒级别时间戳
    }

    return transactionRecord
}

