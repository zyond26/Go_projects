package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)


// Lập trình bài toán đường hẹp (xe chỉ có thể đi theo 1 chiều cùng 1 lúc - không đủ để tránh nhau).
// Dùng 2 threads cho 2 gác cổng ở 2 đầu đường. Mỗi gác cổng có nhiệm vụ cho xe đi qua hoặc dừng lại theo các yêu cầu sau:
// 1) Nếu đường trống và đang có xe đợi, báo bên kia đóng không cho xe qua
// 2) Cho maximum n xe chạy qua (n do người dùng nhập). Khi đạt max thì dừng không cho xe qua tiếp và báo cho bên kia để cho
// xe chạy ngược lại. Nếu chưa đạt max nhưng không còn xe đợi, cũng thực hiện tương tự
// Dùng 2 threads để tạo thêm xe vào hai đầu, các xe xuất hiện theo thời gian random 1-100 giây; khởi động ban đầu mỗi bên
// có 20 xe đợi; ưu tiên cổng trái chạy trước.
// Giả định thời gian để giải phóng con đường là 3 giây (từ lúc báo bên kia để mở đường).
// Yêu cầu file code, video, thuyết minh; dùng Python, C++ hoặc Java.

// translate
// given 2 guards, at a time, either guard can work, a guard can only do max N tasks at a time in max 3 seconds.
// after done, notify the other guard to start
// the number of tasks for each guard is infinite


type guard struct {
	name                    string
	turn                    int
	maxSecondEachTurnCanRun int
	myPhone                 chan bool
	peerPhone               chan bool
	tasks                   chan string
}

func newGuard(name string, myPhone, peerPhone chan bool, tasks chan string) *guard {
	return &guard{name: name, myPhone: myPhone, peerPhone: peerPhone, tasks: tasks, maxSecondEachTurnCanRun: 3}
}

// work can only be called one at a time, not in parallel because turn is not guarded
func (g *guard) work(parallel int) {
	for {
		<-g.myPhone
		g.doMaxNTasks(parallel)
		g.peerPhone <- true
	}
}

func (g *guard) doMaxNTasks(n int) {
	g.turn++
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(g.maxSecondEachTurnCanRun)*time.Second)
	defer cancel()
	for i := 0; i < n; i++ {
		select {
		case task := <-g.tasks:
			fmt.Printf("%s's %dth turn: task %dth: %s\n", g.name, g.turn, i, task)
		case <-ctx.Done():
			return
		}
	}
}

func continuouslyGenerateTasks(id string, tasks chan string) {
	for {
		tasks <- fmt.Sprintf("%s generated task %d", id, time.Now().Unix())
		waitTime := 1 + rand.Intn(3) // to be changed to 100
		time.Sleep(time.Duration(waitTime) * time.Second)
	}
}

func main() {
	fmt.Print("Enter max number of people each guard can allow: ")
	var input string
	fmt.Scanln(&input)
	parallel, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Fatalf("cannot parse input %s to number. please check again. err %v", input, err)
	}

	rand.Seed(time.Now().Unix())
	leftPhone, rightPhone := make(chan bool, 1), make(chan bool, 1)
	leftTasks, rightTasks := make(chan string, 20), make(chan string, 20)
	leftGuard, rightGuard := newGuard("left", leftPhone, rightPhone, leftTasks), newGuard("right", rightPhone, leftPhone, rightTasks)

	for i := 0; i < 20; i++ {
		leftTasks <- fmt.Sprintf("LEFT initial task %d", i)
		rightTasks <- fmt.Sprintf("RIGHT initial task %d", i)
	}
	leftPhone <- true

	go leftGuard.work(int(parallel))
	go rightGuard.work(int(parallel))
	go continuouslyGenerateTasks("LEFT", leftTasks)
	go continuouslyGenerateTasks("RIGHT", rightTasks)
	select {}
}