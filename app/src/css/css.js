import { ref } from 'vue';

/**
 * 获取 CSS 变量的值
 * @param {string} variableName CSS 变量名，如 '--my-var'
 * @returns {string} 变量值
 */
function getCssVariable(variableName) {
  return getComputedStyle(document.documentElement)
    .getPropertyValue(variableName)
    .trim();
}

export function useCssVariables() {
  const minorBgColor = ref(getCssVariable('--billadm-color-minor-backgroud-color'));
  const hoverBgColor = ref(getCssVariable('--billadm-color-icon-hover-bg-color'));
  const iconColor = ref(getCssVariable('--billadm-color-icon-color'));

  return {
    minorBgColor,
    hoverBgColor,
    iconColor
  };
}