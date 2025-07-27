---
title: Using Svelte in Markdown
description: Learn how to use your Svelte components in Markdown documents!
date: 2023-12-22
updated:
aliases: Using Svelte in Markdown and Markdown in Svelte
tags: [markdown, svelte, web-dev, docs, learn]
---
**Ref:** [Mdsvex](https://mdsvex.pngwn.io/docs#install-it)

Here are some examples illustrating how to use Mdsvex in a Svelte application:

**Example 1**: Basic Markdown with Svelte Component
Create a file named example.svx:

```markdown
---
title: "Interactive Markdown Example"
---

<script>
  import Counter from '../components/Counter.svelte';
</script>

# {title}

This is an example of combining Markdown with a Svelte component:

<Counter />
```

In this example:

- The frontmatter (--- sections) defines variables like title.
- A Svelte component Counter is imported and used inside the Markdown.

**Example 2**: Custom Layouts with Mdsvex
Assuming you have a layout component at src/lib/layouts/BlogLayout.svelte:
  
```svelte
<!-- BlogLayout.svelte -->
<script>
  export let title;
</script>

<div class="blog-post">
  <h1>{title}</h1>
  <slot />
</div>
```

Now, to use this layout in your Markdown:

```markdown
---
title: "My Favorite Layout"
layout: "../lib/layouts/BlogLayout.svelte"
---

## Markdown with Custom Layout

This Markdown file will be wrapped by the `BlogLayout`.
```

**Example 3:** Using Frontmatter Variables in Markdown

```markdown
---
author: "John Doe"
date: "2024-11-15"
---

# Blog Post

By {author} on {date}

Here's some markdown content. You can reference frontmatter values directly in the body.
```

**Example 4**: Interactive Elements in Markdown

```markdown
---
title: "Interactive Chart"
---

<script>
  import { Chart } from '../components/Chart.svelte';
</script>

# {title}

Below is an interactive chart:

<Chart />
```

## Setting Up Mdsvex

To make these work, you need to configure your SvelteKit project:

1. Install Mdsvex:

```bash
npm install -D mdsvex
```

2. Configure SvelteKit:

In your svelte.config.js:

```javascript
import adapter from '@sveltejs/adapter-auto';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';
import { mdsvex } from 'mdsvex';

/** @type {import('mdsvex').MdsvexOptions} */
const mdsvexOptions = {
  extensions: ['.svx'],
};

/** @type {import('@sveltejs/kit').Config} */
const config = {
  extensions: ['.svelte', '.svx'],
  preprocess: [
    vitePreprocess(),
    mdsvex(mdsvexOptions),
  ],
  kit: {
    adapter: adapter()
  }
};

export default config;
```

3. Create a Route for Markdown Files:

Place your .svx files in the src/routes directory or subdirectories, and SvelteKit will automatically handle them as routes.

These examples show how you can integrate Mdsvex into your Svelte application, combining the simplicity of Markdown with the interactivity of Svelte components. Remember, any Svelte component you want to use within Markdown must be exported from a .svelte file and imported in your .svx file.
