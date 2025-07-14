#!/usr/bin/env bash
# example-setup-script.sh - Example of how to use the Instruqt libraries
# This demonstrates how a typical setup script would look after refactoring

# Load all Instruqt libraries
source "$(dirname "${BASH_SOURCE[0]}")/instruqt-all.sh"

# Example 1: Simple shell environment setup
example_shell_setup() {
    echo "=== Example 1: Shell Environment Setup ==="
    
    # Single function call replaces ~30 lines of duplicated code
    setup_instruqt_environment "shell"
    
    # Script-specific logic would go here
    echo "Shell environment ready for use"
}

# Example 2: Cluster environment setup
example_cluster_setup() {
    echo "=== Example 2: Cluster Environment Setup ==="
    
    # Setup cluster environment (includes Kubernetes config)
    setup_instruqt_environment "cluster"
    
    # Wait for Kubernetes API to be available
    wait_for_kubernetes_api
    
    # Script-specific logic would go here
    echo "Cluster environment ready for use"
}

# Example 3: Node environment setup
example_node_setup() {
    echo "=== Example 3: Node Environment Setup ==="
    
    # Setup node environment (includes KOTS CLI installation)
    setup_instruqt_environment "node"
    
    # Script-specific logic would go here
    echo "Node environment ready for use"
}

# Example 4: Custom setup using individual functions
example_custom_setup() {
    echo "=== Example 4: Custom Setup ==="
    
    # Initialize with debug mode
    init_setup_script true
    
    # Setup SSH with custom configuration
    setup_ssh_config "LogLevel DEBUG"
    
    # Create custom tmux session
    ensure_tmux_session "custom-session" "replicant"
    
    # Setup custom directory structure
    create_directory "/home/replicant/custom" "replicant:replicant" "755"
    
    # Create configuration file
    create_yaml_file "/home/replicant/custom/config.yaml" "
apiVersion: v1
kind: Config
metadata:
  name: custom-config
data:
  environment: instruqt
" "replicant"
    
    echo "Custom setup completed"
}

# Example 5: Application deployment workflow
example_app_deployment() {
    echo "=== Example 5: Application Deployment ==="
    
    # Setup environment for application deployment
    setup_instruqt_environment "shell"
    
    # Setup Slackernews application
    setup_slackernews
    
    # Update Helm dependencies
    update_helm_dependencies "/home/replicant/slackernews"
    
    # Package and release
    package_and_release "/home/replicant/slackernews" "" "Example release" "Unstable"
    
    echo "Application deployment completed"
}

# Example 6: Service waiting and health checks
example_service_monitoring() {
    echo "=== Example 6: Service Monitoring ==="
    
    # Wait for multiple services
    wait_for_port "localhost" 22 30
    wait_for_http_endpoint "http://localhost:8080/health" 200 60
    
    # Perform health check
    health_check "ssh" "kubernetes-api" "replicated-api"
    
    echo "Service monitoring completed"
}

# Example 7: Cleanup workflow
example_cleanup() {
    echo "=== Example 7: Cleanup ==="
    
    # Cleanup environment
    cleanup_instruqt_environment "shell"
    
    # Cleanup release artifacts
    cleanup_release_artifacts
    
    echo "Cleanup completed"
}

# Main execution
main() {
    echo "Instruqt Libraries Example Script"
    echo "=================================="
    
    # Display library information
    display_library_info
    
    # Run examples (comment out as needed)
    example_shell_setup
    echo ""
    
    example_cluster_setup
    echo ""
    
    example_node_setup
    echo ""
    
    example_custom_setup
    echo ""
    
    example_app_deployment
    echo ""
    
    example_service_monitoring
    echo ""
    
    example_cleanup
    echo ""
    
    echo "All examples completed successfully!"
}

# Run main function if script is executed directly
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    main "$@"
fi