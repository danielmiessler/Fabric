# Project Restructuring Plan

Based on discussion in <https://github.com/danielmiessler/fabric/issues/1127>

This plan synthesizes the proposal by `ksylvan` with key clarifications and additions from `jaredmontoya`, `eugeis`, and others in the thread. The goal is to reorganize the project to align with standard Go conventions, reduce root-level clutter, and improve overall clarity for developers.

**Revision 2 Changes:** Added support for additional binary tools (`code_helper` and `to_pdf`) in the `cmd/` directory structure.

---

## Rationale for Restructuring

The current project structure mixes application code, web assets, scripts, and configuration files at the top level. This creates several challenges:

* **Top-Level Clutter:** A large number of files and directories in the root makes it difficult to quickly understand the project's structure and entry points.
* **Non-Idiomatic Go Structure:** Go source code is spread across multiple top-level directories. Standard Go practice places main application code in `cmd/` and private, non-reusable package code in `internal/`.
* **Mixed Concerns:** Application code (Go), web frontend code (Svelte), data processing scripts (Python), and infrastructure configuration (`nix`, `Dockerfile`) are intermingled, obscuring the separation of concerns.

The proposed restructure addresses these issues by organizing the project by function, adhering to community best practices.

---

## Proposed Final Directory Structure

This is the high-level view of the proposed structure. It incorporates the core plan and ensures that technically required files like `flake.nix` remain in the root directory.

```markdown
.
├── cmd
│   ├── fabric
│   │   └── main.go              # Main application entrypoint
│   ├── code_helper
│   │   ├── main.go              # Code analysis helper tool
│   │   └── code.go              # Supporting code for code_helper
│   └── to_pdf
│       └── main.go              # LaTeX to PDF conversion tool (renamed from to_pdf.go)
├── internal
│   ├── cli                      # All CLI-related code
│   ├── core                     # Core application logic (e.g., chatter)
│   ├── domain                   # Domain types, moved from 'common'
│   ├── patterns                 # Logic for loading/managing patterns
│   ├── plugins                  # All plugin logic (ai, db, etc.)
│   ├── server                   # The 'restapi' code, renamed for clarity
│   ├── tools                    # Non-binary tool utilities (converter, jina, youtube, etc.)
│   └── util                     # Specific, shared utilities (to be used sparingly)
├── data
│   ├── patterns/                # All pattern markdown files
│   └── strategies/              # All strategy json files
├── scripts
│   ├── docker
│   │   ├── Dockerfile
│   │   ├── docker-compose.yml
│   │   ├── start-docker.sh      # Helper script to start docker-compose stack
│   │   └── README.md            # Docker deployment documentation
│   ├── python_ui
│   │   ├── streamlit.py
│   │   └── requirements.txt
│   ├── pattern_generation
│   │   ├── extract_patterns.py
│   │   └── ...
│   └── setup_fabric.bat         # Windows setup script
├── docs
│   ├── images/
│   ├── NOTES.md
│   └── Pattern_Descriptions/    # Documentation about patterns
├── web/                         # (Svelte frontend, unchanged)
├── completions/                 # (Shell completions, unchanged)
├── nix/                         # (Nix environment, unchanged)
├── go.mod
├── go.sum
├── LICENSE
├── README.md
├── .gitignore
├── .envrc                       # (Must remain in root for direnv)
├── flake.nix                    # (Must remain in root for Nix)
└── flake.lock                   # (Must remain in root for Nix)
```

---

### Key Changes Explained

1. ✅ **Introduction of `cmd/` Directory**
    * **What:** All executable entry points are moved to `cmd/` subdirectories:
        * ✅ `cmd/fabric/main.go` - Main application entrypoint
        * ✅ `cmd/code_helper/` - Code analysis helper tool (moved from `plugins/tools/code_helper/`)
        * ✅ `cmd/to_pdf/` - LaTeX to PDF conversion tool (moved from `plugins/tools/to_pdf/`)
    * **Why:** This is a standard Go convention that clearly separates executable code from library code. It immediately shows new developers where all application entry points are and allows for additional binaries to be added cleanly in the future.

