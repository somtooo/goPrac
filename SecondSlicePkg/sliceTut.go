package main

import "fmt"

type user struct {
	likes int
}

func main() {

	//dangerousSliceSituation()
	anotherDangerousSituation()

}



func dangerousSliceSituation() {

	users := make([]user,3)
	shareUser := &users[1]

	shareUser.likes++

	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

	users = append(users, user{})
	fmt.Println("\n")

	shareUser.likes++ // likes should increase from useres perspective

	for i := range users {
		fmt.Printf("User: %d, Likes: %d\n",i,users[i].likes)
	}

	//Very dangerous situation.
}

func anotherDangerousSituation() {
	names := make([]string,5)
	names[0] = "Annie"
	names[1] = "Betty"
	names[2] = "Charley"
	names[3] = "Doug"
	names[4] = "Edward"

	for i := range names {
		name := names[:2]
		fmt.Println("v[%s]\n", names[i], name)
	}
}