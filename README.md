# Wpress-Extractor Windows/Mac

A simple windows app that allows you to extract .wpress files created by the awesome [All-in-one-Wp-Migration Wordpress plugin](https://wordpress.org/plugins/all-in-one-wp-migration/)

## Credits

The original extractor source code: [https://github.com/yani-/wpress](https://github.com/yani-/wpress). I have made a small modification to the reader.go file to take a output path as parameter.
The original Wpress-Extractor code: [https://github.com/fifthsegment/Wpress-Extractor](https://github.com/fifthsegment/Wpress-Extractor) Abdullah Irfan has changed the reader.go og the reader.go to run it on windows. And he has written the wpress-extractor.go.

## Download link

[Windows - Download now](https://github.com/mabakach/Wpress-Extractor/raw/master/dist/wpress-extractor.exe)  

[Mac - Download now](https://github.com/mabakach/Wpress-Extractor/blob/master/dist/mac/wpress-extractor?raw=true)  
*IMPORTANT FOR MAC: Don't forget to make the binary executable by running a  `chmod +x wpress-extractor` on the downloaded file via the Terminal.

## How to extract/open .wpress files?

Simply provide a path to your downloaded .wpress file as the first commandline argument to the program. Optionally you can provide a output path as a second parameter.  
Missing parts of the output directory are automatically created by wpress-extractor.

### Mac

Syntax: `./wpress-extractor </path/to/my/backup.wpress> [/output/path]`  
Example: `./wpress-extractor /User/test/Download/backup.wpress\ /User/test/Documents/Wordpress/backup`

### Windows

Syntax: `./wpress-extractor.exe /path/to/my/backup.wpress`  
Example: `wpress-extractor.exe C:\\temp\\backup.wpress C:\\temp\\Wordpress\\backup`
