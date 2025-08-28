# Go-LocalSearchEngine
A lightweight CLI search engine built in Go. Search words across files (recursively) from the current folder using an in-memory (xml save coming soon) inverted index.
## How to use
1. Clone the repository
```
git clone https://github.com/joaberch/Go-LocalSearchEngine.git
```
2. Build the project (a batch will do it in the near futur
```
go build
```
3. Add the exe to the system path to use it everywhere (a batch will do it in the near futur)
4. Enjoy!

## Example
Display the help :
```
gosearch help
```

Display the CLI version :
```
gosearch version
```

Search any word :
```
gosearch word
```
## Requirements
- Go 1.25 or higher
- Windows (tested on amd64)
