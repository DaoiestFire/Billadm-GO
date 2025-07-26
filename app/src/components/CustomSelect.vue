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
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
}

.select-button {
    height: 28px;
    width: 100%;
    text-align: left;
    border: 1px solid var(--billadm-color-window-border-color);
    border-radius: 4px;
    background: white;
    cursor: pointer;
}

.dropdown {
    position: absolute;
    width: 100%;
    overflow-y: auto;
    border: 1px solid var(--billadm-color-window-border-color);
    border-radius: 4px;
    background: white;
}

.dropdown-down {
    top: 110%;
}

.dropdown-up {
    bottom: 120%;
}

.option {
    padding: 5px;
    cursor: pointer;
}

.option:hover {
    background-color: var(--billadm-color-icon-hover-bg-color);
}
</style>