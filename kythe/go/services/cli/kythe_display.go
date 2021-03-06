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

package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"kythe.io/kythe/go/services/web"
	"kythe.io/kythe/go/services/xrefs"
	"kythe.io/kythe/go/util/kytheuri"
	"kythe.io/kythe/go/util/markedsource"
	"kythe.io/kythe/go/util/schema/edges"
	"kythe.io/kythe/go/util/schema/facts"

	"bitbucket.org/creachadair/stringset"
	"github.com/golang/protobuf/proto"

	cpb "kythe.io/kythe/proto/common_proto"
	ftpb "kythe.io/kythe/proto/filetree_proto"
	gpb "kythe.io/kythe/proto/graph_proto"
	xpb "kythe.io/kythe/proto/xref_proto"
)

var (
	logRequests = flag.Bool("log_requests", false, "Log all requests to stderr as JSON")
	displayJSON = flag.Bool("json", false, "Display results as JSON")
	out         = os.Stdout
)

var jsonMarshaler = web.JSONMarshaler

func init() { jsonMarshaler.Indent = "  " }

func logRequest(req proto.Message) {
	if *logRequests {
		str, err := jsonMarshaler.MarshalToString(req)
		if err != nil {
			log.Fatalf("Failed to encode request for logging %v: %v", req, err)
		}
		log.Printf("%s: %s", baseTypeName(req), string(str))
	}
}

func baseTypeName(x interface{}) string {
	ss := strings.SplitN(fmt.Sprintf("%T", x), ".", 2)
	if len(ss) == 2 {
		return ss[1]
	}
	return ss[0]
}

