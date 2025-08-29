package main

import "fmt"

func main() {
	/*test := []int{2, 1, 3, 2, 1, 3, 3, 6, 9, 8, 5, 4, 1, 0, 2}
	sort(test, 0, len(test))
	fmt.Println(test)

	board := [][]byte{
		[]byte{'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'},
		[]byte{'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b'},
		[]byte{'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'},
		[]byte{'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b'},
		[]byte{'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'},
		[]byte{'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b'},
		[]byte{'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'},
		[]byte{'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b'},
		[]byte{'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a'},
		[]byte{'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b'},
	}
	words := []string{"ababababaa", "ababababab", "ababababac", "ababababad", "ababababae", "ababababaf", "ababababag", "ababababah", "ababababai", "ababababaj", "ababababak", "ababababal", "ababababam", "ababababan", "ababababao", "ababababap", "ababababaq", "ababababar", "ababababas", "ababababat", "ababababau", "ababababav", "ababababaw", "ababababax", "ababababay", "ababababaz", "ababababba", "ababababbb", "ababababbc", "ababababbd", "ababababbe", "ababababbf", "ababababbg", "ababababbh", "ababababbi", "ababababbj", "ababababbk", "ababababbl", "ababababbm", "ababababbn", "ababababbo", "ababababbp", "ababababbq", "ababababbr", "ababababbs", "ababababbt", "ababababbu", "ababababbv", "ababababbw", "ababababbx", "ababababby", "ababababbz", "ababababca", "ababababcb", "ababababcc", "ababababcd", "ababababce", "ababababcf", "ababababcg", "ababababch", "ababababci", "ababababcj", "ababababck", "ababababcl", "ababababcm", "ababababcn", "ababababco", "ababababcp", "ababababcq", "ababababcr", "ababababcs", "ababababct", "ababababcu", "ababababcv", "ababababcw", "ababababcx", "ababababcy", "ababababcz", "ababababda", "ababababdb", "ababababdc", "ababababdd", "ababababde", "ababababdf", "ababababdg", "ababababdh", "ababababdi", "ababababdj", "ababababdk", "ababababdl", "ababababdm", "ababababdn", "ababababdo", "ababababdp", "ababababdq", "ababababdr", "ababababds", "ababababdt", "ababababdu", "ababababdv", "ababababdw", "ababababdx", "ababababdy", "ababababdz", "ababababea", "ababababeb", "ababababec", "ababababed", "ababababee", "ababababef", "ababababeg", "ababababeh", "ababababei", "ababababej", "ababababek", "ababababel", "ababababem", "ababababen", "ababababeo", "ababababep", "ababababeq", "ababababer", "ababababes", "ababababet", "ababababeu", "ababababev", "ababababew", "ababababex", "ababababey", "ababababez", "ababababfa", "ababababfb", "ababababfc", "ababababfd", "ababababfe", "ababababff", "ababababfg", "ababababfh", "ababababfi", "ababababfj", "ababababfk", "ababababfl", "ababababfm", "ababababfn", "ababababfo", "ababababfp", "ababababfq", "ababababfr", "ababababfs", "ababababft", "ababababfu", "ababababfv", "ababababfw", "ababababfx", "ababababfy", "ababababfz", "ababababga", "ababababgb", "ababababgc", "ababababgd", "ababababge", "ababababgf", "ababababgg", "ababababgh", "ababababgi", "ababababgj", "ababababgk", "ababababgl", "ababababgm", "ababababgn", "ababababgo", "ababababgp", "ababababgq", "ababababgr", "ababababgs", "ababababgt", "ababababgu", "ababababgv", "ababababgw", "ababababgx", "ababababgy", "ababababgz", "ababababha", "ababababhb", "ababababhc", "ababababhd", "ababababhe", "ababababhf", "ababababhg", "ababababhh", "ababababhi", "ababababhj", "ababababhk", "ababababhl", "ababababhm", "ababababhn", "ababababho", "ababababhp", "ababababhq", "ababababhr", "ababababhs", "ababababht", "ababababhu", "ababababhv", "ababababhw", "ababababhx", "ababababhy", "ababababhz", "ababababia", "ababababib", "ababababic", "ababababid", "ababababie", "ababababif", "ababababig", "ababababih", "ababababii", "ababababij", "ababababik", "ababababil", "ababababim", "ababababin", "ababababio", "ababababip", "ababababiq", "ababababir", "ababababis", "ababababit", "ababababiu", "ababababiv", "ababababiw", "ababababix", "ababababiy", "ababababiz", "ababababja", "ababababjb", "ababababjc", "ababababjd", "ababababje", "ababababjf", "ababababjg", "ababababjh", "ababababji", "ababababjj", "ababababjk", "ababababjl", "ababababjm", "ababababjn", "ababababjo", "ababababjp", "ababababjq", "ababababjr", "ababababjs", "ababababjt", "ababababju", "ababababjv", "ababababjw", "ababababjx", "ababababjy", "ababababjz", "ababababka", "ababababkb", "ababababkc", "ababababkd", "ababababke", "ababababkf", "ababababkg", "ababababkh", "ababababki", "ababababkj", "ababababkk", "ababababkl", "ababababkm", "ababababkn", "ababababko", "ababababkp", "ababababkq", "ababababkr", "ababababks", "ababababkt", "ababababku", "ababababkv", "ababababkw", "ababababkx", "ababababky", "ababababkz", "ababababla", "abababablb", "abababablc", "ababababld", "abababable", "abababablf", "abababablg", "abababablh", "ababababli", "abababablj", "abababablk", "ababababll", "abababablm", "ababababln", "abababablo", "abababablp", "abababablq", "abababablr", "ababababls", "abababablt", "abababablu", "abababablv", "abababablw", "abababablx", "abababably", "abababablz", "ababababma", "ababababmb", "ababababmc", "ababababmd", "ababababme", "ababababmf", "ababababmg", "ababababmh", "ababababmi", "ababababmj", "ababababmk", "ababababml", "ababababmm", "ababababmn", "ababababmo", "ababababmp", "ababababmq", "ababababmr", "ababababms", "ababababmt", "ababababmu", "ababababmv", "ababababmw", "ababababmx", "ababababmy", "ababababmz", "ababababna", "ababababnb", "ababababnc", "ababababnd", "ababababne", "ababababnf", "ababababng", "ababababnh", "ababababni", "ababababnj", "ababababnk", "ababababnl", "ababababnm", "ababababnn", "ababababno", "ababababnp", "ababababnq", "ababababnr", "ababababns", "ababababnt", "ababababnu", "ababababnv", "ababababnw", "ababababnx", "ababababny", "ababababnz", "ababababoa", "ababababob", "ababababoc", "ababababod", "ababababoe", "ababababof", "ababababog", "ababababoh", "ababababoi", "ababababoj", "ababababok", "ababababol", "ababababom", "ababababon", "ababababoo", "ababababop", "ababababoq", "ababababor", "ababababos", "ababababot", "ababababou", "ababababov", "ababababow", "ababababox", "ababababoy", "ababababoz", "ababababpa", "ababababpb", "ababababpc", "ababababpd", "ababababpe", "ababababpf", "ababababpg", "ababababph", "ababababpi", "ababababpj", "ababababpk", "ababababpl", "ababababpm", "ababababpn", "ababababpo", "ababababpp", "ababababpq", "ababababpr", "ababababps", "ababababpt", "ababababpu", "ababababpv", "ababababpw", "ababababpx", "ababababpy", "ababababpz", "ababababqa", "ababababqb", "ababababqc", "ababababqd", "ababababqe", "ababababqf", "ababababqg", "ababababqh", "ababababqi", "ababababqj", "ababababqk", "ababababql", "ababababqm", "ababababqn", "ababababqo", "ababababqp", "ababababqq", "ababababqr", "ababababqs", "ababababqt", "ababababqu", "ababababqv", "ababababqw", "ababababqx", "ababababqy", "ababababqz", "ababababra", "ababababrb", "ababababrc", "ababababrd", "ababababre", "ababababrf", "ababababrg", "ababababrh", "ababababri", "ababababrj", "ababababrk", "ababababrl", "ababababrm", "ababababrn", "ababababro", "ababababrp", "ababababrq", "ababababrr", "ababababrs", "ababababrt", "ababababru", "ababababrv", "ababababrw", "ababababrx", "ababababry", "ababababrz", "ababababsa", "ababababsb", "ababababsc", "ababababsd", "ababababse", "ababababsf", "ababababsg", "ababababsh", "ababababsi", "ababababsj", "ababababsk", "ababababsl", "ababababsm", "ababababsn", "ababababso", "ababababsp", "ababababsq", "ababababsr", "ababababss", "ababababst", "ababababsu", "ababababsv", "ababababsw", "ababababsx", "ababababsy", "ababababsz", "ababababta", "ababababtb", "ababababtc", "ababababtd", "ababababte", "ababababtf", "ababababtg", "ababababth", "ababababti", "ababababtj", "ababababtk", "ababababtl", "ababababtm", "ababababtn", "ababababto", "ababababtp", "ababababtq", "ababababtr", "ababababts", "ababababtt", "ababababtu", "ababababtv", "ababababtw", "ababababtx", "ababababty", "ababababtz", "ababababua", "ababababub", "ababababuc", "ababababud", "ababababue", "ababababuf", "ababababug", "ababababuh", "ababababui", "ababababuj", "ababababuk", "ababababul", "ababababum", "ababababun", "ababababuo", "ababababup", "ababababuq", "ababababur", "ababababus", "ababababut", "ababababuu", "ababababuv", "ababababuw", "ababababux", "ababababuy", "ababababuz", "ababababva", "ababababvb", "ababababvc", "ababababvd", "ababababve", "ababababvf", "ababababvg", "ababababvh", "ababababvi", "ababababvj", "ababababvk", "ababababvl", "ababababvm", "ababababvn", "ababababvo", "ababababvp", "ababababvq", "ababababvr", "ababababvs", "ababababvt", "ababababvu", "ababababvv", "ababababvw", "ababababvx", "ababababvy", "ababababvz", "ababababwa", "ababababwb", "ababababwc", "ababababwd", "ababababwe", "ababababwf", "ababababwg", "ababababwh", "ababababwi", "ababababwj", "ababababwk", "ababababwl", "ababababwm", "ababababwn", "ababababwo", "ababababwp", "ababababwq", "ababababwr", "ababababws", "ababababwt", "ababababwu", "ababababwv", "ababababww", "ababababwx", "ababababwy", "ababababwz", "ababababxa", "ababababxb", "ababababxc", "ababababxd", "ababababxe", "ababababxf", "ababababxg", "ababababxh", "ababababxi", "ababababxj", "ababababxk", "ababababxl", "ababababxm", "ababababxn", "ababababxo", "ababababxp", "ababababxq", "ababababxr", "ababababxs", "ababababxt", "ababababxu", "ababababxv", "ababababxw", "ababababxx", "ababababxy", "ababababxz", "ababababya", "ababababyb", "ababababyc", "ababababyd", "ababababye", "ababababyf", "ababababyg", "ababababyh", "ababababyi", "ababababyj", "ababababyk", "ababababyl", "ababababym", "ababababyn", "ababababyo", "ababababyp", "ababababyq", "ababababyr", "ababababys", "ababababyt", "ababababyu", "ababababyv", "ababababyw", "ababababyx", "ababababyy", "ababababyz", "ababababza", "ababababzb", "ababababzc", "ababababzd", "ababababze", "ababababzf", "ababababzg", "ababababzh", "ababababzi", "ababababzj", "ababababzk", "ababababzl", "ababababzm", "ababababzn", "ababababzo", "ababababzp", "ababababzq", "ababababzr", "ababababzs", "ababababzt", "ababababzu", "ababababzv", "ababababzw", "ababababzx", "ababababzy", "ababababzz"}
	fmt.Println(findWords(board, words))

	fmt.Println((combine(4, 3)))
	*/

	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
}

