import {fileURLToPath, URL} from 'node:url'

import {defineConfig, loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'
import {ImportMetaEnv} from './env'

// https://vite.dev/config/
export default defineConfig((config)=>{
   const env = loadEnv(config.mode,"./") as ImportMetaEnv
    console.log(env.VITE_SERVER_URL)
    return {
        plugins: [
            vue(),
            vueDevTools(),
        ],
        css: {
            preprocessorOptions: {
                less: {
                    additionalData: '@import "@/assets/var.less";',
                    javascriptEnabled: true,
                }
            }
        },
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url))
            },
        },
        server: {
            host: "0.0.0.0",
            port: 80,
            proxy: {
                "/api": {
                    target: env.VITE_SERVER_URL,
                    changeOrigin: true,
                }
            }
        }
    }
})
