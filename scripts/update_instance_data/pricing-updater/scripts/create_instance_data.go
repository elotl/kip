package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const instanceDataOutputPath =  "pkg/util/instanceselector/instance_data.go"

// Reads all .json files in the current folder
// and encodes them as strings literals in kip/pkg/util/instanceselector/instance_data.go

// Hopefully we will get rid of this scripts once go start supporting embedding files natively
func main() {
	topDir := os.Getenv("TOP_KIP_DIR")
	if topDir == "" {
		panic("please set TOP_KIP_DIR env var")
	}
	outputAbsolutePath := filepath.Join(topDir, instanceDataOutputPath)
	pricingUpdaterPath := filepath.Join(topDir, "scripts/update_instance_data/pricing-updater")
	fs, _ := ioutil.ReadDir(pricingUpdaterPath)
	currentWorkingDirectory, _ := os.Getwd()
	fmt.Printf("current dir: %s\n", currentWorkingDirectory)
	err := os.Remove(outputAbsolutePath)
	if err != nil {
		panic(err)
	}
	out, err := os.Create(outputAbsolutePath)
	if err != nil {
		log.Fatal(err)
	}
	_, err = out.Write([]byte("package instanceselector \n\nconst (\n"))
	if err != nil {
		panic(err)
	}
	for _, f := range fs {
		if !strings.HasSuffix(f.Name(), ".json") {
			fmt.Printf("skipping %s\n", f.Name())
			continue
		}
		fmt.Printf("reading %s\n", f.Name())
		// this will create buildawsInstanceJson,
		out.Write([]byte("    " + "build" + strings.TrimSuffix(f.Name(), ".json") + " = `"))
		fmt.Printf("opening %s\n", f.Name())
		source, err := os.Open(filepath.Join(pricingUpdaterPath, f.Name()))
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(out, source)
		out.Write([]byte("`\n"))
	}
	out.Write([]byte(")\n"))
	fmt.Printf("saving %s\n", out.Name())
	out.Close()

}
