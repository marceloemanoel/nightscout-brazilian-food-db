package main

import (
	"crypto/sha1"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ColumnName         = 0
	ColumnUsualMeasure = 1
	ColumnWeight       = 2
	ColumnCarbs        = 3
	ColumnKiloCalories = 4
	ColumnKiloJoules   = 5
	ColumnCategorie    = 6
	ColumnSubCategorie = 7
	ColumnParty        = 8
)

type FoodEntry struct {
	Name         string
	UsualMeasure string
	Weight       float64
	Unit         string
	Carbs        float64
	KiloCalories float64
	KiloJoules   float64
	Categorie    string
	SubCategorie string
	Party        string
}

func NewFoodEntry(record []string) (*FoodEntry, error) {

	if len(record) != 9 {
		return nil, fmt.Errorf("invalid record, 9 columns required, %d provided", len(record))
	}

	carbs, err := strconv.ParseFloat(strings.ReplaceAll(record[ColumnCarbs], ",", "."), 64)
	if err != nil {
		return nil, err
	}

	var weight float64
	if record[ColumnWeight] == "" {
		weight = 0
	} else {
		weight, err = strconv.ParseFloat(strings.ReplaceAll(record[ColumnWeight], ",", "."), 64)
		if err != nil {
			return nil, err
		}
	}

	kcal, err := strconv.ParseFloat(strings.ReplaceAll(record[ColumnKiloCalories], ",", "."), 64)
	if err != nil {
		return nil, err
	}

	kj, err := strconv.ParseFloat(strings.ReplaceAll(record[ColumnKiloJoules], ",", "."), 64)
	if err != nil {
		return nil, err
	}

	unit := "g"
	if record[ColumnWeight] == "" || record[ColumnWeight] == "0" {
		unit = "pcs"
	}

	return &FoodEntry{
		Name:         record[ColumnName],
		UsualMeasure: record[ColumnUsualMeasure],
		Weight:       weight,
		Unit:         unit,
		Carbs:        carbs,
		KiloCalories: kcal,
		KiloJoules:   kj,
		Categorie:    record[ColumnCategorie],
		SubCategorie: record[ColumnSubCategorie],
	}, nil
}

func main() {
	// open file
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		line, _ := csvReader.FieldPos(0)

		// skip headers
		if line == 1 {
			continue
		}

		food, err := NewFoodEntry(rec)
		if err != nil {
			log.Fatal(err)
		}

		form := url.Values{}
		form.Add("type", "food")
		form.Add("category", food.Categorie)
		form.Add("subcategory", food.SubCategorie)
		form.Add("name", food.Name)
		form.Add("portion", fmt.Sprintf("%.2f", food.Weight))
		form.Add("carbs", fmt.Sprintf("%.2f", food.Carbs))
		// form.Add("fat", "0")
		// form.Add("protein", "0")
		form.Add("unit", food.Unit)
		form.Add("energy", fmt.Sprintf("%.2f", food.KiloJoules))

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/food/", os.Getenv("NS_BASE_URL")), strings.NewReader(form.Encode()))
		if err != nil {
			log.Fatal(err)
		}
		hasher := sha1.New()
		hasher.Write([]byte(os.Getenv("NS_API_SECRET")))
		apiSecret := hex.EncodeToString(hasher.Sum(nil))

		req.Header.Add("Api-Secret", apiSecret)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, _ := http.DefaultClient.Do(req)
		if res.StatusCode == http.StatusOK {
			log.Default().Println("Inserted record ", line)
			continue
		} else {

			body, _ := io.ReadAll(res.Body)
			log.Default().Println("StatusCode: ", res.Status)
			log.Default().Println(string(body))
		}

		time.Sleep(200 * time.Millisecond)
	}
}
