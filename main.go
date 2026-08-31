package main

import "fmt";

func getShard(userId int, numberOfShards int) int {
	return userId % numberOfShards;
}

func main() {
	userId:= 3;
	numberOfShards:= 4;

	shard := getShard(userId, numberOfShards);

	fmt.Println("Shard: ", shard);
}
