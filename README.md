# colorswap

Reads from STDIN, looking for HEX color codes or RGB color codes. It will swap
all occurences in each line to the other type. It prints the modified files to
STDOUT.

```sh
colorswap < sample.txt
```

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
Written and maintained by Dakota Walsh.
Up-to-date sources can be found at https://git.sr.ht/~kota/colorswap/

# License
GNU GPL version 3 or later, see LICENSE.
Copyright 2022 Dakota Walsh
