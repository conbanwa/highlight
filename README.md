# highlight

string highlight: a simple, extensive string colorizer for Go.

Basically, turns string output into this:

![dark](https://raw.githubusercontent.com/reconquest/colorgful/master/dark.png)

Or this:

![light](https://raw.githubusercontent.com/reconquest/colorgful/master/light.png)


# Example

```go
package main

import (
	"github.com/kovetskiy/lorg"
	"github.com/reconquest/colorgful"
)

func main() {
	log := lorg.NewLog()

	log.SetFormat(colorgful.MustApplyDefaultTheme(
		`* ${time} ${level} %s`,
		colorgful.Light,
	))

	log.SetLevel(lorg.LevelTrace)

	log.Trace("tracing dead status")

	log.Debug("debuggin dead status")

	log.Info("soon you will be dead")

	log.Warning("you are not prepared to be dead")

	log.Error("you're dead!")

	log.Fatal("stopping")
}
```
