import {createLedgerByName, queryAllLedgers, getLedgerById} from "@/backend/ledger.js";
import {describe, it} from 'vitest';

describe('Ledger Utils', () => {
    it('test ledger utils', async () => {
        const resp1 = await createLedgerByName('test-ledger-name')
        console.log(resp1)
        const resp2 = await createLedgerByName('test-ledger-name2')
        console.log(resp2)
        const resp3 = await queryAllLedgers()
        console.log(resp3)
        const resp4 = await getLedgerById([resp1.data])
        console.log(resp4)
    })
})