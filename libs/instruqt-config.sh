#!/usr/bin/env bash
# instruqt-config.sh - Environment and configuration management
# Version: 1.0.0

# Library metadata
INSTRUQT_CONFIG_VERSION="1.0.0"
INSTRUQT_CONFIG_LOADED=true

# Setup common environment variables
setup_common_environment() {
    echo "Setting up common environment variables..."
    
    # Get values from header.sh functions (if available)
    local api_token=""
    local app_slug=""
    local username=""
    local password=""
    
    if command -v get_api_token &> /dev/null; then
        api_token=$(get_api_token)
    fi
    
    if command -v get_app_slug &> /dev/null; then
        app_slug=$(get_app_slug)
    fi
    
    if command -v get_username &> /dev/null; then
        username=$(get_username)
    fi
    
    if command -v get_password &> /dev/null; then
        password=$(get_password)
    fi
    
    # Set agent variables
    if [[ -n "$api_token" ]]; then
        agent variable set REPLICATED_API_TOKEN "$api_token"
        export REPLICATED_API_TOKEN="$api_token"
    fi
    
    if [[ -n "$app_slug" ]]; then
        agent variable set REPLICATED_APP "$app_slug"
        export REPLICATED_APP="$app_slug"
    fi
    
    if [[ -n "$username" ]]; then
        agent variable set USERNAME "$username"
        export USERNAME="$username"
    fi
    
    if [[ -n "$password" ]]; then
        agent variable set PASSWORD "$password"
        export PASSWORD="$password"
    fi
    
    echo "Common environment variables setup completed"
}

# Setup Kubernetes configuration
setup_kubernetes_config() {
    local home_dir=${1:-"${HOME_DIR:-/home/replicant}"}
    local cluster_server=${2:-"https://cluster:6443"}
    local kubeconfig_path="$home_dir/.kube/config"
    
    echo "Setting up Kubernetes configuration..."
    
    # Check if kubeconfig exists
    if [[ ! -f "$kubeconfig_path" ]]; then
        echo "ERROR: Kubernetes config not found at $kubeconfig_path"
        return 1
    fi
    
    # Update cluster server
    if ! yq -i ".clusters[0].cluster.server = \"$cluster_server\"" "$kubeconfig_path"; then
        echo "ERROR: Failed to update cluster server in kubeconfig"
        return 1
    fi
    
    # Set proper ownership
    chown -R replicant "$home_dir/.kube"
    
    echo "Kubernetes configuration setup completed"
    return 0
}

