import api_client from "@/backend/api/api_client.js";
import {isResponseSuccess} from "@/backend/functions.js";

export async function openWorkspace(workspace) {
    const resp = await api_client.post('/v1/workspace/open_workspace', {'workspaceDir': workspace});
    if (!isResponseSuccess(resp)) {
        throw "openWorkspace 响应错误"
    }
    return resp.code
}

export async function hasOpenedWorkspace() {
    const resp = await api_client.post('/v1/workspace/is_opened');
    if (!isResponseSuccess(resp)) {
        throw "hasOpenedWorkspace 响应错误"
    }
    return {
        opened: resp.code === 0,
        workspace: resp.data,
    }
}