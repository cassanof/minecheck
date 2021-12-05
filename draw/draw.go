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
		tm.Print(tm.MoveTo(box.String(), i*30|tm.PCT, 5|tm.PCT))
	}

	tm.Println()
	tm.Flush()
}

func MakeCryptonoteBox(user *cryptonote.User) *tm.Box {
	box := tm.NewBox(25|tm.PCT, 10, 0)

	fmt.Fprintf(box, "cryptonote.social\n")
	fmt.Fprintf(box, "Name: %s\n", user.Name)
	fmt.Fprintf(box, "Hash: %d H/s\n", user.HashRate)
	fmt.Fprintf(box, "Progress: %.4f%%\n", user.RewardProgress)
	fmt.Fprintf(box, "Owed: %.6f XMR\n", user.Owed)
	fmt.Fprintf(box, "Paid: %.6f XMR\n", user.Paid)

	return box
}

func MakeTwominersBox(user *twominers.User) *tm.Box {
	box := tm.NewBox(25|tm.PCT, 10, 0)

	fmt.Fprintf(box, "2miners.com\n")
	fmt.Fprintf(box, "Current Hash: %d H/s\n", user.HashRate)
	fmt.Fprintf(box, "Average Hash: %d H/s\n", user.AvgHashRate)
	fmt.Fprintf(box, "Paid: %.6f NANO\n", user.Paid)
	fmt.Fprintf(box, "Workers On: %d\n", user.WorkerOn)
	fmt.Fprintf(box, "Shares Validated: %d\n", user.SharesValid)

	return box
}

func HashRounder(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}