# Get application information from API
get_app_info() {
    local api_token=${1:-""}
    
    if [[ -z "$api_token" ]]; then
        if command -v get_api_token &> /dev/null; then
            api_token=$(get_api_token)
        else
            echo "ERROR: API token not provided and get_api_token not available"
            return 1
        fi
    fi
    
    echo "Retrieving application information..."
    
    local app_info
    if ! app_info=$(curl --silent --header "Accept: application/json" \
                         --header "Authorization: $api_token" \
                         https://api.replicated.com/vendor/v3/apps); then
        echo "ERROR: Failed to retrieve application information"
        return 1
    fi
    
    echo "$app_info"
    return 0
}

# Setup application-specific environment variables
setup_app_environment() {
    local api_token=${1:-""}
    
    echo "Setting up application-specific environment..."
    
    if [[ -z "$api_token" ]]; then
        if command -v get_api_token &> /dev/null; then
            api_token=$(get_api_token)
        else
            echo "ERROR: API token not available"
            return 1
        fi
    fi
    
    # Get app information
    local app_info
    if ! app_info=$(get_app_info "$api_token"); then
        echo "ERROR: Failed to get application information"
        return 1
    fi
    
    # Extract app details
    local app_slug=$(echo "$app_info" | jq -r '.apps[0].slug')
    local app_id=$(echo "$app_info" | jq -r '.apps[0].id')
    
    # Set environment variables
    if [[ -n "$app_slug" && "$app_slug" != "null" ]]; then
        agent variable set REPLICATED_APP "$app_slug"
        export REPLICATED_APP="$app_slug"
    fi
    
    if [[ -n "$app_id" && "$app_id" != "null" ]]; then
        agent variable set REPLICATED_APP_ID "$app_id"
        export REPLICATED_APP_ID="$app_id"
    fi
    
    # Setup domain if function is available
    if command -v get_slackernews_domain &> /dev/null; then
        local domain=$(get_slackernews_domain)
        if [[ -n "$domain" ]]; then
            agent variable set SLACKERNEWS_DOMAIN "$domain"
            export SLACKERNEWS_DOMAIN="$domain"
        fi
    fi
    
    echo "Application environment setup completed"
}

# Setup customer-specific environment variables
setup_customer_environment() {
    local customer_name=${1:-""}
    
    echo "Setting up customer-specific environment..."
    
    if [[ -z "$customer_name" ]]; then
        echo "WARNING: No customer name provided"
        return 0
    fi
    
    # Get customer information if functions are available
    local customer_id=""
    local license_id=""
    
    if command -v get_customer_id &> /dev/null; then
        customer_id=$(get_customer_id)
    fi
    
    if command -v get_license_id &> /dev/null; then
        license_id=$(get_license_id "${INSTRUQT_PARTICIPANT_ID}")
    fi
    
    # Set customer variables
    if [[ -n "$customer_id" ]]; then
        agent variable set CUSTOMER_ID "$customer_id"
        export CUSTOMER_ID="$customer_id"
    fi
    
    if [[ -n "$license_id" ]]; then
        agent variable set LICENSE_ID "$license_id"
        export LICENSE_ID="$license_id"
    fi
    
    echo "Customer environment setup completed"
}

# Configure registry authentication
setup_registry_auth() {
    local registry=${1:-"registry.replicated.com"}
    local username=${2:-""}
    local password=${3:-""}
    
    echo "Setting up registry authentication for $registry..."
    
    # Get credentials if not provided
    if [[ -z "$username" ]] && command -v get_username &> /dev/null; then
        username=$(get_username)
    fi
    
    if [[ -z "$password" ]] && command -v get_password &> /dev/null; then
        password=$(get_password)
    fi
    
    if [[ -z "$username" || -z "$password" ]]; then
        echo "ERROR: Registry credentials not available"
        return 1
    fi
    
    # Login to registry
    if ! helm registry login "$registry" --username "$username" --password "$password"; then
        echo "ERROR: Failed to login to registry $registry"
        return 1
    fi
    
    echo "Registry authentication setup completed"
}

# Update configuration file with yq
update_config_value() {
    local config_file=$1
    local yq_expression=$2
    local value=$3
    
    if [[ -z "$config_file" || -z "$yq_expression" || -z "$value" ]]; then
        echo "ERROR: Configuration file, yq expression, and value are required"
        return 1
    fi
    
    echo "Updating configuration: $config_file"
    
    # Backup original file
    if ! cp "$config_file" "${config_file}.backup"; then
        echo "ERROR: Failed to create backup of $config_file"
        return 1
    fi
    
    # Update configuration
    if ! yq -i "$yq_expression = \"$value\"" "$config_file"; then
        echo "ERROR: Failed to update configuration in $config_file"
        # Restore backup
        cp "${config_file}.backup" "$config_file"
        return 1
    fi
    
    echo "Configuration updated successfully"
    return 0
}

# Load configuration from file
load_config_file() {
    local config_file=$1
    local config_type=${2:-"yaml"}
    
    if [[ -z "$config_file" ]]; then
        echo "ERROR: Configuration file not specified"
        return 1
    fi
    
    if [[ ! -f "$config_file" ]]; then
        echo "ERROR: Configuration file not found: $config_file"
        return 1
    fi
    
    echo "Loading configuration from: $config_file"
    
    case "$config_type" in
        "yaml"|"yml")
            if ! yq eval '.' "$config_file" > /dev/null; then
                echo "ERROR: Invalid YAML configuration file"
                return 1
            fi
            ;;
        "json")
            if ! jq '.' "$config_file" > /dev/null; then
                echo "ERROR: Invalid JSON configuration file"
                return 1
            fi
            ;;
        *)
            echo "WARNING: Unknown configuration type: $config_type"
            ;;
    esac
    
    echo "Configuration loaded successfully"
    return 0
}

