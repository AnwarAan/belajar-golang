package go_routine

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello")
}

func TestRunHelloWorld(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("done")
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("count", number)
}

func TestManyNumber(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(2 * time.Second)
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	// defer close(channel)

	go func() {
		time.Sleep(1 * time.Second)
		channel <- "anwar"
		fmt.Println("finish")
	}()

	data := <-channel

	fmt.Println(data)

	time.Sleep(1 * time.Second)

}

func GiveMeResponse(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "anwaraan"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)
	time.Sleep(1 * time.Second)

}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
	close(channel)
}

func TestBufferChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "mca"

	fmt.Println(cap(channel))
	fmt.Println(len(channel))
}

func TestTangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "loop" + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("get data", data)
	}
}

func TestSelecctMulipleChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("ch 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("ch 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("ch 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("ch 2", data)
			counter++
		default:
			fmt.Println("wait for")

		}
		if counter == 2 {
			break
		}
	}
}

func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("count", x)
}

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("count", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()

}
func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}
func TestReadWriteMutex(t *testing.T) {
	account := BankAccount{}
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println("account", account.GetBalance())
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("balance", account.GetBalance())
}

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transefer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("lock", user1.Name)
	user1.Change(amount)
	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("lock", user1.Name)
	user2.Change(amount)
	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}
func TestDeadLock(t *testing.T) {
	user1 := UserBalance{
		Name:    "person",
		Balance: 100000,
	}
	user2 := UserBalance{
		Name:    "human",
		Balance: 100000,
	}

	fmt.Println("user", user1.Name, "balance", user1.Balance)
	fmt.Println("user", user2.Name, "balance", user2.Balance)

	go Transefer(&user1, &user2, 1000)
	go Transefer(&user1, &user2, 1000)

	time.Sleep(1 * time.Second)
}

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)

	fmt.Println("hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
}
