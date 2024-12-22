#!/bin/bash

# Check for correct number of arguments
if [ "$#" -lt 2 ]; then
  echo "Usage: $0 <count> <program> [args...]"
  exit 1
fi

# Read arguments
count=$1
shift
program="$@"

# Initialize timing variables
total_time=0
min_time=999999999
max_time=0

# Run the program the specified number of times
for ((i = 1; i <= count; i++)); do
  start=$(date +%s.%N)
  $program >/dev/null 2>&1
  end=$(date +%s.%N)

  elapsed=$(echo "$end - $start" | bc)

  # Accumulate total time
  total_time=$(echo "$total_time + $elapsed" | bc)

  # Update min and max times
  if (( $(echo "$elapsed < $min_time" | bc -l) )); then
    min_time=$elapsed
  fi

  if (( $(echo "$elapsed > $max_time" | bc -l) )); then
    max_time=$elapsed
  fi

done

# Calculate average time
average_time=$(echo "$total_time / $count" | bc -l)

# Format all times with leading zero and 9 decimal places
average_time=$(echo "$average_time" | awk '{printf "%.9f", $1}')
min_time=$(echo "$min_time" | awk '{printf "%.9f", $1}')
max_time=$(echo "$max_time" | awk '{printf "%.9f", $1}')

# Output results
echo "Average Time: $average_time seconds"
echo "Minimum Time: $min_time seconds"
echo "Maximum Time: $max_time seconds"
