import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
    plugins: [react()],
    build: {
        outDir: 'dist',
        assetsDir: 'assets',
        sourcemap: false,
        minify: 'esbuild', // Используем esbuild вместо terser
        rollupOptions: {
            output: {
                manualChunks: undefined,
            }
        }
    },
    server: {
        port: 3000,
        host: true
    }
})