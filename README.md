# cadence-go-demo
Cadence demo for GO meetups

# GO Cadence Sample for GO Meetup
This package contains a sample demoed at [Go's 10th Anniversary Seattle Meetup](https://www.meetup.com/golang/events/265858683/)

More Cadence info at:

* [Cadence Service](https://github.com/uber/cadence)
* [Go Cadence Client](https://github.com/uber-go/cadence-client)
* [Cadence Java Client](https://github.com/uber/cadence-java-client)

## Overview of the Samples

### Money Transfer Sample

Demonstrates a simple transfer from one account to another. 

## Get the Samples

Run the following commands:

      git clone git@github.com:samarabbas/cadence-go-demo.git
      cd cadence-go-demo

## Build the Samples

      make

## Run Cadence Server

Run Cadence Server using Docker Compose:

    curl -O https://raw.githubusercontent.com/uber/cadence/master/docker/docker-compose.yml
    docker-compose up

If this does not work, see the instructions for running Cadence Server at https://github.com/uber/cadence/blob/master/README.md.

## Install Cadence CLI

[Command Line Interface Documentation](https://mfateev.github.io/cadence/docs/08_cli)

## Register the Domain

To register the *sample* domain, run the following command once before running any samples:

    cadence --do samples d re --rd 2 --oe "user@email" --desc "samples domain"

## See Cadence UI

The Cadence Server running in a docker container includes a Web UI.

Connect to [http://localhost:8088](http://localhost:8088).

Enter the *samples* domain. You'll see a "No Results" page. After running any sample, change the 
filter in the
top right corner from "Open" to "Closed" to see the list of the completed workflows.

Click on a *RUN ID* of a workflow to see more details about it. Try different view formats to get a different level
of details about the execution history.

## Run the samples

Each sample has specific requirements for running it. The following sections contain information about
how to run each of the samples after you've built them using the preceding instructions.


### Money Transfer Sample

Workflow Worker:
```
./bins/workflowWorker
```
Activities Worker:
```
./bins/activityWorker
```
Initiate Transfer:
```
cadence --do samples wf start --tl samples_workflow_tl --wt transfer --wid transferAToB-1 --dt 2 --et 1200 --wrp 1 --if ./transferRequest.json
```

### Batch Transfer Sample

Workflow Worker:
```
./bins/workflowWorker
```
Activities Worker:
```
./bins/activityWorker
```

Initiate Batch Transfer:
```
cadence --do samples wf start --tl samples_workflow_tl --wt batch-transfer --wid batch-transfer-1 --dt 2 --et 1200 --wrp 1 --if ./batchTransferRequest.json
```

Send Signal For First Withdrawal:
```
cadence --do samples wf signal -w batch-transfer-1 --name withdraw --if ./withdrawSignalPayload1.json
```

Send Signal For Second Withdrawal:
```
cadence --do samples wf signal -w batch-transfer-1 --name withdraw --if ./withdrawSignalPayload2.json
```

Send Signal For Third Withdrawal:
```
cadence --do samples wf signal -w batch-transfer-1 --name withdraw --if ./withdrawSignalPayload3.json
```

Query Withdrawal Count:
```
./cadence --do samples wf query -w batch-transfer-1 --qt get-count
```

Query Balance:
```
./cadence --do samples wf query -w batch-transfer-1 --qt get-balance
```