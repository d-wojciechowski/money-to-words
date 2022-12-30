run:
	go run main.go
bench:
	go test -v -bench . -run=^#
bench5:
	go test -v -bench . -count 5 -run=^#
test:
	go test -v