import {defineConfig} from 'vite';
import vue from '@vitejs/plugin-vue';
import vitePluginRaw from 'vite-plugin-raw';
import {resolve} from 'path';
import AutoImport from 'unplugin-auto-import/vite';
import Components from 'unplugin-vue-components/vite';
import {ElementPlusResolver} from 'unplugin-vue-components/resolvers';

export default defineConfig({
    plugins: [
        vue(),
        vitePluginRaw({
            match: /\.svg$/,
            exclude: [new RegExp(resolve(__dirname, './src/assets/icons'))]
        }),
        AutoImport({
            resolvers: [ElementPlusResolver()],
        }),
        Components({
            resolvers: [ElementPlusResolver()],
        }),
    ],
    resolve: {
        alias: {
            '@': resolve(__dirname, './src'),
        }
    },
    test: {
        environment: 'jsdom',
        globals: true,
    }
});
