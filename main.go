package main

import (
    "fmt"
    "unicode"
)

func validatePasswordLength(password string) bool {
    if len(password) >= 6 {
        return true
    }

    return false
}

func validatePasswordCharacter(password string) bool {
    var (
        hasUppercase = false
        hasLowercase = false
        hasNumber = false
        hasSpecialChar = false
    )

    for _, c := range password {
        switch {
        case unicode.IsNumber(c):
            hasNumber = true
        case unicode.IsUpper(c):
            hasUppercase = true
        case unicode.IsLower(c):
            hasLowercase = true
        case unicode.IsPunct(c) || unicode.IsSymbol(c):
            hasSpecialChar = true
        }
    }

    return hasUppercase && hasLowercase && hasNumber && hasSpecialChar
}

func validateInputedPassword(password string) bool {
    if validatePasswordLength(password) {
        return validatePasswordCharacter(password)
    }

    return false
}

func main() {
    fmt.Println(validateInputedPassword("MyPassword123$$"))
}