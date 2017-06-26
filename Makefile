all: account dist/signin.html
	docker-compose up --build

clean:
	rm -rf account dist/*

.PHONY: clean

account:
	go build

dist/signin.html: frontend/node_modules
	cd frontend && yarn build

frontend/node_modules:
	cd frontend && yarn install
