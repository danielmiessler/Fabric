# prism-svelte

Syntax highlighting for svelte code with [prismjs].

## install

```bash
npm i prism-svelte # or yarn add prism-svelte
```

## Usage

Import `prismjs` then import `prism-svelte` (the order is very important) and it should work:

```js
import Prism from 'prismjs';
import 'prism-svelte';

const source = `
<script>
  let count = 0;
</script>

<button on:click={ () => count++ }>Hello</button>

<h1>{ count }</h1>

<ul>
  {#each Array(10).map((_, i) => i) as }
    <li on:click={() => count = i}>Set count to {i}</li>
  {/each}
</ul>
`;

const highlighted = Prism.highlight(source, Prism.languages.svelte, 'svelte');
```

[prismjs]: https://prismjs.com/
