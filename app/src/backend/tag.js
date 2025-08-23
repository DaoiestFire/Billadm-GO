import api_client from "@/backend/api_client.js";
import {isResponseSuccess} from "@/backend/functions.js";

export async function queryAllTag() {
    const resp = await api_client.post('/v1/tag/query_all_tag');
    if (!isResponseSuccess(resp)) {
        throw "queryAllTag 响应错误"
    }
    return resp.data
}