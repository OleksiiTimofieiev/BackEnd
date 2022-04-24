TARGET=archiver


init:
	@ export GOPATH=/home/user/Desktop/BackEnd/Go/

deploy: 
	@ go install
	@ cp ./../../bin/archiver /home/user/bin
	@ rm ./../../bin/archiver

clean:
	@ rm -rf SN.zip