2. ✅ **Introduction of `internal/` Directory**
    * **What:** The majority of the Go packages (`cli`, `core`, `plugins`, `restapi`, `common`) are moved under `internal/`.
    * **Why:** The Go toolchain enforces that code within an `internal` directory can only be imported by code within the same repository. This makes the application's core logic private and prevents other projects from creating unintended dependencies on it, clarifying that the project is an application, not a public library.

3. ✅ **Reorganizing and Renaming Packages**
    * ✅ **`restapi` -> `internal/server`**: The package is renamed to describe its function (providing an HTTP server) rather than its implementation detail (REST).
    * ✅ **Dissolving `common`**: The `common` package has been broken up. Core data structures moved to a dedicated `internal/domain` package. Utility functions moved closer to the packages that use them, with truly shared utilities placed in `internal/util`.
    * ✅ **`patterns` logic**: Code for loading and managing patterns consolidated into `internal/patterns`.
    * ✅ **`plugins/tools` -> `internal/tools`**: Non-binary tool utilities (converter, jina, youtube, etc.) moved to `internal/tools` while binary tools moved to `cmd/`.

4. ✅ **Consolidating Data, Scripts, and Docs**
    * ✅ **`data/`**: The `patterns/` and `strategies/` directories, which are data assets consumed by the application, moved into a `data/` directory to distinguish them from source code.
    * ✅ **`scripts/`**: Helper scripts (Python, shell, batch files, Docker, etc.) grouped under `scripts/` to clarify their role as auxiliary tools:
        * ✅ `scripts/docker/` - Docker deployment files and helper scripts
        * ✅ `scripts/python_ui/` - Streamlit UI and Python dependencies
        * ✅ `scripts/pattern_generation/` - Pattern extraction and generation tools
    * ✅ **`docs/`**: Miscellaneous markdown files (`NOTES.md`, etc.) and related assets like images moved to `docs/` for better organization.

---

### Step-by-Step Migration Plan

1. ✅ **Create New Directories:** Create the new top-level directories: `cmd/fabric`, `cmd/code_helper`, `cmd/to_pdf`, `internal`, `data`, `scripts`, and `docs`.

2. ✅ **Move Binary Tools:**
    * ✅ Move `plugins/tools/code_helper/` to `cmd/code_helper/`
    * ✅ Move `plugins/tools/to_pdf/to_pdf.go` to `cmd/to_pdf/main.go` (rename file)
    * ✅ Move remaining non-binary tools from `plugins/tools/` to `internal/tools/`

3. ✅ **Move Go Packages:** Move the existing Go package directories (`cli`, `core`, `plugins`, `restapi`) into the new `internal/` directory.

4. ✅ **Refactor and Rename Go Packages:**
    * ✅ Rename `internal/restapi` to `internal/server`.
    * ✅ Break apart the `common` package, moving its contents into appropriate new locations like `internal/domain` and `internal/util`.

5. ✅ **Move Main Entry Point:** Move `main.go` to `cmd/fabric/main.go`.

6. ✅ **Update Go Imports:** This is a critical step. Use an IDE or tools like `goimports` to update all import paths in all `.go` files to reflect the new structure:
    * ✅ `.../fabric/cli` becomes `.../fabric/internal/cli`
    * ✅ `.../fabric/plugins/tools/...` becomes `.../fabric/internal/tools/...`
    * ✅ Update imports in all three binary tools (`fabric`, `code_helper`, `to_pdf`)

7. ✅ **Move Data Assets:** Move the `patterns/` and `strategies/` directories into the `data/` directory. Update the application code to read from these new paths.

