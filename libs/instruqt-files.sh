#!/usr/bin/env bash
# instruqt-files.sh - Directory and file management functions
# Version: 1.0.0

# Library metadata
INSTRUQT_FILES_VERSION="1.0.0"
INSTRUQT_FILES_LOADED=true

# Create directory with proper permissions
create_directory() {
    local dir_path=$1
    local owner=${2:-"replicant:replicant"}
    local permissions=${3:-"755"}
    
    if [[ -z "$dir_path" ]]; then
        echo "ERROR: Directory path not specified"
        return 1
    fi
    
    echo "Creating directory: $dir_path"
    
    # Create directory
    if ! mkdir -p "$dir_path"; then
        echo "ERROR: Failed to create directory $dir_path"
        return 1
    fi
    
    # Set ownership
    if ! chown "$owner" "$dir_path"; then
        echo "ERROR: Failed to set ownership of $dir_path to $owner"
        return 1
    fi
    
    # Set permissions
    if ! chmod "$permissions" "$dir_path"; then
        echo "ERROR: Failed to set permissions of $dir_path to $permissions"
        return 1
    fi
    
    echo "Directory created successfully: $dir_path"
    return 0
}

# Setup standard release directory
setup_release_directory() {
    local home_dir=${1:-"${HOME_DIR:-/home/replicant}"}
    local release_dir="$home_dir/release"
    
    echo "Setting up release directory..."
    
    create_directory "$release_dir" "replicant:replicant" "755"
    
    echo "Release directory setup completed: $release_dir"
}

# Create YAML file with content and proper ownership
create_yaml_file() {
    local filepath=$1
    local content=$2
    local owner=${3:-"replicant"}
    
    if [[ -z "$filepath" || -z "$content" ]]; then
        echo "ERROR: Filepath and content are required"
        return 1
    fi
    
    echo "Creating YAML file: $filepath"
    
    # Create parent directory if it doesn't exist
    local parent_dir=$(dirname "$filepath")
    if [[ ! -d "$parent_dir" ]]; then
        mkdir -p "$parent_dir"
    fi
    
    # Create file with content
    cat <<EOF > "$filepath"
$content
EOF
    
    # Set ownership
    if ! chown "$owner" "$filepath"; then
        echo "ERROR: Failed to set ownership of $filepath to $owner"
        return 1
    fi
    
    echo "YAML file created successfully: $filepath"
    return 0
}

# Create configuration file from template
create_config_file() {
    local filepath=$1
    local template_content=$2
    local owner=${3:-"replicant"}
    local permissions=${4:-"644"}
    
    if [[ -z "$filepath" || -z "$template_content" ]]; then
        echo "ERROR: Filepath and template content are required"
        return 1
    fi
    
    echo "Creating configuration file: $filepath"
    
    # Create parent directory if needed
    local parent_dir=$(dirname "$filepath")
    if [[ ! -d "$parent_dir" ]]; then
        mkdir -p "$parent_dir"
    fi
    
    # Create file
    echo "$template_content" > "$filepath"
    
    # Set ownership and permissions
    chown "$owner" "$filepath"
    chmod "$permissions" "$filepath"
    
    echo "Configuration file created: $filepath"
    return 0
}

# Setup file permissions recursively
setup_file_permissions() {
    local target_path=$1
    local owner=${2:-"replicant:replicant"}
    local dir_permissions=${3:-"755"}
    local file_permissions=${4:-"644"}
    
    if [[ -z "$target_path" ]]; then
        echo "ERROR: Target path not specified"
        return 1
    fi
    
    echo "Setting up file permissions for: $target_path"
    
    # Set ownership recursively
    if ! chown -R "$owner" "$target_path"; then
        echo "ERROR: Failed to set ownership of $target_path"
        return 1
    fi
    
    # Set directory permissions
    find "$target_path" -type d -exec chmod "$dir_permissions" {} \; 2>/dev/null
    
    # Set file permissions
    find "$target_path" -type f -exec chmod "$file_permissions" {} \; 2>/dev/null
    
    echo "File permissions setup completed for: $target_path"
}

# Copy file or directory with proper permissions
copy_with_permissions() {
    local source=$1
    local destination=$2
    local owner=${3:-"replicant:replicant"}
    local preserve_permissions=${4:-false}
    
    if [[ -z "$source" || -z "$destination" ]]; then
        echo "ERROR: Source and destination are required"
        return 1
    fi
    
    echo "Copying $source to $destination"
    
    # Create destination parent directory
    local dest_parent=$(dirname "$destination")
    if [[ ! -d "$dest_parent" ]]; then
        mkdir -p "$dest_parent"
    fi
    
    # Copy with or without preserving permissions
    if [[ "$preserve_permissions" == "true" ]]; then
        cp -R "$source" "$destination"
    else
        cp -R "$source" "$destination"
        chown -R "$owner" "$destination"
    fi
    
    echo "Copy completed: $source -> $destination"
}

# Create temporary file with cleanup
create_temp_file() {
    local prefix=${1:-"instruqt"}
    local suffix=${2:-".tmp"}
    
    local temp_file=$(mktemp "/tmp/${prefix}_XXXXXX${suffix}")
    
    if [[ -z "$temp_file" ]]; then
        echo "ERROR: Failed to create temporary file"
        return 1
    fi
    
    echo "$temp_file"
    return 0
}

# Clean up temporary files
cleanup_temp_files() {
    local pattern=${1:-"instruqt_*"}
    
    echo "Cleaning up temporary files with pattern: $pattern"
    
    # Remove temporary files
    find /tmp -name "$pattern" -type f -delete 2>/dev/null || true
    
    echo "Temporary file cleanup completed"
}

# Backup file before modification
backup_file() {
    local filepath=$1
    local backup_suffix=${2:-".backup"}
    
    if [[ -z "$filepath" ]]; then
        echo "ERROR: Filepath not specified"
        return 1
    fi
    
    if [[ ! -f "$filepath" ]]; then
        echo "ERROR: File does not exist: $filepath"
        return 1
    fi
    
    local backup_path="${filepath}${backup_suffix}"
    
    echo "Creating backup: $filepath -> $backup_path"
    
    if ! cp "$filepath" "$backup_path"; then
        echo "ERROR: Failed to create backup of $filepath"
        return 1
    fi
    
    echo "Backup created: $backup_path"
    return 0
}

# Restore file from backup
restore_file() {
    local filepath=$1
    local backup_suffix=${2:-".backup"}
    
    if [[ -z "$filepath" ]]; then
        echo "ERROR: Filepath not specified"
        return 1
    fi
    
    local backup_path="${filepath}${backup_suffix}"
    
    if [[ ! -f "$backup_path" ]]; then
        echo "ERROR: Backup does not exist: $backup_path"
        return 1
    fi
    
    echo "Restoring from backup: $backup_path -> $filepath"
    
    if ! cp "$backup_path" "$filepath"; then
        echo "ERROR: Failed to restore from backup"
        return 1
    fi
    
    echo "File restored from backup: $filepath"
    return 0
}

# Display file operations status
files_info() {
    echo "Instruqt Files Library v${INSTRUQT_FILES_VERSION}"
    echo "HOME_DIR: ${HOME_DIR:-/home/replicant}"
    echo "Release directory: $(ls -ld ${HOME_DIR:-/home/replicant}/release 2>/dev/null || echo 'Not found')"
    echo "Temporary files: $(ls -l /tmp/instruqt_* 2>/dev/null | wc -l) found"
}