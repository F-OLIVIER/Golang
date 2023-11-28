# Forum

### Descriptif
_______
Ce programme est un outil simple de complétion/édition/autocorrection de texte.

- occurrence de `(hex)` doit remplacer le mot précédent par la version décimale du mot (dans ce cas, le mot sera toujours un nombre hexadécimal). (Ex : "1E (hex) files were added" -> "30 files were added")
- occurrence de `(bin)` doit remplacer le mot précédent par la version décimale du mot (dans ce cas, le mot sera toujours un nombre binaire). (Ex : "It has been 10 (bin) years" -> "It has been 2 years")
- occurrence de `(up)` convertit le mot précédent avec sa version en majuscules. (Ex : "Ready, set, go (up) !" -> "Ready, set, GO !")
- occurrence de `(low)` convertit le mot précédent avec la version minuscule de celui-ci. (Ex : "I should stop SHOUTING (low)" -> "I should stop shouting")
- occurrence de `(cap)` convertit le mot précédent avec sa version en majuscules. (Ex : "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge")

- Pour `(low)`, `(up)`, `(cap)` si un nombre apparaît à côté, comme par exemple : `(low, <number>)`, le nombre de mots spécifié précédemment sera mis en minuscules, en majuscules ou en capitales. (Ex : "This is so exciting (up, 2)" -> "This is SO EXCITING")

- Toutes les ponctuations `.`, `,`, `!`, `?`, `:` et `;` sont mise proches du mot précédent et espacées du suivant. (Ex : "J'étais assis là-bas ,et puis BAMM ! !!" -> "J'étais assis là, et puis BAMM ! !!").
  - Sauf s'il y a des groupes de ponctuation comme : `...` ou `!?`. Dans ce cas, le programme formate le texte comme dans l'exemple suivant : "Je pensais que... Vous aviez raison" -> "Je pensais... Vous aviez raison".
- Le signe de ponctuation `'` sera toujours trouvé avec une autre instance de celui-ci et ils doivent être placés à droite et à gauche du mot au milieu d'eux, sans aucun espace. (Ex : "I am exactly how they describe me : ' awesome '" -> "I am exactly how they describe me : 'awesome'")
  - S'il y a plus d'un mot entre les deux marques `' '`, le programme doit placer les marques à côté des mots correspondants (Ex : "As Elton John said : ' I am the most well-known homosexual in the world '" -> "As Elton John said : 'I am the most well-known homosexual in the world'").
- Chaque occurrence de `a` est transformée en `an` si le mot suivant commence par une voyelle (`a`, `e`, `i`, `o`, `u`) ou un `h`. (Ex : "There it was. A amazing rock !" -> "There it was. An amazing rock !").

### Usage
______
Le fichier `sample.txt` contient le texte à modifier, après exécution, le fichier `result.txt` contient la sortie du programme.
```go
go run . sample.txt result.txt
```

### Dépendence
_______
Le serveur est en Go version 1.20

### Authors
_______
+ Fabien OLIVIER

