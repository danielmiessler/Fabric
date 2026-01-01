# @tailwindcss/forms

A plugin that provides a basic reset for form styles that makes form elements easy to override with utilities.

## Installation

Install the plugin from npm:

```sh
npm install -D @tailwindcss/forms
```

Then add the plugin to your `tailwind.config.js` file:

```js
// tailwind.config.js
module.exports = {
  theme: {
    // ...
  },
  plugins: [
    require('@tailwindcss/forms'),
    // ...
  ],
}
```

## Basic usage

[**View the live demo**](https://tailwindcss-forms.vercel.app/)

All of the basic form elements you use will now have some simple default styles that are easy to override with utilities.

Currently we add basic utility-friendly form styles for the following form element types:

- `input[type='text']`
- `input[type='password']`
- `input[type='email']`
- `input[type='number']`
- `input[type='url']`
- `input[type='date']`
- `input[type='datetime-local']`
- `input[type='month']`
- `input[type='week']`
- `input[type='time']`
- `input[type='search']`
- `input[type='tel']`
- `input[type='checkbox']`
- `input[type='radio']`
- `select`
- `select[multiple]`
- `textarea`

Every element has been normalized/reset to a simple visually consistent style that is easy to customize with utilities, even elements like `<select>` or `<input type="checkbox">` that normally need to be reset with `appearance: none` and customized using custom CSS:

```html
<!-- You can actually customize padding on a select element now: -->
<select class="rounded-full px-4 py-3">
  <!-- ... -->
</select>

<!-- Or change a checkbox color using text color utilities: -->
<input type="checkbox" class="rounded text-pink-500" />
```

More customization examples and best practices coming soon.

### Using classes to style

In addition to the global styles, we also generate a set of corresponding classes which can be used to explicitly apply the form styles to an element. This can be useful in situations where you need to make a non-form element, such as a `<div>`, look like a form element.

```html
<input type="email" class="form-input rounded-full px-4 py-3" />

<select class="form-select rounded-full px-4 py-3">
  <!-- ... -->
</select>

<input type="checkbox" class="form-checkbox rounded text-pink-500" />
```

Here is a complete table of the provided `form-*` classes for reference:

| Base                      | Class              |
| ------------------------- | ------------------ |
| `[type='text']`           | `form-input`       |
| `[type='email']`          | `form-input`       |
| `[type='url']`            | `form-input`       |
| `[type='password']`       | `form-input`       |
| `[type='number']`         | `form-input`       |
| `[type='date']`           | `form-input`       |
| `[type='datetime-local']` | `form-input`       |
| `[type='month']`          | `form-input`       |
| `[type='search']`         | `form-input`       |
| `[type='tel']`            | `form-input`       |
| `[type='time']`           | `form-input`       |
| `[type='week']`           | `form-input`       |
| `textarea`                | `form-textarea`    |
| `select`                  | `form-select`      |
| `select[multiple]`        | `form-multiselect` |
| `[type='checkbox']`       | `form-checkbox`    |
| `[type='radio']`          | `form-radio`       |

### Using only global styles or only classes

Although we recommend thinking of this plugin as a "form reset" rather than a collection of form component styles, in some cases our default approach may be too heavy-handed, especially when integrating this plugin into existing projects.

If generating both the global (base) styles and classes doesn't work well with your project, you can use the `strategy` option to limit the plugin to just one of these approaches.

```js
// tailwind.config.js
plugins: [
  require("@tailwindcss/forms")({
    strategy: 'base', // only generate global styles
    strategy: 'class', // only generate classes
  }),
],
```

When using the `base` strategy, form elements are styled globally, and no `form-{name}` classes are generated.

When using the `class` strategy, form elements are not styled globally, and instead must be styled using the generated `form-{name}` classes.
