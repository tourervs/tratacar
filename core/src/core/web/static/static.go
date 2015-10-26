package static

import "core/web"

type ObjectStatic struct {

     web.Object
     App       string
     FilePath  string
     Type      string
     // not forget use "<link" for "css"  and "<script" for "js"

}


func Collect(path string) ( static []*ObjectStatic ) {

    return

}


//func ( file *File ) 
