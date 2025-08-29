# Publishing Fabric to WinGet - Complete Guide

## Overview

This document provides comprehensive instructions for publishing Fabric releases to the Windows Package Manager (WinGet) using a dedicated repository approach. This method keeps WinGet publishing separate from the main build process while automating package submissions.

## Table of Contents

1. [Background](#background)
2. [Prerequisites](#prerequisites)
3. [Approach: Dedicated fabric-winget Repository](#approach-dedicated-fabric-winget-repository)
4. [Setup Instructions](#setup-instructions)
5. [Workflow Configurations](#workflow-configurations)
6. [Alternative Approaches](#alternative-approaches)
7. [Troubleshooting](#troubleshooting)
8. [References](#references)

## Background

WinGet is Microsoft's official package manager for Windows 10 (1809+), Windows 11, and Windows Server 2025. Publishing to WinGet requires:

- Creating manifest files in YAML format
- Submitting pull requests to [microsoft/winget-pkgs](https://github.com/microsoft/winget-pkgs)
- Following Microsoft's package submission guidelines

## Prerequisites

### Required

1. **GitHub Account** with a forked [microsoft/winget-pkgs](https://github.com/microsoft/winget-pkgs) repository
2. **GitHub Personal Access Token (PAT)** with `public_repo` scope
3. **Existing Package in WinGet** - At least one version of Fabric must already exist in winget-pkgs
4. **Published Releases** - Fabric releases must be published (not draft) on GitHub

### Supported Installer Formats

- `.exe` - Windows executables
- `.msi` - Windows Installer packages
- `.msix` - Modern Windows app packages
- `.appx` - Windows app packages

Note: Script-based installers and fonts are not currently supported.

## Approach: Dedicated fabric-winget Repository

### Why This Approach?

**Advantages:**

- ✅ Complete separation from main build process
- ✅ No modifications needed in upstream fabric repository
- ✅ Easy monitoring and debugging
- ✅ Flexible triggering (scheduled, webhook, or manual)
- ✅ Reusable pattern for other projects

### Repository Structure

```
ksylvan/fabric-winget/
├── .github/
│   └── workflows/
│       ├── winget-publish.yml       # Main automated workflow
│       ├── monitor-releases.yml     # Poll-based monitoring
│       └── manual-publish.yml       # Manual trigger backup
├── README.md                         # Documentation
└── LICENSE                          # License file
```

## Setup Instructions

### Step 1: Fork microsoft/winget-pkgs

1. Navigate to [microsoft/winget-pkgs](https://github.com/microsoft/winget-pkgs)
2. Click "Fork" in the top-right corner
3. Select your account (e.g., `ksylvan`)
4. Keep the default name `winget-pkgs`

### Step 2: Create GitHub Personal Access Token

1. Go to GitHub Settings → [Developer settings → Personal access tokens → Tokens (classic)](https://github.com/settings/tokens)
2. Click "Generate new token (classic)"
3. Name: `WINGET_TOKEN`
4. Select scope: `public_repo`
5. Generate and copy the token

### Step 3: Create fabric-winget Repository

1. Create new repository: `ksylvan/fabric-winget`
2. Add repository secret:
   - Go to Settings → Secrets and variables → Actions
   - Add secret: `WINGET_TOKEN` with your PAT value

### Step 4: Add Workflow Files

Create the following workflow files in `.github/workflows/`:

## Workflow Configurations

### Option 1: Poll-Based Monitoring (Recommended)

**File:** `.github/workflows/monitor-releases.yml`

```yaml
name: Monitor and Publish Fabric Releases to WinGet
on:
  schedule:
    # Check every 6 hours for new releases
    - cron: '0 */6 * * *'

  workflow_dispatch:
    inputs:
      version:
        description: 'Specific version to publish (e.g., v1.4.302)'
        required: false
        type: string

jobs:
  check-and-publish:
    runs-on: windows-latest

    steps:
      - name: Check latest Fabric release
        id: check_release
        uses: actions/github-script@v7
        with:
          script: |
            const { data: release } = await github.rest.repos.getLatestRelease({
              owner: 'danielmiessler',
              repo: 'fabric'
            });

            console.log(`Latest release: ${release.tag_name}`);
            core.setOutput('tag', release.tag_name);
            core.setOutput('url', release.html_url);
            return release.tag_name;

      - name: Publish to WinGet
        uses: vedantmgoyal9/winget-releaser@v2
        with:
          identifier: danielmiessler.Fabric
          version: ${{ github.event.inputs.version || steps.check_release.outputs.tag }}
          token: ${{ secrets.WINGET_TOKEN }}
          fork-user: ksylvan
          max-versions-to-keep: 5  # Keep only 5 latest versions
```

### Option 2: Webhook-Triggered (Immediate Response)

**File:** `.github/workflows/winget-publish.yml`

```yaml
name: Publish Fabric to WinGet on Release
on:
  # Triggered via repository_dispatch from upstream
  repository_dispatch:
    types: [fabric-release]

  # Manual trigger as backup
  workflow_dispatch:
    inputs:
      tag:
        description: 'Release tag (e.g., v1.4.302)'
        required: true
        type: string
      url:
        description: 'Release URL (optional)'
        required: false
        type: string

jobs:
  publish-to-winget:
    runs-on: windows-latest

    steps:
      - name: Extract release information
        id: release_info
        run: |
          if [ "${{ github.event_name }}" == "repository_dispatch" ]; then
            echo "tag=${{ github.event.client_payload.tag }}" >> $GITHUB_OUTPUT
            echo "url=${{ github.event.client_payload.url }}" >> $GITHUB_OUTPUT
          else
            echo "tag=${{ github.event.inputs.tag }}" >> $GITHUB_OUTPUT
            echo "url=${{ github.event.inputs.url }}" >> $GITHUB_OUTPUT
          fi
        shell: bash

      - name: Publish to WinGet
        uses: vedantmgoyal9/winget-releaser@v2
        with:
          identifier: danielmiessler.Fabric
          release-tag: ${{ steps.release_info.outputs.tag }}
          token: ${{ secrets.WINGET_TOKEN }}
          fork-user: ksylvan
          installers-regex: '\.exe$|\.msi$'  # Only Windows installers
          max-versions-to-keep: 5
```

To trigger from upstream, add this to the fabric release workflow:

```yaml
- name: Trigger WinGet Publishing
  if: success()
  run: |
    curl -X POST \
      -H "Authorization: token ${{ secrets.WINGET_DISPATCH_TOKEN }}" \
      -H "Accept: application/vnd.github.v3+json" \
      https://api.github.com/repos/ksylvan/fabric-winget/dispatches \
      -d '{
        "event_type": "fabric-release",
        "client_payload": {
          "tag": "${{ github.ref_name }}",
          "url": "${{ github.event.release.html_url }}"
        }
      }'
```

### Option 3: Using WinGet-Releaser Action (Simplest)

**File:** `.github/workflows/auto-publish.yml`

```yaml
name: Auto-Publish to WinGet
on:
  workflow_dispatch:
    inputs:
      release_url:
        description: 'GitHub Release URL'
        required: true
        type: string

jobs:
  publish:
    runs-on: windows-latest
    steps:
      - name: Parse Release URL
        id: parse
        run: |
          URL="${{ github.event.inputs.release_url }}"
          # Extract tag from URL like https://github.com/danielmiessler/fabric/releases/tag/v1.4.302
          TAG=$(echo $URL | sed 's/.*\/tag\///')
          echo "tag=$TAG" >> $GITHUB_OUTPUT
        shell: bash

      - name: Publish to WinGet
        uses: vedantmgoyal9/winget-releaser@v2
        with:
          identifier: danielmiessler.Fabric
          release-tag: ${{ steps.parse.outputs.tag }}
          token: ${{ secrets.WINGET_TOKEN }}
          fork-user: ksylvan
```

## Alternative Approaches

### 1. GoReleaser Native WinGet Support

If you prefer to use GoReleaser's built-in WinGet support, you can configure it in `.goreleaser.yaml`:

```yaml
winget:
  - name: Fabric
    publisher: danielmiessler
    license: MIT
    homepage: https://github.com/danielmiessler/fabric
    short_description: open-source AI framework for augmenting humans
    repository:
      owner: ksylvan
      name: winget-pkgs
      branch: "ksylvan.Fabric-{{.Version}}"
      pull_request:
        enabled: true
        draft: true
        base:
          owner: microsoft
          name: winget-pkgs
          branch: master
```

**Note:** This requires running GoReleaser from your fork with appropriate credentials.

### 2. Manual Submission Using winget-create

Install and use Microsoft's official tool:

```bash
# Install winget-create
winget install Microsoft.WingetCreate

# Create new manifest
wingetcreate new --urls https://github.com/danielmiessler/fabric/releases/download/v1.4.302/fabric_Windows_x86_64.zip

# Update existing manifest
wingetcreate update --urls https://github.com/danielmiessler/fabric/releases/download/v1.4.302/fabric_Windows_x86_64.zip --version 1.4.302
```

### 3. Using Komac (Modern Alternative)

[Komac](https://github.com/russellbanks/Komac) is a modern, fast manifest creator written in Rust:

```bash
# Install Komac
winget install RussellBanks.Komac

# Update package
komac update --id danielmiessler.Fabric --version 1.4.302 --urls https://github.com/danielmiessler/fabric/releases/download/v1.4.302/fabric_Windows_x86_64.zip
```

## Troubleshooting

### Common Issues

1. **"Package not found" error**
   - Ensure at least one version exists in winget-pkgs
   - First submission must be done manually

2. **"Fork not found" error**
   - Verify fork exists under specified username
   - Check `fork-user` parameter matches your GitHub username

3. **"Unauthorized" error**
   - Verify PAT has `public_repo` scope
   - Check token hasn't expired
   - Ensure secret name matches workflow

4. **"No matching installers" error**
   - Check release has Windows binaries (.exe, .msi, .zip)
   - Verify `installers-regex` pattern if specified

### Validation

Test your setup:

```bash
# Check if package exists
winget search fabric

# View package info
winget show danielmiessler.Fabric

# Test installation
winget install danielmiessler.Fabric
```

## References

### Official Documentation

- [Windows Package Manager Documentation](https://learn.microsoft.com/en-us/windows/package-manager/)
- [Submit packages to Windows Package Manager](https://learn.microsoft.com/en-us/windows/package-manager/package/)
- [WinGet Manifest Schema](https://learn.microsoft.com/en-us/windows/package-manager/package/manifest)
- [microsoft/winget-pkgs Repository](https://github.com/microsoft/winget-pkgs)
- [microsoft/winget-cli Repository](https://github.com/microsoft/winget-cli)

### Tools

- [WinGet-Releaser Action](https://github.com/vedantmgoyal9/winget-releaser) - GitHub Action for automated publishing
- [WinGet-Releaser Marketplace](https://github.com/marketplace/actions/winget-releaser) - Action marketplace page
- [winget-create](https://github.com/microsoft/winget-create) - Microsoft's official manifest creator
- [Komac](https://github.com/russellbanks/Komac) - Modern manifest creator in Rust

### Community Resources

- [WinGet Package Request](https://github.com/microsoft/winget-pkgs/issues/new/choose) - Request new packages
- [WinGet CLI Releases](https://github.com/microsoft/winget-cli/releases) - Latest WinGet versions
- [GoReleaser WinGet Documentation](https://goreleaser.com/customization/winget/) - GoReleaser's WinGet support

### Related Projects Using WinGet-Releaser

- [PowerShell/Win32-OpenSSH](https://github.com/PowerShell/Win32-OpenSSH/blob/latestw_all/.github/workflows/winget.yml)
- [fastfetch-cli/fastfetch](https://github.com/fastfetch-cli/fastfetch/blob/dev/.github/workflows/winget.yml)
- [Chatterino/chatterino2](https://github.com/Chatterino/chatterino2/blob/master/.github/workflows/winget.yml)

## Summary

The recommended approach is to create a dedicated `fabric-winget` repository that monitors the main fabric repository for new releases and automatically publishes them to WinGet. This keeps the concerns separated, makes debugging easier, and requires no changes to the upstream repository.

Key steps:

1. Fork microsoft/winget-pkgs
2. Create fabric-winget repository
3. Add GitHub PAT as secret
4. Deploy monitoring workflow
5. Test with manual trigger

This approach has been successfully used by many projects and provides a reliable, maintainable solution for WinGet publishing.
