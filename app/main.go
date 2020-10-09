package main

import "fmt"

func main() {
	app := NewApp(API_KEY, DB_PATH)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Fatal error: ", r)
		}
		app.Close()
	}()

	executeDataset(app, "dataset1.csv")
	executeDataset(app, "dataset2.csv")
}

func executeDataset(app *App, datasetName string) {
	flights, cities := app.HandleDataSet(DATASETS_PATH + datasetName)
	app.QueryWeather(cities)
	PrintWeather(flights)
}
