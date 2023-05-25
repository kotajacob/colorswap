# colorswap

Finds color codes from STDIN and replaces them with a new format:
```sh
$ echo 'rgb(155,112,255)' | colorswap -hex
#9b70ff
```

You can pass in a huge file intermingled with text, code, and colors. The
output (and detectable input) formats are:
```
hex:  #9b70ff
rgb:  rgb(155,112,255)
rgba: rgba(155,112,255,128)
vec3: vec3(0.607843,0.439216,1.000000)
vec4: vec4(0.607843,0.439216,0.500000)
```
Capitalization doesn't matter for hex inputs, and the shorthand form `#EEE` is
accepted. For the other formats, spaces are accepted after the commas and you
can use less precision in your vecs.

# Install
```
make all
sudo make install
```

# Uninstall
```
sudo make uninstall
```

# Author
Written and maintained by Dakota Walsh.\
Up-to-date sources can be found at https://git.sr.ht/~kota/colorswap/.

# License
Copyright 2023 Dakota Walsh\
GNU GPL version 3 or later, see LICENSE.
