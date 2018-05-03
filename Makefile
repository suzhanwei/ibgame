BIN := `pwd | awk -F"/" '{print $$NF}'`

default: dev
dev:
	go build -gcflags "-N -l" -i -x -o ./bin/$(BIN) && (ps -ef | grep $(BIN) | grep -v grep | awk '{print $$2}' | while read f; do kill -9 $$f ;done ) && (nohup ./bin/$(BIN) &)

