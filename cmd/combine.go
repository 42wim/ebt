package cmd

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/vault/shamir"
	"github.com/spf13/cobra"
)

func combineCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "combine [flags] FILE [FILE...]",
		Short:   "Combine the shamir secret sharded parts to restore a private key.",
		Example: `ebt combine shard1.txt shard2.txt shard3.txt`,
		Args:    cobra.MinimumNArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			result, err := combineParts(args)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(result)
		},
	}
}

func combineParts(args []string) (string, error) {
	var resparts [][]byte

	for _, f := range args {
		part, err := ioutil.ReadFile(f)
		if err != nil {
			return "", fmt.Errorf("reading file %s failed: %s", f, err)
		}

		bpart, err := base64.StdEncoding.DecodeString(string(part))
		if err != nil {
			return "", fmt.Errorf("reading base64 decode %s failed: %s", f, err)
		}

		resparts = append(resparts, bpart)
	}

	combined, err := shamir.Combine(resparts)
	if err != nil {
		return "", fmt.Errorf("combination failed: %s", err)
	}

	return string(combined), nil
}
