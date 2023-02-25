
all: host target

host: 
	@ ${CC} -O0 -g3 -ggdb main.c  -o syscall.arm64.debug
target:
	@ ${CC} main.c -o syscall.arm64
clean:
	@ rm -rf *.o
	@ rm target_bin