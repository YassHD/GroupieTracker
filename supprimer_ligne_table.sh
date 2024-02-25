#!/bin/bash

function aide {
echo "supprimer_ligne_table.sh [-h] -t table -i id-ligne"
echo " "
echo "Supprimme la ligne de la table choisi à l'ID choisi"
echo " "
echo "options"
echo " "
echo " -t nom de la table "
echo " "
echo " -i ID de la ligne à supprimer"
echo " "
echo " -h affiche ce message d'aide"
}

TABLE=""
#variable de la table

ID=""
#variable de l'ID

while getopts "ht:i:" options
#Arguments de notre script Bash

do
	case $options in

	i)
		ID=$OPTARG
		#affecte l'argument i à la variable ID
		;;
	t)
		TABLE=$OPTARG
		#affecte l'argument t à la variable TABLE
		;;
	h)
	aide
	#lance la fonction aide
	exit 0
	;;

	esac
done

mysql -u root -h "127.0.0.1" -P "3306"   -D "groupie_tracker" << EOF
#Joint la base de données

DELETE FROM $TABLE WHERE ID=$ID;
EOF

if [ "$?" -eq 0 ]; then
	echo "La ligne de la base de données Groupie_Tracker a été supprimée"
else
	echo "La suppression de la ligne de la base de données a échouée"
fi