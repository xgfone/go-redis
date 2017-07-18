package redis

import (
	"fmt"
)

func ExampleRedis_ClientGetName() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-clientgetname"

	r.ClientSetName(key)
	v, _ := r.ClientGetName()
	fmt.Println(v)

	// Output:
	// test-clientgetname
}

func ExampleRedis_ClientList() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	v, _ := r.ClientList()
	fmt.Println(len(v) != 0)

	// Output:
	// true
}

func ExampleRedis_ClientReply() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	if err := r.ClientReply("ON"); err == nil {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleRedis_CommandCount() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	if v, err := r.CommandCount(); err == nil && v != 0 {
		fmt.Println(v != 0)
	}

	// Output:
	// true
}

func ExampleRedis_CommandGetKeys() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	v, _ := r.CommandGetKeys("MSET", "a", "b", "c", "d", "e", "f")
	fmt.Println(v)
	v, _ = r.CommandGetKeys("EVAL", "not consulted", 3, "key1", "key2",
		"key3", "arg1", "arg2", "arg3", "argN")
	fmt.Println(v)
	v, _ = r.CommandGetKeys("SORT", "mylist", "ALPHA", "STORE", "outlist")
	fmt.Println(v)

	// Output:
	// [a c e]
	// [key1 key2 key3]
	// [mylist outlist]
}

func ExampleRedis_ConfigGet() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	r.ConfigSet("save", "900 1 300 10")
	vs, _ := r.ConfigGet("save")
	if len(vs) == 2 { // For Windows Redis: ["save", "jd 900 jd 300"]
		if vs[0] == "save" && vs[1] == "jd 900 jd 300" {
			fmt.Println("OK")
		} else {
			fmt.Println("ERR")
		}
	} else { // For Linux Redis: ["save", "900", "1", "300", "10"]
		if vs[0] == "save" && vs[1] == "900" && vs[2] == "1" &&
			vs[3] == "300" && vs[4] == "10" {
			fmt.Println("OK")
		} else {
			fmt.Println("ERR")
		}
	}

	// Output:
	// OK
}

func ExampleRedis_ConfigResetStat() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	if err := r.ConfigResetStat(); err == nil {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleRedis_ConfigRewrite() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	if err := r.ConfigRewrite(); err == nil {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleRedis_DBSize() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	if _, err := r.DBSize(); err == nil {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleRedis_Info() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	v, _ := r.Info()
	fmt.Println(len(v) != 0)

	// Output:
	// true
}

func ExampleRedis_LastSave() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	v, _ := r.LastSave()
	fmt.Println(v != 0)

	// Output:
	// true
}

func ExampleRedis_Save() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	if err := r.Save(); err == nil {
		fmt.Println("OK")
	}

	// Output:
	// OK
}

func ExampleRedis_SlaveOf() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	if err := r.SlaveOf("no", "one"); err == nil {
		fmt.Println("no one")
	}

	if err := r.SlaveOf("not no", "not one"); err != nil {
		fmt.Println("Argument Error")
	}

	// Output:
	// no one
	// Argument Error
}

func ExampleRedis_Time() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	s, m, _ := r.Time()
	fmt.Println(s != 0, m != 0)

	// Output:
	// true true
}
