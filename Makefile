build: 
	sudo docker build -t=gdhameeja/books .

run: 
	sudo docker run -d -p 127.0.0.1:8080:8888 --name test gdhameeja/books /bin/bash

remove:
	sudo docker rm test
