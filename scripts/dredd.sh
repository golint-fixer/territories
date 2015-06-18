#!/bin/sh
if [ "$CI" = "true" ]
then
    npm update -g npm
    npm install -g dredd@0.3.8
fi
contact -m &
sleep 3
PID=$!
dredd apiary.apib http://localhost:8080/
RESULT=$?
kill -9 $PID
exit $RESULT
