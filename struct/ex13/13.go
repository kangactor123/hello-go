package main

import "fmt"

type People struct {
	age int
	name string
}

type User struct {
	name string
	age int
	ID int
}

// 다른 구조체를 포함한 구조체
type VIPUser struct {
	userInfo User
	VIPLevel int
	price int
}

// 필드명을 생략하면 바로 User 필드에 접근이 가능
type NormalUser struct {
	User // 포함된 필드
	level int
}

type NonUser struct {
	User
	ID int
}

func main() {
	var johns People

	johns.age = 18
	johns.name = "johns"

	fmt.Println(johns.name, johns.age)

	// 구조체 선언과 초기화
	var dh People = People{13, "dh"}

	// 구조체 선언과 초기화
	var yj People = People{
		13,
		"yj",
	}

	// 일부 필드만 초기화
	var jw People = People{
		age: 31,
	}

	var u1 User = User{
		name: "u1",
		age: 22,
		ID: 1,
	}

	var u2 VIPUser = VIPUser{
		userInfo: User{
			name: "u2",
			age: 23,
			ID: 2,
		},
		price: 12200,
		VIPLevel: 3,
	}

	u3 := NormalUser{
		User{
			name: "u3",
			age: 24,
			ID: 4,
		},
		4,
	}

	u4 := NonUser{
		User{
			name: "u4",
			age: 25,
			ID: 5,
		},
		3,
	}

	// u4.ID // NonUser 의 필드
	// u4.User.ID // User 의 필드

}