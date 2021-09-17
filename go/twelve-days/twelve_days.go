package twelve

import (
	"fmt"
	"strings"
)

type SongContent struct {
	date string
	gave string
}



func (s *SongContent) setDate(d int) {
	switch d {
	case 1:
		s.date = "first"
	case 2:
		s.date = "second"
	case 3:
		s.date = "third"
	case 4:
		s.date = "fourth"
	case 5:
		s.date = "fifth"
	case 6:
		s.date = "sixth"
	case 7:
		s.date = "seventh"
	case 8:
		s.date = "eighth"
	case 9:
		s.date = "ninth"
	case 10:
		s.date = "tenth"
	case 11:
		s.date = "eleventh"
	case 12:
		s.date = "twelfth"
	}

}

func (s *SongContent) setGave(n int) {
	switch n {
	case 1:
		s.gave = "a Partridge in a Pear Tree."
	case 2:
		s.gave = "two Turtle Doves, and "
	case 3:
		s.gave = "three French Hens, "
	case 4:
		s.gave = "four Calling Birds, "
	case 5:
		s.gave = "five Gold Rings, "
	case 6:
		s.gave = "six Geese-a-Laying, "
	case 7:
		s.gave = "seven Swans-a-Swimming, "
	case 8:
		s.gave = "eight Maids-a-Milking, "
	case 9:
		s.gave = "nine Ladies Dancing, "
	case 10:
		s.gave = "ten Lords-a-Leaping, "
	case 11:
		s.gave = "eleven Pipers Piping, "
	case 12:
		s.gave = "twelve Drummers Drumming, "
	}

}

func Verse(d int) string {
	const s = "On the %v day of Christmas my true love gave to me: %v"
	var gaveContent string
	var ans string

	for i := 1; i <= 12; i++ {
		sc := SongContent{}
		sc.setDate(i)
		sc.setGave(i)
		gaveContent = sc.gave+gaveContent
		ans = fmt.Sprintf(s,sc.date,gaveContent)
		if i == d {
			return ans
		}
	}
	return ans
}

func Song() string{
	const s = "On the %v day of Christmas my true love gave to me: %v\n"
	var gaveContent string
	var ans string
	for i := 1; i <= 12; i++ {
		sc := SongContent{}
		sc.setDate(i)
		sc.setGave(i)
		gaveContent = sc.gave+gaveContent
		ans = ans+fmt.Sprintf(s,sc.date,gaveContent)
	}
	return strings.Trim(ans,"\n")
}
