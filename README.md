# goroutines-phylogenomics-filter

- go phylogenomics routines.
- takes a alignment file, a header to write and merges all the alignment with the specified header and deletes indels

```
[gauravsablok@fedora]~/Desktop/codecreatede/goroutines-phylogenomics-filter% \
go run main.go -h
Usage of /tmp/go-build430267477/b001/exe/main:
  -alignmentfile string
        file (default "path to the alignment file")
  -title string
        gene name (default "name of the header of the gene"
[gauravsablok@fedora]~/Desktop/codecreatede/goroutines-phylogenomics-filter% \
    go run main.go -alignmentfile ./sample-files/test.fasta -title filteredphylo
> filteredphylo
 TCAGTATCTCTATCTC

```

Gaurav Sablok

