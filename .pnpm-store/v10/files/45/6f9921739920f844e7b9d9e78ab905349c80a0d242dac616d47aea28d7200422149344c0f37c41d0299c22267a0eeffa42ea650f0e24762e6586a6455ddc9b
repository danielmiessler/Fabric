![](https://cdn.sanity.io/images/mbh58i22/production/1f71d5306d82ec00b28d884a1d5482b92205988a-2560x1280.png)

# Svelte Reveal

![npm](https://img.shields.io/npm/v/svelte-reveal) ![npm](https://img.shields.io/npm/dw/svelte-reveal) ![GitHub](https://img.shields.io/github/license/davekeehl/svelte-reveal)

Svelte Reveal is a library created with the purpose of helping [Svelte](https://svelte.dev/) users add reveal on scroll animations to their web applications in the easiest way possible. This library leverages the [Intersection Observer API](https://developer.mozilla.org/en-US/docs/Web/API/Intersection_Observer_API) in order to know when to trigger the animations.

## Features

- ⚡️ Near zero config
- 👀 Intersection Observer
- 🧩 Customizable transitions
- 🔌 Extensive API
- 📚 Exhaustive documentation
- 🔥 100% TypeScript

## Table of Contents

1. [Demo](#demo)
1. [Usage](#usage)
1. [Why Svelte Reveal](#why-svelte-reveal)
1. [SvelteKit](#sveltekit)
1. [Options](#options)
1. [Global config](#global-config)
1. [API](#api)
1. [Suggestions](#suggestions)
1. [Troubleshooting](#troubleshooting)
1. [Funding](#funding)
1. [Versioning](#versioning)
1. [Changelog](#changelog)
1. [License](#license)

## Demo

You can see Svelte Reveal in action [in this StackBlitz project](https://stackblitz.com/edit/svelte-reveal?file=src%2FApp.svelte).

## Usage

1. Install the library:

   ```bash
   # npm
   npm install -D svelte-reveal
   
   # yarn
   yarn add -D svelte-reveal
   
   # pnpm
   pnpm add -D svelte-reveal

   # bun
   bun add -D svelte-reveal
   ```

2. Import the library in your Svelte component:

   ```svelte
   <script>
     import { reveal } from 'svelte-reveal';
   </script>
   ```

3. Add the imported `reveal` action to the DOM element you want to transition:

   ```svelte
   <h1 use:reveal>Hello world</h1>
   <p use:reveal={{ transition: "slide" }}>A paragraph</p>
   ```

   If you want to use the action on a Svelte component, you need to pass the reveal options via props:

   ```svelte
   // App.svelte
   
   <script>
     import Heading from './Heading.svelte';
   </script>
   
   <Heading useReveal={{ transition: "slide" }}>Hello world</Heading>
   ```
   
   ```svelte
   // Heading.svelte
   
   <script lang="ts">
     import { reveal, type RevealOptions } from 'svelte-reveal';
     export let useReveal: RevealOptions;
   </script>
   
   <h1 use:reveal={ useReveal }>
     <slot />
   </h1>
   ```
   
   Using [SvelteKit](https://kit.svelte.dev/)? Please refer to the ["SvelteKit"](#sveltekit) section.

## Why Svelte Reveal

If you happened to scout the internet for other similar libraries, you might have noticed that other authors have decided to create their own library using Svelte [slots](https://svelte.dev/docs#template-syntax-slot) (similar to [React children](https://react.dev/learn/passing-props-to-a-component#passing-jsx-as-children)). There is nothing wrong with that approach, but in my opinion it goes a bit against one of Svelte's core purpose: writing more concise code. Having to wrap every to-be-transitioned component adds at least 2 extra lines of code each time, making your files unnecessarily bloated for such a simple add-on.

You might have also noticed people adding event listeners to the window object in order to transition elements, but in terms of performance it [doesn't scale very well](https://itnext.io/1v1-scroll-listener-vs-intersection-observers-469a26ab9eb6).

Instead, I decided to use Svelte [actions](https://svelte.dev/docs#template-syntax-element-directives-use-action), which are functions you can attach to a DOM element and that allow you to get access to that particular element and hook into its lifecycle. They take up considerably fewer lines of code, and so far I haven't encountered any big obstacle or performance drawback. Morever, this library is backed by the Intersection Observer API, which is great for performance.

## SvelteKit

> ⚠️ This is an active area of research. Please [submit a bug report](https://github.com/DaveKeehl/svelte-reveal/issues/new) in case of issues.

> 📅 Last update: April 2024

Since Svelte actions were conceived to operate in a client-side environment, they don't always work 100% in SvelteKit and SSR (server-side rendering) out of the box. Svelte Reveal is no exception, as it needs DOM access, and in order not to incur in weird animation behaviors some small setup is required by the end-users. Out of the following two methods, pick the one that most suit your project requirements.

### Without SSR

If your page doesn't need to be server-side rendered then the fix is very trivial. Turn off `ssr` in your `+page.ts` file as follows.

```typescript
// +page.ts
export const ssr = false;
```

### With SSR

If your page does need to leverage server-side rendering, the setup remains easy but it requires a few more steps.

1. Import the bundled stylesheet in your page or layout
   ```svelte
   // +layout.svelte
   
   <script lang="ts">
     import "svelte-reveal/styles.css";
     ...
   </script>
   
   ...
   ```

2. Add the `sr__hide` css class to every element targeted by Svelte Reveal with `use:reveal`. This will prevent the elements to flicker as soon as the page is hydrated and the DOM is accessible to the library.
   ```svelte
   // +page.svelte
   
   <script lang="ts">
     import { reveal } from 'svelte-reveal';
   </script>
   
   <h1 use:reveal class="sr__hide">Hello world</h1>
   ```

## Options

Depending on the use case, you can either use this library as-is (which applies some [default options](./src/internal/default/options.ts)), or customize it to your liking. If you choose to do so, you can pass an object to this action containing your own options to be applied.

Keep in mind that these options are applied to the single DOM element you add Svelte Reveal to. For global and more in-depth settings, refer to the [API](#api) section.

| Name       | Type                        | Default             | Description                                                  |
| ---------- | --------------------------- | ------------------- | ------------------------------------------------------------ |
| `disable`    | `boolean`                   | `false`             | When set to false, the transition is disabled for the target element. |
| `root`       | `Element \| Document \| null` | `null`              | The [root](https://developer.mozilla.org/en-US/docs/Web/API/IntersectionObserver/root) element used by the Intersection Observer. |
| `rootMargin` | `string`                    | `"0px 0px 0px 0px"` | The [root margin](https://developer.mozilla.org/en-US/docs/Web/API/IntersectionObserver/rootMargin) property of the Intersection Observer. |
| `threshold`  | `number`                    | `0.6`               | The [threshold](https://developer.mozilla.org/en-US/docs/Web/API/IntersectionObserver/thresholds) (in percentage from `0.0` to `1.0`) property used by the Intersection Observer to know when its target element is considered visible. |
| `preset` | `"fade" \| "slide" \| "fly" \| "spin" \| "blur" \| "scale"`       | `"fade"`            | The transition preset that should be applied. Check out the ["presets"](#presets) subsection for more info. |
| `reset`      | `boolean`                   | `false`             | When set to `true`, the node transitions out when out of view, and is revealed again when back in view.<br /><br />⚠️ Be careful not to overuse this option, as it prevents the Intersection Observer to stop observing the target node. Performance is therefore not guaranteed when many elements have `reset` set to `true`. |
| `duration`   | `number`                    | `800`               | How long the transition lasts (in ms).                       |
| `delay`      | `number`                    | `0`                 | How long the transition is delayed (in ms) before being triggered. |
| `easing`     | `Easing`                    | `"easeInOutCubic"`  | The easing function to use. [Check out](./src/internal/types/easing.ts) the full list of available easing functions and [this other website](https://cubic-bezier.com/) to preview timing functions. |
| `x`          | `number`                    | `0`              | The starting offset position in pixels on the x-axis. |
| `y`          | `number`                    | `0`               | The starting offset position in pixels on the y-axis. |
| `rotate`     | `number`                    | `0`              | The starting rotation offset in degrees. |
| `opacity`    | `number`                    | `0`                 | The starting opacity value. |
| `blur`       | `number`                    | `0`              | The starting blur value in pixels. |
| `scale`      | `number`                    | `1`                | The starting scale value in percentage. `1` corresponds to `100%`. |

### Presets

> ⚠️ All presets have the `"fade"` preset backed in

Presets are sets of options with predefined values, packaged under a name to achieve a certain transition effect. The following table shows the presets that come bundled with Svelte Reveal and which options they map to.

| Name      | Options                       | Description                                                  |
| --------- | ----------------------------- | ------------------------------------------------------------ |
| `"fade"`  | `{ opacity: 0 }`              | The element fades in gracefully.<br />In practice: `opacity: 0 -> 1` |
| `"fly"`   | `{ opacity: 0, y: -20 }`      | The element fades in and moves along a translation on the y-axis.<br />In practice: `opacity: 0 -> 1` + `transform: translateY(-20px -> 0px) ` |
| `"slide"` | `{ opacity: 0, x: -20 }`      | The element fades in and performs a translation on the x-axis.<br />In practice: `opacity: 0 -> 1` + `transform: translateX(-20px -> 0px)` |
| `"blur"`  | `{ opacity: 0, blur: 2 }`     | The element fades in and becomes unblurred.<br />In practice: `opacity: 0 -> 1` + `filter: blur(2px -> 0px)` |
| `"scale"` | `{ opacity: 0, scale: 0.8 }`  | The element fades in and gets to the original size.<br />In practice: `opacity: 0 -> 1` + `transform: scale(0.8 -> 1)`<br /><br />⚠️ In order to use this transition it is required to use the `width` CSS property on the element to reveal. If you are not already using this property for other things, you can set it to `width: fit-content` . |
| `"spin"`  | `{ opacity: 0, rotate: -10 }` | The element fades in and gets to the original rotation degree.<br />In practice: `opacity: 0 -> 1` + `transform: rotate(-10 -> 0)`<br /><br />⚠️ In order to use this transition it is required to use the `width` CSS property on the element to reveal. If you are not already using this property for other things, you can use set it to `width: fit-content` . |

### Callbacks

The following table shows all the callback functions we provide to you.

| Name            | Args                  | Return | Description                                                  |
| --------------- | --------------------- | ------ | ------------------------------------------------------------ |
| `onRevealStart` | `(node: HTMLElement)` | `void` | Function that gets fired when the node starts being revealed. |
| `onRevealEnd`   | `(node: HTMLElement)` | `void` | Function that gets fired when the node is fully revealed.    |
| `onResetStart`  | `(node: HTMLElement)` | `void` | Function that gets fired when the` reset` option is set to `true` and the node starts transitioning out. |
| `onResetEnd`    | `(node: HTMLElement)` | `void` | Function that gets fired when the `reset` option is set to `true` and the node has fully transitioned out. |
| `onMount`       | `(node: HTMLElement)` | `void` | Function that gets fired when the node is mounted on the DOM. |
| `onUpdate`      | `(node: HTMLElement)` | `void` | Function that gets fired when the action options are updated. |
| `onDestroy`     | `(node: HTMLElement)` | `void` | Function that gets fired when the node is unmounted from the DOM. |

## Global config

The following table shows how this library is globally configured right of out the box.

| Parameter    | (children) | (children)   | Type           | Default | Description                                                  |
| ------------ | ---------- | ------------ | -------------- | ------- | ------------------------------------------------------------ |
| `once`       |            |              | `boolean`      | `false` | Whether the reveal effect runs only once (i.e. it doesn't re-run on page reload). |
| `responsive` |            |              | `Responsive`   |         | Specifies how the library handles responsiveness. It can be used to enable/disable the reveal effect on some devices. |
|              | `mobile`   |              | `DeviceConfig` |         | Configuration of mobile devices.                             |
|              |            | `enabled`    | `boolean`      | `true`  | Whether the reveal effect is performed on mobile devices.    |
|              |            | `breakpoint` | `number`       | `425`   | The max viewport width of mobile devices.                    |
|              | `tablet`   |              | `DeviceConfig` |         | Configuration of tablet devices.                             |
|              |            | `enabled`    | `boolean`      | `true`  | Whether the reveal effect is performed on tablet devices.    |
|              |            | `breakpoint` | `number`       | `768`   | The max viewport width of tablet devices.                    |
|              | `laptop`   |              | `DeviceConfig` |         | Configuration of laptop devices.                             |
|              |            | `enabled`    | `boolean`      | `true`  | Whether the reveal effect is performed on laptop devices.    |
|              |            | `breakpoint` | `number`       | `1440`  | The max viewport width of laptop devices.                    |
|              | `desktop`  |              | `DeviceConfig` |         | Configuration of desktop devices.                            |
|              |            | `enabled`    | `boolean`      | `true`  | Whether the reveal effect is performed on desktop devices.   |
|              |            | `breakpoint` | `number`       | `2560`  | The max viewport width of desktop devices.                   |

## API

> ⚠️ If you want to customise the behavior of a single DOM node, you are supposed to use [options](#options).

Svelte Reveal also exposes several functions you can call to change the [default options](./src/internal/default/options.ts) and [global configuration](./src/internal/default/config.ts) of this library. Since these functions operate on a global level across all components using Svelte Reveal, you are supposed to only call them from a single file, otherwise you'll keep overriding the default options and global config from multiple points.

| Name                    | Args                                           | Return                       | Description                                                  |
| ----------------------- | ---------------------------------------------- | ---------------------------- | ------------------------------------------------------------ |
| `setOnce`               | `(once: boolean)`                              | `RevealConfig`               | Sets the reveal animations activation status on page reload. |
| `setDeviceStatus`       | `(device: Device, status: boolean)`            | `RevealConfig`               | Sets the status of a device.                                 |
| `setDevicesStatus`      | `(devices: Device[], status: boolean)`         | `RevealConfig`               | Sets the status of multiple devices.                         |
| `setDeviceBreakpoint`   | `(device: Device, breakpoint: number)`         | `RevealConfig`               | Sets the breakpoint of a device.                             |
| `setDevice`             | `(device: Device, settings: IDevice)`          | `RevealConfig`               | Sets the settings of a device.                               |
| `setResponsive`         | `(responsive: Responsive)`                     | `RevealConfig`               | Updates how responsiveness is handled by the library.        |
| `setObserverRoot`       | `(root: Element \| Document \| null)`            | `IntersectionObserverConfig` | Sets the Intersection Observer `root` element.               |
| `setObserverRootMargin` | `(rootMargin: string)`                         | `IntersectionObserverConfig` | Sets the Intersection Observer `rootMargin` property.        |
| `setObserverThreshold`  | `(threshold: number)`                          | `IntersectionObserverConfig` | Sets the Intersection Observer `threshold` property.         |
| `setObserverConfig`     | `(observerConfig: IntersectionObserverConfig)` | `IntersectionObserverConfig` | Sets the Intersection Observer configuration.                |
| `setConfig`             | `(userConfig: RevealConfig)`                   | `RevealConfig`               | Updates the library global configuration.                    |
| `setDefaultOptions`     | `(options: RevealOptions)`                     | `RevealOptions`              | Updates the default options to be used for the reveal effect. |

## Suggestions

If you need to considerably customize the behavior of this library, I suggest creating a dedicated file and to import it from the top-most component in the components tree of your project. Within that file you can then call the API functions to set global settings or shared transition properties.

```typescript
// reveal.config.ts

import { setDefaultOptions } from 'svelte-reveal';

setDefaultOptions({
  blur: 20,
  x: -50,
  duration: 3000
});
```

```svelte
// App.svelte

<script>
  import 'reveal.config';
</script>

<div>{ your markup goes here }</div>
```

## Troubleshooting

Feel free to [open a new issue](https://github.com/DaveKeehl/svelte-reveal/issues/new/choose) in case of any problems.

## Funding

[Want to buy me a coffee?](https://ko-fi.com/davekeehl)

## Versioning

This project uses [Semantic Versioning](https://semver.org/) to keep track of its version number.

## Changelog

[CHANGELOG](./CHANGELOG.md)

## License

[MIT](./LICENSE)
