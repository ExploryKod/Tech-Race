# HETIC Projet Final: “ Tech-Race “

<img src="https://img.shields.io/badge/golang-%5E1.22-blue">
<img src="https://img.shields.io/badge/python-3.10-yellow">

<div align="center">
   <img src="https://github.com/NasssDev/Tech-Race/assets/167258734/8022059e-d34b-422f-9010-bf8d8fdd7132" alt="image" width="300" height="200"/>
</div>

## Description du Projet

Le projet final HETIC, "Tech Race", vise à développer une application mobile permettant de contrôler un véhicule à distance via un réseau sans fil. L'application offre aux utilisateurs la possibilité de piloter le véhicule, d'accéder aux données de télémétrie en temps réel et de participer à des courses autonomes.

Ce projet se compose de plusieurs repositories toutes hébergés sur Github : 
- [API Tech Race](https://github.com/NasssDev/Tech-Race)
- [App Mobile](https://github.com/Hetic-Team/tech_race_8_2024)
- [Programme de la voiture](https://github.com/ExploryKod/freenove_esp32_wrover)
- [Modèle de la voiture](https://www.amazon.fr/Freenove-ESP32-WROVER-Contained-Compatible-Expressions/dp/B08X6PTQFM/ref=sr_1_5?__mk_fr_FR=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=1NFTVTE5M400B&dib=eyJ2IjoiMSJ9.ouyBflLDqHVkfViARMLD6Bn9gOI47kLGrM-5LMAbtJPAUgPogSQ1tQyH60VxNGSHTf-JIYDTkVL4RJ2a7-L92dQ5aqD8IliDd4MzLvffNmw65QxSItZh_qi-vPHXgzjBhvcW8Vy00EckrayFx_47OCj3W4K6Y1W0jHZgIDF7DAvRTI9XcC7oRK8T9xeUORe35q6RJ29TNUuhLCcN5fXl-WqLhsgNb2JA0XzHwnqwHaBBwj-xZ77ohEfVpUYfdyOMWf1wO01Fa42MzKl0b-UGD6PwYD-kBCJYQS3J9twWSGs.OrlAkZRIvlaYtQ2-9pywcADOLR7VY4iRx_9Ps1DkMnk&dib_tag=se&keywords=esp32+car&qid=1715602634&sprefix=esp+32+car,aps,125&sr=8-5)

## Equipe Backend (ce repository) : 

- [Amaury FRANSSEN](https://github.com/ExploryKod) 
- [Nassim AISSAOUI](https://github.com/NasssDev)
- [Justin LELUC](https://github.com/Jykiin)

## Equipe frontend : 
- [Reewaz Maskey](https://github.com/reewaz001)
- [Alexandre VISAGE](https://github.com/Aleex470)
- [Khalifa boubacar DIONE](https://github.com/khalifadione)
- [Achraf CHARDOUDI](https://github.com/Achkey)

Ce projet se compose de plusieurs repositories toutes hébergés sur Github :
- [API Tech Race](https://github.com/NasssDev/Tech-Race)
- [App Mobile](https://github.com/Hetic-Team/tech_race_8_2024)
- [Site web de partage des vidéos](https://site-a-venir)
- [Programme de la voiture](https://github.com/ExploryKod/freenove_esp32_wrover)
- [Modèle de la voiture](https://www.amazon.fr/Freenove-ESP32-WROVER-Contained-Compatible-Expressions/dp/B08X6PTQFM/ref=sr_1_5?__mk_fr_FR=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=1NFTVTE5M400B&dib=eyJ2IjoiMSJ9.ouyBflLDqHVkfViARMLD6Bn9gOI47kLGrM-5LMAbtJPAUgPogSQ1tQyH60VxNGSHTf-JIYDTkVL4RJ2a7-L92dQ5aqD8IliDd4MzLvffNmw65QxSItZh_qi-vPHXgzjBhvcW8Vy00EckrayFx_47OCj3W4K6Y1W0jHZgIDF7DAvRTI9XcC7oRK8T9xeUORe35q6RJ29TNUuhLCcN5fXl-WqLhsgNb2JA0XzHwnqwHaBBwj-xZ77ohEfVpUYfdyOMWf1wO01Fa42MzKl0b-UGD6PwYD-kBCJYQS3J9twWSGs.OrlAkZRIvlaYtQ2-9pywcADOLR7VY4iRx_9Ps1DkMnk&dib_tag=se&keywords=esp32+car&qid=1715602634&sprefix=esp+32+car,aps,125&sr=8-5)

## 🎯 Objectif

L'objectif principal est de créer une interface intuitive pour les pilotes afin de contrôler le véhicule et de visualiser ses performances. Deux niveaux de courses sont organisés : l'un où les pilotes contrôlent directement le véhicule depuis l'application mobile, et l'autre où le véhicule doit naviguer de manière autonome en suivant une ligne au sol.

## 💻 Fonctionnalités Principales

- Contrôle à distance du véhicule via l'application mobile.
- Visualisation en temps réel des données de télémétrie (vitesse, orientation, etc.).
- Mode suiveur de ligne autonome pour les courses sans intervention humaine.

## ⚙️ Configuration locale de l'API Backend (ce repository)

#### 🧮 Base générale

1. Cloner ce repository 
2. Avoir installé Golang sur sa machine locale [Installation de Go](https://go.dev/doc/install)
3. Avoir docker sur sa machine locale et avoir installé l'image docker de mosquitto : [Cliquez ici](https://github.com/ExploryKod/mosquitto-docker) 
4. Avoir postgresSQL (optionnel car on peux passer via docker)
5. Installer MakeFile si vous êtes sur une platforme qui ne l'a pas nativement 
6. Avoir configurer votre véhicule freenov avec un esp32 ayant le programme avec les bons IP liés à vos réseaux dessus <br> 
[voir la base de code ici](https://github.com/ExploryKod/freenove_esp32_wrover) 

#### 🎥 Pour faire fonctionner le service de vidéos
1. S'inscrire gratuitement sur Cloudinary : [S'inscrire sur Cloudinary](https://cloudinary.com/)
2. Récupérer son cloudinary ID et le cloudinary URL depuis son compte
3. Configurer le service cloudinarace (service cloudinary de Tech Race) :
   **Depuis la racine, se rendre dans `pkg/other/cloudinary/.env` :**

```
CLOUDINARY_ID=mon-id-cloudinary-présente-sur-mon-compte
CLOUDINARY_URL=mon-url-cloudinary-présente-sur-mon-compte
GOOS=linux (ou d'autres os si vous n'avez pas linux) > doc de Golang
GOARCH=amd64 (vérifier aussi que c'est bien la bonne architecture pour vous aussi) > doc de Golang
```

#### 📍Configurations du corps de l'API
1. Configurer le `.env` du projet à la racine :
```
IP=127.0.0.1
BOUNDARY=--123456789000000000000987654321
UPLOAD_VIDEO_URL=http://localhost:8083/upload-video
CLOUDINARY_URL=xxxx > votre url de cloudinary
CLOUDINARY_ID=xxxxx > votre id cloudinary
PortVideo=7000
ESP32_ADDRESS=192.168.104.10
ESP32_PORT=7000
RELAY_ADDRESS=:8080
STREAM_BOUNDARY=123456789000000000000987654321
DATABASE_URL=postgresql://root:xxxxx@xxxx/tech_race > demandez-nous l'url de la bdd 
```

2. 🛵 Aller à la racine du projet et lancer ces commandes :

**Usage de nos scripts si Linux (ou wsl) (Attention cela kill des ports sur votre machine)**<br><br>
> Vous devez pouvoir utiliser le mode `sudo` 
> sans Linux, travaillez dans un environnement virtuel ou via WSL si vous êtes sur Windows

***Il est nécessaire d'installer make si vous avez Windows. Il est natif sur Mac et Linux.***
- Lancer le script `./start_app.sh` va ouvrir les ports utilisés par l'API et lancer les container docker puis l'app (`go run`)
- Sur un autre terminal : `make cloudinarace` ou sans make :  `cd pkg/other/cloudinary && go run main/main.go --port=8083`

***Installation sous Windows de Make :***
- `make` : un utilitaire pour executer le script du Makefile.
 
  Sur Windows il n'est pas nativement installé
  > Si vous avez chocolatey : choco install make 
  Info si vous n'avez pas chocolatey : https://earthly.dev/blog/makefiles-on-windows/

  > Sinon explorez ceci: https://gnuwin32.sourceforge.net/packages/make.html

  > Vérifiez enfin la compatibilité du script avec Windows : préferez peut-être alors l'usage de `wsl` ou d'un ubuntu avec la commande `sudo` disponible.

**Manuellement :**
- Vérifier que les ports 1883 (mqtt), 8083 (cloudinarace), 9000, 8888 (pgAdmin), 8889 (adminer) et 5432 (postgres) sont libres.
- Lancer docker : `docker compose up -d`
- Installer les dépendances : `go mod download` et `go mod tidy`
- Lancer le projet : `go run cmd/api/main.go`
- Lancer cloudinarace pour la gestion vidéo (si besoin):
    ```bash
    cd pkg/other/cloudinary && go run main/main.go --port=8083`
    ```
**Pour travailler sur le style: tailwind**

- Pour travailler en phase de developpement :
```sh
    npm run watch
```
- Pour builder avant mise en prod :
```sh
    npm run build
```

3. 📊 Configurer la base de donnée:
- Se rendre sur [adminer](http://localhost:8089) ou [pgAdmin](http://localhost:8888)
- Regarder les credentials présents dans le fichier `docker-compose.yaml`
- Choisir PostgresSQL et utiliser ces credentials :

```
Bdd : tech_race  
User : root
Sur PgAdmin, default email : tech@race.com
Mot de passe : password
serveur : db 
```

## 🧰 Matériel et Logiciel du véhicule lié à cet API
=======
- Importer le dump de la base de donnée présent sur `dump/tech_race_bdd.sql`

### Matériel et capteurs présent sur le véhicule
- Architecture ESP32
- 4 roues motrices pilotées indépendamment
- Connexion sans fil wifi/bluetooth
- Capteur de distance
- Capteur de suivi de ligne
- Caméra embarquée

### Logiciel 
Voir le repository : [cliquez ici](https://github.com/ExploryKod/freenove_esp32_wrover)
- APIs de contrôle pour les roues et les capteurs
- Données de télémétrie fournies par les capteurs
- Communication en temps réel entre l'application mobile et le véhicule

## Contribuer

Les contributions sont les bienvenues ! Pour contribuer à ce projet, veuillez suivre ces étapes :
1. Forker le projet
2. Créer une branche pour votre fonctionnalité (`git checkout -b feature/AmazingFeature`)
3. Commiter vos modifications (`git commit -m 'Ajouter une fonctionnalité incroyable'`)
4. Pousser vers la branche (`git push origin feature/AmazingFeature`)
5. Ouvrir une Pull Request

## Licence

Ce projet est sous licence MIT. Pour plus d'informations, veuillez consulter le fichier [LICENSE](LICENSE).
