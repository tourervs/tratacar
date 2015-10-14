package web

const (
        FB_INDIFFERENT  =  0
        FB_POSITIVE     =  1
        FB_NEGATIVE     =  2
        )

type User  struct {

    Id uint32

}

type FeedBack struct {

    FB_type    int
    Text       string
    Author_id  int

}

type Request struct {




}
