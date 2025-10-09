import api_client from "@/backend/api/api_client.js";
import {isResponseSuccess} from "@/backend/functions.js";

export async function queryCategory(trType) {
    const resp = await api_client.post(`/v1/category/query_category/${trType}`);
    if (!isResponseSuccess(resp)) {
        throw "queryCategory 响应错误"
    }
    return resp.data
}