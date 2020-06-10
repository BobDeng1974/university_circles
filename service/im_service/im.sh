#!/bin/bash

#echo pwd
#chmod -R u=rwx,g=rwx,o=x ./pb/home
#chmod -R 777 ./pb/home
#echo date +"%Y%m%d%H%M%S";
gsed -i "s/,omitempty//g" $(grep ",omitempty" -rl ./pb/home/*)