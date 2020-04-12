#!/bin/bash

projectName=${PWD##*/}   

echo $projectName

find . -type f -not -path '*/\.*' -print0 | xargs -0 sed -i "s/{{projectName}}/$projectName/g"

#rm -- "$0"
