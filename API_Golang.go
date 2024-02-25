package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// const (
// 		baseURL           = "http://localhost/3636/"
// 		endpointArtist    = "Artist/"
// 		endpointDates     = "dates/"
// 		endpointLocations = "Locations/"
// 	)

var db *sql.DB
var err error

type Relations struct {
	IDartist    int `json:"IDartist"`
	Idlocations int `json:"Idlocations"`
	Iddates     int `json:"Iddates"`
}

type Locations struct {
	ID  int    `json:"ID"`
	Nom string `json:"Nom"`
}

type Artist struct {
	ID           int    `json:"ID"`
	Nom          string `json:"Nom"`
	Membres      string `json:"Membres"`
	CreationDate int    `json:"CreationDate"`
	FirstAlbum   int    `json:"FirstAlbum"`
	Image        string `json:"Image"`
	Locations    string `json:"Locations"`
	Date         string `json:"Date"`
}

type Dates struct {
	ID          int    `json:"ID"`
	Lastconcert string `json:"Lastconcert"`
}

func openBase() {
	//Ouvre la connexion à la base de données
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/groupie_tracker")
	if err != nil {
		fmt.Println(err)
	}
}

// func associateRelation() {
// 	for i := 0; i < 14; i++ {
// 		db.Exec("INSERT INTO relation(IDartist,Idlocations,Iddates) VALUES(?,?,?)", i, i, i)

// 	}
// }

