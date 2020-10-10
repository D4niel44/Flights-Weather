

createDB:
	cd scripts && go build && ./scripts

loadDatasets:
	mkdir -p bin/datasets && cp datasets/* bin/datasets

clean:
	rm -rf bin && rm -f scripts/scripts

compile:
	go build  -o bin/weather -i myp/Tarea01/app

test:
	make clean
	make createDB
	go test -v myp/Tarea01/...

build:
	make clean
	make createDB
	make loadDatasets
	make test
	make compile
