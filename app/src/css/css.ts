import {ref} from 'vue'

/**
 * 获取 CSS 变量的值
 */
function getCssVariable(variableName: string) {
    return getComputedStyle(document.documentElement)
        .getPropertyValue(variableName)
        .trim()
}

export function useCssVariables() {
    const minorBgColor = ref(getCssVariable('--billadm-color-minor-background-color'))
    const hoverBgColor = ref(getCssVariable('--billadm-color-icon-hover-bg-color'))
    const iconColor = ref(getCssVariable('--billadm-color-icon-color'))
    const iconActiveFgColor = ref(getCssVariable('--billadm-color-icon-active-fg-color'))
    const positiveColor = ref(getCssVariable('--billadm-color-positive-color'))
    const negativeColor = ref(getCssVariable('--billadm-color-negative-color'))
    const uiSizeMenuWidth = ref(getCssVariable('--billadm-ui-size-menu-width'))

    return {
        minorBgColor,
        hoverBgColor,
        iconColor,
        iconActiveFgColor,
        positiveColor,
        negativeColor,
        uiSizeMenuWidth,
    }
}