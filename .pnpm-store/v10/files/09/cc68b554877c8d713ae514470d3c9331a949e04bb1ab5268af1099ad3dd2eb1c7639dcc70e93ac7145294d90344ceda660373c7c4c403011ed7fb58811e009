# svelte-youtube-lite

A simple svelte component for creating YouTube embeds with a focus on performance and privacy (GDPR compliant).

- Loads the video thumbnail on page load
- Creates the embed when the thumbnail is clicked

## Installation

```sh
npm install -D svelte-youtube-lite
```

## Usage

_Run the [demo page](#run-demo-locally) locally for a preview_

### Import

```html
<script>
	import { Youtube } from 'svelte-youtube-lite';
</script>
```

### Minimal example

```html
<Youtube id="aYtE6XE6b_s" />
```

### With low quality thumbnail

```html
<Youtube id="aYtE6XE6b_s" thumbnail="mqdefault" />
```

### With custom iframe title

_(YouTube iframe API fallback uses the videos title as iframe title)_

```html
<Youtube id="aYtE6XE6b_s" title="Cute cat video" />
```

### Without title

```html
<Youtube id="aYtE6XE6b_s" showTitle="{false}" />
```

## Run demo locally

```sh
git clone https://github.com/radiofrance/svelte-youtube-lite.git
cd svelte-youtube-lite
npm i
npm run dev
```

## Todo

- [ ] support for full youtube urls (eg: with playlist and start time) ?
- [ ] support for youtube shorts ? change from 16:9 to vertical ?
- [ ] use DNS preconnect for all youtube iframe assets
- [ ] parameter (boolean) : load with intersection observer
- [ ] parameter (number) : start time
- [ ] parameter (string) : playlist id
- [ ] fallback to YouTube iframe API for browsers with bad autoplay support
