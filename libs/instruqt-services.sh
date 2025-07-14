#!/usr/bin/env bash
# instruqt-services.sh - Service and API management functions
# Version: 1.0.0

# Library metadata
INSTRUQT_SERVICES_VERSION="1.0.0"
INSTRUQT_SERVICES_LOADED=true

# Wait for Kubernetes API to be available
wait_for_kubernetes_api() {
    local api_endpoint=${1:-"http://localhost:8001/api"}
    local timeout=${2:-120}
    local elapsed=0
    
    echo "Waiting for Kubernetes API at $api_endpoint..."
    
    while ! curl --fail --silent --output /dev/null "$api_endpoint" 2>/dev/null && [[ $elapsed -lt $timeout ]]; do
        echo "Waiting for Kubernetes API to be available (${elapsed}s elapsed)"
        sleep 1
        ((elapsed++))
    done
    
    if [[ $elapsed -ge $timeout ]]; then
        echo "ERROR: Kubernetes API not available within ${timeout} seconds"
        return 1
    fi
    
    echo "Kubernetes API is available"
    return 0
}

# Wait for a generic service to be available
wait_for_service() {
    local service_name=$1
    local check_command=$2
    local timeout=${3:-60}
    local elapsed=0
    
    if [[ -z "$service_name" || -z "$check_command" ]]; then
        echo "ERROR: Service name and check command are required"
        return 1
    fi
    
    echo "Waiting for service '$service_name' to be available..."
    
    while ! eval "$check_command" &> /dev/null && [[ $elapsed -lt $timeout ]]; do
        echo "Waiting for service '$service_name' (${elapsed}s elapsed)"
        sleep 1
        ((elapsed++))
    done
    
    if [[ $elapsed -ge $timeout ]]; then
        echo "ERROR: Service '$service_name' not available within ${timeout} seconds"
        return 1
    fi
    
    echo "Service '$service_name' is available"
    return 0
}

# Wait for HTTP endpoint to be available
wait_for_http_endpoint() {
    local endpoint=$1
    local expected_status=${2:-200}
    local timeout=${3:-60}
    local elapsed=0
    
    if [[ -z "$endpoint" ]]; then
        echo "ERROR: HTTP endpoint not specified"
        return 1
    fi
    
    echo "Waiting for HTTP endpoint: $endpoint (expecting status $expected_status)"
    
    while [[ $elapsed -lt $timeout ]]; do
        local status_code=$(curl --silent --output /dev/null --write-out "%{http_code}" "$endpoint" 2>/dev/null || echo "000")
        
        if [[ "$status_code" == "$expected_status" ]]; then
            echo "HTTP endpoint is available: $endpoint (status $status_code)"
            return 0
        fi
        
        echo "Waiting for HTTP endpoint (${elapsed}s elapsed, status: $status_code)"
        sleep 1
        ((elapsed++))
    done
    
    echo "ERROR: HTTP endpoint not available within ${timeout} seconds"
    return 1
}

# Check if port is listening
wait_for_port() {
    local host=${1:-"localhost"}
    local port=$2
    local timeout=${3:-60}
    local elapsed=0
    
    if [[ -z "$port" ]]; then
        echo "ERROR: Port not specified"
        return 1
    fi
    
    echo "Waiting for port $port on $host..."
    
    while ! timeout 1 bash -c "</dev/tcp/$host/$port" &> /dev/null && [[ $elapsed -lt $timeout ]]; do
        echo "Waiting for port $port on $host (${elapsed}s elapsed)"
        sleep 1
        ((elapsed++))
    done
    
    if [[ $elapsed -ge $timeout ]]; then
        echo "ERROR: Port $port on $host not available within ${timeout} seconds"
        return 1
    fi
    
    echo "Port $port on $host is available"
    return 0
}

# Wait for systemd service to be active
wait_for_systemd_service() {
    local service_name=$1
    local timeout=${2:-60}
    local elapsed=0
    
    if [[ -z "$service_name" ]]; then
        echo "ERROR: Service name not specified"
        return 1
    fi
    
    echo "Waiting for systemd service '$service_name' to be active..."
    
    while ! systemctl is-active --quiet "$service_name" && [[ $elapsed -lt $timeout ]]; do
        echo "Waiting for service '$service_name' (${elapsed}s elapsed)"
        sleep 1
        ((elapsed++))
    done
    
    if [[ $elapsed -ge $timeout ]]; then
        echo "ERROR: Service '$service_name' not active within ${timeout} seconds"
        return 1
    fi
    
    echo "Service '$service_name' is active"
    return 0
}

