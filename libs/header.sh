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

# Download and cache all required library files
download_libraries() {
    # Try multiple sources for library files
    local base_urls=(
        "https://raw.githubusercontent.com/replicatedhq/replicated-field-labs/main/libs"
        "https://raw.githubusercontent.com/replicatedhq/replicated-field-labs/refactor/crdant/builds-comprehensive-library/libs"
    )
    
    local required_libraries=(
        "instruqt-bootstrap.sh"
        "instruqt-files.sh"
        "instruqt-ssh.sh"
        "instruqt-sessions.sh"
        "instruqt-config.sh"
        "instruqt-services.sh"
        "instruqt-apps.sh"
        "instruqt-checks.sh"
        "instruqt-all.sh"
    )
    
    local downloaded_any=false
    local missing_libraries=()
    
    for lib in "${required_libraries[@]}"; do
        local lib_path="$LIBS_DIR/$lib"
        
        # Check if library exists and is not empty
        if [[ ! -f "$lib_path" ]] || [[ ! -s "$lib_path" ]]; then
            echo "Downloading $lib..."
            
            local download_success=false
            
            # Try each base URL
            for base_url in "${base_urls[@]}"; do
                local attempts=0
                local max_attempts=2
                
                while [[ $attempts -lt $max_attempts ]]; do
                    if curl -fsSL "$base_url/$lib" -o "$lib_path.tmp"; then
                        # Verify download was successful and not empty
                        if [[ -s "$lib_path.tmp" ]]; then
                            mv "$lib_path.tmp" "$lib_path"
                            chmod +x "$lib_path"
                            downloaded_any=true
                            download_success=true
                            echo "  ✓ Downloaded $lib from $(basename "$base_url")"
                            break 2  # Break out of both loops
                        else
                            rm -f "$lib_path.tmp"
                        fi
                    fi
                    ((attempts++))
                done
            done
            
            if [[ "$download_success" == "false" ]]; then
                missing_libraries+=("$lib")
                echo "  ✗ Failed to download $lib from any source"
            fi
        fi
    done
    
    if [[ ${#missing_libraries[@]} -gt 0 ]]; then
        echo ""
        echo "WARNING: Could not download the following libraries:"
        for lib in "${missing_libraries[@]}"; do
            echo "  - $lib"
        done
        echo ""
        echo "This may happen if:"
        echo "  1. The libraries haven't been pushed to the main branch yet"
        echo "  2. Network connectivity issues"
        echo "  3. The repository is private or inaccessible"
        echo ""
        echo "To resolve this:"
        echo "  1. Ensure all library files are in the same directory as header.sh, or"
        echo "  2. Wait for the changes to be merged to the main branch, or"
        echo "  3. Use the complete library directory from the repository"
        echo ""
        return 1
    fi
    
    if [[ "$downloaded_any" == "true" ]]; then
        echo "Library downloads completed successfully"
    fi
    
    return 0
}

# Ensure all libraries are available
if ! download_libraries; then
    echo "WARNING: Running in fallback mode with essential functions only"
    echo "For full functionality, ensure all library files are available locally"
    
    # Provide essential fallback functions from the original header.sh
    get_username() {
        echo "${INSTRUQT_PARTICIPANT_ID}@replicated-labs.com"
    }
    
    get_password() {
        local password=$(echo -n "${INSTRUQT_PARTICIPANT_ID}" | sha256sum)
        echo "${password::20}"
    }
    
    get_slackernews_domain() {
        echo "cluster-30443-${INSTRUQT_PARTICIPANT_ID}.env.play.instruqt.com"
    }
    
    show_credentials() {
        local CYAN='\033[0;36m'
        local GREEN='\033[1;32m'
        local NC='\033[0m'
        local password=$(get_password)
        echo -e "${GREEN}Credentials for ${CYAN}https://vendor.replicated.com"
        echo -e "${GREEN}Username: ${CYAN}${INSTRUQT_PARTICIPANT_ID}@replicated-labs.com"
        echo -e "${GREEN}Password: ${CYAN}${password}${NC}"
    }
    
    echo "Fallback mode active - basic functions available"
    HEADER_FALLBACK_MODE=true
    
else
    # Load all Instruqt libraries in dependency order
    source "$LIBS_DIR/instruqt-bootstrap.sh"    # Core initialization
    source "$LIBS_DIR/instruqt-files.sh"        # File and directory management
    source "$LIBS_DIR/instruqt-ssh.sh"          # SSH configuration
    source "$LIBS_DIR/instruqt-sessions.sh"     # Tmux session management
    source "$LIBS_DIR/instruqt-config.sh"       # Environment and credentials
    source "$LIBS_DIR/instruqt-services.sh"     # Service and API management
    source "$LIBS_DIR/instruqt-apps.sh"         # Application and release management
    source "$LIBS_DIR/instruqt-checks.sh"       # Check script validation functions
fi

# Verify all libraries loaded successfully (skip in fallback mode)
if [[ "$HEADER_FALLBACK_MODE" != "true" ]]; then
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
        [[ "$INSTRUQT_CHECKS_LOADED" != "true" ]] && missing_libraries+=("instruqt-checks")
        
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
else
    echo "Header.sh loaded in fallback mode"
fi

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
    echo "Checks: v${INSTRUQT_CHECKS_VERSION}"
    echo "================================================="
}