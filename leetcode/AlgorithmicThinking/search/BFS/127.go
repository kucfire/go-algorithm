/*
	leetcode tag:127 title:单词接龙
*/
package bfs

import "math"

// method of BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordID := map[string]int{}
	// wordId := map[string]int{}
	graph := [][]int{}
	// graph := [][]int{}

	addWord := func(word string) int {
		id, has := wordID[word]
		if !has {
			id = len(wordID)
			wordID[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	// addWord := func(word string) int {
	//     id, has := wordId[word]
	//     if !has {
	//         id = len(wordId)
	//         wordId[word] = id
	//         graph = append(graph, []int{})
	//     }
	//     return id
	// }

	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}
	// addEdge := func(word string) int {
	//     id1 := addWord(word)
	//     s := []byte(word)
	//     for i, b := range s {
	//         s[i] = '*'
	//         id2 := addWord(string(s))
	//         graph[id1] = append(graph[id1], id2)
	//         graph[id2] = append(graph[id2], id1)
	//         s[i] = b
	//     }
	//     return id1
	// }

	for _, word := range wordList {
		addEdge(word)
	}

	// for _, word := range wordList {
	//     addEdge(word)
	// }

	beginID := addEdge(beginWord)
	endID, has := wordID[endWord]
	if !has {
		return 0
	}

	// beginId := addEdge(beginWord)
	// endId, has := wordId[endWord]
	// if !has {
	//     return 0
	// }

	const inf int = math.MaxInt64
	dist := make([]int, len(wordID))
	for i := range dist {
		dist[i] = inf
	}

	// const inf int = math.MaxInt64
	// dist := make([]int, len(wordId))
	// for i := range dist {
	//     dist[i] = inf
	// }

	dist[beginID] = 0
	queue := []int{beginID}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		if v == endID {
			return dist[endID]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inf {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0

	// dist[beginId] = 0
	// queue := []int{beginId}
	// for len(queue) > 0 {
	//     v := queue[0]
	//     queue = queue[1:]
	//     if v == endId {
	//         return dist[endId]/2 + 1
	//     }
	//     for _, w := range graph[v] {
	//         if dist[w] == inf {
	//             dist[w] = dist[v] + 1
	//             queue = append(queue, w)
	//         }
	//     }
	// }
	// return 0
}
