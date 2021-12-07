package draw

import (
	"fmt"

	tm "github.com/buger/goterm"
	"github.com/elleven11/minecheck/cryptonote"
	"github.com/elleven11/minecheck/twominers"
)

func DrawBoxes(boxes []*tm.Box) {
	tm.Clear()

	for i, box := range boxes {
		tm.Print(tm.MoveTo(box.String(), i*60|tm.PCT, 5|tm.PCT))
	}

	tm.Println()
	tm.Flush()
}

func MakeCryptonoteBox(user *cryptonote.User) *tm.Box {
	box := tm.NewBox(40|tm.PCT, 10, 0)

	fmt.Fprintf(box, "cryptonote.social\n")
	fmt.Fprintf(box, "Name: %s\n", user.Name)
	fmt.Fprintf(box, "Hash: %s/s\n", hashRounder(user.HashRate))
	fmt.Fprintf(box, "Progress: %.4f%%\n", user.RewardProgress)
	fmt.Fprintf(box, "Owed: %.6f XMR\n", user.Owed)
	fmt.Fprintf(box, "Paid: %.6f XMR\n", user.Paid)

	return box
}

func MakeTwominersBox(user *twominers.User) *tm.Box {
	box := tm.NewBox(40|tm.PCT, 10, 0)

	fmt.Fprintf(box, "2miners.com\n")
	fmt.Fprintf(box, "Current Hash: %s/s\n", hashRounder(user.HashRate))
	fmt.Fprintf(box, "Average Hash: %s/s\n", hashRounder(user.AvgHashRate))
	fmt.Fprintf(box, "Paid: %.6f NANO\n", user.Paid)
	fmt.Fprintf(box, "Workers On: %d\n", user.WorkerOn)
	fmt.Fprintf(box, "Shares Validated: %d\n", user.SharesValid)

	return box
}

func hashRounder(hash int) string {
	const unit = 1000
	if hash < unit {
		return fmt.Sprintf("%d H", hash)
	}
	div, exp := int(unit), 0
	for n := hash / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cH",
		float64(hash)/float64(div), "KMGTPE"[exp])
}
