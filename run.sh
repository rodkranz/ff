#!/usr/bin/env bash


list=$(ls ./cmd/)

echo "Building..."
for file in ${list}
do
    go build -ldflags="-X main.Prefix=$file" -o "./tmp/ff_$file" "./cmd/$file/main.go"
    test "$?" = 1 && echo "Build $file Fail!!!" || echo "Build $file has been made successfully"
done
echo "Build finished"

echo "Getting data for tracing..."
for file in ${list}
do
    echo "************************************************************************************************"
    echo -en "\nExecuting ${file}:\n"
    ./tmp/ff_$file  "/Users/rodkranz/Projects/Go/src/git.naspersclassifieds.com/" -text "func main" > ./tmp/output/"${file}_$(date +"%FT%H%MZ").log"
    echo -en "\n"
    echo "************************************************************************************************"
done
echo "Get data data finished"

