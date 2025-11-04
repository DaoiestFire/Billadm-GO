import dayjs from "dayjs";
import type {Category, Tag, TransactionRecord, TrQueryCondition, TrStatistics} from "@/types/billadm";
import {
    createTrForLedger,
    deleteTrById,
    queryTrCountOnCondition,
    queryTrOnCondition,
    queryTrStatisticsOnCondition
} from "@/backend/api/tr.ts";
import NotificationUtil from "@/backend/notification.ts";
import {queryCategory} from "@/backend/api/category.ts";
import {queryTags} from "@/backend/api/tag.ts";

export function formatFloat(num: number): number {
    return parseFloat(num.toFixed(2));
}

/**
 * 将秒级时间戳转换为格式化时间字符串
 * @param timestamp 秒级时间戳
 * @param format 格式，默认为 'YYYY-MM-DD HH:mm:ss'
 * @returns 格式化后的时间字符串
 */
export function formatTimestamp(timestamp: number, format: string = 'YYYY-MM-DD'): string {
    return dayjs(timestamp * 1000).format(format);
}

/**
 * 消费记录
 */
export async function getTrTotalOnCondition(condition: TrQueryCondition): Promise<number> {
    try {
        return await queryTrCountOnCondition(condition);
    } catch (error) {
        NotificationUtil.error('查询消费记录数量失败', `${error}`);
        return 0;
    }
}

export async function getTrOnCondition(condition: TrQueryCondition): Promise<TransactionRecord[]> {
    try {
        return await queryTrOnCondition(condition);
    } catch (error) {
        NotificationUtil.error('查询消费记录失败', `${error}`);
        return [];
    }
}

export async function createTransactionRecord(tr: TransactionRecord) {
    try {
        await createTrForLedger(tr);
    } catch (error) {
        NotificationUtil.error('创建消费记录失败', `${error}`);
    }
}

export async function deleteTransactionRecord(trId: string) {
    try {
        await deleteTrById(trId);
    } catch (error) {
        NotificationUtil.error('删除消费记录失败', `${error}`);
    }
}

export async function updateTransactionRecord(tr: TransactionRecord) {
    try {
        await deleteTrById(tr.transactionId);
        await createTrForLedger(tr);
    } catch (error) {
        NotificationUtil.error('更新消费记录失败', `${error}`);
    }
}

/**
 * 分类与标签
 */
export async function getCategoryByType(trType: string): Promise<Category[]> {
    try {
        return await queryCategory(trType);
    } catch (error) {
        NotificationUtil.error(`查询 ${trType} 消费类型失败`, `${error}`);
        return [];
    }
}

export async function getTagsByCategory(category: string): Promise<Tag[]> {
    try {
        return await queryTags(category);
    } catch (error) {
        NotificationUtil.error(`查询 ${category} 消费标签失败`, `${error}`);
        return [];
    }
}

/**
 * 统计指标
 */
export async function getStatisticsOnCondition(condition: TrQueryCondition): Promise<TrStatistics> {
    try {
        return await queryTrStatisticsOnCondition(condition);
    } catch (error) {
        NotificationUtil.error('查询统计数据失败', `${error}`);
        return {income: 0, expense: 0, transfer: 0};
    }
}
