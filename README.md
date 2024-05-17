# markdown rendering programs

based on the markdown parser from
[link](https://github.com/gomarkdown/markdown)

## md2js
A renderer that parses a md file and converts the resulting ast tree into js objects. The objects can be parsed to create a browser DOM structure with the help of azulLib.js.

## mdjsYaml
A program that parses a modifies markdown file. The program first check whether there is a md file contains a yaml section and then parses the remaining part as a markdown section.
The yaml section is encapsulated by lines with 4 or more = chars ('====').

## mdjson
A renderer that parses a md file and creates a json object. Still needs work!

## mdImg
A program that parses a markdown file and creates a yaml file listing all image references contained in the markdown file.


