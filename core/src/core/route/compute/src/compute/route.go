package compute

type Route struct {

    Price        int
    Value        string
    CheckPoints  []CheckPoint
    TimeTables   []TimeTable

}

type CheckPoint struct {

    Lat      float32
    Lon      float32
    Country  string
    Region   string
    City     string
    Area     string
    Street   string
    House    string
    Source   bool
    Dest     bool

}



type TimeTable struct {



}
