default:
	@ GOARCH=arm go build .

run:
	@ go build .
	@ ./NetworkAvailabilityCheck

install: default
	@ sshpass -p "orangepi" scp NetworkAvailabilityCheck orangepi@192.168.8.102:~/go