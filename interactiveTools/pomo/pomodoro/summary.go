package pomodoro

import "time"

func DailySummary(day time.Time, config *IntervalConfig) ([]time.Duration, error) {
	dPomo, err := config.repo.CategorySummary(day, CategoryPomodoro)
	if err != nil {
		return nil, err
	}
	dBreaks, err := config.repo.CategorySummary(day, "%Break")
	if err != nil {
		return nil, err
	}
	return []time.Duration{
		dPomo,
		dBreaks,
	}, nil
}
