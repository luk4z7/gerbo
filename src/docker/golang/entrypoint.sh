#!/usr/bin/env bash

cd $BUILDPATH
go build -o /gerbo
cp /gerbo /usr/bin/
chmod u+x /usr/bin/gerbo

cron
touch /var/log/cron.log
echo "* * * * * gerbo --sync >> /var/log/cron.log 2>&1" >> mycron
echo "* * * * * sqlite3 /data/twitter-movie-ratings.db < /data/insert.sql && echo 'Generated registers on database sqlite by robots' >> /var/log/cron.log 2>&1" >> mycron
echo "* * * * * sleep 15; sqlite3 /data/twitter-movie-ratings.db < /data/insert.sql && echo 'Generated registers on database sqlite by robots' >> /var/log/cron.log 2>&1" >> mycron
echo "* * * * * sleep 30; sqlite3 /data/twitter-movie-ratings.db < /data/insert.sql && echo 'Generated registers on database sqlite by robots' >> /var/log/cron.log 2>&1" >> mycron
echo "* * * * * sleep 45; sqlite3 /data/twitter-movie-ratings.db < /data/insert.sql && echo 'Generated registers on database sqlite by robots' >> /var/log/cron.log 2>&1" >> mycron

crontab mycron
rm -rf mycron

cd $BUILDPATH
reflex -c /var/exec/reflex.conf &

tail -f /var/log/cron.log

wait $!