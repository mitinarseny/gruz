# gruz
Fast Golang HSE RUZ API Client

## Usage
```go
client := gruz.NewClient(http.DefaultClient)
lessons, err := client.GetSchedule(context.Backgound(), 123456, gruz.StudentPerson, time.Now(),time.Now().AddDate(0, 1, 0), gruz.RussianLanguage)
if err != nil {
    log.Fatalln(err)
}
for _, l := range lessons {
    fmt.Println(l.DiDiscipline)
}
```
