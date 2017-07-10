#!/usr/bin/env bash

# Colors
GREEN='\033[0;32m'
BLACK='\033[0;30m'
DARK_GRAY='\033[1;30m'
RED='\033[0;31m'
LIGHT_RED='\033[1;31m'
GREEN='\033[0;32m'
LIGHT_GREEN='\033[1;32m'
ORANGE='\033[0;33m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
LIGHT_BLUE='\033[1;34m'
PURPLE='\033[0;35m'
LIGHT_PURPLE='\033[1;35m'
CYAN='\033[0;36m'
LIGHT_CYAN='\033[1;36m'
LIGHT_GRAY='\033[0;37m'
WHITE='\033[1;37m'
NC='\033[0m'

clear

# Cabeçalho
echo ' '
printf "${GREEN}https://github.com/luk4z7/gerbo for the canonical source repository \n"
printf "Lucas Alves 2017 (c) Gerbo - Rodent and data extractor\n"
echo ' '

if [ $(uname) == "Darwin" ]; then
    ENVIRONMENT='MAC'
else
    ENVIRONMENT='LINUX'
fi
echo ' '

if [ $ENVIRONMENT == 'LINUX' ]; then

    if which figlet > /dev/null; then
        printf "${GREEN}"
        figlet Gerbo
	printf "${GREEN}Gerbo \n${NC}"
    else
        apt-get install -y figlet
        printf "${GREEN}"
        figlet Gerbo
	printf "${GREEN}Gerbo \n${NC}"
    fi
    echo ' '
    printf "${NC}"
else
	if which figlet > /dev/null; then
		printf "${GREEN}"
		figlet Gerbo
		printf "${GREEN}Gerbo \n${NC}"
	fi
	printf "${NC}"
echo ''
fi


# Docker
if which docker > /dev/null; then
    printf "${ORANGE}DOCKER${NC}\n"
    printf "${LIGHT_PURPLE}Generate new containers ${NC} ${WHITE}[ ${PURPLE}1 ${WHITE}]${NC} \n${LIGHT_PURPLE}Delete all containers ${NC} ${WHITE}  [ ${PURPLE}2 ${WHITE}]${NC} \n${LIGHT_PURPLE}Start new build ${NC} ${WHITE}        [ ${PURPLE}3 ${WHITE}]${NC}\n${LIGHT_PURPLE}Preview the logs ${NC} ${WHITE}       [ ${PURPLE}4 ${WHITE}]${NC}\n${LIGHT_PURPLE}Install dependencies ${NC} ${WHITE}   [ ${PURPLE}5 ${WHITE}]${NC}\n${LIGHT_PURPLE}Access the shell (gerbo) ${NC}${WHITE}[ ${PURPLE}6 ${WHITE}]${NC}\n${LIGHT_PURPLE}Access the shell (mongo) ${NC}${WHITE}[ ${PURPLE}7 ${WHITE}]${NC}\n"
    read gerar

    if [ -n "$gerar" ]; then
        if [ $gerar == '1' ]; then
            printf "${ORANGE}Generating new containers ... ${NC}\n"
            docker-compose -f $(pwd)/src/docker-compose.yml ps
            docker-compose -f $(pwd)/src/docker-compose.yml up -d
            docker-compose -f $(pwd)/src/docker-compose.yml ps
        fi
        if [ $gerar == '2' ]; then
            printf "${ORANGE}Removing all containers ... ${NC}\n"
            docker-compose -f $(pwd)/src/docker-compose.yml kill
            docker-compose -f $(pwd)/src/docker-compose.yml rm
        fi
        if [ $gerar == '3' ]; then
        	printf "${LIGHT_PURPLE}Would you like to start a new compilation with cache?${NC} ${WHITE} [ ${PURPLE}yes ${WHITE}]: ${NC} "
        	read cache

        	printf "${ORANGE}Starting a new build process ... ${NC}\n"
        	if [ -n "$cache" ]; then
				if [ $cache == 'no' ]; then
					docker-compose -f $(pwd)/src/docker-compose.yml build --no-cache
				fi
				if [ $cache == 'yes' ]; then
					docker-compose -f $(pwd)/src/docker-compose.yml build
				fi
        	else
        	    docker-compose -f $(pwd)/src/docker-compose.yml build
        	fi
        fi
	    if [ $gerar == '4' ]; then
            printf "${ORANGE}Preview logs ... ${NC}\n"
            docker-compose -f $(pwd)/src/docker-compose.yml logs -f
        fi
        if [ $gerar == '5' ]; then
            if which glide > /dev/null; then
                printf "${ORANGE}Install dependencies ... ${NC}\n"
                glide install
                rm -rf $(pwd)/src/vendor
                mv ./vendor/ $(pwd)/src/
            else
                printf "${BLUE}Installation of glide not found${NC}\n"
                printf "${BLUE}See more details on: https://glide.sh${NC}\n"
            fi
        fi

        if [ $gerar == '6' ]; then
            clear
            docker exec -it gerbo bash
        fi

        if [ $gerar == '7' ]; then
            clear
            docker exec -it mongo bash
        fi

    fi
    echo ' '
else
    printf "${BLUE}Installation of docker not found${NC}\n"
fi

