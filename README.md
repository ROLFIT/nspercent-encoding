# nspercent-encoding
Fix Non Standard Percent-encoding strings before url.QueryUnescape

1. Find "%uXXXX" 
2. Make rune
3. Encode rune by "%XX%XX"
