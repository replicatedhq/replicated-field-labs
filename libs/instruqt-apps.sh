#!/usr/bin/env bash
# instruqt-apps.sh - Application and release management functions
# Version: 1.0.0

# Library metadata
INSTRUQT_APPS_VERSION="1.0.0"
INSTRUQT_APPS_LOADED=true

# Package Helm chart
package_helm_chart() {
    local chart_path=$1
    local destination=${2:-"./release"}
    local version=${3:-""}
    
    if [[ -z "$chart_path" ]]; then
        echo "ERROR: Chart path not specified"
        return 1
    fi
    
    if [[ ! -d "$chart_path" ]]; then
        echo "ERROR: Chart directory not found: $chart_path"
        return 1
    fi
    
    echo "Packaging Helm chart: $chart_path"
    
    # Ensure destination directory exists
    mkdir -p "$destination"
    
    # Package chart
    local helm_args=("package" "$chart_path" "--destination" "$destination")
    
    if [[ -n "$version" ]]; then
        helm_args+=("--version" "$version")
    fi
    
    if ! helm "${helm_args[@]}"; then
        echo "ERROR: Failed to package Helm chart"
        return 1
    fi
    
    echo "Helm chart packaged successfully"
    return 0
}

# Update Helm chart dependencies
update_helm_dependencies() {
    local chart_path=$1
    
    if [[ -z "$chart_path" ]]; then
        echo "ERROR: Chart path not specified"
        return 1
    fi
    
    if [[ ! -d "$chart_path" ]]; then
        echo "ERROR: Chart directory not found: $chart_path"
        return 1
    fi
    
    echo "Updating Helm chart dependencies: $chart_path"
    
    # Check if Chart.yaml exists
    if [[ ! -f "$chart_path/Chart.yaml" ]]; then
        echo "ERROR: Chart.yaml not found in $chart_path"
        return 1
    fi
    
    # Update dependencies
    if ! helm dependency update "$chart_path"; then
        echo "ERROR: Failed to update Helm chart dependencies"
        return 1
    fi
    
    echo "Helm chart dependencies updated successfully"
    return 0
}

# Add Replicated SDK to Helm chart
add_replicated_sdk() {
    local chart_path=$1
    local sdk_version=${2:-""}
    
    if [[ -z "$chart_path" ]]; then
        echo "ERROR: Chart path not specified"
        return 1
    fi
    
    if [[ ! -f "$chart_path/Chart.yaml" ]]; then
        echo "ERROR: Chart.yaml not found in $chart_path"
        return 1
    fi
    
    # Get SDK version if not provided
    if [[ -z "$sdk_version" ]]; then
        if command -v get_replicated_sdk_version &> /dev/null; then
            sdk_version=$(get_replicated_sdk_version)
        else
            echo "ERROR: SDK version not provided and get_replicated_sdk_version not available"
            return 1
        fi
    fi
    
    echo "Adding Replicated SDK version $sdk_version to chart: $chart_path"
    
    # Backup Chart.yaml
    cp "$chart_path/Chart.yaml" "$chart_path/Chart.yaml.backup"
    
    # Add SDK dependency
    if ! yq -i '.dependencies += [{"name": "replicated", "repository": "oci://registry.replicated.com/library", "version": "'"$sdk_version"'"}]' "$chart_path/Chart.yaml"; then
        echo "ERROR: Failed to add Replicated SDK to Chart.yaml"
        # Restore backup
        cp "$chart_path/Chart.yaml.backup" "$chart_path/Chart.yaml"
        return 1
    fi
    
    echo "Replicated SDK added to chart successfully"
    return 0
}

# Create Replicated release
create_replicated_release() {
    local chart_path=$1
    local app_slug=${2:-""}
    local release_notes=${3:-"Automated release"}
    local channel=${4:-"Unstable"}
    
    if [[ -z "$chart_path" ]]; then
        echo "ERROR: Chart path not specified"
        return 1
    fi
    
    if [[ -z "$app_slug" ]]; then
        if command -v get_app_slug &> /dev/null; then
            app_slug=$(get_app_slug)
        else
            echo "ERROR: App slug not provided and get_app_slug not available"
            return 1
        fi
    fi
    
    echo "Creating Replicated release for app: $app_slug"
    
    # Create release
    if ! replicated release create \
        --app "$app_slug" \
        --chart "$chart_path" \
        --release-notes "$release_notes"; then
        echo "ERROR: Failed to create Replicated release"
        return 1
    fi
    
    echo "Replicated release created successfully"
    return 0
}

