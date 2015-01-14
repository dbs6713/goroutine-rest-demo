default:
	go build rest-basic.go
	go build rest-waitgroup.go
	go build rest-channels.go
	time ./rest-basic
	time ./rest-waitgroup
	time ./rest-channels