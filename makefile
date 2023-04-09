all: proto backend fronted

backend: proto
	cd ./backend
	go build
	cd ..
	echo 'backend complie finished'

fronted: proto
	echo 'fronted complie finished'

proto: 
	echo 'proto complie finished'


.PHONY: clean
clean:
