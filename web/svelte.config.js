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
      // "powershell" is necessary for the code example in about/README.md.
      // Shiki gives an error and shows plain text if a language is not here.
      langs: ['javascript', 'typescript', 'svelte', 'markdown', 'bash', 'go', 'text', 'python', 'rust', 'c', 'c++', 'shell', 'ruby', 'json', 'html', 'css', 'java', 'sql', 'toml', 'yaml', 'powershell']
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
    // The Content Security Policy is a second barrier for the chat view,
    // which shows model output through {@html}. In 'auto' mode, SvelteKit
    // adds a nonce or a hash for its own inline scripts. Inline styles stay
    // permitted, because Tailwind and Shiki use them. The app has no form,
    // thus 'form-action' is 'none'.
    csp: {
      mode: 'auto',
      directives: {
        'script-src': ['self'],
        'object-src': ['none'],
        'base-uri': ['self'],
        'form-action': ['none'],
        'frame-ancestors': ['self'],
        // A remote image in a model reply can send chat data to its host.
        // Only the badge and screenshot hosts on the About and Posts pages
        // can serve images. The shields.io paths do not include /endpoint,
        // which fetches a URL from the request.
        'img-src': [
          'self',
          'data:',
          'https://img.shields.io/badge/',
          'https://img.shields.io/github/',
          'https://img.shields.io/twitter/',
          'https://github.com',
          'https://*.githubusercontent.com'
        ],
        'connect-src': ['self']
      }
    },
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
