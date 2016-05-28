# Wpress-Extractor Windows/Mac
A simple windows app that allows you to extract .wpress files created by the awesome All-in-one-Wp-Migration Wordpress plugin

## Credits
The extractor source code : [https://github.com/yani-/wpress](https://github.com/yani-/wpress). I had to make a tiny modification to their reader.go file to allow it to run on Windows systems.

## Download link
[Windows - Download now](https://github.com/fifthsegment/Wpress-Extractor/raw/master/dist/wpress-extractor.exe)

[Mac - Download now](https://github.com/fifthsegment/Wpress-Extractor/blob/master/dist/mac/wpress_extractor?raw=true)
*IMPORTANT FOR MAC: Don't forget to make the binary executable by running a  `chmod +x wpress_extractor` on the downloaded file via the Terminal.


## How to extract/open .wpress files ?
Simply provide a path to your downloaded .wpress file as the first commandline argument to the program.
`./wpress_extractor /path/to/my/backup.wpress`

## I'm not very technical - How to use this thing?
### Windows Instructions

Simply download the extractor then drop your.wpress file onto the executable (Wpress-extractor.exe). ([Thanks hughc](https://github.com/hughc)!)


OR



1. Download the extractor 
2. Create a directory where you wish your files to be extracted to
3. Copy the downloaded extractor to that directory
4. Copy your .wpress file to that directory as well
5. Open up a command prompt
6. CD into the directory you just created, let's say its C:\Wordpress-Backup. The command you'll run would be `cd C:\Wordpress-Backup`
7. Now run the following command `wpress-extractor <name-of-your.wpress file>`. For example my .wpress file was fifthsegment.wpress so the command I ran was `wpress-extractor fifthsegment.wpress`.
8. You'll find your files extracted into the same directory where the extractor was run. In my case it was `C:\Wordpress-Backup`


