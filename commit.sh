#!/bin/bash

git add .
echo "#####"
echo "ADDED"
echo "###################"

git commit -m "$1"
echo "#######################"
echo "commit message $1 ADDED"
echo "###############################"


git push origin main
echo "#############"
echo "COMMIT PUSHED"
echo "####################"