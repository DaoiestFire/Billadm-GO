import api_client from "@/backend/api_client.js";
import {isResponseSuccess} from "@/backend/functions.js";

export async function exitApp() {
    const resp = await api_client.post('/v1/app/exit_app');
    if (!isResponseSuccess(resp)) {
        throw "exitApp 响应错误"
    }
    return resp.data
}