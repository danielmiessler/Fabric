# Fabric Documentation

Welcome to the Fabric documentation! This directory contains detailed guides and technical documentation for various features and components of Fabric.

## 📚 Available Documentation

### Core Features

**[Automated-Changelog-Usage.md](./Automated-Changelog-Usage.md)**
Complete guide for developers on using the automated changelog system. Covers the workflow for generating PR changelog entries during development, including setup, validation, and CI/CD integration.

**[YouTube-Processing.md](./YouTube-Processing.md)**
Comprehensive guide for processing YouTube videos and playlists with Fabric. Covers transcript extraction, comment processing, metadata retrieval, and advanced yt-dlp configurations.

**[Using-Speech-To-Text.md](./Using-Speech-To-Text.md)**
Documentation for Fabric's speech-to-text capabilities using OpenAI's Whisper models. Learn how to transcribe audio and video files and process them through Fabric patterns.

### User Interface & Experience

**[Desktop-Notifications.md](./Desktop-Notifications.md)**
Guide to setting up desktop notifications for Fabric commands. Useful for long-running tasks and multitasking scenarios with cross-platform notification support.

**[Shell-Completions.md](./Shell-Completions.md)**
Instructions for setting up intelligent tab completion for Fabric in Zsh, Bash, and Fish shells. Includes automated installation and manual setup options.

**[Gemini-TTS.md](./Gemini-TTS.md)**
Complete guide for using Google Gemini's text-to-speech features with Fabric. Covers voice selection, audio generation, and integration with Fabric patterns.

### Development & Architecture

**[Automated-ChangeLog.md](./Automated-ChangeLog.md)**
Technical documentation outlining the automated CHANGELOG system architecture for CI/CD integration. Details the infrastructure and workflow for maintainers.

**[Project-Restructured.md](./Project-Restructured.md)**
Project restructuring plan and architectural decisions. Documents the transition to standard Go conventions and project organization improvements.

**[NOTES.md](./NOTES.md)**
Development notes on refactoring efforts, model management improvements, and architectural changes. Includes technical details on vendor and model abstraction.

### Audio Resources

**[voices/README.md](./voices/README.md)**
Index of Gemini TTS voice samples demonstrating different AI voice characteristics available in Fabric.

## 🗂️ Additional Resources

### Configuration Files

- `./notification-config.yaml` - Example notification configuration

### Images

- `images/` - Screenshots and visual documentation assets
  - `fabric-logo-gif.gif` - Animated Fabric logo
  - `fabric-summarize.png` - Screenshot of summarization feature
  - `svelte-preview.png` - Web interface preview

## 🚀 Quick Start

New to Fabric? Start with these essential docs:

1. **[../README.md](../README.md)** - Main project README with installation and basic usage
2. **[Shell-Completions.md](./Shell-Completions.md)** - Set up tab completion for better CLI experience
3. **[YouTube-Processing.md](./YouTube-Processing.md)** - Learn one of Fabric's most popular features
4. **[Desktop-Notifications.md](./Desktop-Notifications.md)** - Get notified when long tasks complete

## 🔧 For Contributors

Contributing to Fabric? These docs are essential:

1. **[./CONTRIBUTING.md](./CONTRIBUTING.md)** - Contribution guidelines and setup
2. **[Automated-Changelog-Usage.md](./Automated-Changelog-Usage.md)** - Required workflow for PR submissions
3. **[Project-Restructured.md](./Project-Restructured.md)** - Understanding project architecture
4. **[NOTES.md](./NOTES.md)** - Current development priorities and patterns

## 📝 Documentation Standards

When adding new documentation:

- Use clear, descriptive filenames
- Include practical examples and use cases
- Update this README index with your new docs
- Follow the established markdown formatting conventions
- Test all code examples before publication

---

*For general help and support, see [./SUPPORT.md](./SUPPORT.md)*
