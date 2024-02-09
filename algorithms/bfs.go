package algorithms

//import "kubernetes/pkg/util/slice"
//
//func bfs(seller string) bool {
//	queue := make([]string, 0)
//	checked := make([]string, 0)
//
//	friends := map[string][]string{
//		"me": {"1", "2", "3"},
//		"1":  {"4", "5"},
//		"5":  {"6", "1"},
//	}
//	queue = append(queue, friends["me"]...)
//	for len(queue) != 0 {
//		person := queue[0]
//		queue = queue[1:]
//		if person == seller {
//			return true
//		}
//		if slice.ContainsString(checked, person, nil) {
//			continue
//		}
//		checked = append(checked, person)
//		queue = append(queue, friends[person]...)
//	}
//	return false
//}
