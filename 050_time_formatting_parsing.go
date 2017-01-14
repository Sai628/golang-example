package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    t := time.Now()    
    p(t.Format(time.RFC3339))  // 2017-01-14T16:10:00+08:00

    t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
    p(t1)  // 2012-11-01 22:08:41 +0000 +0000

    p(t.Format("3:04PM"))  // 4:10PM
    p(t.Format("Mon Jan _2 15:04:05 2006"))  // Sat Jan 14 16:10:00 2017
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))  // 2017-01-14T16:10:00.527113+08:00
    
    form := "3 04 PM"
    t2, _ := time.Parse(form, "8 41 PM")
    p(t2)  // // 0000-01-01 20:41:00 +0000 UTC

    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())  // 2017-01-14T16:10:00-00:00

    ansic := "Mon Jan _2 15:04:05 2006"
    _, e := time.Parse(ansic, "8:41PM")
    p(e)  // parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
}
