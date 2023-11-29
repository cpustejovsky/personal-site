package lifetogether

import (
	"errors"
	"fmt"
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

// CalculateNow takes Input and returns an Output and error based on the current time
func CalculateNow(in Input) (*Output, error) {
	return Calculate(time.Now(), in)
}

// Calculate takes Input and return a pointer to Output and error based on the provided time
func Calculate(t time.Time, in Input) (*Output, error) {
	err := validateInput(in)
	if validateInput(in) != nil {
		return nil, err
	}

	var out = Output{
		YourName:  in.YourName,
		OtherName: in.OtherName,
	}

	//Calculate duration since meeting
	md, err := calculateDayDuration(t, in.DateMet)
	if err != nil {
		return nil, fmt.Errorf("date met error %w", err)
	}
	out.MetDuration = md
	//Calculate percentage person and other have known each other
	yb, err := calculateDayDuration(t, in.YourBirthday)
	if err != nil {
		return nil, fmt.Errorf("your birthday error %w", err)
	}
	ob, err := calculateDayDuration(t, in.OtherBirthday)
	if err != nil {
		return nil, fmt.Errorf("other birthday error %w", err)
	}
	yourAlive := float64(yb)
	otherAlive := float64(ob)
	metDurationFloat := float64(md)
	out.YourPercentTogether = round(metDurationFloat/yourAlive*100, 2)
	out.OtherPercentTogether = round(metDurationFloat/otherAlive*100, 2)

	//Calculate for optional parameters
	dating, err := calculateDayDuration(t, *in.DateDating)
	if err != nil {
		out.DatingDuration = nil
	} else {
		out.DatingDuration = &dating
	}
	married, err := calculateDayDuration(t, *in.DateMarried)
	if err != nil {
		out.MarriedDuration = nil
	} else {
		out.MarriedDuration = &married
	}
	return &out, nil
}

func validateInput(in Input) error {
	//Validate Mandatory Parameters
	if in.YourName == "" {
		return errors.New("provide your name")
	}

	if in.OtherName == "" {
		return errors.New("provide other person's name")
	}

	// Validate that DateMet is greater than both YourBirthday and OtherBirthday
	if in.DateMet.Before(in.YourBirthday) || in.DateMet.Before(in.OtherBirthday) {
		return errors.New("the date you both met should be after you both were born")
	}

	return nil
}

func calculateDayDuration(t time.Time, dur time.Time) (int, error) {
	if dur.IsZero() {
		return -1, errors.New("duration is zero")
	}
	hoursDur := t.Sub(dur).Hours()
	days := math.Floor(hoursDur / 24.0)
	return int(days), nil
}

func round(num float64, decimals int) float64 {
	precision := math.Pow10(decimals)
	return math.Round(num*precision) / precision
}
