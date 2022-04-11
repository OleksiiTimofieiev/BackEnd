TARGET=GoogleDriveUpload


init:
	@ export GOPATH="/home/user/Desktop/BackEnd/Go/"

deploy: 
	@ go install
	@ cp ./../../bin/GoogleDriveUpload /home/user/bin
	@ rm ./../../bin/GoogleDriveUpload

clean:
	@ rm -rf SN.zip