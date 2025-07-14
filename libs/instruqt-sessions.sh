#!/usr/bin/env bash
# instruqt-sessions.sh - Tmux session management functions
# Version: 1.0.0

# Library metadata
INSTRUQT_SESSIONS_VERSION="1.0.0"
INSTRUQT_SESSIONS_LOADED=true

# Ensure tmux session exists
ensure_tmux_session() {
    local session_name=${1:-"shell"}
    local user=${2:-"replicant"}
    
    echo "Ensuring tmux session '$session_name' exists..."
    
    # Check if session exists
    if tmux has-session -t "$session_name" 2>/dev/null; then
        echo "Tmux session '$session_name' already exists"
        return 0
    fi
    
    # Create new session
    if ! tmux new-session -d -s "$session_name" su - "$user"; then
        echo "ERROR: Failed to create tmux session '$session_name'"
        return 1
    fi
    
    echo "Tmux session '$session_name' created successfully"
    return 0
}

# Cleanup tmux session
cleanup_tmux_session() {
    local session_name=${1:-"shell"}
    
    echo "Cleaning up tmux session '$session_name'..."
    
    # Check if session exists
    if ! tmux has-session -t "$session_name" 2>/dev/null; then
        echo "Tmux session '$session_name' does not exist"
        return 0
    fi
    
    # Clear session history
    tmux clear-history -t "$session_name" 2>/dev/null || true
    
    # Send clear command
    tmux send-keys -t "$session_name" clear ENTER 2>/dev/null || true
    
    echo "Tmux session '$session_name' cleaned up"
    return 0
}

# Reset tmux session to clean state
reset_tmux_session() {
    local session_name=${1:-"shell"}
    local user=${2:-"replicant"}
    
    echo "Resetting tmux session '$session_name'..."
    
    # Kill session if it exists
    if tmux has-session -t "$session_name" 2>/dev/null; then
        tmux kill-session -t "$session_name"
    fi
    
    # Create new session
    ensure_tmux_session "$session_name" "$user"
    
    echo "Tmux session '$session_name' reset successfully"
}

# Send command to tmux session
send_to_tmux_session() {
    local session_name=${1:-"shell"}
    local command=${2:-""}
    local send_enter=${3:-true}
    
    if [[ -z "$command" ]]; then
        echo "ERROR: No command specified for tmux session"
        return 1
    fi
    
    echo "Sending command to tmux session '$session_name': $command"
    
    # Check if session exists
    if ! tmux has-session -t "$session_name" 2>/dev/null; then
        echo "ERROR: Tmux session '$session_name' does not exist"
        return 1
    fi
    
    # Send command
    tmux send-keys -t "$session_name" "$command"
    
    # Send enter if requested
    if [[ "$send_enter" == "true" ]]; then
        tmux send-keys -t "$session_name" ENTER
    fi
    
    return 0
}

# Wait for tmux session to be ready
wait_for_tmux_session() {
    local session_name=${1:-"shell"}
    local timeout=${2:-30}
    local elapsed=0
    
    echo "Waiting for tmux session '$session_name' to be ready..."
    
    while ! tmux has-session -t "$session_name" 2>/dev/null && [[ $elapsed -lt $timeout ]]; do
        echo "Waiting for tmux session '$session_name' (${elapsed}s elapsed)"
        sleep 1
        ((elapsed++))
    done
    
    if [[ $elapsed -ge $timeout ]]; then
        echo "ERROR: Tmux session '$session_name' not ready within ${timeout} seconds"
        return 1
    fi
    
    echo "Tmux session '$session_name' is ready"
    return 0
}

# List all tmux sessions
list_tmux_sessions() {
    echo "Active tmux sessions:"
    if tmux list-sessions 2>/dev/null; then
        return 0
    else
        echo "No active tmux sessions found"
        return 1
    fi
}

# Setup tmux session with specific configuration
setup_tmux_session() {
    local session_name=${1:-"shell"}
    local user=${2:-"replicant"}
    local working_dir=${3:-"/home/$user"}
    
    echo "Setting up tmux session '$session_name' with configuration..."
    
    # Ensure session exists
    ensure_tmux_session "$session_name" "$user"
    
    # Set working directory
    tmux send-keys -t "$session_name" "cd $working_dir" ENTER
    
    # Clear any existing content
    cleanup_tmux_session "$session_name"
    
    echo "Tmux session '$session_name' setup completed"
}

# Attach to tmux session (for interactive use)
attach_tmux_session() {
    local session_name=${1:-"shell"}
    
    echo "Attaching to tmux session '$session_name'..."
    
    if ! tmux has-session -t "$session_name" 2>/dev/null; then
        echo "ERROR: Tmux session '$session_name' does not exist"
        return 1
    fi
    
    tmux attach-session -t "$session_name"
}

# Send environment variable export to tmux session
export_to_tmux_session() {
    local session_name=${1:-"shell"}
    local var_name=$2
    local var_value=$3
    
    if [[ -z "$var_name" ]]; then
        echo "ERROR: Variable name not specified"
        return 1
    fi
    
    if [[ -z "$var_value" ]]; then
        echo "ERROR: Variable value not specified"
        return 1
    fi
    
    echo "Exporting $var_name to tmux session '$session_name'"
    
    # Send export command with proper spacing
    send_to_tmux_session "$session_name" "export ${var_name}='${var_value}'" true
    
    return 0
}

# Send multiple environment variables from agent to tmux session
export_agent_vars_to_tmux() {
    local session_name=${1:-"shell"}
    shift
    local var_names=("$@")
    
    if [[ ${#var_names[@]} -eq 0 ]]; then
        # Default common variables if none specified
        var_names=("REPLICATED_API_TOKEN" "REPLICATED_APP" "USERNAME" "PASSWORD")
    fi
    
    echo "Exporting agent variables to tmux session '$session_name': ${var_names[*]}"
    
    # Ensure session exists first
    ensure_tmux_session "$session_name" "replicant"
    
    # Export each variable from agent
    for var_name in "${var_names[@]}"; do
        local var_value
        var_value=$(agent variable get "$var_name" 2>/dev/null || echo "")
        
        if [[ -n "$var_value" ]]; then
            export_to_tmux_session "$session_name" "$var_name" "$var_value"
        else
            echo "WARNING: Agent variable '$var_name' not found or empty"
        fi
    done
    
    echo "Agent variable export to tmux completed"
    return 0
}

# Comprehensive tmux session setup for Instruqt challenges
setup_instruqt_tmux_session() {
    local session_name=${1:-"shell"}
    local user=${2:-"replicant"}
    local export_vars=${3:-true}
    
    echo "Setting up Instruqt tmux session '$session_name'..."
    
    # Ensure session exists
    ensure_tmux_session "$session_name" "$user"
    
    # Export common environment variables if requested
    if [[ "$export_vars" == "true" ]]; then
        export_agent_vars_to_tmux "$session_name"
    fi
    
    echo "Instruqt tmux session '$session_name' setup completed"
    return 0
}

# Display session management status
sessions_info() {
    echo "Instruqt Sessions Library v${INSTRUQT_SESSIONS_VERSION}"
    echo "Available sessions:"
    list_tmux_sessions
}