/*
   Copyright (C) <2017>  <jimmybot@teknik.io>
   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published
   by the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.
   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed()
	choice()
	roll()
}

func choice() {
	for {
		fmt.Println("Heads or Tails?")
		var input1 string
		fmt.Scanln(&input1)
		if input1 == "heads" {
			fmt.Println("You chose heads")
			break
		} else if input1 == "tails" {
			fmt.Println("You chose tails")
			break
		} else {
			fmt.Println("You said something invalid.")
		}
	}
}

func seed() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func roll() {
	if rand.Intn(2) == 0 {
		fmt.Println("Heads")
	} else {
		fmt.Println("Tails")
	}
}
