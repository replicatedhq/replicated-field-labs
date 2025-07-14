#!/usr/bin/env bash
# instruqt-ssh.sh - SSH configuration and connectivity functions
# Version: 1.0.0

# Library metadata
INSTRUQT_SSH_VERSION="1.0.0"
INSTRUQT_SSH_LOADED=true

# Setup SSH client configuration
setup_ssh_config() {
    local additional_config=${1:-""}
    local ssh_config_path="$HOME/.ssh/config"
    
    echo "Setting up SSH client configuration..."
    
    # Ensure SSH directory exists
    mkdir -p "$HOME/.ssh"
    chmod 700 "$HOME/.ssh"
    
    # Create or append to SSH config
    cat <<EOF >> "$ssh_config_path"
Host *
    StrictHostKeyChecking no
    UserKnownHostsFile /dev/null
    PubkeyAcceptedKeyTypes +ssh-rsa
    ${additional_config}
EOF
    
    chmod 600 "$ssh_config_path"
    echo "SSH client configuration completed"
}

# Setup SSH server configuration for Dropbear
setup_ssh_server_config() {
    local dropbear_config_dir="/etc/dropbear"
    local key_file="$dropbear_config_dir/dropbear_rsa_host_key"
    
    echo "Setting up SSH server configuration..."
    
    # Ensure dropbear directory exists
    mkdir -p "$dropbear_config_dir"
    
    # Generate RSA host key if it doesn't exist
    if [[ ! -f "$key_file" ]]; then
        generate_dropbear_key
    fi
    
    echo "SSH server configuration completed"
}

# Generate RSA key for Dropbear SSH server
generate_dropbear_key() {
    local key_file="/etc/dropbear/dropbear_rsa_host_key"
    
    echo "Generating Dropbear RSA host key..."
    
    # Ensure directory exists
    mkdir -p "$(dirname "$key_file")"
    
    # Generate the key
    if ! ssh-keygen -t rsa -f "$key_file" -N ''; then
        echo "ERROR: Failed to generate Dropbear RSA key"
        return 1
    fi
    
    echo "Dropbear RSA key generated successfully"
}

# Wait for SSH connectivity to a host
wait_for_ssh_connectivity() {
    local host=${1:-"shell"}
    local timeout=${2:-60}
    local elapsed=0
    
    echo "Waiting for SSH connectivity to $host..."
    
    while ! ssh -o ConnectTimeout=5 "$host" true 2>/dev/null && [[ $elapsed -lt $timeout ]]; do
        echo "Waiting for SSH connectivity to $host (${elapsed}s elapsed)"
        sleep 1
        ((elapsed++))
    done
    
    if [[ $elapsed -ge $timeout ]]; then
        echo "ERROR: SSH connectivity to $host not available within ${timeout} seconds"
        return 1
    fi
    
    echo "SSH connectivity to $host established"
    return 0
}

# Test SSH connectivity to multiple hosts
test_ssh_connectivity() {
    local hosts=("$@")
    local failed_hosts=()
    
    if [[ ${#hosts[@]} -eq 0 ]]; then
        hosts=("shell")
    fi
    
    echo "Testing SSH connectivity to: ${hosts[*]}"
    
    for host in "${hosts[@]}"; do
        if ! ssh -o ConnectTimeout=5 "$host" true 2>/dev/null; then
            failed_hosts+=("$host")
        fi
    done
    
    if [[ ${#failed_hosts[@]} -gt 0 ]]; then
        echo "ERROR: SSH connectivity failed for: ${failed_hosts[*]}"
        return 1
    fi
    
    echo "SSH connectivity test passed for all hosts"
    return 0
}

# Setup SSH keys for user authentication
setup_ssh_keys() {
    local user=${1:-"replicant"}
    local key_type=${2:-"rsa"}
    local key_file="$HOME/.ssh/id_$key_type"
    
    echo "Setting up SSH keys for user $user..."
    
    # Ensure SSH directory exists
    mkdir -p "$HOME/.ssh"
    chmod 700 "$HOME/.ssh"
    
    # Generate key if it doesn't exist
    if [[ ! -f "$key_file" ]]; then
        ssh-keygen -t "$key_type" -f "$key_file" -N '' -q
        echo "SSH key generated: $key_file"
    else
        echo "SSH key already exists: $key_file"
    fi
    
    # Set proper permissions
    chmod 600 "$key_file"
    chmod 644 "$key_file.pub"
    
    echo "SSH key setup completed"
}

# Copy SSH keys to remote host
copy_ssh_keys() {
    local host=${1:-"shell"}
    local user=${2:-"replicant"}
    local key_file="$HOME/.ssh/id_rsa.pub"
    
    echo "Copying SSH keys to $host..."
    
    if [[ ! -f "$key_file" ]]; then
        echo "ERROR: SSH public key not found at $key_file"
        return 1
    fi
    
    if ! ssh-copy-id -i "$key_file" "$user@$host" 2>/dev/null; then
        echo "ERROR: Failed to copy SSH keys to $host"
        return 1
    fi
    
    echo "SSH keys copied to $host successfully"
}

# Display SSH configuration status
ssh_info() {
    echo "Instruqt SSH Library v${INSTRUQT_SSH_VERSION}"
    echo "SSH config: $(ls -la ~/.ssh/config 2>/dev/null || echo 'Not found')"
    echo "SSH keys: $(ls -la ~/.ssh/id_* 2>/dev/null || echo 'Not found')"
    echo "Dropbear key: $(ls -la /etc/dropbear/dropbear_rsa_host_key 2>/dev/null || echo 'Not found')"
}