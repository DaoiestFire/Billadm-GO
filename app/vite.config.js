import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import vitePluginRaw from 'vite-plugin-raw';
import { resolve } from 'path';

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vitePluginRaw({
      match: /\.svg$/,
      exclude: [new RegExp(resolve(__dirname, './src/assets/icons'))]
    })],
  resolve: {
    alias: {
      '@': resolve(__dirname, './src'),
    }
  },
});
