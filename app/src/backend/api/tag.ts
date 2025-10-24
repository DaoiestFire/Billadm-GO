import api_client from "@/backend/api/api-client.ts";
import type {Result, Tag} from "@/types/billadm";

export async function queryTags(category: string): Promise<Tag[]> {
    const resp: Result<Tag[]> = await api_client.get(`/v1/tag/query_tag/${category}`);
    api_client.isRespSuccess(resp, 'queryTags错误: ');
    return resp.data;
}