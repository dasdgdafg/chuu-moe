import os
import hashlib

pages = []
navbar = ''
i = 0;
for root, subdirs, files in os.walk("html template"):
    for f in files:
        if os.path.basename(f) == "navbar.html":
            navbar = os.path.join(root, f)
        else:
            pages.append(os.path.join(root, f))

f = open(navbar, 'rb')
navContents = f.read()
f.close()

for p in pages:
    f = open(p, 'rb')
    contents = f.read()
    f.close()
    contents = contents.replace(b"{{ put the nav-bar here }}", navContents)
    _, base = os.path.split(p)
    newFile = os.path.join("generated html", base)
    f = open(newFile, 'wb')
    f.write(contents)
    f.close()
