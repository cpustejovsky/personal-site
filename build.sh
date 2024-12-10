STATIC="./handlers/static"
# sass --no-source-map "$STATIC"/styles.scss:handlers/static/styles.min.css --style compressed
touch "$STATIC"/resources.md
touch "$STATIC"/temp.md
cat "$STATIC"/resources/technology.md > "$STATIC"/temp.md
echo "" >> "$STATIC"/temp.md
cat "$STATIC"/resources/theology.md >> "$STATIC"/temp.md
echo "" >> "$STATIC"/temp.md
cat "$STATIC"/resources/race_relations.md >> "$STATIC"/temp.md
echo "" >> "$STATIC"/temp.md
cat "$STATIC"/resources/mental_health.md >> "$STATIC"/temp.md
cat "$STATIC"/temp.md | sh toc.sh 1 3 > "$STATIC"/toc.md
cat "$STATIC"/toc.md > "$STATIC"/resources.md
cat "$STATIC"/temp.md >> "$STATIC"/resources.md
rm "$STATIC"/temp.md
