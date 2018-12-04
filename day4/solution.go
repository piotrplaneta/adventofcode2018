package day4

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/piotrplaneta/adventofcode2018/utils"
)

type log struct {
	timestamp time.Time
	content   string
}

type logsArray []log
type sleepLog struct {
	guardID string
	minute  int
}

//SolvePart1 returns the answer for part 1 of day 4
func SolvePart1() int {
	return SleeperIDMultipliedByMostFrequentMinute(adventInput())
}

//SolvePart2 returns the answer for part 2 of day 4
func SolvePart2() int {
	return SingularMinuteSleeperIDMultipliedByMinute(adventInput())
}

/*
SleeperIDMultipliedByMostFrequentMinute returns id of the guard sleeping the most
multiplied by the minute in which he/she sleeps the most
*/
func SleeperIDMultipliedByMostFrequentMinute(input []string) int {
	logs := sortLogs(timestampedEntries(input))
	sleepingMinutes := calculateSleepingMinutes(logs)

	sleeperID := mostMinutesAsleep(sleepingMinutes)
	mostFrequentMinute := mostFrequentMinute(sleepingMinutes, sleeperID)

	return stringToInt(sleeperID[1:]) * mostFrequentMinute
}

/*
SingularMinuteSleeperIDMultipliedByMinute returns id of the guard sleeping the most
during singular minute multiplied by the minute in which he/she sleeps the most
*/
func SingularMinuteSleeperIDMultipliedByMinute(input []string) int {
	logs := sortLogs(timestampedEntries(input))
	sleepingMinutes := calculateSleepingMinutes(logs)

	sleeperLog := mostFrequentlyAsleepDuringSingularMinuteLog(sleepingMinutes)

	return stringToInt(sleeperLog.guardID[1:]) * sleeperLog.minute
}

func timestampedEntries(input []string) logsArray {
	result := make([]log, len(input))

	for i, v := range input {
		timeString := strings.Trim(strings.Split(v, "]")[0], "[")
		content := strings.Trim(strings.Split(v, "]")[1], " ")

		result[i] = log{timestamp: parseTime(timeString), content: content}
	}

	return result
}

func sortLogs(logs logsArray) logsArray {
	sortedLogs := make(logsArray, len(logs))
	copy(sortedLogs, logs)
	sort.Slice(sortedLogs, func(i, j int) bool { return sortedLogs[i].timestamp.Before(sortedLogs[j].timestamp) })
	return sortedLogs
}

func calculateSleepingMinutes(logs logsArray) map[sleepLog]int {
	sleepingMinutes := make(map[sleepLog]int)
	for i := 0; i < len(logs); {
		currentLog := logs[i]
		guardID := strings.Split(currentLog.content, " ")[1]

		for j := i + 1; true; j += 2 {
			if j == len(logs) || strings.Contains(logs[j].content, "begins shift") {
				i = j
				break
			}

			sleepStartTime := logs[j].timestamp
			sleepEndTime := logs[j+1].timestamp

			for k := sleepStartTime.Minute(); k < sleepEndTime.Minute(); k++ {
				sleepingMinutes[sleepLog{guardID: guardID, minute: k}]++
			}
		}
	}

	return sleepingMinutes
}

func mostMinutesAsleep(sleepingMinutes map[sleepLog]int) string {
	sumPerID := make(map[string]int)

	for k, v := range sleepingMinutes {
		sumPerID[k.guardID] += v
	}

	max := 0
	maxID := ""

	for k, v := range sumPerID {
		if v > max {
			max = v
			maxID = k
		}
	}

	return maxID
}

func mostFrequentMinute(sleepingMinutes map[sleepLog]int, sleeperID string) int {
	sumPerMinute := make(map[int]int)

	for k, v := range sleepingMinutes {
		if k.guardID == sleeperID {
			sumPerMinute[k.minute] += v
		}
	}

	max := 0
	maxMinute := 61

	for k, v := range sumPerMinute {
		if v > max {
			max = v
			maxMinute = k
		}
	}

	return maxMinute
}

func mostFrequentlyAsleepDuringSingularMinuteLog(sleepingMinutes map[sleepLog]int) sleepLog {
	maxLog := sleepLog{}
	maxSingularMinute := 0

	for k, v := range sleepingMinutes {
		if v > maxSingularMinute {
			maxLog = k
			maxSingularMinute = v
		}
	}

	return maxLog
}

func stringToInt(input string) int {
	result, err := strconv.Atoi(input)

	if err != nil {
		panic(err)
	}

	return result
}

func parseTime(input string) time.Time {
	layout := "2006-01-02 15:04"
	t, err := time.Parse(layout, input)

	if err != nil {
		panic(err)
	}

	return t
}

func adventInput() []string {
	lines, _ := utils.ReadLinesFromFile("day4/input")
	return lines
}
