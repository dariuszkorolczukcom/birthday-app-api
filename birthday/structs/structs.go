package structs

import "time"

type BodyRequest struct {
	BirthdayDate string `json:"birthday"`
}

type Birthday struct {
	Born                        time.Time           `json:"birthdayDate"`
	HoursRoundDecimalBirthday   map[int64]time.Time `json:"hoursDecimal"`
	MinutesRoundDecimalBirthday map[int64]time.Time `json:"minutesDecimal"`
	SecondsRoundDecimalBirthday map[int64]time.Time `json:"secondsDecimal"`
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
	m := make(map[int64]time.Time)
	for i := 1; i < 7; i++ {
		n := int64(i * 100000)
		m[n] = b.Born.Add(time.Hour * time.Duration(n))
	}
	b.HoursRoundDecimalBirthday = m
}

func (b *Birthday) CountMinutesRoundDecimalBirthday() {
	m := make(map[int64]time.Time)
	for i := 1; i < 42; i++ {
		n := int64(i * 1000000)
		m[n] = b.Born.Add(time.Minute * time.Duration(n))
	}
	b.MinutesRoundDecimalBirthday = m
}

func (b *Birthday) CountSecondsRoundDecimalBirthday() {
	m := make(map[int64]time.Time)
	for i := 1; i < 27; i++ {
		n := int64(i * 100000000)
		m[n] = b.Born.Add(time.Second * time.Duration(n))
	}
	b.SecondsRoundDecimalBirthday = m
}
