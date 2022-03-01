go build -o ./output/
PID=$(ps -ef | grep psy-consult-backend | grep -v grep| awk '{print $2}')
kill -9 $PID
./output/psy-consult-backend