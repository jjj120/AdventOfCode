#!/bin/bash

year=$(date '+%Y')
today=$(date '+%d')
wait=false

# Function to print usage
usage() {
    echo "Usage: $0 [-d|--day] <day> [-y|--year] <year> [-w|--wait] [-h|--help]"
    exit 1
}

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -d|--day)
            if [[ -n $2 ]]; then
                today="$2"
                shift 2
            else
                echo "Error: Argument for $1 is missing" >&2
                usage
            fi
            ;;
        -h|--help)
            usage
            ;;
        -w|--wait)
            wait=true
            shift
            ;;
        -y|--year)
            if [[ -n $2 ]]; then
                year="$2"
                shift 2
            else
                echo "Error: Argument for $1 is missing" >&2
                usage
            fi
            ;;
        *)
            echo "Error: Unknown parameter $1" >&2
            usage
            ;;
    esac
done

# Add leading zero if day is a single digit
if [[ "$day" =~ ^[0-9]$ ]]; then
    day=$(printf "%02d" "$day")
fi

echo "Creating directory $today and copying template files"

# Check if the directory already exists
if [ -d ./$today ]; then
    # Ask if the user wants to overwrite the existing directory
    read -p "Directory $today already exists. Do you want to overwrite it? (y/n) " -n 1 -r
    echo
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        rm -r ./$today
    else
        exit 1
    fi
fi


# Copy the template files to the new directory
cp -r ./template ./$today/

# Rename the template files
mv ./$today/tmpl.ex ./$today/$today.ex
mv ./$today/tmpl.go ./$today/$today.go
rm ./$today/tmpl.in

# Replace "tmpl" with the day in the copied go file
sed -i "s/tmpl/$today/g" ./$today/$today.go
sed -i "s/const day = 0/const day = $today/g" ./$today/$today.go

# Download the input file
if [ "$wait" = true ]; then
    echo "Waiting for the input file to be available..."
    aocdl -output "$today/{{printf \"%02d\" .Day}}.in" -wait
    if [ $? != 0 ]; then
        echo "Download failed!"
    fi
else
    aocdl -output "$today/{{printf \"%02d\" .Day}}.in" -day $today -year $year
    if [ $? != 0 ]; then
        echo "Download failed!"
    fi
fi
