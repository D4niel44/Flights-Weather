
test:
	go test -v myp/Tarea01/...

createDB:
	cd scripts && go build && ./scripts

loadDatasets:
	mkdir -p bin/datasets && cp datasets/* bin/datasets

clean:
	rm -r bin && rm scripts/scripts

compile:
	go build  -o bin/weather -i myp/Tarea01/app

build:
	make clean
	make createDB
	make loadDatasets
	make test
	make compile