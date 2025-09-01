# Go-LocalSearchEngine
A lightweight CLI search engine built in Go. Search words recursively across files from a specified folder using an in-memory or an XML inverted index.
<hr>
## Features
- Quick search for words in files and folders.<br>
- Work in-memory with optional XML persistance.<br>
- Registering custom indexes to search through huge amounts of data in a short time.<br>
- Recursive parsing of the file tree with filter<br>
- Simple CLI without dependencies<br>
<hr>
## How to use
1. Clone the repository
```
git clone https://github.com/joaberch/Go-LocalSearchEngine.git
cd Go-LocalSearchEngine
```
2.   Start the setup batch with admin right, in case you don't want to or can't you can manually do it:<br>
2.1  Build the project
```
cd goSearch
go build
```
2.2  Add the exe to the system path to use it everywhere
<hr>
## Usage
### Help & Version
```
gosearch -h  # Show help
gosearch -v  # Show version
```
### Search in current directory (and subfolders)
```
gosearch hello
```
### Save index to XML, in Desktop/utils/index/name.yml
```
gosearch -s            # Save index of current directory
gosearch -s path/to    # Save index of specified path
```
### Use saved index
```
gosearch -u hello temp  # Search for 'hello' in the index named 'temp.xml'
```
### Output Example
```
Found 2 file(s) for "hello":
  C:\Users\you\Documents\project\file1.txt
  C:\Users\you\Documents\project\notes.txt
```
<hr>
## Requirements
- Go 1.25 or higher<br>
- Windows (tested on amd64)<br>
- Admin rights (optional for setup script)
<hr>
## Roadmap
- [x] In-memory indexing<br>
- [x] XML index saving/loading<br>
- [ ] Refactor tags/flags<br>
- [ ] Stemming and lemmatization in Normalization<br>
- [ ] Regex search<br>
- [ ] Cross-platform support (for Linux and macOS)
<hr>
## Contributing
Pull requests are welcome!
<hr>
## License
This project is licensed under the MIT License.<br>
See LICENSE for more information.
