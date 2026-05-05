!#bin/bash

git add .
echo "ADDED"
###################
git commit -m "$1"
echo "commit message $1 ADDED"
###############################
git push origin main
echo "COMMIT PUSHED"
####################