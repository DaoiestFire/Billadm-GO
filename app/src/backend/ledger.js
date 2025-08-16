import api_client from "@/backend/api_client.js";
import {isResponseSuccess} from "@/backend/functions.js";

export async function getAllLedgers() {
    const resp = await api_client.post('/v1/ledger/get', {'ledger_id': 'all'});
    if (!isResponseSuccess(resp)) {
        throw "getAllLedgers 响应错误"
    }
    return resp.data
}

export async function createLedgerByName(name) {
    const resp = await api_client.post('/v1/ledger/post', {name});
    if (!isResponseSuccess(resp)) {
        throw "createLedgerByName 响应错误"
    }
    return resp.data
}

export async function deleteLedgerById(id) {
    const resp = await api_client.post('/v1/ledger/delete', {'ledger_id': id});
    if (!isResponseSuccess(resp)) {
        throw "deleteLedgerById 响应错误"
    }
    return resp.data
}