# Check Replicated API connectivity
check_replicated_api() {
    local api_token=${1:-""}
    local timeout=${2:-30}
    
    if [[ -z "$api_token" ]]; then
        if command -v get_api_token &> /dev/null; then
            api_token=$(get_api_token)
        else
            echo "ERROR: API token not provided and get_api_token not available"
            return 1
        fi
    fi
    
    echo "Checking Replicated API connectivity..."
    
    local response
    if ! response=$(curl --silent --max-time "$timeout" \
                        --header "Accept: application/json" \
                        --header "Authorization: $api_token" \
                        https://api.replicated.com/vendor/v3/apps); then
        echo "ERROR: Failed to connect to Replicated API"
        return 1
    fi
    
    # Check if response contains apps
    if ! echo "$response" | jq -e '.apps' > /dev/null 2>&1; then
        echo "ERROR: Invalid response from Replicated API"
        return 1
    fi
    
    echo "Replicated API connectivity verified"
    return 0
}

# Wait for file to exist
wait_for_file() {
    local filepath=$1
    local timeout=${2:-60}
    local elapsed=0
    
    if [[ -z "$filepath" ]]; then
        echo "ERROR: File path not specified"
        return 1
    fi
    
    echo "Waiting for file: $filepath"
    
    while [[ ! -f "$filepath" ]] && [[ $elapsed -lt $timeout ]]; do
        echo "Waiting for file '$filepath' (${elapsed}s elapsed)"
        sleep 1
        ((elapsed++))
    done
    
    if [[ $elapsed -ge $timeout ]]; then
        echo "ERROR: File '$filepath' not found within ${timeout} seconds"
        return 1
    fi
    
    echo "File found: $filepath"
    return 0
}

# Wait for Kubernetes configuration file
wait_for_kubeconfig() {
    local kubeconfig_path=${1:-"/etc/rancher/k3s/k3s.yaml"}
    local timeout=${2:-120}
    
    echo "Waiting for Kubernetes configuration..."
    
    # Wait for file to exist
    if ! wait_for_file "$kubeconfig_path" "$timeout"; then
        return 1
    fi
    
    # Wait for k3s service to be ready
    echo "Waiting for k3s service to be ready..."
    local start_time=$(date +%s)
    local end_time=$((start_time + timeout))
    
    while [ $(date +%s) -lt $end_time ]; do
        if kubectl --kubeconfig="$kubeconfig_path" cluster-info > /dev/null 2>&1; then
            echo "Kubernetes configuration is ready"
            return 0
        fi
        
        echo "Kubernetes cluster not ready yet, waiting..."
        sleep 5
    done
    
    echo "ERROR: Kubernetes configuration is not valid after $timeout seconds"
    return 1
}

# Health check for multiple services
health_check() {
    local -a services=("$@")
    local failed_services=()
    
    if [[ ${#services[@]} -eq 0 ]]; then
        echo "ERROR: No services specified for health check"
        return 1
    fi
    
    echo "Performing health check for services: ${services[*]}"
    
    for service in "${services[@]}"; do
        case "$service" in
            "kubernetes-api")
                if ! wait_for_kubernetes_api "http://localhost:8001/api" 10; then
                    failed_services+=("$service")
                fi
                ;;
            "replicated-api")
                if ! check_replicated_api "" 10; then
                    failed_services+=("$service")
                fi
                ;;
            "ssh")
                if ! wait_for_port "localhost" 22 10; then
                    failed_services+=("$service")
                fi
                ;;
            *)
                echo "WARNING: Unknown service for health check: $service"
                ;;
        esac
    done
    
    if [[ ${#failed_services[@]} -gt 0 ]]; then
        echo "ERROR: Health check failed for services: ${failed_services[*]}"
        return 1
    fi
    
    echo "Health check passed for all services"
    return 0
}

# Start essential services
start_services() {
    local -a services=("$@")
    
    if [[ ${#services[@]} -eq 0 ]]; then
        services=("ssh" "docker")
    fi
    
    echo "Starting services: ${services[*]}"
    
    for service in "${services[@]}"; do
        echo "Starting service: $service"
        if ! systemctl start "$service"; then
            echo "ERROR: Failed to start service: $service"
            return 1
        fi
    done
    
    echo "All services started successfully"
    return 0
}

# Display service status
services_info() {
    echo "Instruqt Services Library v${INSTRUQT_SERVICES_VERSION}"
    echo "Service status:"
    echo "  SSH: $(systemctl is-active sshd 2>/dev/null || echo 'inactive')"
    echo "  Docker: $(systemctl is-active docker 2>/dev/null || echo 'inactive')"
    echo "  K3s: $(systemctl is-active k3s 2>/dev/null || echo 'inactive')"
    echo "Port checks:"
    echo "  SSH (22): $(timeout 1 bash -c '</dev/tcp/localhost/22' 2>/dev/null && echo 'open' || echo 'closed')"
    echo "  HTTP (80): $(timeout 1 bash -c '</dev/tcp/localhost/80' 2>/dev/null && echo 'open' || echo 'closed')"
}