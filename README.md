## Airport

Small web application using the Go
programming language and Echo
framework.


It consists of a single endpoint that accepts `POST` requests containing a JSON payload.

The payload is an array of pairs (tuples), where each pair represents a flight ticket.

Each pair is structured as `["Source", "Destination"]`. 
For example: `[["LAX","DXB"],["JFK","LAX"], ["SFO","SJC"], ["DXB","SFO"]].`

"Source" is the departure airport code, and "Destination" is the arrival airport code.

The application then reconstruct the complete itinerary from the sequence of ticket tuples provided.

The itinerary should be a sequence of airport codes in the order of travel, starting from the first departure to the final destination.

## Installation

- After installing go, Clone the repository to your local machine.
```git
git clone https://github.com/mahmoudfarouq/airport
```

- Run the project, which should install all project dependencies and run it.
```shell
go run main.go
```

- Post requests to `localhost:1323` following the attached OpenAPI schema. 

For example:
```http request
POST localhost:1323
Content-Type: application/json

[
  [
    "LAX",
    "DXB"
  ],
  [
    "JFK",
    "LAX"
  ],
  [
    "SFO",
    "SJC"
  ],
  [
    "DXB",
    "SFO"
  ]
]
```


- To run the tests:
```shell
go test ./...
```

## Dependency

The project only uses two dependencies. 
- `Echo` the web framework. It's not really needed for such trivial project, generally the standard library has all you'll ever need and if you really need more you can install a routing library.
- `testify` for testing.

Since the project is very small and trivial, there was no need to complicate things and add more folders and structure to it. Ideally a single file should be very decent.
