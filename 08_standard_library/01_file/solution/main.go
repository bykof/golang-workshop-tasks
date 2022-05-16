package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type StatisticRow struct {
	iso_code                string
	continent               string
	location                string
	date                    time.Time
	total_cases             int64
	new_cases               int64
	total_deaths            int64
	new_deaths              int64
	total_vaccinations      int64
	people_vaccinated       int64
	people_fully_vaccinated int64
	total_boosters          int64
	new_vaccinations        int64
}

type StatisticsRow []StatisticRow

func (ssr StatisticsRow) ISOGroups() map[string]StatisticsRow {
	isoGroups := make(map[string]StatisticsRow)
	for _, statisticRow := range ssr {
		if _, ok := isoGroups[statisticRow.iso_code]; !ok {
			isoGroups[statisticRow.iso_code] = StatisticsRow{}
		}
		isoGroups[statisticRow.iso_code] = append(isoGroups[statisticRow.iso_code], statisticRow)
	}
	return isoGroups
}

func (ssr StatisticsRow) LastEntries() StatisticsRow {
	isoGroups := ssr.ISOGroups()
	statisticsRows := StatisticsRow{}
	for _, groups := range isoGroups {
		if groups[0].continent == "" {
			continue
		}

		lastEntry := StatisticRow{}
		lastEntry.continent = groups[0].continent
		lastEntry.iso_code = groups[0].iso_code
		lastEntry.location = groups[0].location
		for _, group := range groups {
			if lastEntry.total_cases < group.total_cases {
				lastEntry.total_cases = group.total_cases
			}

			if lastEntry.total_boosters < group.total_boosters {
				lastEntry.total_boosters = group.total_boosters
			}

			if lastEntry.total_deaths < group.total_deaths {
				lastEntry.total_deaths = group.total_deaths
			}

			if lastEntry.total_vaccinations < group.total_vaccinations {
				lastEntry.total_vaccinations = group.total_vaccinations
			}
		}
		statisticsRows = append(statisticsRows, lastEntry)
	}
	return statisticsRows
}

func (sr StatisticRow) LetalityRate() float64 {
	if sr.total_cases == 0 {
		return 0
	}
	val := float64(sr.total_deaths) / float64(sr.total_cases)
	return val
}

func (ssr StatisticsRow) MostTotalCases() StatisticRow {
	var mostTotalCases StatisticRow

	for _, statisticRow := range ssr {
		if mostTotalCases.total_cases < statisticRow.total_cases {
			mostTotalCases = statisticRow
		}
	}
	return mostTotalCases
}

func (ssr StatisticsRow) MostPeopleVaccinated() StatisticRow {

	var mostPeopleVaccinated StatisticRow
	for _, statisticRow := range ssr {

		if mostPeopleVaccinated.total_vaccinations < statisticRow.total_vaccinations {
			mostPeopleVaccinated = statisticRow
		}
	}
	return mostPeopleVaccinated
}

func (ssr StatisticsRow) LeastLetalityRate() StatisticRow {
	var leastStatisticRow StatisticRow

	for _, statisticRow := range ssr {
		if statisticRow.total_cases == 0 ||
			statisticRow.total_deaths == 0 {
			continue
		}

		if (leastStatisticRow == StatisticRow{}) {
			leastStatisticRow = statisticRow
			continue
		}

		leastLetalityRate := leastStatisticRow.LetalityRate()
		currentLetalityRate := statisticRow.LetalityRate()

		if leastLetalityRate > currentLetalityRate {
			leastStatisticRow = statisticRow
		}

	}
	return leastStatisticRow
}

func stringToInt(value string) (int64, error) {
	if value != "" {
		i, err := strconv.Atoi(value)
		if err != nil {
			return 0, err
		}
		return int64(i), err
	}
	return 0, nil
}

func LineToStatisticRow(line string) (StatisticRow, error) {
	var statisticRow StatisticRow
	line = strings.ReplaceAll(line, "\r", "")
	splittedLine := strings.Split(line, ",")
	statisticRow.iso_code = splittedLine[0]
	statisticRow.continent = splittedLine[1]
	statisticRow.location = splittedLine[2]

	date, err := time.Parse("02/01/06", splittedLine[3])
	if err != nil {
		return statisticRow, err
	}

	statisticRow.date = date
	statisticRow.total_cases, err = stringToInt(splittedLine[4])
	if err != nil {
		return statisticRow, err
	}

	statisticRow.new_cases, err = stringToInt(splittedLine[5])
	if err != nil {
		return statisticRow, err
	}

	statisticRow.total_deaths, err = stringToInt(splittedLine[6])
	if err != nil {
		return statisticRow, err
	}

	statisticRow.new_deaths, err = stringToInt(splittedLine[7])
	if err != nil {
		return statisticRow, err
	}

	statisticRow.total_vaccinations, err = stringToInt(splittedLine[8])
	if err != nil {
		return statisticRow, err
	}

	statisticRow.people_vaccinated, err = stringToInt(splittedLine[9])
	if err != nil {
		return statisticRow, err
	}

	statisticRow.people_fully_vaccinated, err = stringToInt(splittedLine[10])
	if err != nil {
		return statisticRow, err
	}

	statisticRow.total_boosters, err = stringToInt(splittedLine[11])
	if err != nil {
		return statisticRow, err
	}

	statisticRow.new_vaccinations, err = stringToInt(splittedLine[12])
	if err != nil {
		return statisticRow, err
	}
	return statisticRow, nil
}

func ReadFile(filePath string) (StatisticsRow, error) {
	var statisticsRow StatisticsRow
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	for i, line := range lines {
		if i == 0 {
			continue
		}

		statisticRow, err := LineToStatisticRow(line)
		if err != nil {
			return nil, err
		}
		statisticsRow = append(statisticsRow, statisticRow)
	}
	return statisticsRow, nil
}

/*
	Task:
		Create a struct called StatisticRow to fill the data from data.csv
		Create a new user-defined type StatisticsRow, which is of type []StatisticRow.
		Read the file and write your own csv parser to parse the data into StatisticRows.

		Answer following questions by implementing methods for StatisticRows:
		1. Which country had the most total_cases?
		2. Which country had the least letality rate? (total_deaths / total_cases)
		3. Which country had the most people_vaccinated?
*/
func main() {
	statisticRows, err := ReadFile("data.csv")
	if err != nil {
		panic(err)
	}
	fmt.Printf("We have %d rows\n", len(statisticRows))

	lastEntries := statisticRows.LastEntries()

	mostTotalCases := lastEntries.MostTotalCases()
	fmt.Printf("%s had the most total cases with %d cases\n", mostTotalCases.location, mostTotalCases.total_cases)

	leastLetalityRate := lastEntries.LeastLetalityRate()
	fmt.Printf(
		"%s had the least letality rate with %.5f (total deaths: %d, total cases: %d)\n",
		leastLetalityRate.location,
		leastLetalityRate.LetalityRate(),
		leastLetalityRate.total_deaths,
		leastLetalityRate.total_cases,
	)

	mostPeopleVaccinated := lastEntries.MostPeopleVaccinated()
	fmt.Printf("%s had the most total people vaccinated with %d vaccinations\n", mostPeopleVaccinated.location, mostPeopleVaccinated.total_vaccinations)
}
