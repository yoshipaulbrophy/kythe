/*
 * Copyright 2015 Google Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Binary kythe exposes a CLI interface to the xrefs and filetree
// services backed by a combined serving table.
//
// Examples:
//   # Show complete command listing
//   kythe
//
//   # List all corpus root uris
//   kythe --serving_table /path/to/table ls --uris
//
//   # List root directory contents for corpus named 'somecorpus'
//   kythe --serving_table /path/to/table ls kythe://somecorpus
//
//   # List Kythe's kythe/cxx/common directory (as URIs)
//   kythe --serving_table /path/to/table ls --uris kythe://kythe?path=kythe/cxx/common
//
//   # Display all file anchors references for kythe/cxx/common/CommandLineUtils.cc
//   kythe --serving_table /path/to/table refs kythe://kythe?lang=c%2B%2B?path=kythe/cxx/common/CommandLineUtils.cc
//
//   # Show all outward edges for a particular node
//   kythe --serving_table /path/to/table edges kythe:?lang=java#java.util.List
//
//   # Show reverse /kythe/edge/defines edges for a node
//   kythe --serving_table /path/to/table edges --kinds '%/kythe/edge/defines' kythe://kythe?lang=java?path=kythe/java/com/google/devtools/kythe/analyzers/base/EntrySet.java#1887f665ee4c77287d1022c151000a489e17147215309818cf4150c601442cc5
//
//   # Show all facts (except /kythe/text) for a node
//   kythe --serving_table /path/to/table node kythe:?lang=c%2B%2B#StripPrefix%3Acommon%3Akythe%23n%23D%40kythe%2Fcxx%2Fcommon%2FCommandLineUtils.cc%3A167%3A1
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"

	"kythe/go/serving/filetree"
	"kythe/go/serving/xrefs"
	"kythe/go/storage/leveldb"
	"kythe/go/storage/table"
)

var (
	// TODO(schroederc): add an option for a remote xrefs/filetree service
	servingTable = flag.String("serving_table", "", "LevelDB serving table (required)")
)

func globalUsage() {
	fmt.Fprintf(os.Stderr, `Usage: %s <global-flags> <command> <flags>
Example: %[1]s --serving_table /path/to/table ls --uris kythe://kythe?path=kythe/cxx/common

Global Flags:
`, filepath.Base(os.Args[0]))
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nCommands:")
	var cmdNames []string
	for name := range cmds {
		cmdNames = append(cmdNames, name)
	}
	sort.Strings(cmdNames)
	for _, name := range cmdNames {
		cmds[name].Usage()
		fmt.Fprintln(os.Stderr)
	}
}

var cmds = map[string]command{
	"edges":  cmdEdges,
	"ls":     cmdLS,
	"node":   cmdNode,
	"refs":   cmdRefs,
	"source": cmdSource,
	// TODO(schroederc): search command using inverted index to find nodes
}

func init() {
	cmds["help"] = newCommand("help", "[command]",
		"Print help information for the given command",
		func(flag *flag.FlagSet) {}, func(flag *flag.FlagSet) error {
			if len(flag.Args()) == 0 {
				globalUsage()
			} else {
				getCommand(flag.Arg(0)).Usage()
			}
			return nil
		})
	flag.Usage = globalUsage
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	if *servingTable == "" {
		log.Fatal("Missing required --serving_table argument")
	}

	db, err := leveldb.Open(*servingTable, nil)
	if err != nil {
		log.Fatalf("Error opening db at %q: %v", *servingTable, err)
	}
	defer db.Close()

	tbl := &table.KVProto{db}
	xs = &xrefs.Table{tbl}
	ft = &filetree.Table{tbl}

	if err := getCommand(flag.Arg(0)).run(); err != nil {
		log.Fatal("ERROR: ", err)
	}
}

func getCommand(name string) command {
	c, ok := cmds[name]
	if !ok {
		fmt.Fprintf(os.Stderr, "ERROR: unknown command %q\n", name)
		globalUsage()
		os.Exit(1)
	}
	return c
}