# === CREDENTIAL AND AUTHENTICATION FUNCTIONS ===
# (Migrated from header.sh)

# Get username based on Instruqt participant ID
get_username() {
    echo "${INSTRUQT_PARTICIPANT_ID}@replicated-labs.com"
}

# Generate password from Instruqt participant ID
get_password() {
    local password=$(echo -n "${INSTRUQT_PARTICIPANT_ID}" | sha256sum)
    echo "${password::20}"
}


# Generate admin console password
get_admin_console_password() {
    local password=$(echo -n "${INSTRUQT_PARTICIPANT_ID}:${INSTRUQT_CHALLENGE_ID}" | sha256sum)
    echo "${password::20}"
}

# Get Slackernews domain for this Instruqt environment with caching
get_slackernews_domain() {
    local domain=$(agent variable get SLACKERNEWS_DOMAIN)
    if [[ -z "$domain" ]]; then
        domain="cluster-30443-${INSTRUQT_PARTICIPANT_ID}.env.play.instruqt.com"
        agent variable set SLACKERNEWS_DOMAIN "$domain"
    fi
    echo "$domain"
}

# Display login credentials with formatting
show_credentials() {
    local CYAN='\033[0;36m'
    local GREEN='\033[1;32m'
    local NC='\033[0m' # No Color
    local password=$(get_password)
    echo -e "${GREEN}Credentials for ${CYAN}https://vendor.replicated.com"
    echo -e "${GREEN}Username: ${CYAN}${INSTRUQT_PARTICIPANT_ID}@replicated-labs.com"
    echo -e "${GREEN}Password: ${CYAN}${password}${NC}"
}

# === CACHED CREDENTIAL FUNCTIONS ===
# These functions use agent variables as cache and set initial values

# Get customer email with caching
get_customer_email() {
    local customer_email=$(agent variable get CUSTOMER_EMAIL)
    if [[ -z "$customer_email" ]]; then
        customer_email="${INSTRUQT_PARTICIPANT_ID}@omozan.io"
        agent variable set CUSTOMER_EMAIL "$customer_email"
    fi
    echo "$customer_email"
}

# Get registry password with caching
get_registry_password() {
    local registry_password=$(agent variable get REGISTRY_PASSWORD)
    if [[ -z "$registry_password" ]]; then
        # Registry password is the license ID
        local customer_id=$(get_customer_id)
        if [[ -n "$customer_id" ]]; then
            registry_password=$(get_license_id "${INSTRUQT_PARTICIPANT_ID}")
            agent variable set REGISTRY_PASSWORD "$registry_password"
        fi
    fi
    echo "$registry_password"
}

