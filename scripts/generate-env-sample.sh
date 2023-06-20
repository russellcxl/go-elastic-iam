source_file=".env"
destination_file=".env.sample"

if [ -f "$destination_file" ]; then
  rm "$destination_file"
fi

sed 's/=.*$/=/' "$source_file" > "$destination_file"