package main

import "fmt"

func main() {
    // Create an empty array to store user names
    var users []string

    // Add users to the array
    users = append(users, "Alice")
    users = append(users, "Bob")
    users = append(users, "Charlie")

    // Print the users
    fmt.Println("Users:", users)
}
