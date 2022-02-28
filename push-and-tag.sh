cd /root/goweb
git add .
git commit -m $1 && git tag $2 -m $1 && git push origin --tags