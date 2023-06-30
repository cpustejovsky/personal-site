package lifetogether_test

import (
	"github.com/cpustejovsky/personal-site/domain/lifetogether"
	"reflect"
	"testing"
	"time"
)

func TestCalculate(t *testing.T) {
	testTime := time.Date(2023, 6, 19, 0, 0, 0, 0, time.UTC)
	//Input Variables
	ccp := "Charles Pustejovsky"
	cep := "Catherine Pustejovsky"
	dating := time.Date(2014, 3, 10, 0, 0, 0, 0, time.UTC)
	married := time.Date(2018, 1, 6, 0, 0, 0, 0, time.UTC)
	yourBirthday := time.Date(1992, 12, 18, 0, 0, 0, 0, time.UTC)
	otherBirthday := time.Date(1994, 10, 12, 0, 0, 0, 0, time.UTC)
	dateMet := time.Date(2014, 2, 19, 0, 0, 0, 0, time.UTC)

	//Output Variables
	datingDuration := 3388
	marriedDuration := 1990
	metDuration := 3407
	yourPercentTogether := 30.58
	otherPercentTogether := 32.52

	want := &lifetogether.Output{
		YourName:              ccp,
		OtherName:             cep,
		YourPercentTogether:   yourPercentTogether,
		OtherPerecentTogether: otherPercentTogether,
		MetDuration:           metDuration,
		DatingDuration:        &datingDuration,
		MarriedDuration:       &marriedDuration,
	}

	t.Run("Correct Input returns expected output and nil error", func(t *testing.T) {
		in := lifetogether.Input{
			YourName:      ccp,
			OtherName:     cep,
			YourBirthday:  yourBirthday,
			OtherBirthday: otherBirthday,
			DateMet:       dateMet,
			DateDating:    &dating,
			DateMarried:   &married,
		}
		got, err := lifetogether.Calculate(testTime, in)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("\nwanted:\t%+v\ngot:\t%+v", want, got)
		}
	})
	t.Run("No name for yourself returns error", func(t *testing.T) {
		in := lifetogether.Input{
			YourName:      "",
			OtherName:     cep,
			YourBirthday:  yourBirthday,
			OtherBirthday: otherBirthday,
			DateMet:       dateMet,
			DateDating:    &dating,
			DateMarried:   &married,
		}
		in.YourName = ""
		_, err := lifetogether.Calculate(testTime, in)
		if err == nil {
			t.Fatal("expected error")
		}
	})
	t.Run("No name for yourself returns error", func(t *testing.T) {
		in := lifetogether.Input{
			YourName:      ccp,
			OtherName:     "",
			YourBirthday:  yourBirthday,
			OtherBirthday: otherBirthday,
			DateMet:       dateMet,
			DateDating:    &dating,
			DateMarried:   &married,
		}
		_, err := lifetogether.Calculate(testTime, in)
		if err == nil {
			t.Fatal("expected error")
		}
	})
	t.Run("DateMet before BirthDays returns error", func(t *testing.T) {
		dm := time.Date(1980, 2, 19, 0, 0, 0, 0, time.UTC)
		in := lifetogether.Input{
			YourName:      ccp,
			OtherName:     cep,
			YourBirthday:  yourBirthday,
			OtherBirthday: otherBirthday,
			DateMet:       dm,
			DateDating:    &dating,
			DateMarried:   &married,
		}
		_, err := lifetogether.Calculate(testTime, in)
		if err == nil {
			t.Fatal("expected error")
		}
	})
}
