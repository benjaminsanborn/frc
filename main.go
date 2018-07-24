package main

import "fmt"
import "time"

func main() {
  //reference := "https://stevemorse.org/jcal/frules.htm"
  yearOne := time.Date(1792, time.September, 22, 0, 0, 0, 0, time.UTC)

  var months[12]string

  months[0] = "Vendémiaire (vintage)"
  months[1] = "Brumaire (mist)"
  months[2] = "Frimaire (frost)"
  months[3] = "Nivôse (snow)"
  months[4] = "Pluviôse (rain)"
  months[5] = "Ventôse (wind)"
  months[6] = "Germinal (seed-time)"
  months[7] = "Floréal (blossom)"
  months[8] = "Prairial (meadow)"
  months[9] = "Messidor (harvest)"
  months[10] = "Thermidor (heat)"
  months[11] = "Fructidor (fruit)"

  var celebrationDays[6]string

  celebrationDays[0] = "Fete de la vertu"
  celebrationDays[1] = "Fete du genie"
  celebrationDays[2] = "Fete du travail"
  celebrationDays[3] = "Fete de l'opinion"
  celebrationDays[4] = "Fete des recompenses"
  celebrationDays[5] = "Jour de la revolution"

  var dow[10]string

  dow[0] = "primidi"
  dow[1] = "duodi"
  dow[2] = "tridi"
  dow[3] = "quartidi"
  dow[4] = "quintidi"
  dow[5] = "sextidi"
  dow[6] = "septidi"
  dow[7] = "octidi"
  dow[8] = "nonidi"
  dow[9] = "decadi"

  p := fmt.Printf
  secondsPerDay := 86400.0

  for {
    now := time.Now()

    // Date Calculations
    year := now.Year() - yearOne.Year() + 1
    month :=

    // Time of Day Calculations
    seconds := (((now.Hour() * 60) + now.Minute()) * 60) + now.Second()
    pctDayPassed := 100000 * float64(seconds)/secondsPerDay
    decHour := (int(pctDayPassed)/10000 % 10) + 1
    decMinute := (int(pctDayPassed)/100 % 100)
    decSeconds := (int(pctDayPassed) % 100)
    tod := fmt.Sprintf("Hour: %v  Minute: %v Second: %v", decHour, decMinute, decSeconds)

    p("\r%v", tod)
    p("\r%v", year)
    time.Sleep(100 * time.Millisecond)
  }
}
