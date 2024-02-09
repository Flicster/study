package algorithms

type states map[string]struct{}

func GreedyStationsSearch() []string {
	statesNeeded := states{
		"mt": {},
		"wa": {},
		"or": {},
		"id": {},
		"nv": {},
		"ut": {},
		"ca": {},
		"az": {},
	}

	stations := map[string]states{
		"kone": {
			"id": {}, "nv": {}, "ut": {},
		},
		"ktwo": {
			"wa": {}, "id": {}, "mt": {},
		},
		"kthree": {
			"or": {}, "nv": {}, "ca": {},
		},
		"kfour": {
			"nv": {}, "ut": {},
		},
		"kfive": {
			"ca": {}, "az": {},
		},
	}
	finalStations := make([]string, 0)
	for len(statesNeeded) > 0 {
		bestStation := ""
		statesCovered := states{}
		for station, stationStates := range stations {
			covered := make(states)
			for ns := range statesNeeded {
				for ss := range stationStates {
					if ns == ss {
						covered[ns] = struct{}{}
					}
				}
			}
			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = stationStates
			}
		}
		finalStations = append(finalStations, bestStation)
		for sc := range statesCovered {
			delete(statesNeeded, sc)
		}
	}
	return finalStations
}
