package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type restaurant struct {
	Category string `json:"category"`
	Name     string `json:"name"`
}

var restaurants = []restaurant{
	{Category: "정식류", Name: "북청집"},
	{Category: "정식류", Name: "통통낙지마을"},
	{Category: "정식류", Name: "계란집"},
	{Category: "정식류", Name: "더피크닉"},
	{Category: "정식류", Name: "소메랑"},
	{Category: "정식류", Name: "베이징스토리"},
	{Category: "정식류", Name: "MS생감자 치킨"},
	{Category: "정식류", Name: "김밥천국"},
	{Category: "정식류", Name: "명품국수보쌈정식"},
	{Category: "정식류", Name: "쉐누하누"},
	{Category: "일식류", Name: "스시가오"},
	{Category: "일식류", Name: "세이카츠"},
	{Category: "일식류", Name: "카츠쇼쿠도우"},
	{Category: "일식류", Name: "이자카야 단"},
	{Category: "일식류", Name: "고쿠텐"},
	{Category: "국물 및 국밥류", Name: "이태리 부대찌개"},
	{Category: "국물 및 국밥류", Name: "원조 닭한마리 칼국수"},
	{Category: "국물 및 국밥류", Name: "토속 상황 삼계탕"},
	{Category: "국물 및 국밥류", Name: "선순대"},
	{Category: "국물 및 국밥류", Name: "재크와콩나물"},
	{Category: "면류", Name: "포베이"},
	{Category: "면류", Name: "클준빛날영"},
	{Category: "면류", Name: "오한수우육면"},
	{Category: "면류", Name: "상록면"},
	{Category: "면류", Name: "슈퍼타이"},
	{Category: "면류", Name: "메콩타이"},
	{Category: "면류", Name: "시추안하우스"},
	{Category: "면류", Name: "소바니우동"},
	{Category: "양식류", Name: "내니스버거"},
	{Category: "양식류", Name: "바스버거"},
	{Category: "양식류", Name: "라디오베이"},
	{Category: "양식류", Name: "부리팝"},
	{Category: "회식류", Name: "빅사이즈"},
	{Category: "회식류", Name: "일편닭심"},
	{Category: "회식류", Name: "생어거스틴"},
	{Category: "회식류", Name: "차알"},
	{Category: "회식류", Name: "감성타코"},
	{Category: "회식류", Name: "부처스컷"},
	{Category: "회식류", Name: "차고145"},
	{Category: "회식류", Name: "돼지맨숀"},
	{Category: "회식류", Name: "진짜돼지"},
	{Category: "회식류", Name: "즐삼생고기"},
	{Category: "회식류", Name: "이가네양꼬치"},
	{Category: "회식류", Name: "야끼화로"},
}

func main() {
	router := gin.Default()
	router.GET("/restaurants", getRestaurants)
	router.GET("/restaurants/:category", getRestaurantsByCategory)
	router.GET("/restaurants/:category/random", getRandomRestaurantByCategory)

	router.Run("localhost:8080")
}

func getRestaurants(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, restaurants)
}

func sliceRestaurantsByCategory(category string) []restaurant {
	var restaurantsByCategory []restaurant

	for _, restaurant := range restaurants {
		if restaurant.Category == category {
			restaurantsByCategory = append(restaurantsByCategory, restaurant)
		}
	}

	return restaurantsByCategory
}

func getRestaurantsByCategory(c *gin.Context) {
	var restaurantsByCategory = sliceRestaurantsByCategory(c.Param("category"))

	if restaurantsByCategory != nil {
		c.IndentedJSON(http.StatusOK, restaurantsByCategory)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "해당 카테고리의 음식점이 없습니다."})
}

func getRandomRestaurantByCategory(c *gin.Context) {
	var restaurantsByCategory = sliceRestaurantsByCategory(c.Param("category"))

	if restaurantsByCategory != nil {
		c.IndentedJSON(http.StatusOK, restaurantsByCategory[rand.Intn(len(restaurantsByCategory))])
	}
}
