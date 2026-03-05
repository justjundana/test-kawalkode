#!/bin/sh
# Example usage of jsonfmt CLI
echo '{"foo":1,"bar":[2,3]}' > tmp.json
jsonfmt tmp.json
cat tmp.json
rm tmp.json
