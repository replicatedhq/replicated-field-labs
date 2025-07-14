#!/usr/bin/env bash
# instruqt-bootstrap.sh - Core initialization and bootstrap functions
# Version: 1.0.0

# Library metadata
INSTRUQT_BOOTSTRAP_VERSION="1.0.0"
INSTRUQT_BOOTSTRAP_LOADED=true

# Standard initialization for all setup scripts
init_setup_script() {
    local debug_mode=${1:-false}
    
    # Set appropriate error handling based on debug mode
    if [[ "$debug_mode" == "true" ]]; then
        set -euxo pipefail
    else
        set -euo pipefail
    fi
    
    # Set standard home directory
    export HOME_DIR=${HOME_DIR:-/home/replicant}
    
    # Wait for Instruqt bootstrap to complete
    wait_for_instruqt_bootstrap
    
    # Load the existing header library if available
    if [[ -f /etc/profile.d/header.sh ]]; then
        source /etc/profile.d/header.sh
    fi
    
    echo "Setup script initialized successfully"
}

# Wait for Instruqt bootstrap completion with timeout
wait_for_instruqt_bootstrap() {
    local timeout=${1:-300}  # 5 minute default timeout
    local bootstrap_file="/opt/instruqt/bootstrap/host-bootstrap-completed"
    local elapsed=0
    
    echo "Waiting for Instruqt bootstrap to complete..."
    
    while [[ ! -f "$bootstrap_file" ]] && [[ $elapsed -lt $timeout ]]; do
        echo "Waiting for Instruqt to finish booting the VM (${elapsed}s elapsed)"
        sleep 1
        ((elapsed++))
    done
    
    if [[ ! -f "$bootstrap_file" ]]; then
        echo "ERROR: Instruqt bootstrap did not complete within ${timeout} seconds"
        return 1
    fi
    
    echo "Instruqt bootstrap completed successfully"
    return 0
}

# Download and setup the shared header library
setup_header_library() {
    local header_url="https://raw.githubusercontent.com/replicatedhq/replicated-field-labs/main/libs/header.sh"
    local header_path="/etc/profile.d/header.sh"
    
    echo "Downloading shared header library..."
    
    if ! curl -s -o "$header_path" "$header_url"; then
        echo "ERROR: Failed to download header library from $header_url"
        return 1
    fi
    
    # Verify the library was downloaded correctly
    if [[ ! -f "$header_path" ]]; then
        echo "ERROR: Header library not found at $header_path"
        return 1
    fi
    
    # Source the library
    source "$header_path"
    echo "Header library loaded successfully"
    return 0
}

# Validate script prerequisites
validate_prerequisites() {
    local required_tools=("curl" "jq" "yq" "tmux")
    local missing_tools=()
    
    for tool in "${required_tools[@]}"; do
        if ! command -v "$tool" &> /dev/null; then
            missing_tools+=("$tool")
        fi
    done
    
    if [[ ${#missing_tools[@]} -gt 0 ]]; then
        echo "ERROR: Missing required tools: ${missing_tools[*]}"
        return 1
    fi
    
    return 0
}

# Set up error handling and cleanup
setup_error_handling() {
    local cleanup_function=${1:-""}
    
    # Set up trap for cleanup on exit
    if [[ -n "$cleanup_function" ]]; then
        trap "$cleanup_function" EXIT
    fi
    
    # Set up trap for error handling
    trap 'echo "ERROR: Script failed at line $LINENO"' ERR
}

# Display library version and status
bootstrap_info() {
    echo "Instruqt Bootstrap Library v${INSTRUQT_BOOTSTRAP_VERSION}"
    echo "HOME_DIR: ${HOME_DIR}"
    echo "Bootstrap file: $(ls -la /opt/instruqt/bootstrap/host-bootstrap-completed 2>/dev/null || echo 'Not found')"
}