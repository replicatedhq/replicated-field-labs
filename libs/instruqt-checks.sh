#!/usr/bin/env bash
# instruqt-checks.sh - Check script validation functions
# Version: 1.0.0

# Library metadata
INSTRUQT_CHECKS_VERSION="1.0.0"
INSTRUQT_CHECKS_LOADED=true

# Global check result counter
CHECK_RESULT=0

# Initialize check script with proper error handling
init_check_script() {
    set -euo pipefail
    CHECK_RESULT=0
    export HOME_DIR=${HOME_DIR:-/home/replicant}
    
    echo "Initializing check script..."
}

# Add a check failure with message
add_check_failure() {
    local message="$1"
    
    if [[ -n "$message" ]]; then
        fail-message "$message"
    fi
    
    CHECK_RESULT=$((CHECK_RESULT + 1))
}

# Generic condition checker with result increment
check_condition() {
    local condition="$1"
    local error_message="$2"
    
    if ! eval "$condition"; then
        add_check_failure "$error_message"
        return 1
    fi
    
    return 0
}

# Finish check script and exit with accumulated result
finish_check_script() {
    echo "Check completed with $CHECK_RESULT failure(s)"
    exit $CHECK_RESULT
}

# ===== API AND CHANNEL FUNCTIONS =====

# Get current version for a specific channel
get_channel_version() {
    local channel_name="$1"
    local api_token
    
    if [[ -z "$channel_name" ]]; then
        echo "ERROR: Channel name not specified"
        return 1
    fi
    
    api_token=$(get_api_token)
    if [[ -z "$api_token" ]]; then
        echo "ERROR: Could not get API token"
        return 1
    fi
    
    local version
    version=$(curl --silent --header "Accept: application/json" \
        --header "Authorization: ${api_token}" \
        https://api.replicated.com/vendor/v3/apps | \
        jq -r --arg channel "$channel_name" \
        '.apps[0].channels[] | select(.name == $channel) | .currentVersion')
    
    if [[ "$version" == "null" || -z "$version" ]]; then
        return 1
    fi
    
    echo "$version"
}

# Get channel info (name, version, exists)
get_channel_info() {
    local channel_name="$1"
    local api_token
    
    if [[ -z "$channel_name" ]]; then
        echo "ERROR: Channel name not specified"
        return 1
    fi
    
    api_token=$(get_api_token)
    if [[ -z "$api_token" ]]; then
        echo "ERROR: Could not get API token"
        return 1
    fi
    
    curl --silent --header "Accept: application/json" \
        --header "Authorization: ${api_token}" \
        https://api.replicated.com/vendor/v3/apps | \
        jq -r --arg channel "$channel_name" \
        '.apps[0].channels[] | select(.name == $channel)'
}

# Check if channel exists
check_channel_exists() {
    local channel_name="$1"
    local error_message="$2"
    
    local channel_info
    channel_info=$(get_channel_info "$channel_name")
    
    if [[ -z "$channel_info" || "$channel_info" == "null" ]]; then
        add_check_failure "$error_message"
        return 1
    fi
    
    return 0
}

# Verify channel has expected version
check_channel_version() {
    local channel_name="$1"
    local expected_version="$2"
    local error_message="$3"
    
    local current_version
    current_version=$(get_channel_version "$channel_name")
    
    if [[ "$current_version" != "$expected_version" ]]; then
        add_check_failure "$error_message"
        return 1
    fi
    
    return 0
}

# Check multiple channels have the same version
check_channel_promotions() {
    local expected_version="$1"
    local channels="$2"  # Comma-separated list
    
    IFS=',' read -ra CHANNEL_ARRAY <<< "$channels"
    
    for channel in "${CHANNEL_ARRAY[@]}"; do
        channel=$(echo "$channel" | xargs)  # Trim whitespace
        
        local error_msg="Please be sure to promote the application to the \`${channel}\` channel with version \`${expected_version}\`"
        check_channel_version "$channel" "$expected_version" "$error_msg"
    done
}

# ===== TMUX SESSION VALIDATION FUNCTIONS =====

