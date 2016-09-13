
coverage:
	go get github.com/axw/gocov/gocov
	go get gopkg.in/matm/v1/gocov-html
	gocov test . -coverpkg . > coverage.json
	gocov-html coverage.json > coverage.html

.PHONY: coverage
