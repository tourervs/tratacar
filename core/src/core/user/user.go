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
    Prop

}

type Car struct {

    Brand    string
    Model    string
    Capacity int
    Photo    int

}

type Prop struct {

     Smoke              bool
     FuckedSpeedLimits  bool
     LoudMusic          bool

}
