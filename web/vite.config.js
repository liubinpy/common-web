import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import VueDevTools from 'vite-plugin-vue-devtools'

// https://vitejs.dev/config/
export default defineConfig({
  lintOnSave: false, //检查命名规范
  plugins: [vue(), VueDevTools()],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  server: {
    proxy: {
      // 把key的路径代理到target位置
      // detail: https://cli.vuejs.org/config/#devserver-proxy
      [process.env.VITE_BASE_API]: {
        // 需要代理的路径   例如 '/api'
        target: `${process.env.VITE_BASE_PATH}:${process.env.VITE_SERVER_PORT}/`, // 代理到 目标路径
        changeOrigin: false
      }
    }
  }
})
