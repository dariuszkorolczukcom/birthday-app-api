package structs

import "time"

type BodyRequest struct {
	BirthdayDate string `json:"birthday"`
}

type Birthday struct {
	Born                        time.Time   `json:"birthdayDate"`
	HoursRoundDecimalBirthday   []time.Time `json:"hoursDecimal"`
	MinutesRoundDecimalBirthday []time.Time `json:"minutesDecimal"`
	SecondsRoundDecimalBirthday []time.Time `json:"secondsDecimal"`
}

func (b *Birthday) GetBirthday() string {
	return b.Born.Format(time.RFC3339)
}

func (b *Birthday) SetBirthday(date string) (err error) {
	b.Born, err = time.Parse(
		time.RFC3339,
		date)
	return err
}

func (b *Birthday) CountHoursRoundDecimalBirthday() {
	for i := 1; i < 7; i++ {
		b.HoursRoundDecimalBirthday = append(b.HoursRoundDecimalBirthday, b.Born.Add(time.Hour*time.Duration(i*100000)))
	}
}

func (b *Birthday) CountMinutesRoundDecimalBirthday() {
	for i := 1; i < 42; i++ {
		b.MinutesRoundDecimalBirthday = append(b.MinutesRoundDecimalBirthday, b.Born.Add(time.Minute*time.Duration(i*1000000)))
	}
}

func (b *Birthday) CountSecondsRoundDecimalBirthday() {
	for i := 1; i < 27; i++ {
		b.SecondsRoundDecimalBirthday = append(b.SecondsRoundDecimalBirthday, b.Born.Add(time.Second*time.Duration(i*100000000)))
	}
}
