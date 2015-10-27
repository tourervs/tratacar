package web

// import "fmt" // just for debug

import "core/web/tag"

type HtmlReturner func() string

type Object struct {

    Slug                                 string
    Name                                 string
    Id                                   string
    Value                                string
    Content                              string // All element content
    //
    //SelfOpenTag        string
    //SelfCloseTag       string
    DoubleTag                            bool
    //
    //SelfDelimTagOpen    string
    //SelfDelimTagClose   string
    //
    Parent                                *Object
    Params                                []string
    //
    templateTag                           string
    //
    Classes                               []string
    Style                                 string
    ChildsCompilationFunctions            []func()( func () *Object )

}

func (o *Object ) Print () {


}

func (o *Object ) Compile ()  ( func ()(*Object) ) { // need to add returning  second argument . *Object

    var child_content string
    //fmt.Printf("\nRunning compile : %s Childs_count :%d \n", o.Name, len(o.ChildsCompilationFunctions))

    for chi_num := range o.ChildsCompilationFunctions {
        // Collect childs html code
        obj_compile        :=  o.ChildsCompilationFunctions[chi_num]()
        child_content      =   child_content+obj_compile().Content+"\n"
        //fmt.Printf("childs content len %s",child_content)

    }

    var params        string
    for prm_id := range o.Params {
        // Append params to html code
        params+=o.Params[prm_id]+" "
    }
    //fmt.Printf("\nchild_content |%s|\n",child_content)

    return func() *Object {


        //fmt.Printf("\n returned function is started %s\n",o.Name)

        if tag.IsDouble(o.Name) == false  {

            var value string

            if o.Value != "" { value="value="+o.Value } else { value="" }

            o.Content = "<" + o.Name + " " + params + " " + value + " />"

        } else {

            var new_line string
            var space    string

            if child_content != "" { new_line="\n" }
            if params        != "" { space=" " }

            o.Content = "<" + o.Name + space + params + space +">" + new_line + o.Value + child_content + new_line + "</" + o.Name + ">"
        }
        return o
    }
}

func Append ( parent *Object, child *Object ) () {

    parent.ChildsCompilationFunctions =   append( parent.ChildsCompilationFunctions, child.Compile )
    //fmt.Printf("\nAppend: parent: %s child %s\n",parent.Name,child.Name)

}

func AppendToChild ( parent *Object, existing_parent_child_name string  ,child *Object ) () {
  //Example:  parent: html-obj ; existing_parent_child_name: "body" ; child: div-obj



}

func ChildDiveSearch ( parent *Object, existing_parent_child_name string )( co *Object ) {

    childs := parent.ChildsCompilationFunctions

    for i:= range parent.ChildsCompilationFunctions {

        if childs[i]().Name == existing_parent_child_name {


        }

    }
}
