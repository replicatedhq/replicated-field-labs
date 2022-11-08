
show_credentials () {
    CYAN='\033[0;36m'
    GREEN='\033[1;32m'
    NC='\033[0m' # No Color
    password=$(echo -n "${INSTRUQT_PARTICIPANT_ID}" | sha256sum)
    echo "${GREEN}Credentials for ${CYAN}https://vendor.replicated.com"
    echo "${GREEN}Username: ${CYAN}${INSTRUQT_PARTICIPANT_ID}@replicated-labs.com"
    echo "${GREEN}Password: ${CYAN}${password::20}${NC}"
}