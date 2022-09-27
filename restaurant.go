package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type FoodMenu struct {
	FoodMenuID string `json:"foodmenu_id"`
	EaterID    string `json:"eater_id"`
}

type KeyValueMap struct {
	Key   string
	Value int
}

func TopThreeFood() error {
	jsonFile, err := os.Open("restlog.json")
	if err != nil {
		return err
	}
	fmt.Println("Successfully Opened log.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var foodmenu_id []FoodMenu
	var eater_id []FoodMenu

	json.Unmarshal(byteValue, &foodmenu_id)
	json.Unmarshal(byteValue, &eater_id)

	foodmenuMap := make(map[string]int)
	eater_idMap := make(map[string]int)

	for i := 0; i < len(foodmenu_id); i++ {
		foodmenuMap[foodmenu_id[i].FoodMenuID]++
	}

	for i := 0; i < len(eater_id); i++ {
		eater_idMap[eater_id[i].EaterID]++
	}
	sortedMap := sortMap(foodmenuMap)

	fmt.Println("----Top 3 FoodMenuIDs----")
	for i, kv := range sortedMap {
		if i == 3 {
			break
		}
		fmt.Println("FoodMenuIDs: " + kv.Key + " Count: " + fmt.Sprint(kv.Value))

	}
	var idArray []string
	for _, tempData := range foodmenu_id {
		idArray = append(idArray, tempData.EaterID)
	}
	unique(foodmenu_id)

	return nil
}

func sortMap(foodmenuMap map[string]int) []KeyValueMap {

	var ss []KeyValueMap
	for k, v := range foodmenuMap {
		ss = append(ss, KeyValueMap{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}

func unique(s []FoodMenu) []FoodMenu {
	inResult := make(map[FoodMenu]bool)
	var result []FoodMenu
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		} else {
			fmt.Println("Error EaterID: " + str.EaterID + " has FoodMenuID: " + str.FoodMenuID + " more than once")
		}
	}
	return result
}

func main() {
	err := TopThreeFood()
	if err != nil {
		fmt.Println(err)
	}

}