func findMin(nums []int) int {
	index := findHelper(nums, 0, len(nums)-1)
	return nums[index]
}
func findHelper(nums []int, l, r int) int {
	if l == r {
		return r
	}

	if r == l+1 {
		if nums[l] > nums[r] {
			return r
		} else {
			return l
		}
	}

	mid := (l + r) >> 1

	if nums[mid] > nums[l] {
		return findHelper(nums, mid+1, r)
	}

	return findHelper(nums, l, mid)
}

func combine(n int, k int) [][]int {
	result := [][]int{}

	var dfs func(nn []int, i int)
	dfs = func(nn []int, i int) {
		if i == k {
			result = append(result, nn[1:])
			return
		}

		for j := nn[len(nn)-1] + 1; j <= n; j++ {
			r := append(nn, j)
			dfs(r, i+1)
		}
	}

	dfs([]int{0}, 0)

	return result
}

type Trie struct {
	neib   []*Trie
	letter byte
}

type LetterIndex struct {
	alphabit [26][]*Trie
}

func BuildTrieFromMatrix(matrix [][]byte) *LetterIndex {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return &LetterIndex{}
	}

	n, m := len(matrix), len(matrix[0])
	index := &LetterIndex{}
	for i := 0; i < 26; i++ {
		index.alphabit[i] = make([]*Trie, 0)
	}

	nodes := make([][]*Trie, n)
	for i := 0; i < n; i++ {
		nodes[i] = make([]*Trie, m)
		for j := 0; j < m; j++ {
			node := &Trie{
				neib:   make([]*Trie, 0),
				letter: matrix[i][j],
			}
			nodes[i][j] = node

			idx := matrix[i][j] - 'a'
			index.alphabit[idx] = append(index.alphabit[idx], node)
		}
	}

	// 构建邻接关系（上下左右四个方向）
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			currentNode := nodes[i][j]

			for _, dir := range directions {
				ni, nj := i+dir[0], j+dir[1]
				if ni >= 0 && ni < n && nj >= 0 && nj < m {
					neighbor := nodes[ni][nj]
					currentNode.neib = append(currentNode.neib, neighbor)
				}
			}
		}
	}

	return index
}

