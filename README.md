# Projet Red Cyberpunk

Projet Red Cyberpunk est un jeu textuel en Go dans un univers cyberpunk sombre et futuriste. Le joueur crée un personnage, choisit une classe, gère ses ressources, affronte des ennemis, visite un marchand et progresse dans un système de combat au tour par tour.

Le projet est pensé comme un petit RPG narratif et mécanique, entièrement jouable en console. Il combine plusieurs éléments classiques du jeu : gestion des PV, RAM, monnaie, inventaire, classes, attaques, équipements et logique d’initiative.

## Synopsis

Dans un monde de ruelles sombres, de corporations, de gangs et de technologies corrompues, le personnage du joueur doit survivre dans un environnement hostile. Il doit faire face à des adversaires variés, gagner de l’argent, améliorer ses capacités et choisir les bonnes décisions pour rester en vie.

Le jeu est volontairement simple dans son interface, mais riche dans ses mécanismes de progression.

## Objectif du jeu

L’objectif principal est de survivre, progresser dans le jeu et devenir de plus en plus fort. Le joueur doit :

- créer son personnage,
- choisir une classe adaptée à son style,
- gérer ses points de vie et sa RAM,
- combattre des ennemis,
- acheter des objets et des améliorations,
- réussir à rester vivant face aux attaques adverses.

## Système de jeu

### 1. Création du personnage

Au lancement, le joueur entre :

- un nom,
- une classe parmi les options disponibles.

Chaque classe a des caractéristiques propres et influence le gameplay.

### 2. Classes disponibles

Le jeu propose plusieurs classes de départ, chacune avec un style de combat différent :

- Netrunner
- Assassin
- Berserk

Chaque classe offre une expérience de jeu distincte, notamment au niveau des statistiques, de la durabilité et de la manière d’interagir avec les combats.

### 3. Statistiques du personnage

Le personnage possède plusieurs variables importantes :

- nom,
- classe,
- niveau,
- points de vie (PV),
- RAM,
- argent,
- compétences,
- inventaire,
- capacité d’inventaire,
- niveau d’initiative.

La gestion de ces éléments est centrale pour réussir les combats et gérer les ressources du jeu.

### 4. Combat au tour par tour

Les combats se déroulent en tour par tour. À chaque tour, le joueur peut choisir une action, comme :

- attaquer,
- ouvrir l’inventaire,
- fuir.

Les ennemis attaquent aussi selon leur propre logique, avec des attaques basiques ou spéciales selon leur type.

### 5. Initiative

Le système d’initiative détermine qui commence le combat. Le personnage et l’ennemi possèdent chacun une valeur d’initiative.

- si le personnage a plus d’initiative, il agit en premier,
- si l’ennemi a plus d’initiative, il attaque avant le joueur.

Cela ajoute une couche de stratégie au début de chaque combat.

### 6. Inventaire et objets

Le joueur peut stocker des objets dans son inventaire. Les objets peuvent être :

- des consommables,
- des ressources,
- des équipements,
- des objets utiles en combat ou hors combat.

Le système permet aussi de gérer la capacité de stockage et d’éventuelles améliorations.

### 7. Marchand

Le marchand est un élément clé du jeu. À la première visite, le joueur reçoit un stimulant gratuit. Ensuite, il peut acheter ou vendre des objets selon ses besoins.

Le marchand permet :

- d’acquérir des objets utiles,
- d’améliorer la préparation du personnage,
- d’augmenter sa survie dans les prochaines rencontres.

### 8. Ennemis

Le jeu inclut plusieurs types d’adversaires :

- monstres de base,
- ennemis spécialisés,
- boss plus puissants.

Chaque adversaire dispose de ses propres statistiques et attaques. Certains ont des comportements spécifiques qui influencent les combats.

## Mécaniques principales

### Points de vie

Les PV représentent la santé du personnage. Si ils tombent à 0, le personnage perd le combat.

### RAM

La RAM est une ressource spéciale utilisée pour certaines attaques. Certaines compétences consomment de la RAM, tandis que le personnage peut en récupérer au fil des tours.

### Argent

L’argent sert à acheter des objets et à améliorer la progression. Il est gagné en combattant, en vainquant des ennemis et en réussissant les challenges du jeu.

### Attaques

Le personnage peut utiliser différentes attaques selon sa progression. Chaque attaque a son propre coût et sa propre puissance.

## Déroulement d’une partie

Une partie typique se déroule comme suit :

1. création du personnage,
2. lecture des informations du personnage,
3. accès au menu principal,
4. choix d’une action : informations, marchand, combat, etc.,
5. affrontement contre des adversaires,
6. gain d’argent et de ressources,
7. progression vers des combats plus difficiles,
8. optimisation de l’inventaire et des compétences.

## Lancer le projet

Pour démarrer le jeu, ouvrez un terminal à la racine du projet puis exécutez :

```bash
go run .
```

Si vous êtes déjà dans le dossier `src`, vous pouvez aussi faire :

```bash
go run .
```

## Prérequis

- Go installé sur votre machine,
- un terminal pour lancer le programme,
- un système compatible avec les applications console.

## Structure du projet

Voici la structure principale du dépôt :

- `src/` : code source du jeu,
- `main.go` : point d’entrée du programme,
- `initCharacter.go` : création du personnage,
- `TourParTour.go` : logique des combats et du tour par tour,
- `accessMerchant.go` : accès au marchand,
- `marchand.go` : gestion des achats,
- `displayInfo.go` : affichage des informations,
- `equipement.go` : gestion des équipements,
- `menu.go` : menu principal,
- `premiermonstre.go` : gestion des ennemis et des attaques,
- `spellBook.go` : éventuelle logique sur les sorts ou compétences.

## Idée de progression

Le jeu est conçu pour être facilement extensible. Il peut être enrichi avec :

- de nouveaux ennemis,
- de nouvelles classes,
- plus d’armes et d’armures,
- des quêtes,
- un système de sauvegarde,
- des boss plus complexes,
- des événements narratifs.

## Remarque

Ce projet est un petit RPG textuel conçu pour apprendre et pratiquer la logique de jeu en Go. Il met l’accent sur la simplicité, la lisibilité et la compréhension des mécanismes de jeu.

## Conclusion

Projet Red Cyberpunk est un jeu simple mais complet dans son esprit. Il permet de découvrir les bases du développement d’un RPG en console, avec des éléments comme les statistiques, les combats, les objets, l’initiative et la progression du personnage.

Il est idéal pour un projet personnel, un apprentissage pratique en Go, ou un support de démonstration pour montrer les principes d’un jeu textuel.
