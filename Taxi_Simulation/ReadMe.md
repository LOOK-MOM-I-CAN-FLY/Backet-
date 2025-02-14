# Taxi Finder Simulator

A Go program demonstrating concurrent taxi service search using context cancellation. The program simulates searching for an available taxi across multiple services and selects the first successful response.

## Features

- **Concurrent Search**: Simultaneously checks multiple taxi services in separate goroutines.
- **Context Cancellation**: Stops all ongoing searches once the first taxi is found.
- **Randomized Success Simulation**: Each service has a 25% chance per iteration to "find" a taxi after an initial delay.
- **Graceful Shutdown**: Ensures all goroutines terminate cleanly after cancellation.

## How It Works

### Overview
The program launches a goroutine for each taxi service (Uber, Delimobil, Moscvich, Yandex Go). Each service:
1. Waits 3 seconds before starting availability checks (simulating initial setup time).
2. Repeatedly checks for taxi availability with a 25% success chance per iteration.
3. Sends its name to the results channel upon success or exits if the context is canceled.

The main function:
1. Listens for the first result from any service.
2. Cancels the context to stop all other searches.
3. Waits for all goroutines to exit before logging the winner.

### Concurrency Flow
1. **Goroutines Launch**: Each service starts a goroutine with a 3-second delay.
2. **Result Handling**: The first successful service sends its name to the channel, triggering context cancellation.
3. **Cancellation Propagation**: All other services stop their searches upon receiving the cancellation signal.
4. **Cleanup**: The program waits for all goroutines to finish before exiting.

## Getting Started

### Prerequisites
- Go 1.20 or later.

### Installation
1. Clone the repository or copy the code into a file named `main.go`.
2. Build the program:
   ```bash
   go build main.go
   ```

### Usage
Run the compiled binary:
```bash
./main
```

#### Example Output
```
2023/10/01 12:00:00 Stopped the search in "Uber" (context canceled)
2023/10/01 12:00:00 Stopped the search in "Delimobil" (context canceled)
2023/10/01 12:00:00 Stopped the search in "Moscvich" (context canceled)
2023/10/01 12:00:00 found car in "Yandex Go"
```

## Key Components
- **Context**: Used to propagate cancellation signals across goroutines.
- **WaitGroup**: Ensures the main function waits for all goroutines to finish.
- **Unbuffered Channel**: Synchronizes the first successful result.

## Customization
- **Initial Delay**: Modify the `time.Sleep(3 * time.Second)` in `getTaxi` to adjust the setup time.
- **Success Probability**: Change the `rand.Float64() > 0.75` condition to alter the success rate (e.g., `> 0.5` for 50% chance).

## Notes
- The 3-second delay ensures all services start checking availability simultaneously after the initial setup.
- The winner is non-deterministic due to the random success condition.