func findWords(board [][]byte, words []string) []string {

	index := BuildTrieFromMatrix(board)
	result := []string{}

	for i := range words {
		word, nw := words[i], len(words[i])
		start := index.alphabit[word[0]-'a']
		found := false

		for j := range start {
			visited := make(map[*Trie]bool)
			var dfs func(now *Trie, index int) bool

			dfs = func(now *Trie, index int) bool {
				if now.letter != word[index] {
					return false
				}

				if index == nw-1 {
					return true
				}
				visited[now] = true
				defer delete(visited, now)

				for i := range now.neib {
					if !visited[now.neib[i]] {
						if dfs(now.neib[i], index+1) {
							return true
						}
					}
				}
				return false
			}

			if dfs(start[j], 0) {
				found = true
				break
			}
		}
		if found {
			result = append(result, words[i])
		}
	}

	return result
}

func sort(aa []int, left int, right int) {

	if right-left <= 1 {
		return
	}

	i, j := left, right-1
	for i < j {
		for aa[j] > aa[left] && i < j {
			j--
		}
		for aa[i] <= aa[left] && i < j {
			i++
		}

		aa[i], aa[j] = aa[j], aa[i]
	}
	aa[left], aa[i] = aa[i], aa[left]
	sort(aa, left, i)
	sort(aa, i+1, right)
	return
}