# Promote release to channel
promote_release() {
    local app_slug=${1:-""}
    local channel=${2:-"Unstable"}
    local release_sequence=${3:-""}
    
    if [[ -z "$app_slug" ]]; then
        if command -v get_app_slug &> /dev/null; then
            app_slug=$(get_app_slug)
        else
            echo "ERROR: App slug not provided and get_app_slug not available"
            return 1
        fi
    fi
    
    # Get latest release sequence if not provided
    if [[ -z "$release_sequence" ]]; then
        echo "Getting latest release sequence..."
        local api_token=""
        if command -v get_api_token &> /dev/null; then
            api_token=$(get_api_token)
        else
            echo "ERROR: Cannot get release sequence without API token"
            return 1
        fi
        
        local releases
        if ! releases=$(curl --silent --header "Accept: application/json" \
                           --header "Authorization: $api_token" \
                           "https://api.replicated.com/vendor/v3/apps/$app_slug/releases"); then
            echo "ERROR: Failed to get releases"
            return 1
        fi
        
        release_sequence=$(echo "$releases" | jq -r '.releases[0].sequence')
    fi
    
    echo "Promoting release $release_sequence to channel $channel"
    
    # Promote release
    if ! replicated release promote \
        --app "$app_slug" \
        --sequence "$release_sequence" \
        --channel "$channel"; then
        echo "ERROR: Failed to promote release"
        return 1
    fi
    
    echo "Release promoted successfully"
    return 0
}

# Setup Slackernews application
setup_slackernews() {
    local home_dir=${1:-"${HOME_DIR:-/home/replicant}"}
    
    echo "Setting up Slackernews application..."
    
    # Use get_slackernews function if available
    if command -v get_slackernews &> /dev/null; then
        if ! get_slackernews; then
            echo "ERROR: Failed to setup Slackernews using get_slackernews"
            return 1
        fi
    else
        echo "ERROR: get_slackernews function not available"
        return 1
    fi
    
    echo "Slackernews setup completed"
    return 0
}

# Install tools to embedded cluster
install_to_embedded_cluster() {
    local install_command=$1
    local install_dir=${2:-"/var/lib/embedded-cluster/bin"}
    
    if [[ -z "$install_command" ]]; then
        echo "ERROR: Install command not specified"
        return 1
    fi
    
    echo "Installing tools to embedded cluster: $install_dir"
    
    # Create installation directory
    mkdir -p "$install_dir"
    
    # Set installation path
    export REPL_INSTALL_PATH="$install_dir"
    
    # Execute installation command
    if ! eval "$install_command"; then
        echo "ERROR: Failed to install tools to embedded cluster"
        return 1
    fi
    
    echo "Tools installed to embedded cluster successfully"
    return 0
}

# Install KOTS CLI
install_kots_cli() {
    local install_dir=${1:-"/var/lib/embedded-cluster/bin"}
    
    echo "Installing KOTS CLI..."
    
    install_to_embedded_cluster "curl https://kots.io/install | bash" "$install_dir"
    
    echo "KOTS CLI installation completed"
}

