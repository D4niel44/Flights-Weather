
test:
	cd src/proyecto1/openweathermap/ && go test -v
	cd src/proyecto1/openweathermapcities/ && go test -v
	cd src/proyecto1/app/ && go test -v

createDB:
	cd scripts/ && go build && ./scripts

clean:
	rm -r bin && rm scripts/scripts

compile:
	cd src/proyecto1/app/ && go build
	mv src/proyecto1/app/app bin/app

build:
	make clean
	make test
	make createDB
	make compile