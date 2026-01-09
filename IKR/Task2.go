package main

import (
	"fmt"
	"sort"
)

type BrainrotMeme struct {
	Name        string
	TrendLevel  int
	Category    string
	Views       float64 
}

func FindTopTrending(memes []BrainrotMeme, minViews float64) []BrainrotMeme {
	var filtered []BrainrotMeme
	for _, m := range memes {
		if m.Views > minViews {
			filtered = append(filtered, m)
		}
	}

	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i].TrendLevel > filtered[j].TrendLevel
	})

	return filtered
}

func CalculateCategoryImpact(memes []BrainrotMeme) map[string]float64 {
	impact := make(map[string]float64)
	for _, m := range memes {
		impact[m.Category] += m.Views
	}
	return impact
}

func FilterByComplexCondition(memes []BrainrotMeme) []string {
	var result []string
	for _, m := range memes {
		if m.TrendLevel >= 7 || (m.Views > 50 && m.Category == "Sigma") {
			result = append(result, m.Name)
		}
	}
	return result
}

func main() {
	memes := []BrainrotMeme{
		{"Skibidi Toilet", 9, "Skibidi", 120.5},
		{"Sigma Grindset", 8, "Sigma", 60.0},
		{"Subo Bratik Dance", 6, "Subo Bratik", 45.2},
		{"TUNTUNSAHUR Compilation", 10, "TUNTUNTUNSAHUR", 200.0},
		{"Mewing Tutorial", 5, "Mewing", 30.0},
		{"Sigma Rage", 7, "Sigma", 55.5},
		{"Other Meme", 4, "Other", 10.0},
		{"Ultra Sigma", 6, "Sigma", 70.0}, 
	}

	fmt.Println("Самые трендовые")
	top := FindTopTrending(memes, 50)
	for _, m := range top {
		fmt.Printf("Название: %s, Уровень трендовости: %d, Просмотры: %.1f млн\n", m.Name, m.TrendLevel, m.Views)
	}

	fmt.Println("\nОбщие просмотры по категориям")
	impact := CalculateCategoryImpact(memes)
	for category, views := range impact {
		fmt.Printf("%s: %.1f млн просмотров\n", category, views)
	}

	fmt.Println("\nМемы для сложного условия")
	names := FilterByComplexCondition(memes)
	for _, name := range names {
		fmt.Println("-", name)
	}
}