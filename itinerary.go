package main

func computeItinerary(request [][]string) []string {
	allDestinations := make(map[string]struct{}, len(request))
	for _, r := range request {
		allDestinations[r[1]] = struct{}{}
	}

	graph := make(map[string]string, len(request))
	for _, r := range request {
		graph[r[0]] = r[1]
	}

	startingPoint := getStartingPoint(graph, allDestinations)

	path := make([]string, 0, len(request)+1)
	for startingPoint != "" {
		path = append(path, startingPoint)
		startingPoint = graph[startingPoint]
	}

	return path
}

func getStartingPoint(graph map[string]string, destinations map[string]struct{}) string {
	for k, _ := range graph {
		// If there is a Source that has never been a destination, it is the starting point.
		if _, ok := destinations[k]; !ok {
			return k
		}
	}
	return ""
}
