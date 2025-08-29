# Go-LocalSearchEngine
A lightweight CLI search engine built in Go. Search words across files (recursively) from the current folder using an in-memory (xml save coming soon) inverted index.
## How to use
1. Clone the repository
```
git clone https://github.com/joaberch/Go-LocalSearchEngine.git
```
2.   Start the setup batch with admin right, in case you don't want to or can't you can manually :<br>
2.1  Build the project
```
go build
```
2.2  Add the exe to the system path to use it everywhere

## Example
Display the help :
```
gosearch -h
gosearch --help
```

Display the CLI version :
```
gosearch -v
gosearch --version
```

Search any word :
```
gosearch word
```
## Requirements
- Go 1.25 or higher
- Windows (tested on amd64)
