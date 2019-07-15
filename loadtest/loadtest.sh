#!/bin/bash
#
# Ping for specified HTTP(S) endpoint continuously and display the status code.
#

TARGET=${1:-http://localhost:30000/}
DELAY_SECONDS=${2:-0.1}
i=0

while :
do
    status=`curl --max-time 2 -o /dev/null -sw '%{http_code}' $TARGET`
    if [ $status == "200" ]
    then
        echo "$i:" $status
    else
        echo "$i:      " $status
    fi

    sleep $DELAY_SECONDS
    let "i++"
done
