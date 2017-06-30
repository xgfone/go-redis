package redis

import (
	"fmt"
)

func ExampleRedis_ClientGetName() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-clientgetname"

	r.ClientSetName(key)
	fmt.Println(r.ClientGetName())

	// Output:
	// test-clientgetname
}

func ExampleRedis_ClientList() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	fmt.Println(len(r.ClientList()) != 0)

	// Output:
	// true
}

func ExampleRedis_CommandCount() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	fmt.Println(r.CommandCount() != 0)

	// Output:
	// true
}

// func ExampleRedis_CommandGetKeys() {
// 	r := NewRedis("redis://127.0.0.1:6379/0", 1)
// 	defer r.Close()

// 	fmt.Println(r.CommandGetKeys())

// 	// Output:
// 	// true
// }

func ExampleRedis_ConfigGet() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	r.ConfigSet("save", "900 1 300 10")
	fmt.Println(r.ConfigGet("save"))

	// Output:
	// [save 900 1 300 10]
}

func ExampleRedis_ConfigResetStat() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	r.ConfigResetStat()
	fmt.Println()

	// Output:
	//
}

func ExampleRedis_ConfigRewrite() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	r.ConfigRewrite()
	fmt.Println()

	// Output:
	//
}

func ExampleRedis_DBSize() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	r.DBSize()
	fmt.Println()

	// Output:
	//
}

func ExampleRedis_Info() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	fmt.Println(len(r.Info()) != 0)

	// Output:
	// true
}

func ExampleRedis_LastSave() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	fmt.Println(r.LastSave() != 0)

	// Output:
	// true
}
