#!/usr/bin/env bash
# instruqt-all.sh - Master library that loads all Instruqt functional libraries
# Version: 1.0.0

# Library metadata
INSTRUQT_ALL_VERSION="1.0.0"
INSTRUQT_ALL_LOADED=true

# Get the directory where this script is located
INSTRUQT_LIBS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Load all functional libraries
echo "Loading Instruqt functional libraries..."

# Core libraries (in dependency order)
source "$INSTRUQT_LIBS_DIR/instruqt-bootstrap.sh"
source "$INSTRUQT_LIBS_DIR/instruqt-files.sh"
source "$INSTRUQT_LIBS_DIR/instruqt-ssh.sh"
source "$INSTRUQT_LIBS_DIR/instruqt-sessions.sh"
source "$INSTRUQT_LIBS_DIR/instruqt-config.sh"
source "$INSTRUQT_LIBS_DIR/instruqt-services.sh"
source "$INSTRUQT_LIBS_DIR/instruqt-apps.sh"

# Verify all libraries loaded successfully
check_library_loading() {
    local missing_libraries=()
    
    # Check each library's loaded flag
    [[ "$INSTRUQT_BOOTSTRAP_LOADED" != "true" ]] && missing_libraries+=("instruqt-bootstrap")
    [[ "$INSTRUQT_FILES_LOADED" != "true" ]] && missing_libraries+=("instruqt-files")
    [[ "$INSTRUQT_SSH_LOADED" != "true" ]] && missing_libraries+=("instruqt-ssh")
    [[ "$INSTRUQT_SESSIONS_LOADED" != "true" ]] && missing_libraries+=("instruqt-sessions")
    [[ "$INSTRUQT_CONFIG_LOADED" != "true" ]] && missing_libraries+=("instruqt-config")
    [[ "$INSTRUQT_SERVICES_LOADED" != "true" ]] && missing_libraries+=("instruqt-services")
    [[ "$INSTRUQT_APPS_LOADED" != "true" ]] && missing_libraries+=("instruqt-apps")
    
    if [[ ${#missing_libraries[@]} -gt 0 ]]; then
        echo "ERROR: Failed to load libraries: ${missing_libraries[*]}"
        return 1
    fi
    
    echo "All Instruqt libraries loaded successfully"
    return 0
}

# Check that all libraries loaded
check_library_loading

# Display library information
display_library_info() {
    echo "===== Instruqt Libraries v${INSTRUQT_ALL_VERSION} ====="
    echo "Bootstrap: v${INSTRUQT_BOOTSTRAP_VERSION}"
    echo "Files: v${INSTRUQT_FILES_VERSION}"
    echo "SSH: v${INSTRUQT_SSH_VERSION}"
    echo "Sessions: v${INSTRUQT_SESSIONS_VERSION}"
    echo "Config: v${INSTRUQT_CONFIG_VERSION}"
    echo "Services: v${INSTRUQT_SERVICES_VERSION}"
    echo "Apps: v${INSTRUQT_APPS_VERSION}"
    echo "================================================="
}

# Common setup function that uses multiple libraries
setup_instruqt_environment() {
    local environment_type=${1:-"shell"}
    local debug_mode=${2:-false}
    
    echo "Setting up Instruqt environment: $environment_type"
    
    # Initialize script with bootstrap
    init_setup_script "$debug_mode"
    
    # Setup SSH configuration
    setup_ssh_config
    
    # Generate SSH keys if needed
    generate_dropbear_key
    
    # Ensure tmux session exists
    ensure_tmux_session
    
    # Setup common environment variables
    setup_common_environment
    
    # Setup release directory
    setup_release_directory
    
    # Environment-specific setup
    case "$environment_type" in
        "cluster")
            setup_kubernetes_config
            wait_for_kubernetes_api
            ;;
        "node")
            install_kots_cli
            ;;
        "shell")
            # Shell environment is the default, no additional setup needed
            ;;
        *)
            echo "WARNING: Unknown environment type: $environment_type"
            ;;
    esac
    
    echo "Instruqt environment setup completed"
}

# Common cleanup function
cleanup_instruqt_environment() {
    local environment_type=${1:-"shell"}
    
    echo "Cleaning up Instruqt environment: $environment_type"
    
    # Cleanup tmux session
    cleanup_tmux_session
    
    # Cleanup temporary files
    cleanup_temp_files
    
    # Environment-specific cleanup
    case "$environment_type" in
        "cluster")
            # Add any cluster-specific cleanup here
            ;;
        "node")
            # Add any node-specific cleanup here
            ;;
        "shell")
            # Add any shell-specific cleanup here
            ;;
    esac
    
    echo "Instruqt environment cleanup completed"
}

# Help function showing available functions
show_help() {
    echo "Instruqt Libraries v${INSTRUQT_ALL_VERSION} - Available Functions"
    echo ""
    echo "=== Bootstrap Functions ==="
    echo "  init_setup_script [debug_mode]       - Initialize script with error handling"
    echo "  wait_for_instruqt_bootstrap [timeout] - Wait for Instruqt bootstrap"
    echo "  setup_header_library                  - Download and setup header.sh"
    echo ""
    echo "=== SSH Functions ==="
    echo "  setup_ssh_config [additional_config] - Setup SSH client configuration"
    echo "  generate_dropbear_key                 - Generate RSA key for Dropbear"
    echo "  wait_for_ssh_connectivity [host]     - Wait for SSH connectivity"
    echo ""
    echo "=== Session Functions ==="
    echo "  ensure_tmux_session [name] [user]    - Ensure tmux session exists"
    echo "  cleanup_tmux_session [name]          - Cleanup tmux session"
    echo "  send_to_tmux_session [name] [cmd]    - Send command to tmux session"
    echo ""
    echo "=== File Functions ==="
    echo "  setup_release_directory [home_dir]   - Setup release directory"
    echo "  create_yaml_file [path] [content]    - Create YAML file with content"
    echo "  setup_file_permissions [path]        - Setup file permissions"
    echo ""
    echo "=== Config Functions ==="
    echo "  setup_common_environment              - Setup common environment variables"
    echo "  setup_kubernetes_config [home] [srv] - Setup Kubernetes configuration"
    echo "  setup_registry_auth [reg] [user] [pw] - Setup registry authentication"
    echo ""
    echo "=== Service Functions ==="
    echo "  wait_for_kubernetes_api [endpoint]   - Wait for Kubernetes API"
    echo "  wait_for_service [name] [command]    - Wait for generic service"
    echo "  wait_for_port [host] [port]          - Wait for port to be available"
    echo ""
    echo "=== App Functions ==="
    echo "  package_helm_chart [path] [dest]     - Package Helm chart"
    echo "  update_helm_dependencies [path]      - Update Helm dependencies"
    echo "  create_replicated_release [path]     - Create Replicated release"
    echo "  setup_slackernews [home_dir]         - Setup Slackernews application"
    echo ""
    echo "=== High-Level Functions ==="
    echo "  setup_instruqt_environment [type]    - Complete environment setup"
    echo "  cleanup_instruqt_environment [type]  - Complete environment cleanup"
    echo "  display_library_info                 - Show library version information"
    echo ""
    echo "Environment types: shell, cluster, node"
}

echo "Instruqt Libraries v${INSTRUQT_ALL_VERSION} loaded successfully"
echo "Type 'show_help' for available functions"