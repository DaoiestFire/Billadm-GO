import api_client from "@/backend/api_client.js";
import {dateToUnixTimestamp, isResponseSuccess} from "@/backend/functions.js";

export async function getAllTrsFromLedgerById(id) {
    const resp = await api_client.post('/v1/tr/query_all_trs', {'ledger_id': id});
    if (!isResponseSuccess(resp)) {
        throw "getAllTrsFromLedgerById 响应错误"
    }
    return resp.data
}

export async function queryTrsOnCondition(id, offset, limit) {
    const resp = await api_client.post('/v1/tr/query_trs_on_condition', {
        'ledger_id': id,
        'offset': offset,
        'limit': limit,
    });
    if (!isResponseSuccess(resp)) {
        throw `queryTrsOnCondition 响应错误 ${resp.msg}`
    }
    return resp.data
}

export async function queryTrCountOnCondition(id) {
    const resp = await api_client.post('/v1/tr/query_tr_count_on_condition', {
        'ledger_id': id,
    });
    if (!isResponseSuccess(resp)) {
        throw `queryTrCountOnCondition 响应错误 ${resp.msg}`
    }
    return resp.data
}

export async function createTrForLedger(data) {
    const resp = await api_client.post('/v1/tr/create_tr', data);
    if (!isResponseSuccess(resp)) {
        throw `createTrForLedger 响应错误 ${resp.msg}`
    }
}

export async function deleteTrById(id) {
    const resp = await api_client.post('/v1/tr/delete_tr_by_id', {'tr_id': id});
    if (!isResponseSuccess(resp)) {
        throw "deleteTrById 响应错误"
    }
}

/**
 * 构造符合后端 TransactionRecordDto 的请求对象 表单数据转化为dto
 * @param {Object} data - 输入数据，字段可选
 * @param ledgerId
 * @returns {Object} 符合后端结构的对象
 */
export function trFormToTrDto(data = {}, ledgerId = '') {

    const transactionRecord = {
        ledger_id: data.ledger_id || ledgerId,           // string，默认为空字符串
        transaction_id: data.id || '', // string，默认为空字符串

        // 交易核心信息
        price: typeof data.price === 'number' ? data.price : 0.0,        // float64，默认 0.0
        transaction_type: data.type || '',                   // string，默认空字符串

        // 分类与描述
        category: data.category || '',                                   // string，默认空字符串
        description: data.description || '',                           // string，默认空字符串
        tags: Array.isArray(data.tags) ? data.tags : [],                // []string，默认空数组
        transaction_at: dateToUnixTimestamp(data.time),
    };

    if (data.transaction_at !== undefined && !isNaN(data.transaction_at)) {
        transactionRecord.transaction_at = data.transaction_at // 转换为秒级别时间戳
    }

    return transactionRecord
}

export function trDtoToTrForm(item) {
    return {
        id: item.transaction_id,
        time: new Date(item.transaction_at * 1000),
        type: item.transaction_type,
        category: item.category,
        description: item.description,
        tags: item.tags,
        price: item.price,
    }
}

