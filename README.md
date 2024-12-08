# Advent of Code

These are my solutions for the [advent of code](https://adventofcode.com/) in go. I'm not a go developer, I did the advent to learn go, so don't expect good code. Some of the solutions are inspired by the ones from the [Advent of Code subreddit](https://www.reddit.com/r/adventofcode/), but I tried to do all of it myself. 

There are different branches for the differnt years, the current year is the default branch. 

## Usage

There is a folder for each day. The input file is the *.in file, the examples to test are in the *.ex files. 

To run the code, you need to have go installed. Then you can run the code with `go run <dayNumber>.go` in the day's folder.

To run the code for part 1, you have to checkout the right commit because part 2 is added on top and edits the code for part 1.

## Download script

There is a script to copy the template files for the day. You can run it with `./download.sh` to download the template files for the current day. You can also run it with `./download.sh -d <dayNumber>` or `./download.sh --day <dayNumber>` to download the input file for a specific day. To download the input file for a specific year, you can run `./download.sh -y <year> -d <dayNumber>` or `./download.sh --year <year> --day <dayNumber>`.

The download script needs a session cookie to download the input file. You can get the session cookie from your browser and save it in a file called `.aocdlconfig` as `{session=<sessionCookie>}`. The script also depends on `aocdl` which you can install from [github.com/GreenLightning/advent-of-code-downloader/](https://github.com/GreenLightning/advent-of-code-downloader/).

If you want to start a time to automatically download the input when it gets released, specify the `-w` or `--wait` flag. The script will then wait until the input is released and download it. If this option is used, all other options are ignored.
