# groupie-tracker-visualizations

## Serveur Go

Groupie Tracker est un seveur en Go qui donne accès à une fonction.

Le programme ouvre le port TCP 8080.

Pour lancer ce serveur, il suffit d'ouvrir un terminal à la racine et d'entrer : 

`go run .`

***********************************************************************************************

## Groupie Tracker

Le serveur donne accès à l'API Groupie Tracker qui regroupe des données 

### Fonctionnement

Une fois le serveur lancé, depuis un navigateur, entrée l'adresse suivante : `http://localhost:8080/` dans la barre d'adresse.

Vous avez devant vous la page d'accueil du serveur où vous trouverez la liste des groupes à cliqué et une barre de recherche pour trouver l'artiste souhaité.

Une fois votre groupe sélectionné vous arriverez sur la page `http://localhost:8080/groupe` sur laquel vous trouverez toutes les informations relatives à ce groupe : `membres` `date de création du groupe` `date du premier album`, ainsi que les `locations` et `dates` de leurs concerts.
Une carte peut être affiché pour visualiser la location de concerts.



*****************************************************************************************************
## Erreurs

Le serveur prend en charge les erreurs : 
 - 404 : Mauvaise Adresse
 - 500 : Erreure Interne au Serveur

 *****************************************************************************************************
 ## Auteurs 
 - Fabien Olivier
 - Fabien Fanise