package main

import "fmt"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	/*

	      You are controlling a delivery drone that needs to complete a full delivery route. The route consists of n checkpoints arranged in a circular path.
	   At each checkpoint i:

	   The drone receives charge[i] units of battery power.
	   It then uses usage[i] units of energy to fly to the next checkpoint.



	   The drone starts the journey with an empty battery, but its battery can hold an unlimited amount of energy.
	   Your task is to determine whether the drone can complete the entire circuit starting from one of the checkpoints, such that it never runs out of energy along the way.
	   Return the index of the checkpoint where the drone should start its journey to complete the circuit. If it's not possible to complete the circuit from any checkpoint, return -1.
	   It is guaranteed that if a solution exists, it is unique.



	   Example :



	   Input: charge = [4, 6, 7, 4], usage = [6, 5, 3, 5]
	   Output: 1



	   Explanation:
	   Start at checkpoint 1:
	   - Get 6 energy, spend 5 → battery = 1
	   - Get 7, spend 3 → battery = 5
	   - Get 4, spend 5 → battery = 4
	   - Get 4, spend 6 → battery = 2 → back to start



	   Drone completes the loop successfully.


	   Input: charge = [2,3,4], usage = [3,4,3]
	   Output: -1



	   Constraints:

	   charge.length == usage.length
	   1 <= charge.length <= 10^5
	   0 <= charge[i], usage[i] <= 10^4
	*/

	charge := []int{2, 3, 4}
	usage := []int{3, 4, 3}
	fmt.Println(findStartCheckPoint(charge, usage))
}

func findStartCheckPoint(charge, usage []int) int {
	if len(charge) != len(usage) {
		return -1
	}

	currentCharge := 0
	start := 0
	totalCharge := 0
	totalUsage := 0

	for i := 0; i < len(charge); i++ {
		currentCharge += charge[i] - usage[i]
		totalUsage += usage[i]
		totalCharge += charge[i]
		if currentCharge < 0 {
			start = i + 1
			currentCharge = 0
		}
	}
	if totalCharge < totalUsage {
		return -1
	}

	return start

}
