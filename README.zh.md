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
- [v1.4.378](https://github.com/danielmiessler/fabric/releases/tag/v1.4.378) (2026年1月14日) — **Digital Ocean GenAI 支持**：添加了对 Digital Ocean GenAI 的支持，以及[使用指南](./docs/DigitalOcean-Agents-Setup.md)。
- [v1.4.356](https://github.com/danielmiessler/fabric/releases/tag/v1.4.356) (2025年12月22日) — **完整的国际化支持**。
- [v1.4.350](https://github.com/danielmiessler/fabric/releases/tag/v1.4.350) (2025年12月18日) — **交互式 API 文档**：在 `/swagger/index.html` 添加了 Swagger/OpenAPI UI。
- [v1.4.338](https://github.com/danielmiessler/fabric/releases/tag/v1.4.338) (2025年12月4日) — 添加了 Abacus 供应商支持（[RouteLLM APIs](https://abacus.ai/app/route-llm-apis)）。
- [v1.4.337](https://github.com/danielmiessler/fabric/releases/tag/v1.4.337) (2025年12月4日) — 添加 "Z AI" 供应商支持（[Z AI overview](https://docs.z.ai/guides/overview/overview)）。
- [v1.4.334](https://github.com/danielmiessler/fabric/releases/tag/v1.4.334) (2025年11月26日) — **[Claude Opus 4.5](https://www.anthropic.com/news/claude-opus-4-5)** 支持。
- [v1.4.331](https://github.com/danielmiessler/fabric/releases/tag/v1.4.331) (2025年11月23日) — **GitHub Models 支持**。
- [v1.4.322](https://github.com/danielmiessler/fabric/releases/tag/v1.4.322) (2025年11月5日) — **交互式 HTML 概念图与 Claude Sonnet 4.5**。
- [v1.4.317](https://github.com/danielmiessler/fabric/releases/tag/v1.4.317) (2025年9月21日) — **葡萄牙语变体支持**：支持巴西葡萄牙语 (pt-BR) 和欧洲葡萄牙语 (pt-PT)。
- [v1.4.314](https://github.com/danielmiessler/fabric/releases/tag/v1.4.314) (2025年9月17日) — **Azure OpenAI 迁移**：迁移到官方 `openai-go/azure` SDK。
- [v1.4.311](https://github.com/danielmiessler/fabric/releases/tag/v1.4.311) (2025年9月13日) — **更多国际化支持**：添加了德语、波斯语、法语、意大利语、日语、葡萄牙语、中文。
- [v1.4.309](https://github.com/danielmiessler/fabric/releases/tag/v1.4.309) (2025年9月9日) — **全面的国际化支持**：包含英语和西班牙语语言文件。
- [v1.4.303](https://github.com/danielmiessler/fabric/releases/tag/v1.4.303) (2025年8月29日) — **新二进制版本**：Linux ARM 和 Windows ARM 目标，支持树莓派和 Windows Surface。
- [v1.4.294](https://github.com/danielmiessler/fabric/releases/tag/v1.4.294) (2025年8月20日) — **Venice AI 支持**：添加了隐私优先的开源 AI 提供商 Venice AI（["About Venice"](https://docs.venice.ai/overview/about-venice)）。
- [v1.4.291](https://github.com/danielmiessler/fabric/releases/tag/v1.4.291) (2025年8月18日) — **语音转文字**：添加了 OpenAI 语音转文字支持，包含 `--transcribe-file`、`--transcribe-model` 和 `--split-media-file` 标志。

这些功能代表了我们致力于使 Fabric 成为最强大、最灵活的 AI 增强框架的承诺！

</details>

## 介绍视频

请注意，以下很多视频是在 Fabric 还是基于 Python 的时代录制的，所以请务必使用下方最新的[安装指南](#安装指南)。

- [Network Chuck](https://www.youtube.com/watch?v=UbDyjIIGaxQ)
- [David Bombal](https://www.youtube.com/watch?v=vF-MQmVxnCs)
- [作者本人的工具介绍](https://www.youtube.com/watch?v=wPEyyigh10g)
- [更多关于 Fabric 的 YouTube 视频](https://www.youtube.com/results?search_query=fabric+ai)

## 目录

- [`fabric`](#fabric)
  - [什么是 Fabric 及其原因](#什么是-fabric-及其原因)
  - [更新日志](#更新日志)
    - [最近的主要功能](#最近的主要功能)
  - [介绍视频](#介绍视频)
  - [目录](#目录)
  - [变更记录](#变更记录)
  - [哲学理念](#哲学理念)
    - [将问题拆解成组件](#将问题拆解成组件)
    - [太多的 Prompt](#太多的-prompt-提示词)
  - [安装指南](#安装指南)
    - [一键安装（推荐）](#一键安装推荐)
    - [手动下载二进制文件](#手动下载二进制文件)
    - [使用包管理器](#使用包管理器)
      - [macOS (Homebrew)](#macos-homebrew)
      - [Arch Linux (AUR)](#arch-linux-aur)
      - [Windows](#windows)
      - [Windows (Scoop)](#windows-scoop)
    - [从源码构建](#从源码构建)
    - [Docker](#docker)
    - [环境变量](#环境变量)
    - [配置设置 (Setup)](#配置设置-setup)
    - [支持的 AI 供应商](#支持的-ai-供应商)
    - [按 Pattern 指定模型](#按-pattern-指定模型)
    - [为所有 Pattern 添加别名](#为所有-pattern-添加别名)
    - [迁移](#迁移)
    - [升级](#升级)
    - [Shell 补全](#shell-补全)
  - [使用方法](#使用方法)
    - [调试级别](#调试级别)
    - [演习模式](#演习模式)
    - [扩展](#扩展)
  - [REST API 服务器](#rest-api-服务器)
    - [Ollama 兼容模式](#ollama-兼容模式)
  - [我们的 Prompting 方法](#我们的-prompting-方法)
  - [示例](#示例)
  - [直接使用 Patterns](#直接使用-patterns)
    - [提示词策略](#提示词策略)
      - [可用策略](#可用策略)
  - [自定义 Patterns](#自定义-patterns)
    - [设置自定义 Patterns](#设置自定义-patterns)
    - [使用自定义 Patterns](#使用自定义-patterns)
    - [工作原理](#工作原理)
  - [辅助应用 (Helper Apps)](#辅助应用-helper-apps)
    - [`to_pdf`](#to_pdf)
    - [`to_pdf` 安装](#to_pdf-安装)
    - [`code2context`](#code2context)
    - [`generate_changelog`](#generate_changelog)
  - [pbpaste](#pbpaste)
  - [网页界面（Fabric Web App）](#网页界面fabric-web-app)
  - [元数据 (Meta)](#元数据-meta)
    - [主要贡献者](#主要贡献者)
    - [贡献者](#贡献者)
  - [💜 支持本项目](#-支持本项目)

<br />

## 变更记录

Fabric 正在快速演进。

请查阅 [CHANGELOG](./CHANGELOG.md) 了解所有最新变更。

## 哲学理念

### 将问题拆解成组件

我们在日常工作和生活中经常遇到的问题是难以自动化完成一件大而复杂的事情。

以"写一篇文章"为例。这很难做到，即使对 AI 而言也是如此。为什么呢？因为你要写什么？谁是你的受众？写作基调是怎样的？一旦写完，你打算把它放在哪里？你需要配合图片吗？

解决复杂系统的最佳方法是将它们分解成单一职责的组件（模块）。

### 太多的 Prompt (提示词)

这就引出了我们对 Prompt 的处理方法。我们不仅需要将问题分解为组件，我们还需要让这些组件具有独立存在的价值。并且我们需要为这些组件命名，这样我们就能快速找到它们。

在 Fabric 之前，你可能有很多 Prompt 散落在你的笔记、桌面文件中，或者只存在于你的脑海里。

> [!NOTE]
> Fabric 的核心是一个包含独立、针对特定问题的 Markdown 格式 Prompt 库，我们称之为 **`Patterns` (模式)**。

除了这套精心构建的 Prompt 之外，Fabric 还提供了一套原生于 Go 语言（过去是 Python）的命令行工具。

Fabric 拥有适用于各种生活和工作场景的 Patterns，包括：

- 提取 YouTube 视频和播客中最有趣的部分
- 仅凭一个想法，以你自己的写作风格写一篇文章
- 总结晦涩难懂的学术论文
- 为一段文字创作完美匹配的 AI 艺术提示词
- 评估内容质量，帮你决定是否值得阅读/观看全文
- 获取冗长无聊内容的摘要
- 向你解释代码
- 将糟糕的文档转化为可用的文档
- 从任意内容输入创建社交媒体帖子
- 以及更多……

## 安装指南

### 一键安装（推荐）

**Unix/Linux/macOS：**

```bash
curl -fsSL https://raw.githubusercontent.com/danielmiessler/fabric/main/scripts/installer/install.sh | bash
```

**Windows PowerShell：**

```powershell
iwr -useb https://raw.githubusercontent.com/danielmiessler/fabric/main/scripts/installer/install.ps1 | iex
```

> 请参阅 [scripts/installer/README.md](./scripts/installer/README.md) 了解自定义安装选项和故障排除。

### 手动下载二进制文件

最新发布的二进制存档及其 SHA256 哈希值可在 <https://github.com/danielmiessler/fabric/releases/latest> 找到。

### 使用包管理器

**注意：** 使用 Homebrew 或 Arch Linux 包管理器安装时，`fabric` 命令名称为 `fabric-ai`，请在 shell 配置文件中添加以下别名：

```bash
alias fabric='fabric-ai'
```

#### macOS (Homebrew)

```bash
brew install fabric-ai
```

#### Arch Linux (AUR)

```bash
yay -S fabric-ai
```

#### Windows

使用 Microsoft 官方支持的 `Winget` 工具：

```bash
winget install danielmiessler.Fabric
```

#### Windows (Scoop)

```bash
scoop install fabric-ai
```

### 从源码构建

安装 Fabric 前，请先[确保已安装 Go](https://go.dev/doc/install)，然后运行：

```bash
go install github.com/danielmiessler/fabric/cmd/fabric@latest
```

### Docker

使用预构建的 Docker 镜像运行 Fabric：

```bash
# 使用 Docker Hub 的最新镜像
docker run --rm -it kayvan/fabric:latest --version

# 使用 GHCR 的特定版本
docker run --rm -it ghcr.io/ksylvan/fabric:v1.4.305 --version

# 首次运行时进行配置
mkdir -p $HOME/.fabric-config
docker run --rm -it -v $HOME/.fabric-config:/home/appuser/.config/fabric kayvan/fabric:latest --setup

# 使用你的 patterns
docker run --rm -it -v $HOME/.fabric-config:/home/appuser/.config/fabric kayvan/fabric:latest -p summarize

# 运行 REST API 服务器
docker run --rm -it -p 8080:8080 -v $HOME/.fabric-config:/home/appuser/.config/fabric kayvan/fabric:latest --serve
```

**镜像来源：**

- Docker Hub：[kayvan/fabric](https://hub.docker.com/repository/docker/kayvan/fabric/general)
- GHCR：[ksylvan/fabric](https://github.com/ksylvan/fabric/pkgs/container/fabric)

请参阅 [scripts/docker/README.md](./scripts/docker/README.md) 了解自定义镜像构建和高级配置。

### 环境变量

在 Linux 上你可能需要在 `~/.bashrc`，在 macOS 上需要在 `~/.zshrc` 中设置环境变量，以便运行 `fabric` 命令。

Intel Mac 或 Linux：

```bash
# Golang 环境变量
export GOROOT=/usr/local/go
export GOPATH=$HOME/go

# 更新 PATH
export PATH=$GOPATH/bin:$GOROOT/bin:$HOME/.local/bin:$PATH
```

Apple Silicon Mac：

```bash
# Golang 环境变量
export GOROOT=$(brew --prefix go)/libexec
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$GOROOT/bin:$HOME/.local/bin:$PATH
```

Fabric 还支持通过环境变量配置语言和默认模型：

```bash
export FABRIC_LANG="zh"
export FABRIC_DEFAULT_MODEL="claude-sonnet-4-5-20251022"
```

### 配置设置 (Setup)

```bash
fabric --setup
```

### 支持的 AI 供应商

**原生集成：**

- OpenAI（包括 O1 和 O3 序列）
- OpenAI Codex
- Anthropic (Claude)
- Google Gemini
- Ollama（本地模型）
- Azure OpenAI
- Amazon Bedrock
- Vertex AI
- LM Studio
- Perplexity

**OpenAI 兼容供应商：**

- Abacus、AIML、Cerebras、DeepSeek、DigitalOcean、GitHub Models、GrokAI、Groq、Langdock、LiteLLM、MiniMax、Mistral、Novita AI、OpenRouter、SiliconCloud、Together、Venice AI、Z AI

运行 `fabric --setup` 配置首选供应商，或使用 `fabric --listvendors` 查看所有可用供应商。

### 按 Pattern 指定模型

你可以使用环境变量为单个 pattern 配置特定模型，格式为 `FABRIC_MODEL_PATTERN_NAME=vendor|model`。可以在 shell 启动文件中维护这些按 pattern 的模型映射。

### 为所有 Pattern 添加别名

在 `.zshrc` 或 `.bashrc` 中添加以下内容，可以直接使用 pattern 名称作为命令（例如，用 `summarize` 代替 `fabric --pattern summarize`）：

```bash
for pattern_file in $HOME/.config/fabric/patterns/*; do
    pattern_name="$(basename "$pattern_file")"
    alias_name="${FABRIC_ALIAS_PREFIX:-}${pattern_name}"
    alias_command="alias $alias_name='fabric --pattern $pattern_name'"
    eval "$alias_command"
done

yt() {
    if [ "$#" -eq 0 ] || [ "$#" -gt 2 ]; then
        echo "Usage: yt [-t | --timestamps] youtube-link"
        return 1
    fi
    transcript_flag="--transcript"
    if [ "$1" = "-t" ] || [ "$1" = "--timestamps" ]; then
        transcript_flag="--transcript-with-timestamps"
        shift
    fi
    local video_link="$1"
    fabric -y "$video_link" $transcript_flag
}
```

### 迁移

如果你已安装旧版（Python 版），以下是迁移到 Go 版本的步骤：

```bash
# 卸载旧版 Fabric
pipx uninstall fabric

# 清理旧的 Fabric 别名（检查 .bashrc、.zshrc 等）

# 安装 Go 版本
go install github.com/danielmiessler/fabric/cmd/fabric@latest

# 运行新版配置
fabric --setup
```

然后按照上方说明[设置环境变量](#环境变量)。

### 升级

得益于 Go 的特性，升级非常简单，只需运行安装时相同的命令即可获取最新版本：

```bash
go install github.com/danielmiessler/fabric/cmd/fabric@latest
```

### Shell 补全

Fabric 提供了 Zsh、Bash 和 Fish 的 shell 补全脚本，让 CLI 使用更加便捷。

**快速安装（无需克隆仓库）：**

```bash
curl -fsSL https://raw.githubusercontent.com/danielmiessler/Fabric/refs/heads/main/completions/setup-completions.sh | sh
```

**Zsh 补全：**

```bash
mkdir -p ~/.zsh/completions
cp completions/_fabric ~/.zsh/completions/
echo 'fpath=(~/.zsh/completions $fpath)' >> ~/.zshrc
echo 'autoload -Uz compinit && compinit' >> ~/.zshrc
```

**Bash 补全：**

```bash
echo 'source /path/to/fabric/completions/fabric.bash' >> ~/.bashrc
```

**Fish 补全：**

```bash
mkdir -p ~/.config/fish/completions
cp completions/fabric.fish ~/.config/fish/completions/
```

## 使用方法

配置完成后，运行以下命令查看帮助：

```bash
fabric -h
```

处理 YouTube 视频时，还可以使用以下视觉提取选项：

- `--visual`：使用 OCR 和 FFmpeg 从视频中提取视觉信息
- `--visual-sensitivity`：设置 FFmpeg 场景检测的容差（`0.0` - `1.0`）
- `--visual-fps`：按固定每秒帧数提取画面，而不是使用场景检测

将你复制的任何文本流式输入到 `fabric` 并选择你想应用的 Pattern：

```bash
pbpaste | fabric --pattern extract_wisdom
```

### 调试级别

使用 `--debug` 标志控制运行时日志：

- `0`：关闭（默认）
- `1`：基本调试信息
- `2`：详细调试
- `3`：追踪级别

### 演习模式

使用 `--dry-run` 预览将发送给 AI 模型的内容，而不实际发送请求：

```bash
echo "test input" | fabric --dry-run -p summarize
```

这对于调试 pattern、检查提示词构建以及在使用 API 配额之前验证输入格式非常有用。

### 扩展

Fabric 支持可在 pattern 中调用的扩展。请参阅 [Extension Guide](internal/plugins/template/Examples/README.md) 获取完整文档。

**重要提示：** 扩展只能在 pattern 文件中使用，不能通过直接 stdin 使用。

## REST API 服务器

Fabric 内置了 REST API 服务器，通过 HTTP 暴露所有核心功能。启动服务器：

```bash
fabric --serve
```

服务器提供以下端点：

- 带流式响应的聊天补全
- Pattern 管理（创建、读取、更新、删除）
- 上下文和会话管理
- 模型和供应商列表
- YouTube 字幕提取
- 配置管理

有关完整的端点文档、身份验证设置和使用示例，请参阅 [REST API 文档](docs/rest-api.md)。

### Ollama 兼容模式

Fabric 可以通过暴露 Ollama 兼容的 API 端点，作为 Ollama 的直接替代品：

```bash
fabric --serve --serveOllama
```

这将启用以下 Ollama 兼容端点：

- `GET /api/tags` — 将可用 patterns 列为模型
- `POST /api/chat` — 聊天补全
- `GET /api/version` — 服务器版本

配置为使用 Ollama API 的应用程序可以指向你的 Fabric 服务器，Patterns 将显示为模型（例如 `summarize:latest`）。

## 我们的 Prompting 方法

Fabric 的 _Patterns_ 与你见过的大多数提示词有所不同。

- **首先，我们使用 `Markdown` 以确保最大的可读性和可编辑性**。这不仅帮助创建者写出好的 pattern，也方便任何想深入理解它的人——重要的是，这也包括你发送给它的 AI！

- **其次，我们的指令极为清晰**，并使用 Markdown 结构来强调我们希望 AI 做什么以及按什么顺序做。

- **最后，我们几乎只使用提示词的 System 部分**。经过一年多的深入研究，我们发现这种方式效果更好。

## 示例

> 以下示例使用 macOS 的 `pbpaste` 从剪贴板粘贴内容。Windows 和 Linux 的替代方案请参阅下方的 [pbpaste](#pbpaste) 部分。

1. 基于 `stdin` 输入运行 `summarize` Pattern（例如文章正文）：

    ```bash
    pbpaste | fabric --pattern summarize
    ```

2. 使用 `--stream` 选项运行 `analyze_claims` Pattern，获取即时流式结果：

    ```bash
    pbpaste | fabric --stream --pattern analyze_claims
    ```

3. 对任意 YouTube 视频运行 `extract_wisdom` Pattern 并流式输出结果：

    ```bash
    fabric -y "https://youtube.com/watch?v=uXs-zPc63kM" --stream --pattern extract_wisdom
    ```

4. 创建 pattern：在 `~/.config/fabric/patterns/[yourpatternname]` 中创建一个 `.md` 文件即可。

5. 对网站运行 `analyze_claims` pattern（Fabric 使用 Jina AI 将 URL 抓取为 Markdown 格式）：

    ```bash
    fabric -u https://github.com/danielmiessler/fabric/ -p analyze_claims
    ```

## 直接使用 Patterns

如果你不想做任何复杂的事情，只是想要大量优质的提示词，可以直接浏览 [`/patterns`](https://github.com/danielmiessler/fabric/tree/main/data/patterns) 目录！

你可以在任何 AI 应用中使用这些 Patterns，无论是 ChatGPT 还是其他应用或网站。

### 提示词策略

Fabric 还实现了"思维链"或"草稿链"等提示词策略，可以与基本 patterns 结合使用。

每个策略都是 [`/strategies`](https://github.com/danielmiessler/fabric/tree/main/data/strategies) 目录中的一个小型 `json` 文件。

使用 `fabric -S` 并选择安装策略选项，将策略安装到 `~/.config/fabric` 目录。

#### 可用策略

- `cot` — 思维链：逐步推理
- `cod` — 草稿链：迭代起草，每步最多 5 个词
- `tot` — 思维树：生成多条推理路径并选择最佳
- `aot` — 思维原子：将问题分解为最小的独立子问题
- `ltm` — 由易到难：从最简单到最难的子问题依次解决
- `self-consistent` — 自洽性：多条推理路径取共识
- `self-refine` — 自我精炼：回答、批判、精炼
- `reflexion` — 反思：回答、简短批判、提供精炼答案
- `standard` — 标准：直接回答，不作解释

使用 `--strategy` 标志应用策略：

```bash
echo "分析这段代码" | fabric --strategy cot -p analyze_code
```

## 自定义 Patterns

你可能希望使用 Fabric 创建自己的自定义 Patterns，但不与他人分享。没问题！

Fabric 支持专用的自定义 patterns 目录，将你的个人 patterns 与内置的分开保存。这意味着当你更新 Fabric 的内置 patterns 时，你的自定义 patterns 不会被覆盖。

### 设置自定义 Patterns

1. 运行 Fabric 配置：

   ```bash
   fabric --setup
   ```

2. 从工具菜单中选择"Custom Patterns"选项，输入你想要的目录路径（例如 `~/my-custom-patterns`）。

3. 如果目录不存在，Fabric 会自动创建。

### 使用自定义 Patterns

1. 创建自定义 pattern 目录结构：

   ```bash
   mkdir -p ~/my-custom-patterns/my-analyzer
   ```

2. 创建 pattern 文件：

   ```bash
   echo "You are an expert analyzer of ..." > ~/my-custom-patterns/my-analyzer/system.md
   ```

3. **使用你的自定义 pattern：**

   ```bash
   fabric --pattern my-analyzer "分析这段文本"
   ```

### 工作原理

- **优先级系统**：自定义 patterns 优先于同名的内置 patterns
- **无缝集成**：自定义 patterns 与内置 patterns 一起出现在 `fabric --listpatterns` 中
- **更新安全**：你的自定义 patterns 不会被 `fabric --updatepatterns` 影响
- **默认私密**：自定义 patterns 保持私密，除非你主动分享

## 辅助应用 (Helper Apps)

Fabric 还提供了一些核心辅助工具，便于与各种工作流集成。

### `to_pdf`

`to_pdf` 是一个将 LaTeX 文件转换为 PDF 格式的辅助命令：

```bash
to_pdf input.tex
```

也可以与 `write_latex` pattern 结合使用 stdin：

```bash
echo "ai security primer" | fabric --pattern write_latex | to_pdf
```

### `to_pdf` 安装

```bash
go install github.com/danielmiessler/fabric/cmd/to_pdf@latest
```

请确保系统已安装 LaTeX 发行版（如 TeX Live 或 MiKTeX），因为 `to_pdf` 需要 `pdflatex` 在系统路径中可用。

### `code2context`

`code2context` 与 `create_coding_feature` pattern 配合使用。它生成代码目录的 `json` 表示，可以与指令一起输入 AI 模型，用于创建新功能或编辑代码。

详情请参阅 [create_coding_feature Pattern README](./data/patterns/create_coding_feature/README.md)。

```bash
go install github.com/danielmiessler/fabric/cmd/code2context@latest
```

### `generate_changelog`

`generate_changelog` 从 git 提交历史和 GitHub Pull Request 生成变更日志。它遍历仓库的 git 历史，提取 PR 信息，生成格式良好的 Markdown 变更日志。

功能包括：SQLite 缓存（快速增量更新）、GitHub GraphQL API 集成（高效获取 PR）以及使用 Fabric 的可选 AI 增强摘要。

```bash
go install github.com/danielmiessler/fabric/cmd/generate_changelog@latest
```

详细使用说明和选项请参阅 [generate_changelog README](./cmd/generate_changelog/README.md)。

## pbpaste

[示例](#示例)部分使用了 macOS 程序 `pbpaste`，将剪贴板内容通过管道输入到 `fabric`。`pbpaste` 在 Windows 或 Linux 上不可用，但有替代方案。

**Windows：** 使用 PowerShell 命令 `Get-Clipboard`，或在 PowerShell 配置文件中添加别名：

```powershell
Set-Alias pbpaste Get-Clipboard
```

**Linux：** 使用 `xclip -selection clipboard -o`。首先安装 `xclip`：

```sh
sudo apt update
sudo apt install xclip -y
```

然后在 `~/.bashrc` 或 `~/.zshrc` 中添加别名：

```sh
alias pbpaste='xclip -selection clipboard -o'
```

## 网页界面（Fabric Web App）

Fabric 现在包含一个内置的网页界面，提供命令行界面的 GUI 替代方案。请参阅 [Web App README](./web/README.md) 了解安装说明和功能概览。

## 元数据 (Meta)

> [!NOTE]
> 特别感谢以下人员的灵感和贡献！

- _Jonathan Dunn_ — 项目核心开发 MVP，主导了新的 Go 版本和 GUI 开发，同时还是一名全职医生！
- _Caleb Sima_ — 推动了将此项目公开的决定。
- _Eugen Eisler_ 和 _Frederick Ros_ — 对 Go 版本做出了宝贵贡献。
- _David Peters_ — 负责网页界面的工作。
- _Joel Parish_ — 对项目 GitHub 目录结构提供了非常有用的建议。
- _Joseph Thacker_ — 提出了 `-c` 上下文标志的想法。
- _Jason Haddix_ — 提出了使用链式 Pattern（stitch）在本地模型过滤内容后再发送到云端模型的想法。
- _Andre Guerra_ — 协助了众多组件，使事情更简单、更易维护。

### 主要贡献者

<a href="https://github.com/danielmiessler"><img src="https://avatars.githubusercontent.com/u/50654?v=4" title="Daniel Miessler" width="50" height="50" alt="Daniel Miessler"></a>
<a href="https://github.com/xssdoctor"><img src="https://avatars.githubusercontent.com/u/9218431?v=4" title="Jonathan Dunn" width="50" height="50" alt="Jonathan Dunn"></a>
<a href="https://github.com/sbehrens"><img src="https://avatars.githubusercontent.com/u/688589?v=4" title="Scott Behrens" width="50" height="50" alt="Scott Behrens"></a>
<a href="https://github.com/agu3rra"><img src="https://avatars.githubusercontent.com/u/10410523?v=4" title="Andre Guerra" width="50" height="50" alt="Andre Guerra"></a>

### 贡献者

<a href="https://github.com/danielmiessler/fabric/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=danielmiessler/fabric" alt="contrib.rocks" />
</a>

Made with [contrib.rocks](https://contrib.rocks).

`fabric` 由 <a href="https://danielmiessler.com/subscribe" target="_blank">Daniel Miessler</a> 于 2024 年 1 月创建。
<br /><br />
<a href="https://twitter.com/intent/user?screen_name=danielmiessler">![X (formerly Twitter) Follow](https://img.shields.io/twitter/follow/danielmiessler)</a>

## 💜 支持本项目

<div align="center">

<img src="https://img.shields.io/badge/Sponsor-❤️-EA4AAA?style=for-the-badge&logo=github-sponsors&logoColor=white" alt="Sponsor">

**我每年在开源上花费数百小时。如果你想支持这个项目，可以在[这里赞助我](https://github.com/sponsors/danielmiessler)。🙏🏼**

</div>

---

## 许可证

MIT

---

*本文档由 [@JasonYeYuhe](https://github.com/JasonYeYuhe) 翻译并维护。如果您发现任何翻译问题或需要增加新特性说明，欢迎提交 Issue 或与我联系。*
