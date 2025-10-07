<template>
  <div class="app-icon-display">
    <!-- 内框 (包含SVG图标) -->
    <div class="inner-frame">
      <!-- 动态渲染SVG图标 -->
      <div class="icon-content" v-html="iconSvg"/>
    </div>
  </div>
</template>

<script setup>
import {computed} from 'vue'

// 处理传入的SVG字符串
const iconSvg = computed(() => {
  return `<svg
   width="256"
   height="256"
   viewBox="0 0 67.733333 67.733333"
   version="1.1"
   id="svg1"
   inkscape:version="1.3.2 (091e20e, 2023-11-25, custom)"
   sodipodi:docname="billadm.svg"
   xmlns:inkscape="http://www.inkscape.org/namespaces/inkscape"
   xmlns:sodipodi="http://sodipodi.sourceforge.net/DTD/sodipodi-0.dtd"
   xmlns:xlink="http://www.w3.org/1999/xlink"
   xmlns="http://www.w3.org/2000/svg"
   xmlns:svg="http://www.w3.org/2000/svg">
  <sodipodi:namedview
     id="namedview1"
     pagecolor="#000000"
     bordercolor="#000000"
     borderopacity="0.25"
     inkscape:showpageshadow="2"
     inkscape:pageopacity="0.0"
     inkscape:pagecheckerboard="true"
     inkscape:deskcolor="#d1d1d1"
     inkscape:document-units="px"
     inkscape:zoom="3.2480499"
     inkscape:cx="169.33237"
     inkscape:cy="136.54347"
     inkscape:window-width="3840"
     inkscape:window-height="2054"
     inkscape:window-x="-11"
     inkscape:window-y="-11"
     inkscape:window-maximized="1"
     inkscape:current-layer="layer1"
     borderlayer="true" />
  <defs
     id="defs1">
    <color-profile
       name="Agfa-:-Swop-Standard"
       xlink:href="file:///C:/WINDOWS/system32/spool/drivers/color/RSWOP.icm"
       id="color-profile2" />
  </defs>
  <g
     inkscape:label="图层 1"
     inkscape:groupmode="layer"
     id="layer1">
    <rect
       style="fill:#409eff;fill-opacity:1;stroke:#409eff;stroke-width:0;stroke-dasharray:none;stroke-opacity:1"
       id="rect1"
       width="67.73333"
       height="67.73333"
       x="1.6135253e-06"
       y="1.6135253e-06"
       rx="11.90625"
       ry="11.90625" />
    <text
       xml:space="preserve"
       style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:67.7333px;font-family:Arial;-inkscape-font-specification:Arial;fill:#ffffff;fill-opacity:1;stroke:#409eff;stroke-width:0;stroke-dasharray:none;stroke-opacity:1"
       x="10.599871"
       y="58.1091"
       id="text3"><tspan
         sodipodi:role="line"
         id="tspan3"
         style="font-style:normal;font-variant:normal;font-weight:normal;font-stretch:normal;font-size:67.7333px;font-family:Arial;-inkscape-font-specification:Arial;fill:#ffffff;fill-opacity:1;stroke-width:0"
         x="10.599871"
         y="58.1091">B</tspan></text>
  </g>
</svg>`
})
</script>

<style scoped>
.app-icon-display {
  /* 外框容器 */
  display: inline-flex;
  justify-content: center;
  align-items: center;
  width: 40px;
  height: 40px;
}

.inner-frame {
  /* 内框 - 包含图标的区域 */
  width: 26px;
  height: 26px;
  display: flex;
  justify-content: center;
  align-items: center;
  box-sizing: border-box;
}

.icon-content {
  /* 图标内容容器 */
  width: 100%; /* 图标占内框的80%，可根据需要调整 */
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
}

.icon-content svg {
  /* 确保SVG填满容器 */
  width: 100%;
  height: 100%;
  fill: currentColor; /* 让SVG填充色跟随当前文本颜色 */
}
</style>