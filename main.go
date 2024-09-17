package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
 Author Gaurav Sablok
 Universitat Potsdam
 Date 2024-9-17


 Check out the complete alignmentGo package for dealing with all alignment functions.
*/

func main() {
	alignment := flag.String("alignmentfile", "path to the alignment file", "file")
	title := flag.String("title", "name of the header of the gene", "gene name")

	flag.Parse()

	type alignmentID struct {
		id string
	}

	type alignmentSeq struct {
		seq string
	}

	type alignmentFilterID struct {
		id string
	}

	type alignmentFilterSeq struct {
		seq string
	}

	fOpen, err := os.Open(*alignment)
	if err != nil {
		log.Fatal(err)
	}

	// file opening and storing into structs
	fRead := bufio.NewScanner(fOpen)
	alignmentReadID := []alignmentID{}
	alignmentReadSeq := []alignmentSeq{}
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignmentReadID = append(alignmentReadID, alignmentID{
				id: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignmentReadSeq = append(alignmentReadSeq, alignmentSeq{
				seq: string(line),
			})
		}
	}

	var seqFilter string

	for i := range alignmentReadSeq {
		for j := range alignmentReadSeq[i].seq {
			if string(alignmentReadSeq[i].seq[j]) == "-" {
				continue
			} else {
				seqFilter += string(alignmentReadSeq[i].seq[j])
			}
		}
	}

	fmt.Println(">", *title, "\n", seqFilter)
}