# Package and release workflow
package_and_release() {
    local chart_path=$1
    local app_slug=${2:-""}
    local release_notes=${3:-"Automated release"}
    local channel=${4:-"Unstable"}
    local version=${5:-""}
    
    if [[ -z "$chart_path" ]]; then
        echo "ERROR: Chart path not specified"
        return 1
    fi
    
    echo "Starting package and release workflow..."
    
    # Update dependencies
    if ! update_helm_dependencies "$chart_path"; then
        echo "ERROR: Failed to update dependencies"
        return 1
    fi
    
    # Package chart
    if ! package_helm_chart "$chart_path" "${HOME_DIR:-/home/replicant}/release" "$version"; then
        echo "ERROR: Failed to package chart"
        return 1
    fi
    
    # Create release
    if ! create_replicated_release "$chart_path" "$app_slug" "$release_notes" "$channel"; then
        echo "ERROR: Failed to create release"
        return 1
    fi
    
    # Promote release
    if ! promote_release "$app_slug" "$channel"; then
        echo "ERROR: Failed to promote release"
        return 1
    fi
    
    echo "Package and release workflow completed successfully"
    return 0
}

# Clean up release artifacts
cleanup_release_artifacts() {
    local release_dir=${1:-"${HOME_DIR:-/home/replicant}/release"}
    
    echo "Cleaning up release artifacts in: $release_dir"
    
    if [[ -d "$release_dir" ]]; then
        # Remove .tgz files
        find "$release_dir" -name "*.tgz" -type f -delete 2>/dev/null || true
        
        # Remove charts directory
        rm -rf "$release_dir/charts" 2>/dev/null || true
        
        echo "Release artifacts cleaned up"
    else
        echo "Release directory not found: $release_dir"
    fi
}

# === VERSION MANAGEMENT FUNCTIONS ===
# (Migrated from header.sh)

