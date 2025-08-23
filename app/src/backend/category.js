import api_client from "@/backend/api_client.js";
import {isResponseSuccess} from "@/backend/functions.js";

export async function queryAllCategory() {
    const resp = await api_client.post('/v1/category/query_all_category');
    if (!isResponseSuccess(resp)) {
        throw "queryAllCategory 响应错误"
    }
    return resp.data
}