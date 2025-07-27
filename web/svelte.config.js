import adapter from '@sveltejs/adapter-auto';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';
import { mdsvex } from 'mdsvex';
import rehypeSlug from 'rehype-slug';
import rehypeAutolinkHeadings from 'rehype-autolink-headings';
import rehypeExternalLinks from 'rehype-external-links';
import rehypeUnwrapImages from 'rehype-unwrap-images';
import { escapeSvelte } from 'mdsvex';
//import { fileURLToPath } from 'url';
//import { dirname, join } from 'path';
import { getSingletonHighlighter } from 'shiki'
import dracula from 'shiki/themes/dracula.mjs'

//const __filename = fileURLToPath(import.meta.url);
//const __dirname = dirname(__filename);

// Initialize Shiki highlighter
const initializeHighlighter = async () => {
  try {
    return await getSingletonHighlighter({
      themes: ['dracula'],
      langs: ['javascript', 'typescript', 'svelte', 'markdown', 'bash', 'go', 'text', 'python', 'rust', 'c', 'c++', 'shell', 'ruby', 'json', 'html', 'css', 'java', 'sql', 'toml', 'yaml']
    });
  } catch (error) {
    console.error('Failed to initialize Shiki highlighter:', error);
    return null;
  }
};

let shikiHighlighterPromise = initializeHighlighter();

/** @type {import('mdsvex').MdsvexOptions} */
const mdsvexOptions = {
  extensions: ['.md', '.svx'],
  smartypants: {
    quotes: true,
    ellipses: true,
    backticks: true,
    dashes: true,
  },
  highlight: {
    highlighter: async (code, lang) => {
      try {
        const highlighter = await shikiHighlighterPromise;
        if (!highlighter) {
          console.warn('Shiki highlighter not available, falling back to plain text');
          return `<pre><code>${code}</code></pre>`;
        }
        const html = escapeSvelte(highlighter.codeToHtml(code, { lang, theme: dracula }));
        return `{@html \`${html}\`}`;
      } catch (error) {
        console.error('Failed to highlight code:', error);
        return `<pre><code>${code}</code></pre>`;
      }
    }
  },
  rehypePlugins: [
    rehypeSlug,
    rehypeUnwrapImages,
    [rehypeAutolinkHeadings, {behavior: 'wrap'}],
    [rehypeExternalLinks, {
      target: '_blank',
      rel: ['nofollow', 'noopener', 'noreferrer']
    }]
  ],
};

/** @type {import('@sveltejs/kit').Config} */
const config = {
  extensions: ['.svelte', '.md', '.svx'],
  kit: {
    adapter: adapter({
      // You can add adapter-specific options here
      pages: 'build',
      assets: 'build',
      fallback: null,
      precompress: false,
      strict: true
    }),
    prerender: {
      handleHttpError: ({ path, referrer, message }) => {
        // Log the error for debugging
        console.warn(`HTTP error during prerendering: ${message}\nPath: ${path}\nReferrer: ${referrer}`);
        
        // ignore 404 for specific case
        if (path === '/not-found' && referrer === '/') {
          return;
        }

        // otherwise fail
        throw new Error(message);
      },
    },
  },
  preprocess: [
    vitePreprocess({
      script: true,
    }),
    mdsvex(mdsvexOptions)
  ],
};

export default config;
