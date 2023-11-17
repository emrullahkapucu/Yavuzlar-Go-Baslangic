package main
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)
func main() {
    maxLoginAttempts := 5
    userType := -1
    for maxLoginAttempts > 0 {
        fmt.Println("Lütfen giriş türünüzü seçin:")
        fmt.Println("0 - Admin Girişi")
        fmt.Println("1 - Öğrenci Girişi")

        var username, password string
        fmt.Scanln(&userType)

        if userType == 0 {
            fmt.Print("Admin kullanıcı adı: ")
            fmt.Scanln(&username)
            fmt.Print("Admin şifre: ")
            fmt.Scanln(&password)

            if isAdminLoginValid(username, password) {
                logLogin("admin", "Başarılı")
                adminMenu()
                break
            } else {
                logLogin("admin", "Başarısız")
                maxLoginAttempts--
                fmt.Printf("Hatalı giriş. Kalan deneme hakkı: %d\n", maxLoginAttempts)
            }
        } else if userType == 1 {
            fmt.Print("Öğrenci kullanıcı adı: ")
            fmt.Scanln(&username)
            fmt.Print("Öğrenci şifre: ")
            fmt.Scanln(&password)
            if isStudentLoginValid(username, password) {
                logLogin("student", "Başarılı")
                studentMenu()
                break
            } else {
                logLogin("student", "Başarısız")
                maxLoginAttempts--
                fmt.Printf("Hatalı giriş. Kalan deneme hakkı: %d\n", maxLoginAttempts)
            }
        } else {
            fmt.Println(" Program kapanıyor.")
            break
        }
    }
    if maxLoginAttempts == 0 {
        fmt.Println(" Program kapanıyor.")
    }
}
func isAdminLoginValid(username, password string) bool {
    return username == "admin" && password == "admin"
}
func isStudentLoginValid(username, password string) bool {
    return username == "student" && password == "student"
}
func adminMenu() {
    for {
        fmt.Println("Lütfen bir işlem seçin:")
        fmt.Println("0 - Logları Görüntüle")
        fmt.Println("1 - Çıkış Yap")

        var choice int
        fmt.Scanln(&choice)

        if choice == 0 {
            displayLogs()
        } else if choice == 1 {
            fmt.Println("Çıkış yapılıyor.")
            break
        } else {
            fmt.Println("Geçersiz işlem .Lütfen tekrar deneyin.")
        }
    }
}
func studentMenu() {
    for {
        fmt.Println("Lütfen bir işlem seçin:")
        fmt.Println("0 - Çıkış")
        var choice int
        fmt.Scanln(&choice)
        switch choice {
        case 0:
            fmt.Println("Çıkış yapılıyor.")
            return
        default:
            fmt.Println("Geçersiz işlem . Lütfen tekrar deneyin.")
        }
    }
}
func displayLogs() {
    file, err := os.Open("logs.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
func logLogin(username, status string) {
    currentDateTime := time.Now().Format("2006-01-02 15:04:05")
    logEntry := fmt.Sprintf("Kullanıcı Adı: %s\nGiriş Tarihi ve Saati: %s\nGiriş Durumu: %s\n\n", username, currentDateTime, status)
    file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    _, err = file.WriteString(logEntry)
    if err != nil {
        log.Fatal(err)
    }
}
