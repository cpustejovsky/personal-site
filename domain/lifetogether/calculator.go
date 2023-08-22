package lifetogether

import (
	"errors"
	"math"
	"time"
)

type Input struct {
	YourName  string
	OtherName string
	//TODO: better variable names than YourBirthday and OtherBirthday
	YourBirthday  time.Time
	OtherBirthday time.Time
	DateMet       time.Time
	DateDating    *time.Time
	DateMarried   *time.Time
}

type Output struct {
	YourName             string
	OtherName            string
	YourPercentTogether  float64
	OtherPercentTogether float64
	MetDuration          int
	DatingDuration       *int
	MarriedDuration      *int
}

//TODO: create error structs to store information

func CalculateNow(in Input) (*Output, error) {
	return Calculate(time.Now(), in)
}

// Calculate takes Input and return a pointer to Output along with an error
func Calculate(t time.Time, in Input) (*Output, error) {
	//Validate Mandatory Parameters
	if in.YourName == "" {
		return nil, errors.New("provide your name")
	}

	if in.OtherName == "" {
		return nil, errors.New("provide other person's name")
	}

	// Validate that DateMet is greater than both YourBirthday and OtherBirthday
	if in.DateMet.Before(in.YourBirthday) || in.DateMet.Before(in.OtherBirthday) {
		return nil, errors.New("the date you both met should be after you both were born")
	}

	var out = Output{
		YourName:  in.YourName,
		OtherName: in.OtherName,
	}

	//Calculate duration since meeting
	out.MetDuration = CalculateDayDuration(t, in.DateMet)

	//Calculate percentage person and other have known each other
	metDurationFloat := float64(out.MetDuration)
	yourAlive := float64(CalculateDayDuration(t, in.YourBirthday))
	otherAlive := float64(CalculateDayDuration(t, in.OtherBirthday))
	out.YourPercentTogether = round(metDurationFloat/yourAlive*100, 2)
	out.OtherPercentTogether = round(metDurationFloat/otherAlive*100, 2)

	//Calculate for optional parameters
	if in.DateDating != nil {
		d := CalculateDayDuration(t, *in.DateDating)
		out.DatingDuration = &d
	}
	if in.DateMarried != nil {
		d := CalculateDayDuration(t, *in.DateMarried)
		out.MarriedDuration = &d
	}
	return &out, nil
}

func round(num float64, decimals int) float64 {
	precision := math.Pow10(decimals)
	return math.Round(num*precision) / precision
}

func CalculateDayDuration(t time.Time, dur time.Time) int {
	hoursDur := t.Sub(dur).Hours()
	metdays := math.Floor(hoursDur / 24.0)
	return int(metdays)
}
