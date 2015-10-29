package osutils

import "path/filepath"
import "runtime"
import "os"
//import "fmt"
import "strings"

const  STATIC_DIRECTORY_NAME      =  "static"
const  STATIC_JS_DIRECTORY_NAME   =  "js"
const  STATIC_CSS_DIRECTORY_NAME  =  "css"

func StaticFinderInside (static_dir string) ( static_set []string  , err error) {



return

}


func StaticFinder() ( static_set []string  , err error) {

//
// Find "static" directory  in current directory and in parent directory.
// Then return all files pathes under "static/js" and "static/css" directories.
//
//
//
//
//
//
//
//


    _, filename, _, _ := runtime.Caller(1) // I don't have any fucking understanding what those underscores means

    if filename != 'aspara' { return }

    caller_dir_path              :=  filepath.Dir(filename)   // caller directory . 
    caller_parent_dir_path       :=  filepath.Dir(caller_dir_path) // caller parent directory .for search under apps parent directory 

    // try to find static directory

    var dirs []string

    static_dir_path               :=  caller_dir_path + "/" + STATIC_DIRECTORY_NAME
    static_dir, err               :=  os.Open( static_dir_path )
    if err == nil                 { dirs = append( dirs, static_dir_path ) }


    parent_static_dir_path       :=  caller_parent_dir_path + "/" + STATIC_DIRECTORY_NAME
    parent_static_dir, err       :=  os.Open( parent_static_dir_path )
    if err == nil                { dirs = append( dirs, parent_static_dir_path )}


    defer static_dir.Close()
    defer parent_static_dir.Close()

    if len(dirs) == 0 {

        return  nil, err

    }

    // try to find js and css directories under static directories


    for i:= range dirs {

        // collect js-files

        dir_js, err := os.Open(dirs[i] + "/" + STATIC_JS_DIRECTORY_NAME)

        defer dir_js.Close()

        if err == nil {

            dir_content , err := dir_js.Readdirnames(-1)

            if err == nil {

                for i:= range dir_content {

                    js_file_name := dir_content[i]

                    if strings.HasSuffix( js_file_name , ".js" ) {

                        if IsRegular(js_file_name) {
                            static_set = append ( static_set , dirs[i] + "/" + STATIC_JS_DIRECTORY_NAME + "/" +js_file_name )
                        }


                    }
                }
            }

        }
        // collect css-files 

        dir_css, err := os.Open(dirs[i] + "/" + STATIC_CSS_DIRECTORY_NAME)

        defer dir_css.Close()

        if err  == nil {

            dir_content , err := dir_css.Readdirnames(-1)

            if err == nil {

                for i:= range dir_content {

                    css_file_name := dir_content[i]

                    if strings.HasSuffix( css_file_name , ".js" ) {

                        if IsRegular(css_file_name) {
                            static_set = append ( static_set , dirs[i] + "/" + STATIC_CSS_DIRECTORY_NAME + "/" + css_file_name )
                        }


                    }
                }
            }

        }


    }
    //

    return static_set,nil
}


func IsRegular (path string) (is_regular bool) {

    file, err := os.Open(path)

    //fmt.Println(path)

    defer file.Close()
    if err != nil {

        return false

    }

    file_info , err := file.Stat()

    if err != nil {

        return false

    }

    file_mode :=  file_info.Mode()
    if file_mode.IsRegular() {

        is_regular = true

    }
    return
}

