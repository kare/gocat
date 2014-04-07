#!/bin/sh

rm tests/rand.txt
BYTES=`echo "79" | bc`
for a in `seq 1 500`; do
    LINE=`head -c $BYTES /dev/random | base64`
    echo $a $LINE
    echo "$LINE" >> tests/rand.txt
done
