#!/usr/bin/env bash
# test-refactored-header.sh - Test refactored header.sh system
# This script validates that the refactored header.sh provides all expected functions

# Test header.sh loading and function availability
test_header_functionality() {
    echo "Testing refactored header.sh functionality..."
    
    # Source the refactored header.sh
    source "$(dirname "${BASH_SOURCE[0]}")/header.sh"
    
    # List of functions that should be available from original header.sh
    local expected_functions=(
        # Credential functions (moved to instruqt-config.sh)
        "get_username"
        "get_password"
        "get_api_token"
        "get_admin_console_password"
        "get_slackernews_domain"
        "show_credentials"
        
        # Application functions (moved to instruqt-apps.sh)
        "get_replicated_sdk_version"
        "get_embedded_cluster_version"
        "get_app_slug"
        "get_app_id"
        "get_customer_id"
        "get_license_id"
        "get_slackernews"
        
        # New header function
        "header_info"
    )
    
    local missing_functions=()
    
    # Check each function
    for func in "${expected_functions[@]}"; do
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

# Test that library versions are correct
test_library_versions() {
    echo "Testing library versions..."
    
    local expected_version="1.0.0"
    local header_version="2.0.0"
    
    # Check header version
    if [[ "$HEADER_VERSION" != "$header_version" ]]; then
        echo "✗ Header version mismatch: expected $header_version, got $HEADER_VERSION"
        return 1
    fi
    
    # Check library versions
    local version_vars=(
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
        echo "✓ All library versions are correct"
        return 0
    else
        echo "✗ Version mismatches: ${version_errors[*]}"
        return 1
    fi
}

# Test function categories work correctly
test_function_categories() {
    echo "Testing function categories..."
    
    # Test credential functions
    if ! command -v get_username &> /dev/null; then
        echo "✗ Credential functions not available"
        return 1
    fi
    
    # Test app functions  
    if ! command -v get_app_slug &> /dev/null; then
        echo "✗ Application functions not available"
        return 1
    fi
    
    # Test new utility functions
    if ! command -v setup_instruqt_environment &> /dev/null; then
        echo "✗ New utility functions not available"
        return 1
    fi
    
    echo "✓ All function categories are working"
    return 0
}

# Test backward compatibility
test_backward_compatibility() {
    echo "Testing backward compatibility..."
    
    # Test that functions work the same way as before
    # (using safe functions that don't require API access)
    
    # Test domain function
    local domain_pattern="cluster-30443-.*\.env\.play\.instruqt\.com"
    local test_domain=""
    
    # Set a test participant ID for testing
    export INSTRUQT_PARTICIPANT_ID="test123"
    
    if test_domain=$(get_slackernews_domain) && [[ "$test_domain" =~ $domain_pattern ]]; then
        echo "✓ get_slackernews_domain works correctly"
    else
        echo "✗ get_slackernews_domain failed or returned unexpected format: $test_domain"
        return 1
    fi
    
    # Test username function
    local expected_username="test123@replicated-labs.com"
    local test_username=""
    
    if test_username=$(get_username) && [[ "$test_username" == "$expected_username" ]]; then
        echo "✓ get_username works correctly"
    else
        echo "✗ get_username failed or returned unexpected value: $test_username"
        return 1
    fi
    
    echo "✓ Backward compatibility maintained"
    return 0
}

# Main test execution
main() {
    echo "Refactored Header.sh Test Suite"
    echo "==============================="
    echo ""
    
    local test_results=()
    
    # Run all tests
    test_header_functionality && test_results+=("✓ Header Functionality") || test_results+=("✗ Header Functionality")
    echo ""
    
    test_library_versions && test_results+=("✓ Library Versions") || test_results+=("✗ Library Versions")
    echo ""
    
    test_function_categories && test_results+=("✓ Function Categories") || test_results+=("✗ Function Categories")
    echo ""
    
    test_backward_compatibility && test_results+=("✓ Backward Compatibility") || test_results+=("✗ Backward Compatibility")
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
        echo "🎉 All tests passed! Refactored header.sh is working correctly."
        echo ""
        echo "📋 Migration Summary:"
        echo "  ➤ Original header.sh backed up to header-original.sh"
        echo "  ➤ Functions migrated to appropriate libraries"
        echo "  ➤ Backward compatibility maintained"
        echo "  ➤ All existing scripts will continue to work"
        return 0
    else
        echo "❌ $failed_tests test(s) failed. Please check the refactored implementation."
        return 1
    fi
}

# Run tests if script is executed directly
if [[ "${BASH_SOURCE[0]}" == "${0}" ]]; then
    main "$@"
fi