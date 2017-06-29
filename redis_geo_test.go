package redis

import "fmt"

func ExampleRedis_GeoRadius() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-georadius"
	r.Del(key)

	fmt.Println(r.GeoAdd(key, "13.361389", "38.115556", "Palermo",
		"15.087269", "37.502669", "Catania"))
	fmt.Println(r.GeoDist(key, "Palermo", "Catania"))
	fmt.Println(r.GeoDist(key, "Palermo", "Catania", "km"))
	fmt.Println(r.GeoDist(key, "Palermo", "Catania", "mi"))
	fmt.Println(r.GeoDist(key, "foo", "bar"))

	fmt.Println(r.GeoRadius(key, 15, 37, 100, "km"))
	fmt.Println(r.GeoRadius(key, 15, 37, 200, "km"))
	fmt.Println(r.GeoRadius(key, 15, 37, 200, "km", "WITHDIST"))

	// Output:
	// 2
	// 166274.1516
	// 166.2742
	// 103.3182
	// -1
	// [Catania]
	// [Palermo Catania]
	// [[Palermo 190.4424] [Catania 56.4413]]
}

func ExampleRedis_GeoHash() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-geohash"
	r.Del(key)

	fmt.Println(r.GeoAdd(key, "13.361389", "38.115556", "Palermo",
		"15.087269", "37.502669", "Catania"))
	fmt.Println(r.GeoHash(key, "Palermo", "Catania"))

	// Output:
	// 2
	// [sqc8b49rny0 sqdtr74hyu0]
}

func ExampleRedis_GeoPos() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-geopos"
	r.Del(key)

	fmt.Println(r.GeoAdd(key, "13.361389", "38.115556", "Palermo",
		"15.087269", "37.502669", "Catania"))
	vs := r.GeoPos(key, "Palermo", "Catania", "NonExisting")
	fmt.Println(vs[0][0][:17])
	fmt.Println(vs[0][1][:17])
	fmt.Println(vs[1][0][:17])
	fmt.Println(vs[1][1][:17])

	// Output:
	// 2
	// 13.36138933897018
	// 38.11555639549629
	// 15.08726745843887
	// 37.50266842333162
}

func ExampleRedis_GeoRadiusByMember() {
	r := NewRedis("redis://127.0.0.1:6379/0", 1)
	defer r.Close()

	key := "test-georadiusbymember"
	r.Del(key)

	r.GeoAdd(key, "13.583333", "37.316667", "Agrigento")
	r.GeoAdd(key, "13.361389", "38.115556", "Palermo",
		"15.087269", "37.502669", "Catania")

	fmt.Println(r.GeoRadiusByMember(key, "Agrigento", 100, "km"))

	// Output:
	// [Agrigento Palermo]
}
