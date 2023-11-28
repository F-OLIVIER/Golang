# Liste des commandes :
#   ls
#   ls Readme.md
#   ls package
#   ls -l
#   ls -l Readme.md
#   ls -l package
#   ls -l /usr/bin
#   ls -R
#   ls -a
#   ls -r
#   ls -t
#   ls -la
#   ls -l -t package
#   ls -lRt package
#   ls -l package -a audit.sh
#   ls -lR /dev///bus///usb/002/
#   ls -la /dev
#   ls -alRrt package
#   ls -
#   ls -l "Lien vers imageascii.txt"
#   ls -l "Lien vers Bureau"

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                                        ls                                            │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                                    ls <file name>                                    │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls Readme.md
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                                 ls <directory name>                                  │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls package
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                                       ls -l                                          │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -l
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                                 ls -l <file name>                                    │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -l Readme.md
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                              ls -l <directory name>                                  │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -l package
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                                  ls -l /usr/bin                                      │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -l /usr/bin
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                                       ls -R                                          │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -R
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                                       ls -a                                          │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -a
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                                       ls -r                                          │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -r
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                                       ls -t                                          │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -t
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                                       ls -la                                         │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -la
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                            ls -l -t <directory name>                                 │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -l -t package
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                            ls -lRr <directory name>                                  │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -lRr package
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                         ls -l <directory name> -a <file name>                        │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -l package -a audit.sh
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│ -lR <directory name>///<sub directory name>///<directory name>/<sub directory name>/ │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -lR /dev///bus///usb/002/
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                                    ls -la /dev                                       │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -la /dev
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                             ls -alRrt <directory name>                               │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -alRrt package
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                                        ls -                                          │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                                 ls -l <symlink file>/                                │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -l "Lien vers imageascii.txt"
echo

echo "┌──────────────────────────────────────────────────────────────────────────────────────┐"
echo "│                                 ls -l <symlink dir>/                                 │"
echo "└──────────────────────────────────────────────────────────────────────────────────────┘"
ls -l "Lien vers Bureau"
echo
