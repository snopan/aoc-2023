.PHONY: run
run: 
	go run ./src/cmd $(DAY) $(PART) < ./input/day$(DAY).txt