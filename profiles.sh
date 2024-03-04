#!/bin/sh

# Granted: https://docs.commonfate.io/granted/getting-started/
# brew tap common-fate/granted and brew install granted.
alias assume="source assume"
PROFILE="TEST-01 TEST-02 TEST-03"

# Display menu to the user.
select profile in $PROFILE; do
case $profile in
TEST-01|TEST-02|TEST-03)

# Print the current date.
DATE=$( date +%Y-%m-%d-%H:%M:%S )
echo "Date/time: $DATE"

# Logic for opening AWS Console and setting credentials in your terminal..
if ! assume -c "$profile"; then { echo "! Operation error SSO: GetRoleCredentials, StatusCode: 403. ! Api error ForbiddenException: No access.
! Check your profiles in .config! You are not authorized or not logged in to this account!  " ; exit 1; }
fi
if ! assume "$profile"; then { echo "! Operation error SSO: GetRoleCredentials, StatusCode: 403. ! Api error ForbiddenException: No access.
! Check your profiles in .config! You are not authorized or not logged in to this account!  " ; exit 1; }
fi

# Check for the correct AWS account.
ID=$( aws sts get-caller-identity --query Account --output text)
ALIAS=$( aws iam list-account-aliases --query AccountAliases --output text)
echo "Check for correct Account ID and alias: "  "id = $ID" " | " "alias = $ALIAS"
echo

break
;;
*)
echo "You selected an invalid profile. Please select a valid profile."
;;
esac
done