# Get API token with caching (improved version)
get_api_token() {
    local access_token=$(agent variable get REPLICATED_API_TOKEN)
    if [[ -z "$access_token" ]]; then
        set +eu
        sleep 5
        local password=$(get_password)
        local login=$(jq -n -c --arg email "${INSTRUQT_PARTICIPANT_ID}@replicated-labs.com" --arg password "${password}" '$ARGS.named')
        set +eu pipefail
        local token=$(curl -s -H "Content-Type: application/json" --request POST -d "$login" https://api.replicated.com/vendor/v1/login | jq -r ".token")
        set -eu pipefail
        
        local i=0
        while [[ ( -z "$token" || "$token" == "null" ) && $i -lt 20 ]]; do
            sleep $((i*5))
            set +eu pipefail
            token=$(curl -s -H "Content-Type: application/json" --request POST -d "$login" https://api.replicated.com/vendor/v1/login | jq -r ".token")
            set -eu pipefail
            i=$((i+1))
        done

        local UUID=$(cat /proc/sys/kernel/random/uuid)
        local apiToken=$(jq -n -c --arg name "instruqt-${UUID}" --argjson read_only false '$ARGS.named')
        access_token=$(curl -s -H "Content-Type: application/json" -H "Authorization: $token" --request POST -d "$apiToken" https://api.replicated.com/vendor/v1/user/token | jq -r ".access_token")

        agent variable set REPLICATED_API_TOKEN "$access_token"
        set +eu
    fi
    echo "$access_token"
}

# Get app ID with caching
get_app_id() {
    local app_id=$(agent variable get APP_ID)
    if [[ -z "$app_id" ]]; then
        local api_token=$(get_api_token)
        app_id=$(curl --header 'Accept: application/json' --header "Authorization: ${api_token}" https://api.replicated.com/vendor/v3/apps | jq -r --arg application "Slackernews" '.apps[] | select( .name | startswith( $application )) | .id')
        agent variable set APP_ID "$app_id"
    fi
    echo "$app_id"
}

# Get customer ID with caching
get_customer_id() {
    local customer_id=$(agent variable get CUSTOMER_ID)
    if [[ -z "$customer_id" ]]; then
        local api_token=$(get_api_token)
        local app_id=$(get_app_id)
        customer_id=$(curl --header 'Accept: application/json' --header "Authorization: ${api_token}" https://api.replicated.com/vendor/v3/app/${app_id}/customers | jq -r --arg name "${INSTRUQT_PARTICIPANT_ID}" '.customers[] | select( .name == $name ) | .id')
        agent variable set CUSTOMER_ID "$customer_id"
    fi
    echo "$customer_id"
}

# Get app slug with caching
get_replicated_app() {
    local app_slug=$(agent variable get REPLICATED_APP)
    if [[ -z "$app_slug" ]]; then
        local api_token=$(get_api_token)
        app_slug=$(curl --header 'Accept: application/json' --header "Authorization: ${api_token}" https://api.replicated.com/vendor/v3/apps | jq -r '.apps[0].slug')
        agent variable set REPLICATED_APP "$app_slug"
    fi
    echo "$app_slug"
}

# Get Replicated SDK version with caching
get_replicated_sdk_version() {
    local sdk_version=$(agent variable get REPLICATED_SDK_VERSION)
    if [[ -z "$sdk_version" ]]; then
        # Get the latest SDK version from the registry
        local token=$(curl --silent "https://registry.replicated.com/v2/token?scope=repository:library/replicated:pull&service=registry.replicated.com" | jq -r .token)
        sdk_version=$(curl --silent -H "Authorization: Bearer $token" "https://registry.replicated.com/v2/library/replicated/tags/list" | jq -r '.tags[]' | sort -r | awk -F '[.-]' '{
            # Extract version components
            major=$1;
            minor=$2;
            patch=$3;
            prerelease=$4;
            prerelease_number=$5;

            # Assign priority to pre-release versions
            if (prerelease == "alpha") {
              prerelease_priority = 1;
            } else if (prerelease == "beta") {
              prerelease_priority = 2;
            } else {
              prerelease_priority = 3;
            }

            # Handle missing pre-release number
            if (prerelease_number == "") {
              prerelease_number = 0;
            }

            # Format output to aid sorting
            printf "%04d%04d%04d%02d%04d-%s\n", major, minor, patch, prerelease_priority, prerelease_number, $0
          }' | sort -r | head -1 | sed 's/^[0-9]*-//')
        agent variable set REPLICATED_SDK_VERSION "$sdk_version"
    fi
    echo "$sdk_version"
}

# Display configuration status
config_info() {
    echo "Instruqt Configuration Library v${INSTRUQT_CONFIG_VERSION}"
    echo "Environment variables:"
    echo "  REPLICATED_API_TOKEN: ${REPLICATED_API_TOKEN:+[SET]}"
    echo "  REPLICATED_APP: ${REPLICATED_APP:-[NOT SET]}"
    echo "  USERNAME: ${USERNAME:-[NOT SET]}"
    echo "  HOME_DIR: ${HOME_DIR:-[NOT SET]}"
    echo "Kubernetes config: $(ls -la ~/.kube/config 2>/dev/null || echo 'Not found')"
}