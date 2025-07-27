---
title: Obsidian
description: Create and manage your notes with Obsidian!
aliases: Obsidian
date: 2024-11-16
updated: 2024-12-08
tags: 
  - obsidian
---
<div align="center">
<a href="https://obsidian.md" target="_blank" rel="noopener noreferrer">
  <img src="/obsidian-logo.png" alt="Obsidian Logo" style="max-width: 20%; height: auto;" />
</a>
</div>
While you don't need to use Obsidian to write your blog posts, it's a great tool for managing your notes.

## What is Obsidian?

Obsidian is a powerful knowledge base that works on top of a local folder of plain text Markdown files. It's designed for creating and managing interconnected notes.

### Key Features

- 📝 **Plain Text** - All notes are stored as local Markdown files, i.e. they last forever
- 🔗 **Bidirectional Linking** - Create connections between notes like wikilinks
- 🎨 **Graph View** - Visualize your knowledge network with knowledge graphs
- 🧩 **Plugin System** - Extend functionality with community plugins

### Example Note Structure

```md
# Project Planning

## Goals
- Define project scope
- Set milestones
- Assign resources

## Links
[[Resources]]
[[Timeline]]
[[TeamMembers]]

```

### Why Use Obsidian?

1. **Privacy First** - Your data stays on your device
2. **Future Proof** - Plain text files never become obsolete
3. **Flexible** - Adapt it to your workflow
4. **Community Driven** - Rich ecosystem of themes and plugins

## Getting Started

1. **Install Obsidian** - Download from [obsidian.md](https://obsidian.md/)
2. **Create a Vault** - Create a local folder for your notes. Try opening a vault in the `src/lib/content` directory.
3. **Start Writing** - Create your first note

## Next Steps

1. **Organize Your Notes** - Create folders for potential posts. All posts are currently placed in the `posts` folder of the `src/lib/content/{posts}` directory. This can be chaged in the `src/lib/content/posts/index.ts` file.
2. **Develop Templates** - Develop templates for your posts
3. **Develop SvelteKit Components** - Build components for displaying content using metadata and templates
4. **Publish Notes** - Place notes into the posts folder in Obsidian to publish them
5. **Test and Refine Workflow** - Regularly test the publishing process by running the SvelteKit development server
6. **Share and Collaborate**

Check out some examples of the posts you can make over at [Posts](/posts)
