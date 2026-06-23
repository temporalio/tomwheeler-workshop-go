# Tom Wheeler's Go Workshop (June 2026)
This repository contains code for hands-on exercises used in a
Temporal workshop that Tom Wheeler is teaching on June 23rd and 
24th.

## Prerequisites
To run these exercises locally, you'll need three things:

1. A clone of this repository to your laptop
2. A working installation of [Go](https://go.dev/dl/), version 1.24 or higher
3. A working installation of the `temporal` command-line tool, also known as
   the Temporal CLI. If you haven't installed this, follow the instructions
   in the [Install the Temporal CLI](https://docs.temporal.io/cli#install) 
   section of the Temporal documentation. 

You will also need a standard Web browser, such as Chrome, Firefox, or 
Safari.

Before beginning the first exercise, run `temporal server start-dev` to start
a local Temporal Service. Leave it running for the duration of the workshop,
since it's required for the subsequent exercises too.



## Hands-On Exercises

Directory Name                        | Exercise
:------------------------------------ | :----------------------------
`exercises/farewell-workflow`         | [Exercise 1](exercises/farewell-workflow/README.md)
`exercises/durable-execution`         | [Exercise 2](exercises/durable-execution/README.md)
`exercises/sending-signals-client`    | [Exercise 3](exercises/sending-signals-client/README.md)
`exercises/sending-signals-external`  | [Exercise 4](exercises/sending-signals-external/README.md) (Optional)
`exercises/querying-workflows`        | [Exercise 5](exercises/querying-workflows/README.md)
`exercises/non-retryable-error-types` | [Exercise 6](exercises/non-retryable-error-types/README.md)
`exercises/rollback-with-saga`        | [Exercise 7](exercises/rollback-with-saga/README.md) (Optional)


Each of these exercise directories contains two subdirectories: 
1. `practice` - This is where you'll modify the code as instructed
2. `solution` - This is the completed version for comparison


## Reference
The following links provide additional information that you may find 
helpful as you work through the hands-on exercises.
* [General Temporal Documentation](https://docs.temporal.io/)
* [General Go Language Documentation](https://go.dev/doc/)
* [Temporal Go Developer Guide](https://docs.temporal.io/develop/go/)
* [Temporal Go SDK API Documentation](https://pkg.go.dev/go.temporal.io/sdk)
