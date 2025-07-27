#!/usr/bin/env bash

set -e

# Get the directory where this script is located
top_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
base_name="$(basename "$top_dir")"
cd "$top_dir"/../.. || exit 1

# Check if bash version supports associative arrays
if [[ ${BASH_VERSION%%.*} -lt 4 ]]; then
    echo "This script requires bash 4.0 or later for associative arrays."
    echo "Current version: $BASH_VERSION"
    exit 1
fi

IMAGE_NAME="fabric-test-setup"
ENV_DIR="scripts/${base_name}/env"

# Test case descriptions
declare -A test_descriptions=(
    ["no-config"]="No APIs configured"
    ["gemini-only"]="Only Gemini configured (reproduces original issue)"
    ["openai-only"]="Only OpenAI configured"
    ["ollama-only"]="Only Ollama configured"
    ["bedrock-only"]="Only Bedrock configured"
    ["mixed"]="Mixed configuration (Gemini + OpenAI + Ollama)"
)

# Test case order for consistent display
test_order=("no-config" "gemini-only" "openai-only" "ollama-only" "bedrock-only" "mixed")

build_image() {
    echo "=== Building Docker image ==="
    docker build -f "${top_dir}/base/Dockerfile" -t "$IMAGE_NAME" .
    echo
}

check_env_file() {
    local test_name="$1"
    local env_file="$ENV_DIR/env.$test_name"

    if [[ ! -f "$env_file" ]]; then
        echo "Error: Environment file not found: $env_file"
        exit 1
    fi
}

run_test() {
    local test_name="$1"
    local description="${test_descriptions[$test_name]}"
    local env_file="$ENV_DIR/env.$test_name"

    check_env_file "$test_name"

    echo "===================="
    echo "Test: $description"
    echo "Config: $test_name"
    echo "Env file: $env_file"
    echo "===================="

    echo "Running test..."
    if docker run --rm \
        -e HOME=/home/testuser \
        -e USER=testuser \
        -v "$(pwd)/$env_file:/home/testuser/.config/fabric/.env:ro" \
        "$IMAGE_NAME" --listmodels 2>&1; then
        echo "✅ Test completed"
    else
        echo "❌ Test failed"
    fi
    echo
}

shell_into_env() {
    local test_name="$1"
    local description="${test_descriptions[$test_name]}"
    local env_file="$ENV_DIR/env.$test_name"

    check_env_file "$test_name"

    echo "===================="
    echo "Shelling into: $description"
    echo "Config: $test_name"
    echo "Env file: $env_file"
    echo "===================="
    echo "You can now run 'fabric -S' to configure, or 'fabric --listmodels' or 'fabric -L' to test."
    echo "Changes to .env will persist in $env_file"
    echo "Type 'exit' to return to the test runner."
    echo

    docker run -it --rm \
        -e HOME=/home/testuser \
        -e USER=testuser \
        -v "$(pwd)/$env_file:/home/testuser/.config/fabric/.env" \
        --entrypoint=/bin/sh \
        "$IMAGE_NAME"
}

interactive_mode() {
    echo "=== Interactive Mode ==="
    echo "Available test cases:"
    echo
    local i=1
    local cases=()
    for test_name in "${test_order[@]}"; do
        echo "$i) ${test_descriptions[$test_name]} ($test_name)"
        cases[i]="$test_name"
        ((i++))
    done
    echo "$i) Run all tests"
    echo "0) Exit"
    echo
    echo "Add '!' after number to shell into test environment (e.g., '1!' to shell into no-config)"
    echo

    while true; do
        read -r -p "Select test case (0-$i) [or 1!, etc. to shell into test environment]: " choice

        # Check for shell mode (! suffix)
        local shell_mode=false
        if [[ "$choice" == *"!" ]]; then
            shell_mode=true
            choice="${choice%!}"  # Remove the ! suffix
        fi

        if [[ "$choice" == "0" ]]; then
            if [[ "$shell_mode" == true ]]; then
                echo "Cannot shell into exit option."
                continue
            fi
            echo "Exiting..."
            exit 0
        elif [[ "$choice" == "$i" ]]; then
            if [[ "$shell_mode" == true ]]; then
                echo "Cannot shell into 'run all tests' option."
                continue
            fi
            echo "Running all tests..."
            run_all_tests
            break
        elif [[ "$choice" -ge 1 && "$choice" -lt "$i" ]]; then
            local selected_test="${cases[$choice]}"
            if [[ "$shell_mode" == true ]]; then
                echo "Shelling into: ${test_descriptions[$selected_test]}"
                shell_into_env "$selected_test"
            else
                echo "Running: ${test_descriptions[$selected_test]}"
                run_test "$selected_test"
            fi

            read -r -p "Continue testing? (y/n): " again
            if [[ "$again" != "y" && "$again" != "Y" ]]; then
                break
            fi
            echo
        else
            echo "Invalid choice. Please select 0-$i (optionally with '!' for shell mode)."
        fi
    done
}

run_all_tests() {
    echo "=== Testing PR #1645: Conditional API initialization ==="
    echo

    for test_name in "${test_order[@]}"; do
        run_test "$test_name"
    done

    echo "=== Test run complete ==="
    echo "Review the output above to check:"
    echo "1. No Ollama connection errors when OLLAMA_URL not set"
    echo "2. No Bedrock authentication errors when BEDROCK_AWS_REGION not set"
    echo "3. Only configured services appear in model listings"
}

show_help() {
    echo "Usage: $0 [OPTIONS] [TEST_CASE]"
    echo
    echo "Test PR #1645 conditional API initialization"
    echo
    echo "Options:"
    echo "  -h, --help       Show this help message"
    echo "  -i, --interactive Run in interactive mode"
    echo "  -b, --build-only  Build image only, don't run tests"
    echo "  -s, --shell TEST  Shell into test environment"
    echo
    echo "Test cases:"
    for test_name in "${test_order[@]}"; do
        echo "  $test_name: ${test_descriptions[$test_name]}"
    done
    echo
    echo "Examples:"
    echo "  $0                    # Run all tests"
    echo "  $0 -i                 # Interactive mode"
    echo "  $0 gemini-only        # Run specific test"
    echo "  $0 -s gemini-only     # Shell into gemini-only environment"
    echo "  $0 -b                 # Build image only"
    echo
    echo "Environment files are located in $ENV_DIR/ and can be edited directly."
}

# Parse command line arguments
if [[ $# -eq 0 ]]; then
    build_image
    run_all_tests
elif [[ "$1" == "-h" || "$1" == "--help" ]]; then
    show_help
elif [[ "$1" == "-i" || "$1" == "--interactive" ]]; then
    build_image
    interactive_mode
elif [[ "$1" == "-b" || "$1" == "--build-only" ]]; then
    build_image
elif [[ "$1" == "-s" || "$1" == "--shell" ]]; then
    if [[ -z "$2" ]]; then
        echo "Error: -s/--shell requires a test case name"
        echo "Use -h for help."
        exit 1
    fi
    if [[ -z "${test_descriptions[$2]}" ]]; then
        echo "Error: Unknown test case: $2"
        echo "Use -h for help."
        exit 1
    fi
    build_image
    shell_into_env "$2"
elif [[ -n "${test_descriptions[$1]}" ]]; then
    build_image
    run_test "$1"
else
    echo "Unknown test case or option: $1"
    echo "Use -h for help."
    exit 1
fi