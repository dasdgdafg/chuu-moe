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

f = open(navbar, 'r')
navContents = f.read()
f.close()

for p in pages:
    f = open(p, 'r')
    contents = f.read()
    f.close()
    contents = contents.replace("{{ put the nav-bar here }}", navContents)
    _, base = os.path.split(p)
    newFile = os.path.join("generated html", base)
    f = open(newFile, 'w')
    f.write(contents)
    f.close()
