# Text To ASCII Art

*The ASCII Art Generator is a program written in Go. This program takes text input and converts it into a graphical representation using ASCII characters based on selected banner styles. Users can choose from various banner styles and create the text in different graphical designs.*

## Features  
- Convert text inputs to ASCII art. 
-  Select from different banner styles. 
-  Adjust text alignment. 
-  Save ASCII art to a file.

### Banner Styles
`shadow`: Stylish text with a shadow effect. 
`standard`: Standard ASCII art representation.

## How It Works
1.  **Command-Line Arguments**: The program takes command-line arguments to specify the input text, banner style, output file (optional), and alignment (optional).
    
2.  **Parsing Command-Line Arguments**: It parses the command-line arguments to extract the input text, banner style, output file name (if specified), and alignment (if specified).
    
3.  **Reading Banner Files**: The program reads the ASCII art templates from separate text files corresponding to the chosen banner style.
    
4.  **Generating ASCII Art**: Based on the input text and banner style, the program generates ASCII art by matching each character in the text with its corresponding ASCII representation from the banner file.
    
5.  **Text Alignment**: If alignment is specified, the program aligns the generated ASCII art accordingly (left, right, center, or justify).
    
6.  **Output Handling**: It outputs the generated ASCII art to the console or saves it to a file if an output file name is specified.
    
7.  **Terminal Size Detection**: The program automatically detects the terminal size to adjust the alignment if necessary.
    
8.  **Error Handling**: It handles errors gracefully, such as missing input text, invalid banner styles, or failed file operations.

## Usage
```bash
$ go run . [OPTIONS] [TEXT] [BANNER]
```
Example 1:
```bash
$ go run . "test" | cat -e
```
Result 1:
```
 _                  _    $
| |                | |   $
| |_    ___   ___  | |_  $
| __|  / _ \ / __| | __| $
\ |_  |  __/ \__ \ \ |_  $
 \__|  \___| |___/  \__| $
                         $
                         $
```
Example 2:
```bash
$ go run . "hello\nthere" shadow | cat -e
```
Result 2:
```
                                 $
_|                _| _|          $
_|_|_|     _|_|   _| _|   _|_|   $
_|    _| _|_|_|_| _| _| _|    _| $
_|    _| _|       _| _| _|    _| $
_|    _|   _|_|_| _| _|   _|_|   $
                                 $
                                 $
                                             $
  _|     _|                                  $
_|_|_|_| _|_|_|     _|_|   _|  _|_|   _|_|   $
  _|     _|    _| _|_|_|_| _|_|     _|_|_|_| $
  _|     _|    _| _|       _|       _|       $
    _|_| _|    _|   _|_|_| _|         _|_|_| $
                                             $
                                             $
```
Example 3:
```bash
$ go run . --output=result.txt "HeLL0" standard
```
Result 3:
```
$ cat result.txt -e
 _    _          _        _               $
| |  | |        | |      | |        ___   $
| |__| |   ___  | |      | |       / _ \  $
|  __  |  / _ \ | |      | |      | | | | $
| |  | | |  __/ | |____  | |____  | |_| | $
|_|  |_|  \___| |______| |______|  \___/  $
                                          $
                                          $
$
```
Example 4:
```bash
$ go run . --align=center "Hello There" standard
```
Result 4:
(based on your terminal size)
```
                                            _    _          _   _                 _______   _                                                                      
                                           | |  | |        | | | |               |__   __| | |                                                                     
                                           | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___                                             
                                           |  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \                                            
                                           | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/                                            
                                           |_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|                                            
                                                                                                                                                                   
                                                                                                                                                                   
```
Example 5:
```bash
$ go run . --align=right "Hello There" standard
```
Result 5:
(based on your terminal size)



                                                                                                                                 _    _          _   _                 _______   _
                                                                                                                                | |  | |        | | | |               |__   __| | |
                                                                                                                                | |__| |   ___  | | | |   ___            | |    | |__     ___   _ __    ___
                                                                                                                                |  __  |  / _ \ | | | |  / _ \           | |    |  _ \   / _ \ | '__|  / _ \
                                                                                                                                | |  | | |  __/ | | | | | (_) |          | |    | | | | |  __/ | |    |  __/
                                                                                                                                |_|  |_|  \___| |_| |_|  \___/           |_|    |_| |_|  \___| |_|     \___|




---
#### Author
[Ahat Demirezen](https://github.com/ahatdemirezen)

#### LinkedIn
[Ahat Demirezen](https://www.linkedin.com/in/ahat-demirezen-8bb5262a9/)

 ## License
 [MIT](https://choosealicense.com/licenses/mit/)