package structs

import "time"

type BodyRequest struct {
	BirthdayDate string `json:"birthday"`
}

type Birthday struct {
	Born                        time.Time
	HoursRoundDecimalBirthday   []time.Time
	MinutesRoundDecimalBirthday []time.Time
	SecondsRoundDecimalBirthday []time.Time
}

func (b *Birthday) getBirthday() string {
	return b.Born.Format(time.RFC3339)
}

func (b *Birthday) setBirthday(date string) (err error) {
	b.Born, err = time.Parse(
		time.RFC3339,
		date)
	return err
}

func (b *Birthday) DecimalBirthday() {
	for i := 1; i < 7; i++ {
		b.HoursRoundDecimalBirthday = append(b.HoursRoundDecimalBirthday, b.Born.Add(time.Hour*time.Duration(i*100000)))
	}
	for i := 1; i < 42; i++ {
		b.MinutesRoundDecimalBirthday = append(b.MinutesRoundDecimalBirthday, b.Born.Add(time.Minute*time.Duration(i*1000000)))
	}
	for i := 1; i < 27; i++ {
		b.SecondsRoundDecimalBirthday = append(b.SecondsRoundDecimalBirthday, b.Born.Add(time.Second*time.Duration(i*100000000)))
	}
}
