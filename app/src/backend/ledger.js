import api_client from "@/backend/api_client.js";

export async function getAllLedgers() {
    return await api_client.get('/v1/ledger', {'ledger_id': 'all'});
}

export async function getLedgerById(ids) {
    return await api_client.get('/v1/ledger', {'ledger_id': ids.join(',')});
}

export async function createLedgerByName(name) {
    return await api_client.post('/v1/ledger', {name});
}