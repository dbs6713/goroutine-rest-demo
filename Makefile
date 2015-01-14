default:
	# go run conflates build time and runtime. We do this in two steps so we
	# can measure runtime independently.
	go build rest-basic.go
	go build rest-waitgroup.go
	go build rest-channels.go
	time ./rest-basic
	time ./rest-waitgroup
	time ./rest-channels
