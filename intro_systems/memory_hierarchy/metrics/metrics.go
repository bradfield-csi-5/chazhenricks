package metrics

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

type UserId int
type UserMap map[UserId]*User

type Address struct {
	fullAddress string
	zip         int
}

type DollarAmount struct {
	dollars, cents uint64
}

type Payment struct {
	amount DollarAmount
	time   time.Time
}

type User struct {
	id       UserId
	name     string
	age      int
	address  Address
	payments []Payment
}

// new type - all we care about for these calculations is age and payment averages
type Users struct {
	ages     []float64
	payments []uint64
}

func AverageAge2(users Users) float64{
	average, count := 0.0, 0.0
	for _, a := range users.ages {
		count += 1
		average += (a - average) / count
	}
	return average
}
func AverageAge(users UserMap) float64 {
	average, count := 0.0, 0.0
	for _, u := range users {
		count += 1
		average += (float64(u.age) - average) / count
	}
	return average
}

func AveragePaymentAmount(users UserMap) float64 {
	average, count := 0.0, 0.0
	for _, u := range users {
		for _, p := range u.payments {
			count += 1
			amount := float64(p.amount.dollars) + float64(p.amount.cents)/100
			average += (amount - average) / count
		}
	}
	return average
}
func AveragePaymentAmount2(users Users) float64 {
	average, count := 0.0, 0.0
		for _, p := range users.payments {
			count += 1
			amount := float64(p / 100)
			average += (amount - average) / count
		}
	return average
}

// Compute the standard deviation of payment amounts
// per initial benchmarks, this takes the most time, so focus attention here to start.
// variance == how spread out data is: [5,5,5] variance is 0. [5,10,15,20] variance is higher
// variance = average(x^2 for x in data) - average(x for x in data)^2
func StdDevPaymentAmount(users UserMap) float64 {
	mean := AveragePaymentAmount(users)
	squaredDiffs, count := 0.0, 0.0
	for _, u := range users {
		for _, p := range u.payments {
			count += 1
			amount := float64(p.amount.dollars) + float64(p.amount.cents)/100
			diff := amount - mean
			squaredDiffs += diff * diff
		}
	}
	return math.Sqrt(squaredDiffs / count)
}
func StdDevPaymentAmount2(users Users) float64 {
	mean := AveragePaymentAmount2(users)
	squaredDiffs, count := 0.0, 0.0
		for _, p := range users.payments {
			count += 1
			amount := float64(p / 100)
			diff := amount - mean
			squaredDiffs += diff * diff
		}
	return math.Sqrt(squaredDiffs / count)
}

func LoadData2() Users {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Unable to read users.csv", err)
	}
	reader := csv.NewReader(f)
	userLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse users.csv as csv", err)
	}

	ages_read := make([]float64, 0)
	for _, line := range userLines {
		age, _ := strconv.Atoi(line[2])
		ages_read = append(ages_read, float64(age))
	}


	f, err = os.Open("payments.csv")
	if err != nil {
		log.Fatalln("Unable to read payments.csv", err)
	}
	reader = csv.NewReader(f)
	paymentLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse payments.csv as csv", err)
	}

	payments := make([]uint64,0)
	for _, line := range paymentLines {
		paymentCents, _ := strconv.Atoi(line[0])
		payments = append(payments, uint64(paymentCents))
	}

	users := Users{ages: ages_read, payments: payments}
	return users

}

func LoadData() UserMap {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Unable to read users.csv", err)
	}
	reader := csv.NewReader(f)
	userLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse users.csv as csv", err)
	}

  fmt.Printf("USERLINES: %d\n", len(userLines))
	users := make(UserMap, len(userLines))
	for _, line := range userLines {
		id, _ := strconv.Atoi(line[0])
		name := line[1]
		age, _ := strconv.Atoi(line[2])
		address := line[3]
		zip, _ := strconv.Atoi(line[3])
		users[UserId(id)] = &User{UserId(id), name, age, Address{address, zip}, []Payment{}}
	}

	f, err = os.Open("payments.csv")
	if err != nil {
		log.Fatalln("Unable to read payments.csv", err)
	}
	reader = csv.NewReader(f)
	paymentLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse payments.csv as csv", err)
	}

	for _, line := range paymentLines {
		userId, _ := strconv.Atoi(line[2])
		paymentCents, _ := strconv.Atoi(line[0])
		datetime, _ := time.Parse(time.RFC3339, line[1])
		users[UserId(userId)].payments = append(users[UserId(userId)].payments, Payment{
			DollarAmount{uint64(paymentCents / 100), uint64(paymentCents % 100)},
			datetime,
		})
	}

  fmt.Printf("I AM ORIG USERS. AGE COUNT IS %d\n", len(users))
	return users
}