func displayCorpusRoots(cr *ftpb.CorpusRootsReply) error {
	if *displayJSON {
		return jsonMarshaler.Marshal(out, cr)
	}

	for _, c := range cr.Corpus {
		for _, root := range c.Root {
			var err error
			if lsURIs {
				uri := kytheuri.URI{
					Corpus: c.Name,
					Root:   root,
				}
				_, err = fmt.Fprintln(out, uri.String())
			} else {
				_, err = fmt.Fprintln(out, filepath.Join(c.Name, root))
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func displayDirectory(d *ftpb.DirectoryReply) error {
	if *displayJSON {
		return jsonMarshaler.Marshal(out, d)
	}

	for _, d := range d.Subdirectory {
		if !lsURIs {
			uri, err := kytheuri.Parse(d)
			if err != nil {
				return fmt.Errorf("received invalid directory uri %q: %v", d, err)
			}
			d = filepath.Base(uri.Path) + "/"
		}
		if _, err := fmt.Fprintln(out, d); err != nil {
			return err
		}
	}
	for _, f := range d.File {
		if !lsURIs {
			uri, err := kytheuri.Parse(f)
			if err != nil {
				return fmt.Errorf("received invalid file ticket %q: %v", f, err)
			}
			f = filepath.Base(uri.Path)
		}
		if _, err := fmt.Fprintln(out, f); err != nil {
			return err
		}
	}
	return nil
}

func displaySource(decor *xpb.DecorationsReply) error {
	if *displayJSON {
		return jsonMarshaler.Marshal(out, decor)
	}

	_, err := out.Write(decor.SourceText)
	return err
}

func displayDecorations(decor *xpb.DecorationsReply) error {
	if *displayJSON {
		return jsonMarshaler.Marshal(out, decor)
	}

	nodes := xrefs.NodesMap(decor.Nodes)

	for _, ref := range decor.Reference {
		nodeKind := factValue(nodes, ref.TargetTicket, facts.NodeKind, "UNKNOWN")
		subkind := factValue(nodes, ref.TargetTicket, facts.Subkind, "")

		var targetDef string
		if ref.TargetDefinition != "" {
			targetDef = ref.TargetDefinition
			// TODO(schroederc): fields from decor.DefinitionLocations
			// TODO(zarko): fields from decor.ExtendsOverrides
		}

		r := strings.NewReplacer(
			"@target@", ref.TargetTicket,
			"@edgeKind@", ref.Kind,
			"@nodeKind@", nodeKind,
			"@subkind@", subkind,
			"@^offset@", itoa(ref.Span.Start.ByteOffset),
			"@^line@", itoa(ref.Span.Start.LineNumber),
			"@^col@", itoa(ref.Span.Start.ColumnOffset),
			"@$offset@", itoa(ref.Span.End.ByteOffset),
			"@$line@", itoa(ref.Span.End.LineNumber),
			"@$col@", itoa(ref.Span.End.ColumnOffset),
			"@targetDef@", targetDef,
		)
		if _, err := r.WriteString(out, refFormat+"\n"); err != nil {
			return err
		}
	}

	return nil
}

func itoa(n int32) string { return strconv.Itoa(int(n)) }

func displayEdges(reply *gpb.EdgesReply) error {
	if *displayJSON {
		return jsonMarshaler.Marshal(out, reply)
	}

	for source, es := range reply.EdgeSets {
		if _, err := fmt.Fprintln(out, "source:", source); err != nil {
			return err
		}
		for kind, g := range es.Groups {
			for _, edge := range g.Edge {
				if _, err := fmt.Fprintf(out, "%s\t%s\n", kind, edge.TargetTicket); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func displayTargets(edges map[string]*gpb.EdgeSet) error {
	var targets stringset.Set
	for _, es := range edges {
		for _, g := range es.Groups {
			for _, e := range g.Edge {
				targets.Add(e.TargetTicket)
			}
		}
	}

	if *displayJSON {
		return json.NewEncoder(out).Encode(targets.Elements())
	}

	for target := range targets {
		if _, err := fmt.Fprintln(out, target); err != nil {
			return err
		}
	}
	return nil
}

func displayEdgeGraph(reply *gpb.EdgesReply) error {
	nodes := xrefs.NodesMap(reply.Nodes)
	esets := make(map[string]map[string]stringset.Set)

	for source, es := range reply.EdgeSets {
		for gKind, g := range es.Groups {
			for _, edge := range g.Edge {
				tgt := edge.TargetTicket
				src, kind := source, gKind
				if edges.IsReverse(kind) {
					src, kind, tgt = tgt, edges.Mirror(kind), src
				}
				groups, ok := esets[src]
				if !ok {
					groups = make(map[string]stringset.Set)
					esets[src] = groups
				}
				targets, ok := groups[kind]
				if ok {
					targets.Add(tgt)
				} else {
					groups[kind] = stringset.New(tgt)
				}

			}
		}
	}
	if _, err := fmt.Println("digraph kythe {"); err != nil {
		return err
	}
	for ticket, node := range nodes {
		if _, err := fmt.Printf(`	%q [label=<<table><tr><td colspan="2">%s</td></tr>`, ticket, html.EscapeString(ticket)); err != nil {
			return err
		}
		var factNames []string
		for fact := range node {
			if fact == facts.Code {
				continue
			}
			factNames = append(factNames, fact)
		}
		sort.Strings(factNames)
		for _, fact := range factNames {
			if _, err := fmt.Printf("<tr><td>%s</td><td>%s</td></tr>", html.EscapeString(fact), html.EscapeString(string(node[fact]))); err != nil {
				return err
			}
		}
		if _, err := fmt.Println("</table>> shape=plaintext];"); err != nil {
			return err
		}
	}
	if _, err := fmt.Println(); err != nil {
		return err
	}

	for src, groups := range esets {
		for kind, targets := range groups {
			for tgt := range targets {
				if _, err := fmt.Printf("\t%q -> %q [label=%q];\n", src, tgt, kind); err != nil {
					return err
				}
			}
		}
	}
	if _, err := fmt.Println("}"); err != nil {
		return err
	}
	return nil
}

func displayEdgeCounts(edges *gpb.EdgesReply) error {
	counts := make(map[string]int)
	for _, es := range edges.EdgeSets {
		for kind, g := range es.Groups {
			counts[kind] += len(g.Edge)
		}
	}

	if *displayJSON {
		return json.NewEncoder(out).Encode(counts)
	}

	for kind, cnt := range counts {
		if _, err := fmt.Fprintf(out, "%s\t%d\n", kind, cnt); err != nil {
			return err
		}
	}
	return nil
}

func displayNodes(nodes map[string]*cpb.NodeInfo) error {
	if *displayJSON {
		return json.NewEncoder(out).Encode(nodes)
	}

	for ticket, n := range nodes {
		if _, err := fmt.Fprintln(out, ticket); err != nil {
			return err
		}
		for name, value := range n.Facts {
			if len(value) <= factSizeThreshold {
				if _, err := fmt.Fprintf(out, "  %s\t%s\n", name, value); err != nil {
					return err
				}
			} else {
				if _, err := fmt.Fprintf(out, "  %s\n", name); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func displayDocumentation(reply *xpb.DocumentationReply) error {
	// TODO(zarko): Emit formatted data for -json=false.
	return json.NewEncoder(out).Encode(reply)
}

func factValue(m map[string]map[string][]byte, ticket, factName, def string) string {
	if n, ok := m[ticket]; ok {
		if val, ok := n[factName]; ok {
			return string(val)
		}
	}
	return def
}

func displayXRefs(reply *xpb.CrossReferencesReply) error {
	if *displayJSON {
		return json.NewEncoder(os.Stdout).Encode(reply)
	}

	for _, xr := range reply.CrossReferences {
		if _, err := fmt.Fprintln(out, "Cross-References for ", showSignature(xr.MarkedSource), xr.Ticket); err != nil {
			return err
		}
		if err := displayRelatedAnchors("Definitions", xr.Definition); err != nil {
			return err
		}
		if err := displayRelatedAnchors("Declarations", xr.Declaration); err != nil {
			return err
		}
		if err := displayRelatedAnchors("References", xr.Reference); err != nil {
			return err
		}
		if err := displayRelatedAnchors("Callers", xr.Caller); err != nil {
			return err
		}
		if len(xr.RelatedNode) > 0 {
			if _, err := fmt.Fprintln(out, "  Related Nodes:"); err != nil {
				return err
			}
			for _, n := range xr.RelatedNode {
				var nodeKind, subkind string
				if node, ok := reply.Nodes[n.Ticket]; ok {
					for name, value := range node.Facts {
						switch name {
						case facts.NodeKind:
							nodeKind = string(value)
						case facts.Subkind:
							subkind = string(value)
						}
					}
				}
				if nodeKind == "" {
					nodeKind = "UNKNOWN"
				} else if subkind != "" {
					nodeKind += "/" + subkind
				}
				if _, err := fmt.Fprintf(out, "    %s %s [%s]\n", n.Ticket, n.RelationKind, nodeKind); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func showPrintable(printable *xpb.Printable) string {
	if printable == nil {
		return "(nil)"
	}
	return printable.RawText
}

func showSignature(signature *xpb.MarkedSource) string {
	if signature == nil {
		return "(nil)"
	}
	return markedsource.Render(signature)
}

func displayRelatedAnchors(kind string, anchors []*xpb.CrossReferencesReply_RelatedAnchor) error {
	if len(anchors) > 0 {
		if _, err := fmt.Fprintf(out, "  %s:\n", kind); err != nil {
			return err
		}

		for _, a := range anchors {
			pURI, err := kytheuri.Parse(a.Anchor.Parent)
			if err != nil {
				return err
			}
			if _, err := fmt.Fprintf(out, "    %s\t%s\t[%d:%d-%d:%d)\n      %q\n",
				pURI.Path, showSignature(a.MarkedSource),
				a.Anchor.Span.Start.LineNumber, a.Anchor.Span.Start.ColumnOffset,
				a.Anchor.Span.End.LineNumber, a.Anchor.Span.End.ColumnOffset,
				string(a.Anchor.Snippet)); err != nil {
				return err
			}
			for _, site := range a.Site {
				if _, err := fmt.Fprintf(out, "      [%d:%d-%d-%d)\n        %q\n",
					site.Span.Start.LineNumber, site.Span.Start.ColumnOffset,
					site.Span.End.LineNumber, site.Span.End.ColumnOffset,
					string(site.Snippet)); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
