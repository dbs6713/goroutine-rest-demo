default:
	go build rest-basic.go
	go build rest-goroutines.go
	go build rest-channels.go
	time ./rest-basic
	time ./rest-goroutines
	time ./rest-channels