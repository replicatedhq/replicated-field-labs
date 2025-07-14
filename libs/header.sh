#!/usr/bin/env bash
# header.sh - Refactored shared library loader for Replicated Field Labs
# Version: 2.0.0 
# 
# This file has been refactored to use the new modular Instruqt libraries.
# All functions have been moved to appropriate functional libraries.

# Library metadata
HEADER_VERSION="2.0.0"
HEADER_LOADED=true

# Get the directory where this script is located
LIBS_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

echo "Loading Replicated Field Labs libraries v${HEADER_VERSION}..."

# Load all Instruqt libraries in dependency order
source "$LIBS_DIR/instruqt-bootstrap.sh"    # Core initialization
source "$LIBS_DIR/instruqt-files.sh"        # File and directory management
source "$LIBS_DIR/instruqt-ssh.sh"          # SSH configuration
source "$LIBS_DIR/instruqt-sessions.sh"     # Tmux session management
source "$LIBS_DIR/instruqt-config.sh"       # Environment and credentials
source "$LIBS_DIR/instruqt-services.sh"     # Service and API management
source "$LIBS_DIR/instruqt-apps.sh"         # Application and release management

# Verify all libraries loaded successfully
check_libraries() {
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
    
    return 0
}

# Check that all libraries loaded
if ! check_libraries; then
    echo "FATAL: Some libraries failed to load. Cannot continue."
    exit 1
fi

echo "All Replicated Field Labs libraries loaded successfully"

# Also load the high-level utility functions from instruqt-all.sh
# This ensures header.sh provides everything the old version did plus new utilities
if [[ -f "$LIBS_DIR/instruqt-all.sh" ]]; then
    # Extract just the high-level functions from instruqt-all.sh
    # (skip the library loading since we already did that)
    
    # setup_instruqt_environment function
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
    
    # cleanup_instruqt_environment function
    cleanup_instruqt_environment() {
        local environment_type=${1:-"shell"}
        
        echo "Cleaning up Instruqt environment: $environment_type"
        
        # Cleanup tmux session
        cleanup_tmux_session
        
        # Cleanup temporary files
        cleanup_temp_files
        
        echo "Instruqt environment cleanup completed"
    }
fi

# ===== FUNCTION MIGRATION REFERENCE =====
# The following functions have been moved to new libraries:
#
# Moved to instruqt-config.sh:
#   - get_username()
#   - get_password() 
#   - get_api_token()
#   - get_admin_console_password()
#   - get_slackernews_domain()
#   - show_credentials()
#
# Moved to instruqt-apps.sh:
#   - get_replicated_sdk_version()
#   - get_embedded_cluster_version()
#   - get_app_slug()
#   - get_app_id()
#   - get_customer_id()
#   - get_license_id()
#   - get_slackernews()
#
# All functions maintain the same API and behavior.
# Scripts using header.sh will continue to work without modification.
# ==========================================

# Display version information for all loaded libraries
header_info() {
    echo "===== Replicated Field Labs Libraries v${HEADER_VERSION} ====="
    echo "Bootstrap: v${INSTRUQT_BOOTSTRAP_VERSION}"
    echo "Files: v${INSTRUQT_FILES_VERSION}"
    echo "SSH: v${INSTRUQT_SSH_VERSION}"
    echo "Sessions: v${INSTRUQT_SESSIONS_VERSION}"
    echo "Config: v${INSTRUQT_CONFIG_VERSION}"
    echo "Services: v${INSTRUQT_SERVICES_VERSION}"
    echo "Apps: v${INSTRUQT_APPS_VERSION}"
    echo "================================================="
}