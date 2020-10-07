
test:
	cd src/proyecto1/openweathermap/ && go test -v
	cd src/proyecto1/app/ && go test -v

createDB:
	cd scripts/ && go build && ./scripts

clean:
	rm -r bin && rm scripts/scripts