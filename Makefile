build:
	docker build -t my-app .

run:
	docker run --publish 2133:2133 --name test --rm my-app

stop:
	docker stop test