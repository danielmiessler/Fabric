# Lucide Svelte

Implementation of the lucide icon library for svelte applications.

> What is lucide? Read it [here](https://github.com/lucide-icons/lucide#what-is-lucide).

## Installation

```sh
yarn add lucide-svelte
```

or

```sh
npm install lucide-svelte
```

## How to use

All the icons are Svelte components, that ouputs Svg elements. So each icon can be imported and used as a component. This also helps with the use of threeshaking so you only import the icons you use.

### Example

Default usage:

```sv
<script>
import { Skull } from 'lucide-svelte'
</script>

<Skull/>
```

You can pass additional props to adjust the icon.

```sv
<script>
import { Camera } from 'lucide-svelte'
</script>

<Camera />
```

### Available props

| name          | type     | default      |
| ------------- | -------- | ------------ |
| `size`        | _Number_ | 24           |
| `color`       | _String_ | currentColor |
| `strokeWidth` | _Number_ | 2            |
| `*<SVGProps>` | _String_ | -            |

- All SVGProps are available to style the svgs. See the list of SVG Presentation Attributes on [MDN](https://developer.mozilla.org/en-US/docs/Web/SVG/Attribute/Presentation)

### Example of custom props

```sv
<script>
import { Phone } from 'lucide-svelte'
</script>

<Phone fill="#333"/>
```

This results a filled phone icon.

### One generic icon component

It is possible to create one generic icon component to load icons.

> :warning: Example below importing all EsModules, caution using this example, not recommended when you bundle your application,the build size will grow strongly. Because it will import all the icons.

#### Icon Component Example

```svelte
<script>
  import * as icons from "lucide-svelte";
  export let name;
</script>

<svelte:component this={icons[name]} {...$$props}/>
```

##### Then you can use it like this

```svelte
<script>
  import LucideIcon from "./LucideIcon";
</script>

<LucideIcon name="Menu" />
```
