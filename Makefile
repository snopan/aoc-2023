.PHONY: run
run: 
	go run ./cmd/day$(DAY).go ./output/day$(DAY) < ./input/day$(DAY)