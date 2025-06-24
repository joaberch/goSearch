# Go-LocalSearchEngine
## The first input is either 1 or 2.
input '1': 
Select a folder, then a temporary index in XML and a 'permanent' gzip of it will be created in the same place as the exe. (both are overwritten if a new selection is made)
The temporary file will be used to make the research, the 'permanent' one to regenerate temporary file.

input '2':
Use the already existing gzip to generate a temporary file and use it's index.

## The second input is either '!select', '!help', !exit', '!displayTree' or any word
input '!select' :
Go back to the input '1' in the first input, select new folder to generate index.

input '!help' :
Display the command possible.

input '!exit' :
Close the application and delete the temp index.

input '!displayTree' :
Display the current tree structure in a user friendly way.
WARNING : The file may be showed in the wrong folder

input <any word> :
Make a research of this word in the index this is an inverted index, each word found get the file in which they are, example :
```
<entry word="salut">
    <file>C:\Users\px59nyu\Desktop\temp\goTest1\Nouveau Document texte (2).txt</file>
    <file>C:\Users\px59nyu\Desktop\temp\goTest1\Nouveau Document texte.txt</file>
    <file>C:\Users\px59nyu\Desktop\temp\goTest2\Nouveau Document texte.txt</file>
</entry>
```