# Get the latest Replicated SDK version
get_replicated_sdk_version() {
    set +eu pipefail
    local replicated_sdk_version=$(agent variable get REPLICATED_SDK_VERSION)

    # if we don't already have a version, fetch one
    if [[ -z "$replicated_sdk_version" ]]; then
        set -eu pipefail
        local token=$(curl --silent "https://registry.replicated.com/v2/token?scope=repository:library/replicated:pull&service=registry.replicated.com" | jq -r .token)
        replicated_sdk_version=$(curl --silent -H "Authorization: Bearer ${token}" https://registry.replicated.com/v2/library/replicated/tags/list | jq -r '.tags[]' | awk -F '[.-]' '{
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
    fi

    set -eu
    echo "${replicated_sdk_version}"
}

# Get the latest Embedded Cluster version
get_embedded_cluster_version() {
    set +eu pipefail
    local embedded_cluster_version=$(agent variable get EMBEDDED_CLUSTER_VERSION)

    # if we don't already have a version, fetch one
    if [[ -z "$embedded_cluster_version" ]]; then
        embedded_cluster_version=$(curl -s "https://api.github.com/repos/replicatedhq/embedded-cluster/releases/latest" | jq -r .tag_name)
    fi

    set -eu pipefail
    echo "${embedded_cluster_version}"
}

# === APPLICATION AND API FUNCTIONS ===
# (Migrated from header.sh)

# Get application slug from Replicated API
get_app_slug() {
    local application=${1:-"Slackernews"}
    local access_token=$(get_api_token)
    local app_slug=$(curl --header 'Accept: application/json' --header "Authorization: ${access_token}" https://api.replicated.com/vendor/v3/apps | jq -r --arg application "${application}" '.apps[] | select( .name | startswith( $application )) | .slug')
    echo "${app_slug}"
}

# Get application ID from Replicated API
get_app_id() {
    local application=${1:-"Slackernews"}
    local access_token=$(get_api_token)
    local app_id=$(curl --header 'Accept: application/json' --header "Authorization: ${access_token}" https://api.replicated.com/vendor/v3/apps | jq -r --arg application "${application}" '.apps[] | select( .name | startswith( $application )) | .id')
    echo "${app_id}"
}

# Get customer ID with caching - this function is now defined in instruqt-config.sh
# but kept here for backwards compatibility
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

# Get license ID from Replicated API
get_license_id() {
    local customer=${1}
    local access_token=$(get_api_token)
    local app_id=$(get_app_id)
    local license_id=$(curl --header 'Accept: application/json' --header "Authorization: ${access_token}" https://api.replicated.com/vendor/v3/app/${app_id}/customers | jq -r --arg name "${customer}" '.customers[] | select ( .name == $name ) | .installationId')
    echo "${license_id}"
}

# === COMPLETE APPLICATION SETUP ===
# (Migrated from header.sh)

# Complete Slackernews application setup and configuration
get_slackernews() {
    echo "Setting up Slackernews application..."
    
    # get the app slug, since there's only one app created by the automation, just grab the first in the list
    local app_slug=$(get_app_slug)

    # grab the sources for the Helm chart using a community license
    helm registry login chart.slackernews.io --username marc@replicated.com --password 2ViYIi8SDFubA8XwQRhJtcrwn4C
    helm pull --untar oci://chart.slackernews.io/slackernews/slackernews

    # specify the nodeport for NGINX so we get a consistent and addressable endpoint
    # TODO: Update upstream to take this as a value
    sed -i '17 a\    nodePort: 30443' slackernews/templates/nginx-service.yaml

    # remove the Replicated SDK dependency, if we add more dependencies to
    # Slackernews this will need to be revised
    yq -i 'del(.dependencies)' slackernews/Chart.yaml

    # start version numbers over to simplify the lab text
    yq -i '.version = "0.1.0"' slackernews/Chart.yaml

    # get rid of troubleshoot files since learners will create their own
    rm -rf slackernews/troubleshoot slackernews/templates/preflights.yaml slackernews/templates/support-bundle.yaml

    # set the values file to use the right proxy image URI
    local web_image=$(yq .images.slackernews.repository slackernews/values.yaml)
    local rewritten_web_image=${web_image//images.slackernews.io/proxy.replicated.com}
    rewritten_web_image=${rewritten_web_image//proxy\/slackernews/proxy\/${app_slug}}
    yq -i ".images.slackernews.repository = \"${rewritten_web_image}\"" slackernews/values.yaml

    local nginx_image=$(yq .images.nginx.repository slackernews/values.yaml)
    local rewritten_nginx_image=${nginx_image//images.slackernews.io/proxy.replicated.com}
    rewritten_nginx_image=${rewritten_nginx_image//proxy\/slackernews/proxy\/${app_slug}}
    yq -i ".images.nginx.repository = \"${rewritten_nginx_image}\"" slackernews/values.yaml

    # add some optional components to make the application a bit more representative
    yq -i '.nginx.enabled = true' slackernews/values.yaml
    yq -i '.postgres.deploy_postgres = true' slackernews/values.yaml
    yq -i '.postgres.enabled = true' slackernews/values.yaml
    yq -i '.postgres.password = "thisisasecret"' slackernews/values.yaml

    # address awkward scenario where a TLS cert is required even if TLS isn't enabled
    # TODO: Fix upstream to not require TLS certs unless TLS is enabled
    openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 -keyout server.key -out server.crt -subj "/CN=Slackernews" -addext "subjectAltName = DNS:$(get_slackernews_domain)" \
      && yq -i ".service.tls.key = \"$(cat server.key)\"" slackernews/values.yaml \
      && rm server.key \
      && yq -i ".service.tls.cert = \"$(cat server.crt)\"" slackernews/values.yaml \
      && rm server.crt

    # since we have the certs anyway, let's enable TLS
    yq -i '.service.tls.enabled = true' slackernews/values.yaml

    # let's also delete the values injected by Replicated so users can release
    # the chart without any sort of double injection
    yq -i 'del(.replicated)' slackernews/values.yaml
    yq -i 'del(.global.replicated)' slackernews/values.yaml
    
    echo "Slackernews application setup completed"
}

# Display application management status
apps_info() {
    echo "Instruqt Apps Library v${INSTRUQT_APPS_VERSION}"
    echo "Helm version: $(helm version --short 2>/dev/null || echo 'Not installed')"
    echo "Replicated CLI: $(replicated version 2>/dev/null || echo 'Not installed')"
    echo "KOTS CLI: $(kubectl kots version 2>/dev/null || echo 'Not installed')"
    echo "Release directory: $(ls -ld ${HOME_DIR:-/home/replicant}/release 2>/dev/null || echo 'Not found')"
    echo "Chart files: $(find ${HOME_DIR:-/home/replicant} -name "Chart.yaml" 2>/dev/null | wc -l) found"
}