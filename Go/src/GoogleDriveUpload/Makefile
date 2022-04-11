TARGET=GoogleDriveUpload

default: 
	@ go build -o ${TARGET} 

init:
	@ export GOPATH="/home/user/Desktop/BackEnd/Go/"



run: clean default
	@ ./${TARGET}

clean:
	@ rm ${TARGET}
	@ rm -rf upload