# Rester

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
    restAPI := rester.Rest("User", "https://api.example.org/v1/users")

    user := User{Name: "Iceyer", Email: "me@iceyer.net"}
    newUser := User{}
    req := restAPI.Create(&user).
        AddHeader("Access-Token", token).
        WithPayloadAs(&newUser)

    result, err := req.Go()
    if nil != err {
        fmt.Println("Create User Failed")
    }
    fmt.Println("Create New User:", newUser)
}

````
