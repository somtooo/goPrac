package main

import "fmt"


type user struct {
	name string
	email string
}

type notifier interface {
	notify()
}

type duration int

type printer interface {
	print()
}

type admin struct {
	user
	level string
}

func (u user) print() {
	fmt.Printf("User name: %s\n", u.name)
}

func (d *duration) notify() {
	fmt.Println("Send notification in", *d)
}

func (u *user) notify() {
	fmt.Printf("Sending User Email to %s<%s>\n", u.name,  u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func sendNotification(n notifier) {
	n.notify();
}

func main() {


	// s := 3;
	// duration(s).notify()
		// Values of type user can be used to call methods
	// declared with both value and pointer receivers.
	bill := user{"Bill", "bill@email.com"}
	bill.changeEmail("bill@hotmail.com")
	// sendNotification(bill)

	// Pointers of type user can also be used to call methods
	// declared with both value and pointer receiver.
	joan := &user{"Joan", "joan@email.com"}
	joan.changeEmail("joan@hotmail.com")
	joan.notify()

	fmt.Println("\n Code smell")

	users := []user{
		{"ed","ed@email.com"},
		{"erick","erik@email.com"},
	}

	for _, u := range users {
		u.changeEmail("changed@email.com")
	}

	fmt.Println(users)
	u := user{"Bill","bill@gmail.com"}

	agg := []printer{u,&u}
	fmt.Println(agg)

	ad := admin{
		user{"John","john@yahoo.com"},
		"super",
	}

	ad.notify();

}