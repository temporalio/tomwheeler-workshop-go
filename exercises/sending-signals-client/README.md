# Exercise #3: Sending a Signal from a Client

During this exercise, you will:

- Define and handle a Signal
- Retrieve a handle on the Workflow to Signal
- Use a Temporal Client to send the Signal

Make your changes to the code in the `practice` subdirectory (look for
`TODO` comments that will guide you to where you should make changes to
the code). If you need a hint or want to verify your changes, look at
the complete version in the `solution` subdirectory.

## Part A: Defining a Signal

In this part of the exercise, you will define your Signal.

1. Edit the `workflow.go` file.
2. Between the `import` and `Workflow()` blocks, define a new Signal 
   Type `struct` named `FulfillOrderSignal`. It should contain a single
   variable, of type `bool`, named `Fulfilled`.
3. Save the file.

## Part B: Handling the Signal

You will now handle the Signal you defined in part A, and let the
Workflow know what to do when it encounters the `FulfillOrderSignal`.

1. In `workflow.go`, locate the `signalChan.Receive()` call. This will
   block the Workflow until a Signal is received. 
2. Wrap the `ExecuteActivity()` call and the `logger.Info()` call in a
   test for `if signal.Fulfilled == true`.
3. Save the file.

## Part C: Import the Signal Type into your Signal Client

In this part of the exercise, you will create another Temporal Client
that sends a Signal. `signalclient/main.go` currently contains a
near-empty Temporal Client--it looks like a Starter or a Worker, but
it does not register any Workflows or do anything. You will use this
Client to send a Signal.

1. Edit the `signalclient/main.go` file.
2. Within the `main()` block, use the `FullfillOrderSignal` struct type
   from the `signals` module (i.e., the `workflow.go` file in the 
   parent directory) to create an instance of `FulfillOrderSignal`
   that contains `Fulfilled: true`.
3. Save the file.

## Part D: Signaling Your Workflow

Now you will add the `SignalWorkflow()` call itself.

1. Continue editing the `signalclient/main.go` file.
2. Within the `main()` block, after you create an instance of 
   `FulfillOrderSignal`, call `SignalWorkflow()` to send a Signal to
   your running Workflow. It needs, as arguments, `context.Background()`,
   your Workflow ID, your run ID (which can be an empty string), the
   name of the signal, and the Signal instance. It should assign its
   result to `err` so that it can be checked in the next line.
4. Save the file.

## Part E: Start Your Workflow

At this point, you can run your Workflow. It should run normally but
will wait to receive a Signal.

1. In one terminal, navigate to the `worker` subdirectory and run
   `go run main.go`.
2. In another terminal, navigate to the `starter` subdirectory and 
   run `go run main.go`. You should receive some logging from your 
   Worker along these lines:

```
2025/08/05 08:48:10 INFO  No logger configured for temporal client. Created default one.
2025/08/05 08:48:10 INFO  Started Worker Namespace default TaskQueue signals WorkerID 43388@hostname@
2025/08/05 08:48:21 INFO  Signal workflow started Namespace default TaskQueue signals WorkerID 43388@Omelas@ WorkflowType Workflow WorkflowID signals RunID 905330da-9c0f-490e-bd48-c6a9e8840f7a Attempt 1 input Plain text input
```

3. Your Workflow will not return as you added a blocking Signal call.
In the next step, you will Signal your Workflow.

## Part F: Send Your Signal

1. In a third terminal, navigate to the `signalclient` subdirectory and 
   run `go run main.go`. It will send a Signal to your Workflow. This
   should cause your Workflow to return, and the `Signal workflow completed`
   call to be logged in the terminal running your Worker:

```
2025/08/05 08:48:37 INFO  Signal workflow completed. Namespace default TaskQueue signals WorkerID 43388@hostname@ WorkflowType Workflow WorkflowID signals RunID 905330da-9c0f-490e-bd48-c6a9e8840f7a Attempt 1 result Received Plain text input
```

2. You have successfully sent a Signal from a Temporal Client to a running Workflow.

### This is the end of the exercise.
