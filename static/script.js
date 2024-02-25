//Fonction qui recherche les artistes par leur nom

function search_artists() {
    let input = document.getElementById('searchbar').value // récupère la valeur de la barre de recherche
    input = input.toLowerCase(); // met la valeur de la barre de recherche en minuscule
    let x = document.getElementsByTagName('H2'); // récupère les noms des artistes
    art = document.getElementsByClassName('groupie');

    for (i = 0; i < x.length; i++) { 
        if (!x[i].innerHTML.toLowerCase().includes(input)) { // si la valeur de la barre de recherche n'est pas dans le nom de l'artiste
            art[i].style.display = "none"; // on cache l'artiste
        }
        else {
            art[i].style.display = "list-item"; // sinon on l'affiche
        }
    }
}

//Fonction qui trie les artistes par ordre alphabétique ou anti-alphabétique dès qu'on clique sur le bouton "A-Z"

function sortListDir() {
    var list, i, switching, b, shouldSwitch, dir, switchcount = 0; 
    list = document.getElementById("id01"); 
    switching = true;
    dir = "asc";
    while (switching) { 
        switching = false; 
        b = list.getElementsByClassName("groupie");
        for (i = 0; i < (b.length - 1); i++) { 
            shouldSwitch = false; 
            if (dir == "asc") { 
                if (b[i].getElementsByTagName("h2")[0].innerHTML.toLowerCase() > b[i + 1].getElementsByTagName("h2")[0].innerHTML.toLowerCase()) {
                    shouldSwitch = true;
                    break;
                }
            } else if (dir == "desc") {
                if (b[i].getElementsByTagName("h2")[0].innerHTML.toLowerCase() < b[i + 1].getElementsByTagName("h2")[0].innerHTML.toLowerCase()) {
                    shouldSwitch = true;
                    break;
                }
            }
        }
        if (shouldSwitch) {
            b[i].parentNode.insertBefore(b[i + 1], b[i]); 
            switching = true;
            switchcount++;
        } else {
            if (switchcount == 0 && dir == "asc") {
                dir = "desc";
                switching = true;
            }
        }
    }
    var arrow = document.getElementById("A-Z"); 
    if (dir == "asc") {
        arrow.innerHTML = "Z-A";
    } else {
        arrow.innerHTML = "A-Z";
    }
}

function selectArtist(event){ // récupère le data name de la cible puis créer un URL qui dirige vers cet artiste
    const artistNom = event.currentTarget.getAttribute('data-name');
    console.log(`Selected artist: ${artistNom}`);
    window.location.href = `/artist?name=${artistNom}`;
}
