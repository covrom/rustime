package rustime

import (
	"strconv"
	"strings"
	"time"
)

// д (d) - день месяца (цифрами) без лидирующего нуля
// дд (dd) - день месяца (цифрами) с лидирующим нулем
// ддд (ddd) - краткое название дня недели
// дддд (dddd) - полное название дня недели
// М (M) - номер месяца (цифрами) без лидирующего нуля
// ММ (MM) - номер месяца (цифрами) с лидирующим нулем
// МММ (MMM) - краткое название месяца
// ММММ (MMMM) - полное название месяца
// К (Q) - номер квартала в году
// г (y) - номер года без века и лидирующего нуля
// гг (yy) - номер года без века с лидирующим нулем
// гггг (yyyy) - номер года с веком
// ч (h) - час в 24 часовом варианте без лидирующих нулей
// чч (hh) - час в 24 часовом варианте с лидирующим нулем
// м (m) - минута без лидирующего нуля
// мм (mm) - минута с лидирующим нулем
// с (s) - секунда без лидирующего нуля
// сс (ss) - секунда с лидирующим нулем
// ссс (sss) - миллисекунда с лидирующим нулем
func FormatTimeRu(t time.Time, fmtstr string) string {
	days := [...]string{
		"воскресенье",
		"понедельник",
		"вторник",
		"среда",
		"четверг",
		"пятница",
		"суббота",
	}

	dayssm := [...]string{
		"вс",
		"пн",
		"вт",
		"ср",
		"чт",
		"пт",
		"сб",
	}

	months1 := [...]string{
		"", //0-го не бывает
		"январь",
		"февраль",
		"март",
		"апрель",
		"май",
		"июнь",
		"июль",
		"август",
		"сентябрь",
		"октябрь",
		"ноябрь",
		"декабрь",
	}

	months2 := [...]string{
		"", //0-го не бывает
		"января",
		"февраля",
		"марта",
		"апреля",
		"мая",
		"июня",
		"июля",
		"августа",
		"сентября",
		"октября",
		"ноября",
		"декабря",
	}

	src := []rune(string(fmtstr))
	res := make([]rune, 0, len(src)*2)
	wasday := false
	hour, min, sec := t.Clock()
	y, m, d := t.Date()

	i := 0
	for i < len(src) {
		var s []rune

		if i+4 <= len(src) {
			s = src[i : i+4]
			switch string(s) {
			case "дддд", "dddd":
				res = append(res, []rune(days[t.Weekday()])...)
				i += 4
				continue
			case "ММММ", "MMMM":
				if wasday {
					res = append(res, []rune(months2[int(t.Month())])...)
				} else {
					res = append(res, []rune(months1[int(t.Month())])...)
				}
				i += 4
				continue
			case "гггг", "yyyy":
				res = append(res, []rune(strconv.FormatInt(int64(t.Year()), 10))...)
				i += 4
				continue

			}
		}

		if i+3 <= len(src) {
			s = src[i : i+3]
			switch string(s) {
			case "ддд", "ddd":
				res = append(res, []rune(dayssm[t.Weekday()])...)
				i += 3
				continue
			case "МММ", "MMM":
				if wasday {
					res = append(res, []rune(months2[int(t.Month())])[:3]...)
				} else {
					res = append(res, []rune(months1[int(t.Month())])[:3]...)
				}
				i += 3
				continue
			case "ссс", "sss":
				sm := strconv.FormatInt(int64(t.Nanosecond())/1e6, 10)
				if len(sm) < 3 {
					sm = strings.Repeat("0", 3-len(sm)) + sm
				}
				res = append(res, []rune(sm)...)
				i += 3
				continue
			}
		}

		if i+2 <= len(src) {
			s = src[i : i+2]
			switch string(s) {
			case "дд", "dd":
				sm := strconv.Itoa(d)
				if len(sm) < 2 {
					sm = "0" + sm
				}
				res = append(res, []rune(sm)...)
				i += 2
				wasday = true
				continue
			case "ММ", "MM":
				sm := strconv.Itoa(int(m))
				if len(sm) < 2 {
					sm = "0" + sm
				}
				res = append(res, []rune(sm)...)
				i += 2
				continue
			case "гг", "yy":
				sm := strconv.Itoa(int(y % 100))
				if len(sm) < 2 {
					sm = "0" + sm
				}
				res = append(res, []rune(sm)...)
				i += 2
				continue
			case "чч", "hh":
				sm := strconv.Itoa(int(hour))
				if len(sm) < 2 {
					sm = "0" + sm
				}
				res = append(res, []rune(sm)...)
				i += 2
				continue
			case "мм", "mm":
				sm := strconv.Itoa(int(min))
				if len(sm) < 2 {
					sm = "0" + sm
				}
				res = append(res, []rune(sm)...)
				i += 2
				continue
			case "сс", "ss":
				sm := strconv.Itoa(int(sec))
				if len(sm) < 2 {
					sm = "0" + sm
				}
				res = append(res, []rune(sm)...)
				i += 2
				continue
			}
		}

		c := src[i]
		switch c {
		case 'д', 'd':
			sm := strconv.Itoa(d)
			res = append(res, []rune(sm)...)
			i++
			wasday = true
			continue
		case 'М', 'M':
			sm := strconv.Itoa(int(m))
			res = append(res, []rune(sm)...)
			i++
			continue
		case 'г', 'y':
			sm := strconv.Itoa(int(y % 100))
			res = append(res, []rune(sm)...)
			i++
			continue
		case 'ч', 'h':
			sm := strconv.Itoa(int(hour))
			res = append(res, []rune(sm)...)
			i++
			continue
		case 'м', 'm':
			sm := strconv.Itoa(int(min))
			res = append(res, []rune(sm)...)
			i++
			continue
		case 'с', 's':
			sm := strconv.Itoa(int(sec))
			res = append(res, []rune(sm)...)
			i++
			continue
		case 'К', 'Q':
			sm := strconv.FormatInt(int64(int64(t.Month())/4+1), 10)
			res = append(res, []rune(sm)...)
			i++
			continue
		}
		res = append(res, c)
		i++
	}

	return string(res)
}
