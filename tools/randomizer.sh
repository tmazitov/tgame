#!/bin/bash

# Define the grass tiles
declare -a grass_tiles=("0" "0" "0" "0" "0" "0" "0" "0" "0" "0" "0" "0" "0"  "1" "2" "3" "4")


# Read the tile map file
map_file="../maps/map1/ground_2"
map_content=$(cat "$map_file")

# Split the map content into an array of lines
IFS=$'\n' read -d '' -r -a map_lines <<< "$map_content"

# Iterate over each line in the map
for ((i=0; i<${#map_lines[@]}; i++)); do
	line=${map_lines[i]}
	new_line=""

	# Iterate over each character in the line
	for ((j=0; j<${#line}; j++)); do
		char=${line:j:1}

		# Check if the character is a space (empty tile)
		if [[ "$char" != " " ]]; then
			# Generate a random index to select a grass tile
			random_index=$((RANDOM % ${#grass_tiles[@]}))
			grass_tile=${grass_tiles[random_index]}
			new_line+="$grass_tile"
		else
			new_line+="$char"
		fi
	done

	# Replace the line in the map with the new line
	map_lines[i]=$new_line
done

# Join the modified lines back into a single string
new_map_content=$(IFS=$'\n'; echo "${map_lines[*]}")

# Write the modified map content back to the file
echo "$new_map_content" > "$map_file"
