package osutils

import "path/filepath"
import "runtime"
import "os"
import "fmt"

const STATIC_DIRECTORY_NAME = "static"
const STATIC_JS_DIRECTORY_NAME = "js"
const STATIC_CSS_DIRECTORY_NAME = "css"


func StaticFilesFinder() (err error) {

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

    if err != nil {

        return  err

    }

    _, filename, _, _ := runtime.Caller(1) // I don't have any fucking understanding what those underscores means

    caller_dir_path         :=  filepath.Dir(filename)   // caller directory . 
    caller_parent_dir_path  :=  filepath.Dir(caller_dir) // caller parent directory .  for search under apps parent directory 

    caller_dir,         err_cd        :=  os.Open( caller_dir_path        + "/" + STATIC_DIRECTORY_NAME )
    caller_parent_dir,  err_cpd       :=  os.Open( caller_parent_dir_path + "/" + STATIC_DIRECTORY_NAME )

    defer caller_dir.Close()
    defer caller_parent_dir.Close()

    if err_cd != nil && err_cpd != nil {

        return  err

    }

    fmt.Printf("\n|filename_directory: %s|\n", filepath.Dir(filename))

    defer runtime_dir.Close()

    dir_content , err := runtime_dir.Readdirnames(-1)

    for i:= range dir_content {

        fmt.Printf("\n| %s |\n",dir_content[i])

    }

    if err != nil {
        return  err
    }


    return nil
}

