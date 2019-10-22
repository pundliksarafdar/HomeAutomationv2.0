package schedular

import "github.com/carlescere/scheduler"

type Duration string

const (
	SECOND  Duration = "SECOND"
	MIN     Duration = "MIN"
	HOUR    Duration = "HOUR"
	DAY     Duration = "DAY"
	MONTH   Duration = "MONTH"
)

/*
		This method wil schedule job based on Duration every (quant)  dur(Sec,Min,hour,day,month),
		runNow : Set to true if you want to run it immediately
		job is the function to run on given time
*/
func StartSchedular(dur Duration, qunt int, runNow bool,job func()) {
	sc := scheduler.Every(qunt)
	switch dur {
	case SECOND:
		{
			sc = sc.Seconds()
		}
	case MIN:
		{
			sc = sc.Minutes()
		}
	case HOUR:
		{
			sc = sc.Hours()
		}
	case DAY:
		{
			sc = sc.Day()
		}
	}

	if !runNow {
		sc = sc.NotImmediately()
	}

	sc.Run(job)
}