func specificArtist(id string) Artist {
	//requête sur la table Artiste à un ID donné
	rows, err := db.Query("SELECT *FROM artist WHERE ID=(?);", id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	thisArtist := Artist{}
	for rows.Next() {
		err = rows.Scan(&thisArtist.ID, &thisArtist.Nom, &thisArtist.Membres, &thisArtist.CreationDate, &thisArtist.FirstAlbum, &thisArtist.Image, &thisArtist.Locations, &thisArtist.Date)
		if err != nil {
			panic(err.Error())
		}
	}
	return thisArtist
}

func specificLocations(id string) Locations {
	//requête sur la table Locations à un ID donné
	rows, err := db.Query("SELECT *FROM locations WHERE ID=(?);", id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	thisLocation := Locations{}
	for rows.Next() {
		err = rows.Scan(&thisLocation.ID, &thisLocation.Nom)
		if err != nil {
			panic(err.Error())
		}
	}
	return thisLocation
}

func specificDates(id string) Dates {
	//requête sur la table Dates à un ID donné
	rows, err := db.Query("SELECT *FROM dates WHERE ID=(?);", id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	thisdate := Dates{}
	for rows.Next() {
		err = rows.Scan(&thisdate.ID, &thisdate.Lastconcert)
		if err != nil {
			panic(err.Error())
		}
	}
	return thisdate
}

func selectallArtist() []Artist {
	//requête sur la table Artiste
	var allArtist []Artist
	rows, err := db.Query("SELECT *FROM Artist  ")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	// Parcourt les lignes retournées
	for rows.Next() {
		recupArtist := Artist{}
		err = rows.Scan(&recupArtist.ID, &recupArtist.Nom, &recupArtist.Membres, &recupArtist.CreationDate, &recupArtist.FirstAlbum, &recupArtist.Image, &recupArtist.Locations, &recupArtist.Date)
		if err != nil {
			panic(err.Error())
		}
		allArtist = append(allArtist, recupArtist) //ajoute les données des artistes dans le tableau allArtist

	}
	return allArtist
}

func selectallLocations() []Locations {
	//requête sur la table Locations
	var allLocations []Locations
	rows, err := db.Query("SELECT *FROM locations ")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	// Parcourt les lignes retournées
	for rows.Next() {
		recupLocations := Locations{}
		err = rows.Scan(&recupLocations.ID, &recupLocations.Nom)
		if err != nil {
			panic(err.Error())
		}
		allLocations = append(allLocations, recupLocations) //ajoute les données des lieux dans le tableau allLocations

	}
	return allLocations
}

func selectallRelations() []Relations {
	//requête sur la table Relations
	var allRelations []Relations
	rows, err := db.Query("SELECT *FROM relation ")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	// Parcourt les lignes retournées
	for rows.Next() {
		recupRelations := Relations{}
		err = rows.Scan(&recupRelations.IDartist, &recupRelations.Idlocations, &recupRelations.Iddates)
		if err != nil {
			panic(err.Error())
		}
		allRelations = append(allRelations, recupRelations) //ajoute les données des lieux dans le tableau allrelations
	}
	return allRelations
}

//displayRelations(){

//}

func selectallDates() []Dates {
	//requête sur la table Dates
	var allDates []Dates
	rows, err := db.Query("SELECT *FROM dates ")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	// Parcourt les lignes retournées
	for rows.Next() {
		recupdates := Dates{}
		err = rows.Scan(&recupdates.ID, &recupdates.Lastconcert)
		if err != nil {
			panic(err.Error())
		}
		allDates = append(allDates, recupdates) //ajoute les données des dates dans le tableau allDates
	}
	return allDates
}

func ArtistparNom(nom string) (*Artist, error) { //compare le nom passé en argument et le nom de la variable artist
	for _, artist := range selectallArtist() {
		if artist.Nom == nom {
			return &artist, nil
		}
	}
	return nil, fmt.Errorf("artiste non trouvé")
}

func DatesparID(id int) (*Dates, error) { //compare le nom passé en argument et le nom de la variable dates
	for _, dates := range selectallDates() {
		if dates.ID == id {
			return &dates, nil
		}
	}
	return nil, fmt.Errorf("dates non trouvé")
}

func LocationsparID(id int) (*Locations, error) { //compare le nom passé en argument et le nom de la variable locations
	for _, locations := range selectallLocations() {
		if locations.ID == id {
			return &locations, nil
		}
	}
	return nil, fmt.Errorf("location non trouvé")
}

func main() {

	router := gin.Default()
	// Démarre le serveur

	router.Static("/static", "./static") // charge le css et le script js
	router.LoadHTMLGlob("tmpl/*.html")   // charge les pages HTML

	openBase()                  // ouvre la connexion a la base de données
	artist := selectallArtist() // récupère tous les artistes
	//Retourne l'accueuil du site
	router.GET("/index.html", func(c *gin.Context) {
		// membresListe := strings.Split(artist.Membres, " ")
		// numMembres := len(membresListe)
		c.HTML(http.StatusOK, "index.html", gin.H{"Artist": artist})
	})

	router.GET("/artist", func(c *gin.Context) {
		artistNom := c.Query("name")           // récupère le nom de l'artiste dans l'URL
		artist, err := ArtistparNom(artistNom) //Récupère les valeurs de l'artiste choisi
		if err != nil {
			c.String(http.StatusNotFound, "Artiste non trouvé") //erreur si pas d'artiste correspondant au nom
			return
		}

		dates, err := DatesparID(artist.ID)
		if err != nil {
			c.String(http.StatusNotFound, "Date non trouvé")
			return
		}
		locations, err := LocationsparID(artist.ID)
		if err != nil {
			c.String(http.StatusNotFound, "location non trouvé")
			return
		}
		locationsListe := strings.Split(locations.Nom, ",")   //créer un liste des locations
		lastlocation := locationsListe[len(locationsListe)-1] // récupère la dernière location
		membresListe := strings.Split(artist.Membres, " ")    //créer une liste des membres
		var essai []string                                    //tableau Nom Prénom
		var membre string                                     //variable Nom Prénom
		if len(membresListe) <= 2 {
			essai = append(essai, membresListe[0]+" "+membresListe[1])
		} else {
			for i := 2; i < len(membresListe)+1; i += 2 {
				membre = membresListe[i-2] + " " + membresListe[i-1] // associe nom prénom
				essai = append(essai, membre)                        // créer une liste des noms prénoms
			}
		}

		c.HTML(http.StatusOK, "info.html", gin.H{"artistNom": artist.Nom, "artistimage": artist.Image, "Lmembres": essai, "artistcreation": artist.CreationDate, "datelastconcert": dates.Lastconcert, "locationlastconcert": lastlocation})
	})

	//Retourne l'intégralité des artistes à l'adresse /Artist en JSON
	router.GET("Artist/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": selectallArtist()})
	})
	//Retourne l'intégralité des locations à l'adresse /Locations en JSON
	router.GET("Locations/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": selectallLocations()})
	})

	//Retourne un artiste spécifique en fonction de l'ID en JSON
	router.GET("Artist/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		c.JSON(http.StatusOK, gin.H{"message": specificArtist(id)})
	})
	//Retourne une location spécifique en fonction de l'ID en JSON
	router.GET("Locations/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		c.JSON(http.StatusOK, gin.H{"message": specificLocations(id)})
	})
	//Retourne l'intégralité des relations à l'adresse /Relation en JSON
	router.GET("Relation/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": selectallRelations()})
	})
	//Retourne l'intégralité des dates à l'adresse /Dates en JSON
	router.GET("Dates/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": selectallDates()})
	})
	//Retourne une date spécifique en fonction de l'ID en JSON
	router.GET("Dates/:ID", func(c *gin.Context) {
		id := c.Param("ID")
		c.JSON(http.StatusOK, gin.H{"message": specificDates(id)})
	})

	router.Run(":3636") //ListenAndServe(":3636")

}
