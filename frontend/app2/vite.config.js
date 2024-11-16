import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vite.dev/config/
export default defineConfig({
  base: "/app2/",
  plugins: [react()],
  server: {
    port: 3002,
  },
});
