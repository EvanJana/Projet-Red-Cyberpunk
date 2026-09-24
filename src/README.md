# Red Cyberpunk

RPG en ligne de commande developpe en Go, dans un univers cyberpunk. Le joueur cree un personnage, affronte des ennemis au tour par tour, gagne de l'experience, recupere des ressources, apprend des programmes et ameliore son equipement.

## Sommaire

- [A propos du jeu](#a-propos-du-jeu)
- [Fonctionnalites](#fonctionnalites)
- [Installation](#installation)
- [Lancer le jeu](#lancer-le-jeu)
- [Premiere partie](#premiere-partie)
- [Creation du personnage](#creation-du-personnage)
- [Menu principal](#menu-principal)
- [Systeme de combat](#systeme-de-combat)
- [Classes](#classes)
- [Competences et RAM](#competences-et-ram)
- [Ennemis](#ennemis)
- [Experience et niveaux](#experience-et-niveaux)
- [Inventaire et objets](#inventaire-et-objets)
- [Marchand](#marchand)
- [Cyberforgeron](#cyberforgeron)
- [Equipement](#equipement)
- [Mort et resurrection](#mort-et-resurrection)
- [Architecture du code](#architecture-du-code)
- [Limites actuelles](#limites-actuelles)
- [Pistes d'evolution](#pistes-devolution)

## A propos du jeu

Red Cyberpunk est un jeu solo textuel. Toutes les interactions se font avec le clavier dans le terminal. Il n'y a actuellement ni interface graphique, ni sauvegarde, ni serveur de jeu.

L'objectif est de progresser en combattant, de collecter des materiaux et de construire un personnage suffisamment puissant pour affronter Adam Smasher, le boss qui apparait lors de chaque cinquieme rencontre aleatoire.

## Fonctionnalites

- Creation d'un personnage avec nom et classe.
- Trois classes jouables avec des valeurs de PV et de RAM differentes.
- Combats au tour par tour contre des ennemis normaux, un robot d'entrainement et un boss.
- Attaques consommant de la RAM.
- Experience, montee de niveau et amelioration automatique des statistiques.
- Gain d'argent et drops aleatoires apres les combats.
- Inventaire avec capacite limitee et trois ameliorations possibles.
- Achat de soins, programmes, objets et equipements.
- Apprentissage de programmes via l'ordinateur.
- Fabrication d'equipements a partir de materiaux.
- Equipement de trois emplacements : torse, bottes et gants.
- Systeme de mort avec choix de resurrection ou de fin de partie.

## Installation

### Prerequis

- Go installe, de preference une version recente.
- Un terminal capable d'accepter les saisies clavier.

Verifier l'installation de Go :

```powershell
go version
```

### Recuperer le projet

Depuis le dossier du projet :

```powershell
git clone <URL_DU_DEPOT>
cd Projet-Red-Cyberpunk\src
```

Si le projet est deja present localement :

```powershell
cd c:\Users\tehau\Projet-Red-Cyberpunk\src
```

Le projet ne contient pas encore de fichier `go.mod`. Les fichiers appartiennent tous au package `main` et peuvent donc etre executes directement ensemble.

## Lancer le jeu

Depuis le dossier qui contient les fichiers `.go` :

```powershell
go run *.go
```

Pour compiler un executable Windows :

```powershell
go build -o red-cyberpunk.exe *.go
```

Puis lancer le jeu :

```powershell
.\red-cyberpunk.exe
```

Le programme attend les choix dans le terminal. Entrez un numero lorsque le menu affiche une liste d'actions.

## Premiere partie

1. Lancez le programme.
2. Saisissez un nom compose uniquement de lettres.
3. Choisissez `1`, `2` ou `3` pour selectionner une classe.
4. Consultez la fiche du personnage.
5. Choisissez `4. Combattre` dans le menu principal.
6. Utilisez l'entrainement pour decouvrir le systeme de combat.
7. Utilisez les recompenses pour acheter des objets ou fabriquer de l'equipement.
8. Affrontez les rencontres aleatoires pour gagner de l'XP et progresser.

Le nom est normalise : il est converti en minuscules, puis sa premiere lettre est mise en majuscule.

## Creation du personnage

Le personnage commence avec :

- Niveau 1.
- 0 XP.
- 100 pieces.
- La competence `Coup de Poing`.
- 10 places d'inventaire.
- Aucun equipement.
- Les PV actuels initialises a la moitie des PV maximum, puis remplis au lancement de la partie.

Les choix de classe acceptent le numero ou le nom de la classe, sans distinction entre majuscules et minuscules.

## Menu principal

| Choix | Action |
| --- | --- |
| 1 | Afficher les informations du personnage et les competences |
| 2 | Ouvrir l'inventaire |
| 3 | Entrer chez le cyberforgeron |
| 4 | Choisir un combat |
| 5 | Quitter le jeu |

Le menu de combat propose :

- `Entrainement` : combat contre un robot fixe.
- `Combat aleatoire` : rencontre avec un ennemi normal ou Adam Smasher.

## Systeme de combat

Les combats sont au tour par tour. L'ordre est determine par l'initiative : le joueur commence si son initiative est superieure ou egale a celle de l'ennemi.

Pendant son tour, le joueur peut :

1. Attaquer avec une competence apprise.
2. Ouvrir l'inventaire pour utiliser un objet ou consulter l'ordinateur.
3. Fuir le combat.

Apres un tour complet, le joueur recupere 10 points de RAM, sans depasser sa RAM maximale.

### Degats du joueur

Les degats d'une competence suivent cette formule :

```text
degats = degats_de_base + (niveau - 1) * 5
```

Les degats sont appliques aux PV de l'ennemi, sans descendre sous zero.

### Victoire et recompenses

Lors d'une victoire contre un ennemi aleatoire :

- Le joueur gagne l'XP de l'ennemi.
- Le joueur gagne 20 pieces.
- Le joueur peut recevoir un drop.
- Un drop `dollars` rapporte 40 pieces supplementaires.
- Un autre drop est ajoute a l'inventaire.

Une fuite ne donne pas de recompense. Un personnage tombe a terre lorsque ses PV atteignent zero.

## Classes

| Classe | PV maximum | RAM maximum | Initiative | Style conseille |
| --- | ---: | ---: | ---: | --- |
| Netrunner | 100 | 125 | 2 | Utiliser frequemment les programmes |
| Assassin | 75 | 75 | 2 | Jouer rapidement et economiser ses ressources |
| Berserk | 125 | 50 | 2 | Encaisser les coups et privilegier les degats directs |

Toutes les classes commencent avec `Coup de Poing`.

## Competences et RAM

| Competence | Degats de base | Cout en RAM | Deblocage |
| --- | ---: | ---: | --- |
| Coup de Poing | 10 | 0 | Disponible au debut |
| Surcharge | 15 | 15 | Programme achete au marchand |
| Crash | 20 | 25 | Programme achete au marchand |
| Suicide | 250 | 80 | Programme achete au marchand |

Pour apprendre une competence, achetez le programme correspondant, ouvrez l'ordinateur depuis l'inventaire, puis selectionnez-le. Le programme est consomme apres apprentissage.

## Ennemis

### Robot d'entrainement

| PV | Degats | XP | Recompense |
| ---: | ---: | ---: | --- |
| 40 | 5 | 0 | 40 pieces |

### Ennemis normaux

| Ennemi | PV | Degats | XP | Initiative | Particularite |
| --- | ---: | ---: | ---: | ---: | --- |
| Cyberpsycho | 100 | 5 | 45 | 1 | Degats doubles tous les 3 tours |
| Punk brutal de quartier | 120 | 8 | 40 | 0 | Attaque au corps a corps |
| Punk tireur de quartier | 55 | 12 | 40 | 0 | Tir a distance |
| Punk junkie de quartier | 70 | 6 | 40 | 5 | Peut s'assommer, ou doubler ses degats un tour sur deux |

Les ennemis normaux sont choisis aleatoirement. Leur drop est choisi aleatoirement selon les probabilites suivantes :

- `Acier` : 30 %.
- `Kevlar` : 30 %.
- `Cuir` : 20 %.
- `Puce neuronale` : 20 %.

### Adam Smasher

Adam Smasher apparait a chaque cinquieme rencontre aleatoire.

| PV | Degats de base | XP | Initiative | Particularite |
| ---: | ---: | ---: | ---: | --- |
| 200 | 10 | 200 | 10 | `Skullcrusher` ajoute 20 degats tous les 5 tours |

Son drop est choisi avec les memes probabilites que les ennemis normaux.

## Experience et niveaux

L'XP est conservee pendant la partie. Le seuil du niveau suivant est calcule ainsi :

```text
XP requise pour le niveau N = N * (N - 1) / 2 * 100
```

A chaque niveau gagne :

- `+10` PV maximum.
- `+10` RAM maximum.
- `+50` pieces.
- Les PV et la RAM actuels augmentent egalement de 10.

## Inventaire et objets

La capacite initiale est de 10 objets. Chaque amelioration ajoute 10 places, avec un maximum de trois ameliorations, soit une capacite maximale de 40 objets.

Les objets sont comptes par quantite. Les objets qui ne sont pas des equipements ou des programmes peuvent rester dans l'inventaire sans action directe associee.

### Objets utilisables

- `Stimulant` : restaure 50 PV, sans depasser les PV maximum.
- `Virus` : retire un virus de l'inventaire et applique actuellement 3 degats de 10 au monstre fourni par l'action.

### Materiaux

Les materiaux servent a la fabrication :

- `Acier`.
- `Kevlar`.
- `Cuir`.
- `Puce neuronale`.

## Marchand

Le marchand est accessible depuis l'inventaire. Les prix sont les suivants :

| Objet | Prix | Limite |
| --- | ---: | ---: |
| Stimulant | 30 | Aucune limite specifique |
| Programme : Surcharge | 100 | 1 exemplaire |
| Programme : Crash | 110 | 1 exemplaire |
| Programme : Suicide | 400 | 1 exemplaire |
| IEM | 50 | Aucune limite specifique |
| Armure de combat | 150 | Aucune limite specifique |
| Bottes de soldat | 80 | Aucune limite specifique |
| Gants de precision | 60 | Aucune limite specifique |
| Virus | 40 | Aucune limite specifique |
| Augmentation d'inventaire | 30 | Maximum 3 achats |

Un achat est refuse si :

- Le joueur n'a pas assez d'argent.
- La quantite demandee est invalide.
- L'inventaire n'a pas assez de place.
- La limite d'un programme est depassee.
- La limite de trois ameliorations d'inventaire est atteinte.

## Cyberforgeron

Le cyberforgeron permet de fabriquer un equipement contre des materiaux et 5 pieces.

| Objet fabrique | Materiaux requis | Bonus de PV |
| --- | --- | ---: |
| Armure de combat | 1 Kevlar + 2 Acier | +25 |
| Bottes de soldat | 2 Cuir + 1 Acier | +15 |
| Gants de precision | 1 Cuir + 2 Puces neuronales | +10 |

La fabrication consomme les materiaux uniquement lorsque l'objet peut etre ajoute a l'inventaire et que le joueur peut payer les 5 pieces.

## Equipement

Les equipements occupent un emplacement dedie :

- `Armure de combat` : torse.
- `Bottes de soldat` : bottes.
- `Gants de precision` : gants.

Un seul objet peut etre equipe par emplacement. Equiper un nouvel objet remplace l'ancien, qui revient dans l'inventaire. Le bonus de PV maximum est recalculé pour tenir compte de l'objet equipe.

## Mort et resurrection

Si les PV atteignent zero, le jeu affiche deux options :

1. `Ressusciter` : revenir avec la moitie des PV maximum.
2. `Quitter le jeu` : fermer la partie.

Apres une resurrection, le prochain combat conserve l'etat prevu par le code de reinitialisation du combat. Il n'y a pas encore de sauvegarde automatique ni de reprise apres fermeture du programme.

## Architecture du code

Tous les fichiers sont dans le package Go `main`.

| Fichier | Responsabilite |
| --- | --- |
| `main.go` | Definit la structure du personnage et demarre la partie |
| `initCharacter.go` | Creation, validation et initialisation du personnage |
| `menu.go` | Menu principal et selection du type de combat |
| `TourParTour.go` | Boucle de combat, tours, attaques, XP et RAM |
| `premiermonstre.go` | Structure des ennemis, generation et comportements speciaux |
| `isDead.go` | Gestion de la defaite, mort et resurrection |
| `displayInfo.go` | Affichage de la fiche du personnage |
| `displaySkill.go` | Affichage des competences |
| `accessInventory.go` | Consultation et actions de l'inventaire |
| `marchand.go` | Boutique, prix, achats et drops |
| `ferailleur.go` | Recettes et fabrication d'equipements |
| `equipement.go` | Emplacements et bonus d'equipement |
| `spellBook.go` | Apprentissage des programmes |
| `takePot.go` | Utilisation des stimulants |
| `poisonPot.go` | Effet du virus sur un ennemi |
| `upgradeInventorySlot.go` | Amelioration de la capacite de l'inventaire |
| `limiteInventory.go` | Comptage des objets dans l'inventaire |
| `accessMerchant.go` | Acces a l'interface du marchand |

## Limites actuelles

- Le jeu ne sauvegarde pas la progression.
- Il n'y a pas de fichier `go.mod` ni de dependances declarees.
- L'interface repose sur `fmt.Scan` et ne gere pas encore tous les retours invalides de saisie.
- Certaines actions sont presentes dans le marchand mais leur utilisation est encore limitee ou non implementee, notamment `IEM`.
- Le combat d'entrainement et les combats aleatoires partagent une grande partie de leur logique, mais leurs recompenses sont gerees separement.
- La partie s'arrete lorsque le joueur quitte ou ferme le terminal.

## Pistes d'evolution

- Ajouter un `go.mod` et une version Go ciblee.
- Ajouter un systeme de sauvegarde et de chargement.
- Ajouter plusieurs fichiers de tests automatises.
- Centraliser les constantes de combat, prix et statistiques.
- Ameliorer la validation des saisies utilisateur.
- Ajouter davantage de classes, competences, ennemis et zones.
- Donner une utilisation complete a l'IEM et aux objets actuellement non actifs.
- Ajouter une interface graphique ou une interface terminal plus avancee.

## Licence

Aucune licence n'est actuellement indiquee dans le projet.
