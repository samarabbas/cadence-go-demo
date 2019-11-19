.PHONY: test bins clean
PROJECT_ROOT = github.com/samarabbas/cadence-go-demo

# default target
default: test

PROGS = workflowWorker \
	activityWorker \

TEST_ARG ?= -race -v -timeout 5m

# Automatically gather all srcs
ALL_SRC := $(shell find . -name "*.go")

# all directories with *_test.go files in them
TEST_DIRS=./workflows \
	./activities \

workflowWorker:
	go build -i -o bins/workflowWorker ./cmd/workflowWorker/*.go

activityWorker:
	go build -i -o bins/activityWorker ./cmd/activityWorker/*.go

bins: workflowWorker \
	activityWorker \

test: bins
	@rm -f test
	@rm -f test.log
	@echo $(TEST_DIRS)
	@for dir in $(TEST_DIRS); do \
		go test -coverprofile=$@ "$$dir" | tee -a test.log; \
	done;

clean:
	rm -rf bins