8. ✅ **Move Scripts and Docs:**
    * ✅ Move Docker files (`Dockerfile`, `docker-compose.yml`) to `scripts/docker/` and create helper scripts and documentation
    * ✅ Move Python UI (`streamlit.py` and `requirements.txt`) to `scripts/python_ui/`
    * ✅ Move pattern generation scripts (`extract_patterns.py`, etc.) to `scripts/pattern_generation/`
    * ✅ Move batch files (`setup_fabric.bat`) and other helper scripts into `scripts/`
    * ✅ Move documentation files like `NOTES.md` and the `images` directory into `docs/`

9. ✅ **Update Build and CI/CD Processes:**
    * ✅ Review and update any build scripts, Makefiles, or CI/CD workflows that reference old paths
    * ✅ Update Dockerfile paths in `scripts/docker/Dockerfile` to reference new locations
    * ✅ Update GitHub Actions to build all three binaries: `./cmd/fabric`, `./cmd/code_helper`, `./cmd/to_pdf`
    * ✅ Update installation instructions in README.md to reflect new binary locations and Docker setup
    * ✅ Specifically, update the "Update Version File and Create Tag" GitHub Action to work with the new file structure

10. ✅ **Test and Validate:**
    * ✅ Run `go build ./cmd/fabric` to ensure the main application compiles correctly.
    * ✅ Run `go build ./cmd/code_helper` to ensure the code helper tool compiles correctly.
    * ✅ Run `go build ./cmd/to_pdf` to ensure the PDF tool compiles correctly.
    * ✅ Execute the full test suite with `go test ./...`.
    * ✅ Run all applications and manually test the CLI, API, pattern loading, and helper tools to confirm all functionality is intact.
    * ⚠️ Verify that external packaging and distribution methods, such as the Homebrew package, continue to build correctly after the reorganization. **Note:** The Homebrew formula will need to be updated to build from `./cmd/fabric` instead of the root directory.
    * ⚠️ Test that `go install github.com/danielmiessler/fabric/cmd/fabric@latest` works for all three tools.

---

## ✅ **RESTRUCTURING COMPLETE**

**Status:** All major restructuring tasks have been completed successfully as of the current PR.

**Summary of Achievements:**

* ✅ All 10 migration steps completed
* ✅ All 4 key structural changes implemented
* ✅ All binaries (`fabric`, `code_helper`, `to_pdf`) compile successfully
* ✅ Full test suite passes (all packages)
* ✅ Standard Go project layout achieved
* ✅ `internal/common` package successfully dissolved into `internal/domain` and `internal/util`
* ✅ All import paths updated to reflect new structure
* ✅ GitHub Actions workflows updated for new structure

**Remaining Tasks (⚠️):**

* External packaging verification (Homebrew, etc.) - requires separate testing
* `go install` command verification - requires publishing/tagging

### **Required Homebrew Formula Update**

I have a draft PR ready here: <https://github.com/Homebrew/homebrew-core/pull/229472>

The current Homebrew formula at `https://raw.githubusercontent.com/ksylvan/homebrew-core/refs/heads/main/Formula/f/fabric-ai.rb` will need to be updated to work with the new project structure:

**Current formula build command:**

```ruby
def install
  system "go", "build", *std_go_args(ldflags: "-s -w")
end
```

**Required update for new structure:**

```ruby
def install
  system "go", "build", *std_go_args(ldflags: "-s -w"), "./cmd/fabric"
end
```

**Additional considerations:**

* The formula currently builds from the root `main.go` (which no longer exists)
* After restructuring, it needs to build from `./cmd/fabric`
* The binary name and test commands should remain the same
* All three tools (`fabric`, `code_helper`, `to_pdf`) could potentially be packaged, but the main `fabric` binary is the primary target

**`go install` commands for new structure:**

```bash
# Main fabric tool
go install github.com/danielmiessler/fabric/cmd/fabric@latest

# Additional tools (if desired)
go install github.com/danielmiessler/fabric/cmd/code_helper@latest
go install github.com/danielmiessler/fabric/cmd/to_pdf@latest
```

The project now follows standard Go conventions and is ready for review and merge.
