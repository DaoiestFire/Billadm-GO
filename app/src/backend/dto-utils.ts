import type {TransactionRecord, TrForm} from "@/types/billadm";
import dayjs from "dayjs";

/**
 * 构造符合后端 TransactionRecordDto 的请求对象 表单数据转化为dto
 */
export function trFormToTrDto(data: TrForm, ledgerId: string = ''): TransactionRecord {
    return {
        ledgerId: ledgerId,
        transactionId: data.id,
        price: data.price,
        transactionType: data.type,
        category: data.category,
        description: data.description,
        tags: data.tags,
        transactionAt: data.time.unix(),
    };
}

/**
 * 将后端返回的 TransactionRecord DTO 转换为前端使用的 TrForm 表单对象
 */
export function trDtoToTrForm(dto: TransactionRecord): TrForm {
    return {
        id: dto.transactionId,
        price: dto.price,
        type: dto.transactionType,
        category: dto.category,
        description: dto.description,
        tags: dto.tags,
        time: dayjs(dto.transactionAt * 1000),
    };
}