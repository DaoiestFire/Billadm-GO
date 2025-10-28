import api_client from "@/backend/api/api-client.ts";
import type {Result, TransactionRecord, TrQueryCondition, TrStatistics} from "@/types/billadm";

export async function queryTrOnCondition(condition: TrQueryCondition): Promise<TransactionRecord[]> {
    const resp: Result<TransactionRecord[]> = await api_client.post('/v1/tr/query', condition);
    api_client.isRespSuccess(resp, 'queryTrsOnCondition错误: ');
    return resp.data;
}

export async function queryTrCountOnCondition(condition: TrQueryCondition): Promise<number> {
    const resp: Result<number> = await api_client.post('/v1/tr/query-count', condition);
    api_client.isRespSuccess(resp, 'queryTrCountOnCondition错误: ');
    return resp.data;
}

export async function createTrForLedger(data: TransactionRecord): Promise<string> {
    const resp: Result<string> = await api_client.post('/v1/tr/create', data);
    api_client.isRespSuccess(resp, 'createTrForLedger错误: ');
    return resp.data;
}

export async function deleteTrById(id: string) {
    const resp: Result = await api_client.post('/v1/tr/delete-by-id', {
        'trId': id,
    });
    api_client.isRespSuccess(resp, 'deleteTrById错误: ');
}

export async function queryTrStatisticsOnCondition(condition: TrQueryCondition): Promise<TrStatistics> {
    const resp: Result<TrStatistics> = await api_client.post('/v1/statistics/tr', condition);
    api_client.isRespSuccess(resp, 'queryTrStatisticsOnCondition错误: ');
    return resp.data;
}