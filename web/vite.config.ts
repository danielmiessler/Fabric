import { sveltekit } from '@sveltejs/kit/vite';
import tailwindcss from '@tailwindcss/vite';
import { defineConfig } from 'vite';

// Get the Fabric base URL from environment variable with fallback
const FABRIC_BASE_URL = process.env.FABRIC_BASE_URL || 'http://localhost:8080';

// CAUTION: Lightning CSS minifies the CSS, and it refuses a pseudo-class that
// comes after a pseudo-element. Do not use a Tailwind class that puts the two
// in that order, such as `file:disabled:opacity-50`, because Lightning CSS
// stops the build on the `::file-selector-button:disabled` selector that the
// class produces. See https://github.com/parcel-bundler/lightningcss/issues/1292
// The Skeleton 2 Tailwind plugin made that selector on its own, and Skeleton 5
// does not, so the project no longer needs a different minifier.

export default defineConfig({
  // Tailwind 4 gives a Vite plugin, which replaces the PostCSS setup. It must
  // come before sveltekit() so that it can process the styles in components.
  plugins: [tailwindcss(), sveltekit()],
  define: {
    'process.env': {
      NODE_ENV: JSON.stringify(process.env.NODE_ENV)
    },
    'process.platform': JSON.stringify(process.platform),
    'process.cwd': JSON.stringify('/'),
    'process.browser': true,
    'process': {
      cwd: () => ('/')
    },
    // Inject Fabric configuration for client-side access
    '__FABRIC_CONFIG__': {
      FABRIC_BASE_URL: JSON.stringify(FABRIC_BASE_URL)
    }
  },
  resolve: {
    alias: {
      process: 'process/browser'
    }
  },
  server: {
    fs: {
      allow: ['..']  // allows importing from the parent directory
    },
    proxy: {
      '/api': {
        target: FABRIC_BASE_URL,
        changeOrigin: true,
        timeout: 900000,
        rewrite: (path) => path.replace(/^\/api/, ''),
        configure: (proxy, _options) => {
          proxy.on('error', (err, req, res) => {
            console.log('proxy error', err);
            res.writeHead(500, {
              'Content-Type': 'text/plain',
            });
            res.end('Something went wrong. The backend server may not be running.');
          });
        }
      },
      '^/(patterns|models|sessions)/names': {
        target: FABRIC_BASE_URL,
        changeOrigin: true,
        timeout: 900000,
        configure: (proxy, _options) => {
          proxy.on('error', (err, req, res) => {
            console.log('proxy error', err);
            res.writeHead(500, {
              'Content-Type': 'application/json',
            });
            res.end(JSON.stringify({ error: 'Backend server not running', names: [] }));
          });
        }
      }
    },
    watch: {
      usePolling: true,
      interval: 100,
      ignored: ['**/node_modules/**', '**/dist/**', '**/.git/**', '**/.svelte-kit/**']
    }
  },
  build: {
    commonjsOptions: {
      transformMixedEsModules: true
    },
    target: 'esnext',
    minify: true,
    rollupOptions: {
      output: {
        format: 'es'
      }
    }
  }
});
