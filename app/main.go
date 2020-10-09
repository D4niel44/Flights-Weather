package main

func main() {
	app := NewApp(API_KEY, DB_PATH)
	executeDataset(app, "dataset1.csv")
	executeDataset(app, "dataset2.csv")
}

func executeDataset(app *App, datasetName string) {
	flights, cities := app.HandleDataSet(DATASETS_PATH + datasetName)
	app.QueryWeather(cities)
	PrintWeather(flights)
}
