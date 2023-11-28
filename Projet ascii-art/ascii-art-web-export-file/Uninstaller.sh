#!/bin/bash
# Stoppe le container nommé
echo "****************************************************************************************************"
echo "*************************************** Arrêt du container. ****************************************"
echo "****************************************************************************************************"
docker stop ascii_art_dockerize
echo "****************************************** Container arrêté.****************************************"
echo
# Supprime le container nommé
echo "****************************************************************************************************"
echo "************************************ Suppression du container. *************************************"
echo "****************************************************************************************************"
docker rm ascii_art_dockerize
echo "***************************************** Container supprimé.****************************************"
echo
# Supprime l'image Docker
echo "****************************************************************************************************"
echo "************************************* Suppression de l'image. **************************************"
echo "****************************************************************************************************"
docker image rm -f ascii_art_dockerize
echo "****************************************** Image supprimée.*****************************************"
echo
# Vidage du cache.
echo "****************************************************************************************************"
echo "***************************************** Vidage du cache. *****************************************"
echo "****************************************************************************************************"
docker image prune
echo "******************************************** Cache vidé. *******************************************"
echo
