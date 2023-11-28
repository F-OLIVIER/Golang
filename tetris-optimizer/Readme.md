# Tetris Optimizer

### Descriptif
_______
Le programme permet à partir de tetromino (piéce de Tetris) de remplir automatiquement une grille contenant le moins d'espace possible<br>
Pour indiqué au programme les tetrominos à utiliser, il faut crée un fichier `.txt` qui sera placé dans le dossier `Fichier_Tetrominos` est les tetrominos seront identifier dans un quadrillage de 4 par 4 par des `#`, les espaces vide étant identifié par des `.`

Liste des 5 tetrominos possible<br>
<img src="./readme/tetrominos.png">


Position possible des tetrominos peuvant être utilisé<br>
<img src="./readme/alltetrominos.png">

Exemple de format de tetrominos pour crée un fichier `.txt` :
```
....
.##.
.##.
....
```
```
##..
.#..
.#..
....
```

### Example of result
_______
<table align= "center">
    <thead>
        <th align= "center">Exemple 1</th>
        <th align= "center">Exemple 2</th>
    </thead>
    <tbody>
        <tr>
            <td><img src="./readme/exemple2.png"></td>
            <td><img src="./readme/exemple2.png"></td>
        </tr>    
    </tbody>
</table>

### Usage
_______
```go
go run . "nom du fichier"
```

### Authors
_______
+ Fabien Olivier
