#!/usr/bin/env bash
# test-libraries.sh - Simple test script to validate library functionality
# This script tests basic library loading and function availability

# Get the directory where this script is located
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Test library loading
test_library_loading() {
    echo "Testing library loading..."
    
    # Load all libraries
    source "$SCRIPT_DIR/instruqt-all.sh"
    
    # Check if all libraries loaded
    if [[ "$INSTRUQT_ALL_LOADED" == "true" ]]; then
        echo "✓ All libraries loaded successfully"
        return 0
    else
        echo "✗ Library loading failed"
        return 1
    fi
}

# Test function availability
test_function_availability() {
    echo "Testing function availability..."
    
    local functions_to_test=(
        "init_setup_script"
        "wait_for_instruqt_bootstrap"
        "setup_ssh_config"
        "ensure_tmux_session"
        "setup_release_directory"
        "setup_common_environment"
        "wait_for_kubernetes_api"
        "package_helm_chart"
        "setup_instruqt_environment"
        "display_library_info"
        "show_help"
    )
    
    local missing_functions=()
    
    for func in "${functions_to_test[@]}"; do
        if ! declare -f "$func" > /dev/null 2>&1; then
            missing_functions+=("$func")
        fi
    done
    
    if [[ ${#missing_functions[@]} -eq 0 ]]; then
        echo "✓ All expected functions are available"
        return 0
    else
        echo "✗ Missing functions: ${missing_functions[*]}"
        return 1
    fi
}

# Test basic function execution (safe functions only)
test_basic_functions() {
    echo "Testing basic function execution..."
    
    # Test info functions (these should be safe to run)
    if display_library_info > /dev/null 2>&1; then
        echo "✓ display_library_info works"
    else
        echo "✗ display_library_info failed"
        return 1
    fi
    
    # Test help function
    if show_help > /dev/null 2>&1; then
        echo "✓ show_help works"
    else
        echo "✗ show_help failed"
        return 1
    fi
    
    # Test file creation (in tmp directory)
    local test_file="/tmp/test-instruqt-lib.yaml"
    if create_yaml_file "$test_file" "test: content" "$(whoami)" > /dev/null 2>&1; then
        echo "✓ create_yaml_file works"
        rm -f "$test_file"
    else
        echo "✗ create_yaml_file failed"
        return 1
    fi
    
    return 0
}

# Test library versions
test_library_versions() {
    echo "Testing library versions..."
    
    local expected_version="1.0.0"
    local version_vars=(
        "INSTRUQT_ALL_VERSION"
        "INSTRUQT_BOOTSTRAP_VERSION"
        "INSTRUQT_SSH_VERSION"
        "INSTRUQT_SESSIONS_VERSION"
        "INSTRUQT_FILES_VERSION"
        "INSTRUQT_CONFIG_VERSION"
        "INSTRUQT_SERVICES_VERSION"
        "INSTRUQT_APPS_VERSION"
    )
    
    local version_errors=()
    
    for var in "${version_vars[@]}"; do
        if [[ "${!var}" != "$expected_version" ]]; then
            version_errors+=("$var=${!var}")
        fi
    done
    
    if [[ ${#version_errors[@]} -eq 0 ]]; then
        echo "✓ All library versions are correct (v$expected_version)"
        return 0
    else
        echo "✗ Version mismatches: ${version_errors[*]}"
        return 1
    fi
}

# Test error handling
test_error_handling() {
    echo "Testing error handling..."
    
    # Test function with missing required parameter
    if create_yaml_file "" "content" > /dev/null 2>&1; then
        echo "✗ Error handling failed - function should have returned error"
        return 1
    else
        echo "✓ Error handling works correctly"
        return 0
    fi
}

# Main test execution
main() {
    echo "Instruqt Libraries Test Suite"
    echo "============================"
    echo ""
    
    local test_results=()
    
    # Run all tests
    test_library_loading && test_results+=("✓ Library Loading") || test_results+=("✗ Library Loading")
    echo ""
    
    test_function_availability && test_results+=("✓ Function Availability") || test_results+=("✗ Function Availability")
    echo ""
    
    test_basic_functions && test_results+=("✓ Basic Functions") || test_results+=("✗ Basic Functions")
    echo ""
    
    test_library_versions && test_results+=("✓ Library Versions") || test_results+=("✗ Library Versions")
    echo ""
    
    test_error_handling && test_results+=("✓ Error Handling") || test_results+=("✗ Error Handling")
    echo ""
    
    # Summary
    echo "Test Results Summary:"
    echo "===================="
    for result in "${test_results[@]}"; do
        echo "  $result"
    done
    echo ""
    
    # Check if all tests passed
    local failed_tests=0
    for result in "${test_results[@]}"; do
        if [[ "$result" == *"✗"* ]]; then
            ((failed_tests++))
        fi
    done
    
    if [[ $failed_tests -eq 0 ]]; then
        echo "🎉 All tests passed! Libraries are ready for use."
        return 0
    else
        echo "❌ $failed_tests test(s) failed. Please check the library implementation."
        return 1
    fi
}

# Run tests if script is executed directly
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    main "$@"
fi