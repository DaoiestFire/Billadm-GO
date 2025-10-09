import api_client from "@/backend/api/api_client.js";
import {isResponseSuccess} from "@/backend/functions.js";

export async function queryTag(category) {
    const resp = await api_client.post(`/v1/tag/query_tag/${category}`);
    if (!isResponseSuccess(resp)) {
        throw "queryTag 响应错误"
    }
    return resp.data
}