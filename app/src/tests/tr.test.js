import {describe, it} from "vitest";
import {createLedgerByName} from "@/backend/api/ledger.js";
import {createTrForLedger, deleteTrById, getAllTrsFromLedgerById, trFormToTrDto} from "@/backend/api/tr.js";

describe('Tr Utils', () => {
    it('test tr utils', async () => {
        const resp1 = await createLedgerByName('test-ledger-name')
        console.log(resp1)
        const ledgerId = resp1.data
        console.log("测试账本id：", ledgerId)

        // 创建一个消费记录
        const tr = trFormToTrDto({
            ledger_id: ledgerId,
            price: 1234.5,
            transaction_type: 'expense',
            category: '一日三餐',
            description: '这是一个测试消费记录',
            tags: ['tags1', 'tags2'],
        })
        console.log('构造出的tr为', tr)
        const resp2 = await createTrForLedger(tr)
        console.log(resp2)
        const resp3 = await getAllTrsFromLedgerById(ledgerId)
        console.log(resp3)
        const resp4 = await deleteTrById(resp3.data[0].transaction_id)
        console.log(resp4)
    })
})