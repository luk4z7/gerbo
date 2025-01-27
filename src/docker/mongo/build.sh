#!/bin/bash

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

mongod &

sleep 5

printf "${ORANGE}Creating user....${NC}\n"
/usr/bin/mongo gerbo --eval "db.createUser({ user: 'user.gerbo', pwd: '12345', roles: [ { role: 'dbOwner', db: 'gerbo' } ] })"

mongoimport --verbose --db gerbo --collection movies --file /var/www/insert.json

cron
touch /var/log/cron.log
echo "* * * * * sh /remove.sh && echo 'Removing registers with min value on database sqlite by robots' >> /var/log/cron.log 2>&1" >> mycron
echo "* * * * * sleep 15; sh /remove.sh && echo 'Removing registers with min value on database sqlite by robots' >> /var/log/cron.log 2>&1" >> mycron
crontab mycron
rm -rf mycron


printf "${ORANGE}Configuration Finish!${NC}\n"

tail -f /var/log/cron.log

wait $!