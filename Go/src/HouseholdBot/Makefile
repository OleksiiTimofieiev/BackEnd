default:
	@ GOARCH=arm go build .

run:
	go build .
	./HouseholdBot

install: default
	@ sshpass -p "orangepi" scp HouseholdBot orangepi@192.168.8.102:~/go