<template>
  <a-select @change="handleChange" style="width: 120px">
    <a-select-option
        v-for="option in options"
        :key="option.value"
        :value="option.value">
      {{ option.label }}
    </a-select-option>
  </a-select>
</template>

<script setup lang="ts">
import {computed} from 'vue'
import {useLedgerStore} from '@/stores/ledgerStore'
import type {Ledger} from '@/types/billadm'

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

<style scoped>
.ledger-select-container {
  position: relative;
  display: inline-block;
}

.ledger-select {
  display: flex;
  border: 1px solid var(--billadm-color-window-border-color);
  border-radius: 8px;
  cursor: pointer;
  background-color: var(--billadm-color-major-background-color);
}

.ledger-select:hover {
  border-color: var(--billadm-color-positive-color);
}

.select-left {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 130px;
  user-select: none;
}

.select-right {
  display: flex;
  align-items: center;
  z-index: 1;
}

.dropdown {
  position: absolute;
  top: 110%;
  left: 0;
  border: 1px solid var(--billadm-color-window-border-color);
  border-radius: 4px;
  background-color: var(--billadm-color-major-background-color);
  overflow-y: auto;
  z-index: 3;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.dropdown-item {
  position: relative;
  display: flex;
  height: 30px;
  cursor: pointer;
}

.item-left {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  width: 130px;
}

.item-left:hover::after {
  content: '';
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 126px;
  height: 26px;
  border-radius: 4px;
  transition: background-color 0.3s ease;
  background-color: var(--billadm-color-icon-hover-bg-color);
  z-index: -1;
}

.item-right {
  display: flex;
  align-items: center;
}
</style>