test: memory cpu
	go test memory
	go test cpu

cpu: src/cpu/*

memory: src/memory/memory.go
	go build memory
