<div align="center">
    <a href="https://go.warp.dev/fabric" target="_blank">
        <sup>特别鸣谢：</sup>
        <br>
        <img alt="Warp sponsorship" width="400" src="https://raw.githubusercontent.com/warpdotdev/brand-assets/refs/heads/main/Github/Sponsor/Warp-Github-LG-02.png">
        <br>
        <h>Warp，专为多 AI 智能体编程而生</b>
        <br>
        <sup>支持 macOS, Linux 和 Windows</sup>
    </a>
</div>

<br>

<div align="center">

<img src="./docs/images/fabric-logo-gif.gif" alt="fabriclogo" width="400" height="400"/>

# `fabric`

![Static Badge](https://img.shields.io/badge/mission-human_flourishing_via_AI_augmentation-purple)
<br />
![GitHub top language](https://img.shields.io/github/languages/top/danielmiessler/fabric)
![GitHub last commit](https://img.shields.io/github/last-commit/danielmiessler/fabric)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/danielmiessler/fabric)

<div align="center">
<h4><code>fabric</code> 是一个使用 AI 来增强人类能力的开源框架。</h4>
</div>

<p align="center">
  <a href="README.md">English</a> ·
  <strong>中文</strong>
</p>

![Screenshot of fabric](./docs/images/fabric-summarize.png)

</div>

[更新日志](#更新日志) •
[什么是 Fabric 及其原因](#什么是-fabric-及其原因) •
[哲学理念](#哲学理念) •
[安装指南](#安装指南) •
[使用方法](#使用方法) •
[REST API](#rest-api-服务器) •
[示例](#示例) •
[直接使用 Patterns](#直接使用-patterns) •
[自定义 Patterns](#自定义-patterns) •
[辅助应用](#辅助应用-helper-apps) •
[元数据](#元数据-meta)

## 什么是 Fabric 及其原因

自 2022 年底现代 AI 兴起以来，我们看到了**海量**的用于完成任务的 AI 应用。有成千上万的网站、聊天机器人、移动应用和其他接口可供使用。

这一切都非常令人兴奋且强大，但是_将这些功能整合到我们的生活中并不容易_。

<div class="align center">
<h4>换句话说，AI 没有能力问题——它有的是<em>整合</em>问题。</h4>
</div>

**Fabric 的诞生就是为了解决这个问题，通过创建和组织 AI 的基本单元——Prompt（提示词）本身！**

Fabric 按照现实世界中的任务来组织 Prompt，允许人们在一个地方创建、收集和组织他们最重要的 AI 解决方案，以便在他们最喜欢的工具中使用。如果你喜欢命令行，你甚至可以直接把 Fabric 本身作为接口使用！

## 更新日志

若想深入了解 Fabric 及其内部机制，请阅读 [docs 文件夹](https://github.com/danielmiessler/Fabric/tree/main/docs) 中的文档。这里还有一个极其有用且定期更新的 Fabric [DeepWiki](https://deepwiki.com/danielmiessler/Fabric)。

<details>
<summary>点击查看最近更新</summary>

亲爱的用户：

我们在 Fabric 做了很多激动人心的事情，我想在这里做一个简短的总结，让您感受一下我们的开发速度！

以下是我们添加的**新功能和特性**（最新在前）：

### 最近的主要功能

- [v1.4.437](https://github.com/danielmiessler/fabric/releases/tag/v1.4.437) (2026年3月16日) — **OpenAI Codex 插件**：Fabric 现在支持使用 OpenAI Codex 作为后端（需订阅）！
- [v1.4.417](https://github.com/danielmiessler/fabric/releases/tag/v1.4.417) (2026年2月21日) — **Azure AI Gateway 插件**：添加了 Azure AI Gateway 插件，支持通过统一的 Azure APIM Gateway 和共享订阅密钥身份验证连接多个后端 (AWS Bedrock, Azure OpenAI, Google Vertex AI)。
- [v1.4.416](https://github.com/danielmiessler/fabric/releases/tag/v1.4.416) (2026年2月21日) — **Azure Entra ID 身份验证**：添加了带有共享 Azure 实用程序、Entra ID/MSAL 支持的认证插件，并将通用的 Azure 逻辑提取到可重用的 `azurecommon` 包中。
- [v1.4.380](https://github.com/danielmiessler/fabric/releases/tag/v1.4.380) (2026年1月15日) — **Microsoft 365 Copilot 集成**：添加了对企业版 Microsoft 365 Copilot 的支持。
- [v1.4.378](https://github.com/danielmiessler/fabric/releases/tag/v1.4.378) (2026年1月14日) — **Digital Ocean GenAI 支持**：添加了对 Digital Ocean GenAI 的支持。
- [v1.4.356](https://github.com/danielmiessler/fabric/releases/tag/v1.4.356) (2025年12月22日) — **完整的国际化支持**。
- [v1.4.350](https://github.com/danielmiessler/fabric/releases/tag/v1.4.350) (2025年12月18日) — **交互式 API 文档**：在 `/swagger/index.html` 添加了 Swagger/OpenAPI UI。
- [v1.4.338](https://github.com/danielmiessler/fabric/releases/tag/v1.4.338) (2025年12月4日) — 添加了 Abacus 供应商支持。
- [v1.4.337](https://github.com/danielmiessler/fabric/releases/tag/v1.4.337) (2025年12月4日) — 添加 "Z AI" 供应商支持。
- [v1.4.334](https://github.com/danielmiessler/fabric/releases/tag/v1.4.334) (2025年11月26日) — **Claude Opus 4.5** 支持。
- [v1.4.331](https://github.com/danielmiessler/fabric/releases/tag/v1.4.331) (2025年11月23日) — **GitHub Models 支持**。

这些功能代表了我们致力于使 Fabric 成为最强大、最灵活的 AI 增强框架的承诺！

</details>

## 介绍视频

请注意，以下很多视频是在 Fabric 还是基于 Python 的时代录制的，所以请务必使用下方最新的[安装指南](#安装指南)。

- [Network Chuck](https://www.youtube.com/watch?v=UbDyjIIGaxQ)
- [David Bombal](https://www.youtube.com/watch?v=vF-MQmVxnCs)
- [作者本人的工具介绍](https://www.youtube.com/watch?v=wPEyyigh10g)
- [更多关于 Fabric 的 YouTube 视频](https://www.youtube.com/results?search_query=fabric+ai)

## 哲学理念

### 将问题拆解成组件

我们在日常工作和生活中经常遇到的问题是难以自动化完成一件大而复杂的事情。

以“写一篇文章”为例。这很难做到，即使对 AI 而言也是如此。为什么呢？因为你要写什么？谁是你的受众？写作基调是怎样的？一旦写完，你打算把它放在哪里？你需要配合图片吗？

解决复杂系统的最佳方法是将它们分解成单一职责的组件（模块）。

### 太多的 Prompt (提示词)

这就引出了我们对 Prompt 的处理方法。我们不仅需要将问题分解为组件，我们还需要让这些组件具有独立存在的价值。并且我们需要为这些组件命名，这样我们就能快速找到它们。

在 Fabric 之前，你可能有很多 Prompt 散落在你的笔记、桌面文件中，或者只存在于你的脑海里。

> [!NOTE]
> Fabric 的核心是一个包含独立、针对特定问题的 Markdown 格式 Prompt 库，我们称之为 **`Patterns` (模式)**。

除了这套精心构建的 Prompt 之外，Fabric 还提供了一套原生于 Go 语言（过去是 Python）的命令行工具。

## 安装指南

### 一键安装 (推荐)

最快的启动方式是下载最新的安装脚本。

如果你在 macOS 或 Linux 机器上，只需运行以下命令：

```bash
curl -L https://github.com/danielmiessler/fabric/releases/latest/download/fabric-installer.sh > fabric-installer.sh && chmod +x fabric-installer.sh && ./fabric-installer.sh
```

这会自动：
1. 检测你的系统架构
2. 下载最新的发布版
3. 如果尚未安装，则安装 `ffmpeg` (某些功能需要)
4. 将二进制文件放入你的路径
5. 自动运行 `fabric --setup`

如果你正在使用 Windows 安装：请参考下面的 Windows 安装部分。

### 手动下载二进制文件

如果你不想运行一键脚本，你可以在 [GitHub Releases 页面](https://github.com/danielmiessler/fabric/releases) 直接下载二进制文件。

### 使用包管理器

#### macOS (Homebrew)

```bash
brew tap danielmiessler/fabric
brew install fabric
```

#### Arch Linux (AUR)

你可以使用你的 AUR 助手（如 `paru` 或 `yay`）安装 fabric：

```bash
paru -S fabric-cli
```

或者：

```bash
paru -S fabric-cli-git
```

#### Windows

1. 获取最新的 [Go 语言](https://go.dev/doc/install) 并在你的系统上安装。
2. 安装完 Go 之后，在你的终端中运行：
    ```bash
    go install github.com/danielmiessler/fabric@latest
    ```
3. 在你的路径中添加 `C:\Users\YOUR_USERNAME\go\bin`。
4. 现在你可以运行 `fabric --setup` 开始了。

> **注意：** 在 Windows 下执行 `--setup` 时可能偶尔会超时，这是已知的网络限制问题，重试即可。

### 从源码构建

因为这是原生的 Go 程序，所以安装过程只是常规的 Go 构建过程。

```bash
go build
```

如果你想将其安装到你的系统中：

```bash
go install
```

### Docker

由于所有的逻辑都是在本地执行，Docker 可以提供一个完全隔离的环境。

构建镜像：
```bash
docker build -t fabric .
```

然后作为容器运行：
```bash
docker run -it fabric --setup
```

### 环境变量

Fabric 拥有对环境感知的国际化系统。你可以通过设置以下变量来改变输出语言（例如，对于中文）：

```bash
export FABRIC_LANG="zh"
```

你也可以配置全局的默认模型：

```bash
export FABRIC_DEFAULT_MODEL="claude-3-7-sonnet-20250219"
```

### 配置设置 (Setup)

一旦你运行了 `--setup`，所有的基础环境就会准备就绪。
```bash
fabric --setup
```

### 支持的 AI 供应商

- OpenAI (包括 O1 和 O3 序列)
- Anthropic (Claude 3.5, 3.7)
- Google (Gemini 2.0 Flash / Pro)
- Ollama (本地模型支持)
- LMStudio
- Groq
- 以及数十个企业级网关支持...

## 使用方法

将你复制的任何文本流式输入到 `fabric` 并选择你想应用的 Pattern。

```bash
pbpaste | fabric --pattern extract_wisdom
```

这里有一个将文本总结出来的例子：

```bash
echo "Fabric is an open source framework..." | fabric -p summarize
```

---

## 许可证

MIT

---

*本文档由 [@JasonYeYuhe](https://github.com/JasonYeYuhe) 翻译并维护。如果您发现任何翻译问题或需要增加新特性说明，欢迎提交 Issue 或与我联系。*