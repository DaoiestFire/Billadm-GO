<template>
  <a-select @change="handleChange" style="width: 120px" v-model:value="ledgerStore.currentLedgerName">
    <a-select-option
        v-for="option in options"
        :key="option.value"
        :value="option.value">
      {{ option.label }}
    </a-select-option>
  </a-select>
</template>

<script setup lang="ts">
import {computed} from 'vue';
import {useLedgerStore} from '@/stores/ledgerStore.ts';
import type {Ledger} from '@/types/billadm';

const ledgerStore = useLedgerStore();

const options = computed(() => {
  if (!Array.isArray(ledgerStore.ledgers)) return [];
  return ledgerStore.ledgers.map((ledger: Ledger) => ({
    label: ledger.name,
    value: ledger.id,
  }));
});

const handleChange = (value: string) => {
  ledgerStore.setCurrentLedger(value);
};
</script>
