import api_client from "@/backend/api/api-client.ts";
import type {Ledger, Result} from "@/types/billadm";

export async function queryAllLedgers(): Promise<Ledger[]> {
    const resp: Result<Ledger[]> = await api_client.post('/v1/ledger/query-all', {
        'ledgerId': 'all',
    });
    api_client.isRespSuccess(resp, 'queryAllLedgers错误: ');
    return resp.data;
}

export async function createLedgerByName(name: string): Promise<string> {
    const resp: Result<string> = await api_client.post('/v1/ledger/create-ledger', {name});
    api_client.isRespSuccess(resp, 'createLedgerByName错误: ');
    return resp.data;
}

export async function deleteLedgerById(id: string): Promise<void> {
    const resp: Result = await api_client.post('/v1/ledger/delete-ledger', {
        'ledgerId': id,
    });
    api_client.isRespSuccess(resp, 'deleteLedgerById错误: ');
}