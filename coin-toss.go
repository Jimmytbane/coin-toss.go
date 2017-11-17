/*
BSD 3-Clause License

Copyright (c) 2017, JimmyBot
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of the copyright holder nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

func main() {
	var input2 string
	var input3 string
	var choice string
	var else_eval bool
	

	for {
		rand.Seed(time.Now().UTC().UnixNano())

		fmt.Printf("Heads or Tails?\n-->  ")
		fmt.Scanln(&choice)
		strings.ToLower(choice)

		if (choice == "heads") {
			fmt.Println("You chose Heads")
			break
		} else if choice == "tails" {
			fmt.Println("You chose Tails")
			break
		} else {
			fmt.Println("You said something invalid.")
		}
	}

	if (rand.Intn(2) == 0) {
		fmt.Printf("You Flipped: Heads")
		fmt.Scanln(&input2)
	} else {
		fmt.Printf("You Flipped: Tails")
		fmt.Scanln(&input3)
		else_eval = true
	}

	if (rand.Intn(2) == 0 && choice == "heads") {
		fmt.Println("You won!")
	} else if (rand.Intn(2) == 0 && choice == "tails") {
		fmt.Println("You lost.")
	} else if (rand.Intn(2) == 1 && choice == "tails") {
		fmt.Println("You won!")
	} else if (else_eval == true && choice == "heads") {
		fmt.Println("You lost.")
	}
}
