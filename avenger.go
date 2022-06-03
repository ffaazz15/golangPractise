package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var choice int

type Sheild struct {
	Mission_Name, Mission_Details, Avengers, Status string
}

type Avengers struct {
	Name, personName, Abilities string
}

var sheild [100]Sheild

var avengers [100]Avengers

func recruitAvengers() {
	avengers[0] = Avengers{Name: "Iron Man", personName: "Tony Stark", Abilities: "Highly Intelligent Suit of Armor"}
	avengers[1] = Avengers{Name: "Captain America", personName: "Steve Rogers", Abilities: "Enhanced strength, speed, stamina, durability, agility, reflexes, senses, and mental processing"}
	avengers[2] = Avengers{Name: "Hulk", personName: "Bruce Banner", Abilities: "Immense superhuman strength, speed, stamina, and durability"}
	avengers[3] = Avengers{Name: "Thor", personName: "Thor Odinson", Abilities: "He also possesses superhuman strength, speed, agility, durability and immunity to most diseases."}
	avengers[4] = Avengers{Name: "Black Widow", personName: "Natasha Romanoff", Abilities: "Master in the covert arts of espionage, infiltration, and subterfuge."}
	avengers[5] = Avengers{Name: "Hawkeye", personName: "Clint Barton", Abilities: "Master archer and marksman Expert tactician, acrobat and hand-to-hand combatant Using a variety of trick arrows As Goliath: Superhuman strength and durability Size and mass manipulation."}
}

func checkMissions() {
	if len(sheild) == 0 {
		fmt.Printf("No Mission has been assigned to an Avenger.\n")
	} else {
		fmt.Println("Mission Name \t| Avengers \t| Status")
		for i := 0; i < len(sheild); i++ {

			fmt.Println(sheild[i].Mission_Name, " \t  ", sheild[i].Avengers, " \t  ", sheild[i].Status)
		}
	}
}

func checkNumberOfAvengers(Avengers string) int {
	split := strings.Split(Avengers, ",") //each avenger is given by using "," operator
	return len(split)
}

func checkStructure(split string) bool {
	for i := 0; i < len(avengers); i++ {
		if avengers[i].Name == split {
			return true
		}
	}
	fmt.Printf("Error: Avenger does not exist.\n")
	return false
}

func checkAvengerExists(Avengers string) bool {
	split := strings.Split(Avengers, ",")
	for i := 0; i < len(split); i++ {
		if checkStructure(split[i]) == false {
			return false
		}
	}
	return true
}

func checkAvengerMissions(Avengers string) bool {
	split := strings.Split(Avengers, ",")
	for i := 0; i < len(split); i++ {
		count := 0
		for j := 0; j < len(sheild); j++ {
			if strings.Contains(sheild[j].Avengers, split[i]) {
				count++
			}
			if count == 2 {
				fmt.Println(split[i], " has already taken up two missions.")
				return false
			}
		}
	}
	return true
}

func checkAvengerAvailability(Avengers string) bool {
	if checkNumberOfAvengers(Avengers) > 2 {
		fmt.Println("Error: You cannot assign more than 2 Avengers on a single mission.")
		return false
	}
	if checkAvengerExists(Avengers) && checkAvengerMissions(Avengers) {
		return true
	}
	return false
}

func checkMissionExists(Mission_Name string) bool {
	for i := 0; i < len(sheild); i++ {
		if Mission_Name == sheild[i].Mission_Name {
			fmt.Print("Mission Exists and is already taken up by", sheild[i].Avengers, "\n")
			return true
		}
	}
	return false
}

func assignMission() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Avengers: ")
	avg, _ := consoleReader.ReadString('\n')
	avg = strings.TrimSuffix(avg, "\n")
	if !checkAvengerAvailability(avg) {
		return
	}
	fmt.Print("Enter Mission Name: ")
	Mission_Name, _ := consoleReader.ReadString('\n')
	Mission_Name = strings.TrimSuffix(Mission_Name, "\n")
	if checkMissionExists(Mission_Name) {
		return
	}
	sheild[len(sheild)-1].Avengers = avg
	sheild[len(sheild)-1].Mission_Name = Mission_Name
	fmt.Print("Enter Mission Details: ")
	Mission_Details, _ := consoleReader.ReadString('\n')
	Mission_Details = strings.TrimSuffix(Mission_Details, "\n")
	sheild[len(sheild)-1].Mission_Details = Mission_Details
	sheild[len(sheild)-1].Status = "Assigned"
}

