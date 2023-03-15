package highlight_test

import (
	"highlight"
	"testing"
)

func TestLog(t *testing.T) {
	t.Log(highlight.GREEN("[debug]"), "rainbows are  so beautiful !")
	t.Log(highlight.Concat(highlight.RED("rainbows"), highlight.GREEN("are"), highlight.YELLOW("so"), highlight.BLUE("beautiful"), highlight.MAGENTA("!")))
	t.Log(highlight.Yellow("[warn]"), "quick brown fox jumps over the lazy dog")
	t.Log(highlight.Concat(highlight.RED("quick"), highlight.GREEN("brown"), highlight.YELLOW("fox"), highlight.BLUE("jumps"), highlight.MAGENTA("over"), highlight.CYAN("the"), highlight.WHITE("lazy"), highlight.BLACK("dog")))
	t.Log(highlight.Red("[error]"), "to be or not to be, that is the question")
	t.Log(highlight.Concat(highlight.RED("to"), highlight.GREEN("be"), highlight.YELLOW("or"), highlight.BLUE("not"), highlight.MAGENTA("to"), highlight.CYAN("be,"), highlight.WHITE("that"), highlight.BLACK("is"), highlight.RED("the"), highlight.GREEN("question")))
	t.Log(highlight.Green("[info]"), "you can either choose to use font standard, bold, light, italic, or underline")
	t.Log(highlight.Concat(highlight.RED("you"), highlight.Green("can"), highlight.Yellow("either"), highlight.Blue("choose"), highlight.Magenta("to"), highlight.Cyan("use"), highlight.Blue("font"),highlight.YELLOW("bold"), highlight.Redf("light"), highlight.GREENf("italic"), highlight.Magenta("or"), highlight.CYAN("underline")))
	
}
