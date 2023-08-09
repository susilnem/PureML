import { sveltekit } from "@sveltejs/kit/vite"
import { defineConfig, loadEnv } from "vite"

/** @type {import('vite').UserConfig} */
export default function ({ mode }: any) {
  // Extends 'process.env.*' with VITE_*-variables from '.env.(mode=production|development)'
  process.env = { ...process.env, ...loadEnv(mode, process.cwd()) }

  return defineConfig({
    build: {
      sourcemap: true, // Config vite to generate sourcemap when bundling.
    },
    plugins: [sveltekit()],
  })
}