func checkMissionDetails() {
	var Search string
	fmt.Print("Enter Mission Name: ")
	fmt.Scan(&Search)
	for i := 0; i < len(sheild); i++ {
		if strings.Contains(sheild[i].Mission_Name, Search) {
			fmt.Println("Mission Details: ", sheild[i].Mission_Details, " Mission Status: ", sheild[i].Status, "Avengers: ", sheild[i].Avengers)
			return
		}
	}
	fmt.Print("Not Found")
}

func assigned(Search string) int {
	count := 0
	for i := 0; i < len(sheild); i++ {
		if strings.Contains(sheild[i].Avengers, Search) && sheild[i].Status == "Assigned" {
			count++
		}
	}
	return count
}

func completed(Search string) int {
	count := 0
	for i := 0; i < len(sheild); i++ {
		if strings.Contains(sheild[i].Avengers, Search) && sheild[i].Status == "Completed" {
			count++
		}
	}
	return count
}

func checkAvengerDetails() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Avenger Name: ")
	Search, _ := consoleReader.ReadString('\n')
	Search = strings.TrimSuffix(Search, "\n")
	for i := 0; i < len(avengers); i++ {
		if avengers[i].Name == Search {
			fmt.Println("Person Name: ", avengers[i].personName, " Abilities: ", avengers[i].Abilities, " Assigned : ", assigned(Search), " Mission Completed : ", completed(Search))
			return
		}
	}
	fmt.Println("Avenger not found..!!")
}

func updateStatus() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Mission Name: ")
	Search, _ := consoleReader.ReadString('\n')
	Search = strings.TrimSuffix(Search, "\n")
	for i := 0; i < len(sheild); i++ {
		if strings.Contains(sheild[i].Mission_Name, Search) {
			var status string
			fmt.Println("please enter the new status")
			fmt.Scan(&status)
			sheild[i].Status = status
		}
	}
}

func listAvengers() {
	fmt.Println("Name\t", "|status\t,", "|Mission")
	flag := "1"
	for i := 0; i < len(avengers); i++ {
		for j := 0; j < len(sheild); j++ {
			if strings.Contains(avengers[i].Name, sheild[j].Avengers) {
				fmt.Println(avengers[i].Name, "\t|", sheild[j].Status, "\t|", sheild[j].Mission_Name)
				flag = avengers[i].Name
				break
			}
		}
		if flag != "1" {
			fmt.Printf(avengers[i].Name, "\t|Available\t|", "\n")
		}
		flag = "1"
	}
}

func checkMissionMembers(Mission string, Search string) {
	flag := 0
	for i := 0; i < len(sheild); i++ {
		if strings.Contains(sheild[i].Mission_Name, Search) {
			flag = 1
			if checkNumberOfAvengers(sheild[i].Avengers) == 2 {
				fmt.Print("Error: Already two avangers Present")
				return
			} else if checkNumberOfAvengers(sheild[i].Avengers) == 0 {
				sheild[i].Avengers = Search
			} else {
				sheild[i].Avengers = sheild[i].Avengers + "," + Search
			}
		}
	}
	if flag == 0 {
		fmt.Print("Error: Mission Not Found")
		return
	}

}

func assignAvenger() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Avenger Name: ")
	Search, _ := consoleReader.ReadString('\n')
	Search = strings.TrimSuffix(Search, "\n")
	if !checkAvengerExists(Search) {
		return
	}
	if !checkAvengerMissions(Search) {
		return
	}
	fmt.Print("Enter Mission Name: ")
	Mission, _ := consoleReader.ReadString('\n')
	Mission = strings.TrimSuffix(Mission, "\n")
	checkMissionMembers(Mission, Search)
}

func avenger() {
	recruitAvengers()
	for {
		fmt.Println("=========------S.H.E.I.L.D\t=========\nWelcome to Captain Fury.")
		fmt.Println("1. Check the missions\n2. Assign mission to Avengers\n3. Check mission’s details\n4. Check Avenger’s details\n5. Update Mission’s status\n6. List Avengers\n7. Assign avenger to mission.")
		fmt.Println("Enter the option: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			checkMissions()
		case 2:
			assignMission()
		case 3:
			checkMissionDetails()
		case 4:
			checkAvengerDetails()
		case 5:
			updateStatus()
		case 6:
			listAvengers()
		case 7:
			assignAvenger()
		default:
			fmt.Println("Invalid")
		}
	}
}