# Capture tmux session content
capture_tmux_session() {
    local session_name="${1:-shell}"
    
    # Check if session exists
    if ! tmux has-session -t "$session_name" 2>/dev/null; then
        echo "ERROR: Tmux session '$session_name' does not exist"
        return 1
    fi
    
    # Capture session content
    tmux capture-pane -t "$session_name" -S -
    tmux save-buffer -
}

# Check if variable was exported in tmux session
check_tmux_var_export() {
    local session_name="${1:-shell}"
    local var_name="$2"
    local error_message="$3"
    
    if [[ -z "$var_name" ]]; then
        echo "ERROR: Variable name not specified"
        return 1
    fi
    
    local session_content
    session_content=$(capture_tmux_session "$session_name")
    
    if [[ $? -ne 0 ]]; then
        add_check_failure "Could not capture tmux session '$session_name'"
        return 1
    fi
    
    # Check for variable export pattern
    if ! grep -qE "${var_name}=[A-Za-z0-9 \"']+" <(echo "$session_content"); then
        add_check_failure "$error_message"
        return 1
    fi
    
    return 0
}

# Check multiple variables were exported in tmux session
check_multiple_tmux_vars() {
    local session_name="${1:-shell}"
    shift
    local var_names=("$@")
    
    for var_name in "${var_names[@]}"; do
        local error_msg="Please set the variable \`\$${var_name}\` so you will be able to use it in future steps"
        check_tmux_var_export "$session_name" "$var_name" "$error_msg"
    done
}

# Check if tmux session contains a pattern
tmux_session_contains() {
    local session_name="${1:-shell}"
    local pattern="$2"
    local error_message="$3"
    
    if [[ -z "$pattern" ]]; then
        echo "ERROR: Pattern not specified"
        return 1
    fi
    
    local session_content
    session_content=$(capture_tmux_session "$session_name")
    
    if [[ $? -ne 0 ]]; then
        add_check_failure "Could not capture tmux session '$session_name'"
        return 1
    fi
    
    if ! grep -qE "$pattern" <(echo "$session_content"); then
        add_check_failure "$error_message"
        return 1
    fi
    
    return 0
}

# ===== FILE SYSTEM AND CHART VALIDATION FUNCTIONS =====

# Verify Helm chart version
check_chart_version() {
    local chart_path="$1"
    local expected_version="$2"
    local error_message="$3"
    
    if [[ ! -f "$chart_path" ]]; then
        add_check_failure "Chart file not found: $chart_path"
        return 1
    fi
    
    local current_version
    current_version=$(yq .version "$chart_path" 2>/dev/null)
    
    if [[ "$current_version" != "$expected_version" ]]; then
        add_check_failure "$error_message"
        return 1
    fi
    
    return 0
}

# Verify chart dependency exists
check_chart_dependency() {
    local chart_path="$1"
    local dependency_name="$2"
    local error_message="$3"
    
    if [[ ! -f "$chart_path" ]]; then
        add_check_failure "Chart file not found: $chart_path"
        return 1
    fi
    
    local dependency_check
    dependency_check=$(yq '.dependencies[] | select(.name == "'"$dependency_name"'") | .name' "$chart_path" 2>/dev/null)
    
    if [[ "$dependency_check" != "$dependency_name" ]]; then
        add_check_failure "$error_message"
        return 1
    fi
    
    return 0
}

# Verify dependency file exists
check_dependency_file() {
    local charts_dir="$1"
    local dependency_name="$2"
    local version="$3"
    local error_message="$4"
    
    local dependency_file="${charts_dir}/${dependency_name}-${version}.tgz"
    
    if [[ ! -f "$dependency_file" ]]; then
        add_check_failure "$error_message"
        return 1
    fi
    
    return 0
}

# Verify packaged chart exists
check_packaged_chart() {
    local release_dir="$1"
    local chart_name="$2"
    local version="$3"
    local error_message="$4"
    
    local chart_file="${release_dir}/${chart_name}-${version}.tgz"
    
    if [[ ! -f "$chart_file" ]]; then
        add_check_failure "$error_message"
        return 1
    fi
    
    return 0
}

# ===== KUBERNETES RESOURCE VALIDATION FUNCTIONS =====

