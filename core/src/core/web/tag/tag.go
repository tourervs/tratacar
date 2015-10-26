package tag

var single  =  []string{ "area",
                         "base",
                         "br",
                         "col",
                         "embed",
                         "hr",
                         "img",
                         "input",
                         "keygen",
                         "link",
                         "meta",
                         "param",
                         "source",
                         "track",
                         "wbr" }

func IsDouble ( tag string ) ( double bool ) {

    double = true

    for i:= range single {

        if tag == single[i] { double = false ; break }

    }

    return

}
