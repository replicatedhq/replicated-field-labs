
show_credentials () {
    CYAN='\033[0;36m'
    GREEN='\033[1;32m'
    NC='\033[0m' # No Color
    password=$(get_password)
    echo -e "${GREEN}Credentials for ${CYAN}https://vendor.replicated.com"
    echo -e "${GREEN}Username: ${CYAN}${INSTRUQT_PARTICIPANT_ID}@replicated-labs.com"
    echo -e "${GREEN}Password: ${CYAN}${password}${NC}"
}

get_password () {
    password=$(echo -n "${INSTRUQT_PARTICIPANT_ID}" | sha256sum)
    echo ${password::20}
}