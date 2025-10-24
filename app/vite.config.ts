import {fileURLToPath} from "node:url"
import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {ElementPlusResolver} from 'unplugin-vue-components/resolvers'


export default defineConfig({
    base: '/static',
    server: {
        port: 31945,
    },
    plugins: [
        vue(),
        AutoImport({
            dts: 'src/types/auto-imports.d.ts',
            resolvers: [ElementPlusResolver()]
        }),
        Components({
            dirs: ['src/components'],
            dts: 'src/types/components.d.ts',
            resolvers: [ElementPlusResolver()]
        }),
    ],
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url))
        }
    }
})
