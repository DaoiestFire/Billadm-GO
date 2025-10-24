import api_client from "@/backend/api/api-client.ts";
import type {Result} from "@/types/billadm";

export async function exitApp(): Promise<void> {
    const resp: Result = await api_client.post('/v1/app/exit_app');
    api_client.isRespSuccess(resp, 'exitApp错误: ');
}