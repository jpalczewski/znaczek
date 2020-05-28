package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kyokomi/emoji"
	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v2"
)

// Export exports  all labels from specified repo to json file.
func Export(c *cli.Context) error {
	fmt.Println("I should export something.")
	if c.Bool("debug") {

		emoji.Println(aurora.Yellow("But not yet :cowboy_hat_face:"))
	}
	ctx, client := getClient()

	sourceLabels, _, err := client.Issues.ListLabels(ctx, c.String("owner"), c.String("repository"), nil)
	if err != nil {
		log.Fatal(err)
	}

	exportLabels := make([]Label, len(sourceLabels))
	for i, l := range sourceLabels {
		exportLabels[i] = Label{*l.Name, *l.Description, *l.Color}
	}
	out, _ := json.MarshalIndent(exportLabels, "  ", "    ")

	if !c.Bool(Quiet) {
		fmt.Println(string(out))
	}

	handle, _ := os.Create(c.String(File))
	handle.WriteString(string(out))
	// for i, x := range ret {
	// 	lab := *label.NewLabel(x)
	// 	test, err := json.Marshal(lab)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(i, string(test))
	// }
	return nil
}
