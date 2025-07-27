# The Fabric Web App

- [The Fabric Web App](#the-fabric-web-app)
  - [Installing](#installing)
    - [From Source](#from-source)
      - [TL;DR: Convenience Scripts](#tldr-convenience-scripts)
  - [Tips](#tips)
  - [Obsidian](#obsidian)

This is a web app for Fabric. It was built using [Svelte][svelte], [SkeletonUI][skeleton], and [Mdsvex][mdsvex].

The goal of this app is to not only provide a user interface for Fabric, but also an out-of-the-box website for those who want to get started with web development, blogging, or to just have a web interface for fabric. You can use this app as a GUI interface for Fabric, a ready to go blog-site, or a website template for your own projects.

![Preview](./static/preview.png)

## Installing

There are a few days to install and run the Web UI.

### From Source

#### TL;DR: Convenience Scripts

To install the Web UI using `npm`, from the top-level directory:

```bash
./web/scripts/npm-install.sh
```

To use pnpm (preferred and recommended for a huge speed improvement):

```bash
./web/scripts/pnpm-install.sh
```

The app can be run by navigating to the `web` directory and using `npm install`, `pnpm install`, or your preferred package manager. Then simply run `npm run dev`, `pnpm run dev`, or your equivalent command to start the app. *You will need to run fabric in a separate terminal with the `fabric --serve` command.*

Using npm:

```bash
# Install the GUI and its dependencies
npm install
# Install PDF-to-Markdown components in this order
npm install -D patch-package
npm install -D pdfjs-dist
npm install -D github:jzillmann/pdf-to-markdown#modularize

npx svelte-kit sync

# Now, with "fabric --serve" running already, you can run the GUI
npm run dev
```

Using pnpm:

```bash
# Install the GUI and its dependencies
pnpm install
# Install PDF-to-Markdown components in this order
pnpm install -D patch-package
pnpm install -D pdfjs-dist
pnpm install -D github:jzillmann/pdf-to-markdown#modularize

pnpm exec svelte-kit sync

# Now, with "fabric --serve" running already, you can run the GUI
pnpm run dev
```

## Tips

When creating new posts make sure to include a date, description, tags, and aliases. Only a date is needed to display a note.

You can include images, tags to other articles, code blocks, and more all within your markdown files.

## Obsidian

If you choose to use Obsidian alongside this app,
you can design and order your vault however you like, though a `posts` folder should be kept in your vault to house any articles you'd like to post.

[svelte]: https://svelte.dev/
[skeleton]: https://skeleton.dev/
[mdsvex]: https://mdsvex.pngwn.io/
