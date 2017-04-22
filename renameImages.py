import os
import hashlib

images = []
i = 0;
for root, subdirs, files in os.walk("images"):
    for f in files:
        images.append(os.path.join(root, f))

for i in images:
    f = open(i, 'rb')
    contents = f.read()
    f.close()
    newName = hashlib.md5(contents).hexdigest()
    path, base = os.path.split(i)
    _, ext = os.path.splitext(base)
    os.rename(i, os.path.join(path, newName + ext))
