package main

import (
	"fmt"
	"strconv"
	"strings"
)

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	max, _ := strconv.Atoi(os.Args[1])

// 	var uniques []int
// loop:
// 	for len(uniques) < max {
// 		n := rand.Intn(max) + 1
// 		fmt.Print(n, " ")

// 		for _, u := range uniques {
// 			if u == n {
// 				continue loop
// 			}
// 		}
// 		// uniques[found] = n
// 		uniques = append(uniques, n)

// 	}
// 	fmt.Println("\n\nuniques:", uniques)
// 	sort.Ints(uniques)
// 	fmt.Println("sorted:", uniques)
// }

// func main (){
// 	var todo []string
// 	todo = append(todo, "sing")
// 	todo = append(todo, "run", "code", "play")
// 	// todo = append(todo,"run")
// 	// todo = append(todo,"code")
// 	// todo = append(todo,"play")

// 	tomorrow := []string{"see mom", "learn go"}
// 	todo = append(todo, tomorrow...)

// 	s.Show("todo", todo)
// }

// func main() {
// 	items := []string{
// 		"pacman", "mario", "tetris", "doom",
// 		"galaga", "frogger", "asteroids", "simcity",
// 		"metroid", "defender", "rayman", "tempest",
// 		"ultima",
// 	}

// 	s.MaxPerLine = 4
// 	s.Show("items", items)

// 	top3 := items[0:3]
// 	s.Show("top 3 item: ",top3)
// }

// func main() {
// 	items := []string{
// 		"pacman", "mario", "tetris", "doom",
// 		"galaga", "frogger", "asteroids", "simcity",
// 		"metroid", "defender", "rayman", "tempest",
// 		"ultima",
// 	}

// 	const pageSize = 4

// 	l := len(items)
// 	for from := 0; from < l; from += pageSize {
// 		to := from + pageSize
// 		// fmt.Printf("%d:%d\n",from,to)
// 		if to > len(items){
// 			to = l
// 		}
// 		currentPage := items[from:to]

// 		head := fmt.Sprintf("page #%d", (from/pageSize)+1)
// 		s.Show(head, currentPage)
// 	}
// }

// Backing Array
//type collection []string
//func main(){
//	s.PrintElementAddr = true
//
//	data := collection{"slices", "are", "awesome", "period","!"}
//	change(data)
//	s.Show("main's data", data)
//	fmt.Printf("main's data slice address: %p\n", &data)
//
//	array := [...]string{"slices", "are", "awesome", "period","!"}
//	fmt.Printf("array's size: %d bytes\n", unsafe.Sizeof(array))
//	fmt.Printf("slice's size: %d bytes\n",unsafe.Sizeof(data))
//}
//
//func change(data collection){
//	data[2] = "brilliant!"
//	s.Show("change's data", data)
//	fmt.Printf("main's data slice address: %p\n", &data)
//
//}

//func main (){
//	s.PrintBacking = true
//
//	var nums []int
//	s.Show("no backing array", nums)
//
//	nums = append(nums, 1,3)
//	s.Show("a_",nums)
//
//	nums = append(nums, 2)
//	s.Show("free capacity",nums)
//
//	nums = append(nums, 4)
//	s.Show("no allocation", nums)
//
//	nums = append(nums, nums[2:]...)
//	s.Show("num <-- nums[2:]", nums)
//
//	nums = append(nums[:2], 7, 9)
//	s.Show("num[:2]", nums)
//
//	nums = nums[:6]
//	s.Show("nums:extend", nums)
//
//	var games []string
//	s.Show("games:", games)
//
//	games = []string{}
//	s.Show("games:",games)
//	games = []string{"pacman", "boomz","DBD","mario"}
//	s.Show("games", games)
//}

//func main() {
//	s.PrintBacking = true
//	s.MaxPerLine = 30
//	s.Width = 150
//
//	var nums []int
//
//	screen.Clear()
//	for cap(nums) <= 128 {
//		screen.MoveTopLeft()
//		s.Show("nums", nums)
//		nums = append(nums, rand.Intn(9)+ 1)
//		time.Sleep(time.Second/4)
//	}
//
//
//}

//func main()  {
//	ages, oldCap := []int{1}, 1.
//
//	for len(ages) < 5e5{
//		ages = append(ages, 1)
//
//		c:= float64(cap(ages))
//		if c != oldCap {
//			fmt.Printf("len:%-10d cap:%-10g growth:%.2f\n",
//			len(ages), c,c/oldCap)
//		}
//		oldCap = c
//	}
//
// }
//
//func main()  {
//	s.PrintBacking = true
//	nums := []int{1, 3, 2, 4}
//	odds := nums[:2:2]
//	odds = append(odds, 5, 7)
//
//	s.Show("nums ", nums)
//	s.Show("odds ", odds)
//
//}

//func main() {
//	s.PrintBacking = true
//	s.MaxPerLine = 10
//
//	tasks := []string{"jump", "run", "read"}
//
//	upTasks := make([]string, 0 ,len(tasks))
//	s.Show("upTasks",upTasks)
//
//	//var upTasks []string
//	for _, task := range tasks {
//		upTasks =  append(upTasks, strings.ToUpper(task))
//		s.Show("up tasks", upTasks)
//	}
//
//	//upTasks = append(upTasks,"PLAY")
//	//s.Show("upTasks", upTasks)
//}

//func main()  {
//	data := []float64{10, 25,30, 50}
//
//	//newData := []float64{99, 100}
//	//for i := range newData {
//	//	data[i] = newData[i]
//	//}
//	//copy(data, newData)
//	//data = append(data[:0], []float64{10, 5, 15, 0, 20}...)
//
//	saved := make([]float64, len(data))
//	copy(saved,data)
//
//	s.Show("Probabilities",data)
//
//	fmt.Printf("Is it gonna rain? %.f%% chance. \n",
//		(data[0]+data[1]+data[2]+data[3])/float64(len(data)))
//}

func main() {
	// 1st day: $200, $100
	// 2nd day: $500
	// 3rd day: $50, $25, and $75

	//spendings := [][]int{
	//	{200, 100},
	//	{500},
	//	{50, 25, 75},
	//}

	//spendings := make([][]int, 0,5)
	//
	//spendings = append(spendings, []int{200,100})
	//spendings = append(spendings, []int{25,10,45,60})
	//spendings = append(spendings, []int{5,15,35})
	//spendings = append(spendings, []int{95,10}, []int{50,25})

	spendings := fetch()

	for i, daily := range spendings {
		//fmt.Println(i+1, daily)
		var total int
		for _, spending := range daily {
			total += spending
		}
		fmt.Printf("Day %d: %d\n", i+1, total)
	}
}

func fetch() [][]int {
	//gets automatically converted to [][]int(nil)
	//because the function returns [][]int
	content := `200 100
25 10 45 60
5 15 35
95 10
50 25`
	lines := strings.Split(content,"\n")

	spendings := make([][]int, len(lines))


	for i, line := range lines {
		fmt.Printf("%d: %#v\n", i+1, line)

		fields := strings.Fields(line)

		spendings[i] = make([]int, len(fields))

		for j, field := range fields {
			fmt.Printf("\t%d: %#v\n",j+1,field)
			spending, _ := strconv.Atoi(field)
			spendings[i][j] = spending
		}
	}
	fmt.Printf("spendings: %v\n", spendings)
	return spendings
}
