<template>
    <div class="custom-select">
        <button @click="toggleDropdown" class="select-button">
            {{ selectedLabel || placeholder }}
        </button>
        <div v-show="isOpen" :class="['dropdown', direction === 'up' ? 'dropdown-up' : 'dropdown-down']">
            <div v-for="option in options" :key="option.value" class="option" @click="selectOption(option)">
                {{ option.label }}
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, watch, defineProps } from 'vue';

const selectedValue = defineModel()

// Props
const props = defineProps({
    options: {
        type: Array,
        required: true
    },
    placeholder: {
        type: String,
        default: '请选择'
    },
    direction: {
        type: String,
        default: 'down',
        validator: value => ['down', 'up'].includes(value)
    }
});

// Data
const isOpen = ref(false);
const selectedLabel = ref('');

// Watch modelValue to update selected label
watch(
    () => props.modelValue,
    (newValue) => {
        const selectedOption = props.options.find(opt => opt.value === newValue);
        selectedLabel.value = selectedOption ? selectedOption.label : '';
    },
    { immediate: true }
);

// Toggle dropdown
function toggleDropdown() {
    isOpen.value = !isOpen.value;
}

// Select option
function selectOption(option) {
    selectedValue.value = option.value;
    selectedLabel.value = option.label;
    isOpen.value = false;
    
}
</script>

<style scoped>
.custom-select {
    position: relative;
    width: 200px;
}

.select-button {
    width: 100%;
    padding: 10px;
    text-align: left;
    border: 1px solid #ccc;
    background: white;
    cursor: pointer;
}

.dropdown {
    position: absolute;
    width: 100%;
    max-height: 200px;
    overflow-y: auto;
    border: 1px solid #ccc;
    background: white;
    z-index: 100;
}

.dropdown-down {
    top: 100%;
    border-top: none;
    margin-top: 2px;
}

.dropdown-up {
    bottom: 100%;
    border-bottom: none;
    margin-bottom: 2px;
}

.option {
    padding: 10px;
    cursor: pointer;
}

.option:hover {
    background-color: #f0f0f0;
}
</style>