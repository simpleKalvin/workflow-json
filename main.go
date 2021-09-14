// script_filter.go

package main

// Import name is "aw"
import (
	"bytes"
	"encoding/json"
	"github.com/deanishe/awgo"
)

// aw.Workflow is the main API
var wf *aw.Workflow

func init() {
	// Create a new *Workflow using default configuration
	// (workflow settings are read from the environment variables
	// set by Alfred)
	wf = aw.New()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}

func run() {
	// Create a new item
	//wf.NewItem("123").Copytext("{hello jack}").Arg("{hello jack}")
	argList := wf.Args()
	item := wf.NewItem("print 'enter' copy json result")
	inputString := argList[0]
	var str bytes.Buffer
	_ = json.Indent(&str, []byte(inputString), "", "    ")
	item.Arg(str.String()).Valid(true)
	// And send the results to Alfred
	wf.SendFeedback()
}