# Verify Helm release
check_helm_release() {
    local namespace="$1"
    local expected_chart="$2"
    local error_message="$3"
    local kubeconfig="${4:-/home/replicant/.kube/config}"
    
    local installed_chart
    installed_chart=$(helm list -n "$namespace" -o yaml --kubeconfig "$kubeconfig" | yq '.[0].chart' 2>/dev/null)
    
    if [[ "$installed_chart" != "$expected_chart" ]]; then
        add_check_failure "$error_message"
        return 1
    fi
    
    return 0
}

# Get Helm release information
get_helm_release_info() {
    local namespace="$1"
    local release_name="${2:-}"
    local kubeconfig="${3:-/home/replicant/.kube/config}"
    
    if [[ -n "$release_name" ]]; then
        helm get all "$release_name" -n "$namespace" --kubeconfig "$kubeconfig" 2>/dev/null
    else
        helm list -n "$namespace" -o yaml --kubeconfig "$kubeconfig" 2>/dev/null
    fi
}

# ===== COMPOSITE CHECK FUNCTIONS =====

# Complete chart update validation workflow
check_chart_update_workflow() {
    local chart_path="$1"
    local version="$2"
    local dependency_name="$3"
    local dependency_version="$4"
    local release_dir="$5"
    local chart_name="$6"
    
    # Check chart version
    check_chart_version "$chart_path" "$version" \
        "Please be sure to update the version of the ${chart_name} Helm chart to reflect your changes"
    
    # Check dependency exists
    check_chart_dependency "$chart_path" "$dependency_name" \
        "Please be sure to include the ${dependency_name} dependency in the ${chart_name} Helm chart"
    
    # Check dependency file exists
    local charts_dir="$(dirname "$chart_path")/charts"
    check_dependency_file "$charts_dir" "$dependency_name" "$dependency_version" \
        "Please be sure to update the ${chart_name} Helm chart's dependencies to include the ${dependency_name}"
    
    # Check packaged chart exists
    check_packaged_chart "$release_dir" "$chart_name" "$version" \
        "Please be sure to update and repackage the ${chart_name} Helm chart"
}

# Validate Instruqt environment setup
check_instruqt_environment_setup() {
    local required_vars=("$@")
    
    if [[ ${#required_vars[@]} -eq 0 ]]; then
        required_vars=("REPLICATED_API_TOKEN" "REPLICATED_APP")
    fi
    
    check_multiple_tmux_vars "shell" "${required_vars[@]}"
}

# Display library version and available functions
checks_info() {
    echo "Instruqt Checks Library v${INSTRUQT_CHECKS_VERSION}"
    echo ""
    echo "=== Available Functions ==="
    echo "Check Management:"
    echo "  init_check_script()                           - Initialize check with result=0"
    echo "  add_check_failure(message)                    - Add failure and increment result"
    echo "  check_condition(condition, error_msg)         - Generic condition checker"
    echo "  finish_check_script()                         - Exit with accumulated result"
    echo ""
    echo "Channel Functions:"
    echo "  get_channel_version(channel)                  - Get current channel version"
    echo "  check_channel_exists(channel, error_msg)      - Verify channel exists"
    echo "  check_channel_version(channel, ver, error)    - Verify channel version"
    echo "  check_channel_promotions(version, channels)   - Check multiple channel versions"
    echo ""
    echo "Tmux Functions:"
    echo "  capture_tmux_session(session)                 - Capture session content"
    echo "  check_tmux_var_export(session, var, error)    - Check variable export"
    echo "  check_multiple_tmux_vars(session, vars...)    - Check multiple variables"
    echo "  tmux_session_contains(session, pattern, err)  - Pattern matching"
    echo ""
    echo "File/Chart Functions:"
    echo "  check_chart_version(path, version, error)     - Verify chart version"
    echo "  check_chart_dependency(path, dep, error)      - Verify dependency"
    echo "  check_dependency_file(dir, dep, ver, error)   - Verify dependency file"
    echo "  check_packaged_chart(dir, chart, ver, error)  - Verify packaged chart"
    echo ""
    echo "Kubernetes Functions:"
    echo "  check_helm_release(ns, chart, error)          - Verify Helm release"
    echo "  get_helm_release_info(ns, release)            - Get release info"
    echo ""
    echo "Composite Functions:"
    echo "  check_chart_update_workflow(...)              - Complete chart validation"
    echo "  check_instruqt_environment_setup(vars...)     - Environment validation"
}