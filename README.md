# Rester

![master build test result](https://api.travis-ci.org/Iceyer/rester.svg?branch=master)

Rester is an micro rest api request lib. The base design is from [restit](https://github.com/yookoala/restit).

## Dependencies

## Example

````golang

import github.com/Iceyer/rester

type User struct {
    Name  string
    Email string
}

func main() {
    apiURL = "https://api.iceyer.org/v0/test"
    restAPI := rester.Rest("User", apiURL)

    user := User{Name: "Iceyer", Email: "me@iceyer.net"}
    newUser := User{}
    req := restAPI.Create(&user).
        AddHeader("Access-Token", token).
        WithResponeAs(&newUser)

    result, err := req.Go()
    if nil != err {
        fmt.Println("Create User Failed")
    }
    fmt.Println("Request Status:", result.Status())
    fmt.Println("Create New User:", newUser)
}

````
