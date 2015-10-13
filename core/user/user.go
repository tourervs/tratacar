package user

type User struct {

    UType string // Passenger or Driver
    Radius uint8 // 
    Routes []Route

}

type Driver struct {

    Name     string
    Surname  string
    Age      string
    Mob      string
    Email    string
    Skype    string
    Car

}

type Car struct {

    Brand    string
    Model    string
    Capacity int
    Photo    int




}
