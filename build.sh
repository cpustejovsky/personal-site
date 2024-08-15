STATIC="./handlers/static/"
# sass --no-source-map "$STATIC"/styles.scss:handlers/static/styles.min.css --style compressed
touch "$STATIC"/resource.md
cat "$STATIC"/resources/technology.md > "$STATIC"/resources.md
echo "" >> "$STATIC"/resources.md
cat "$STATIC"/resources/theology.md >> "$STATIC"/resources.md
echo "" >> "$STATIC"/resources.md
cat "$STATIC"/resources/race_relations.md >> "$STATIC"/resources.md
echo "" >> "$STATIC"/resources.md
cat "$STATIC"/resources/mental_health.md >> "$STATIC"/resources.md
pandoc -s --toc "$STATIC"/resources.md -o "$STATIC"/resources.md
pandoc -f markdown -t html "$STATIC"/resources.md -o "$STATIC"/resources.html
