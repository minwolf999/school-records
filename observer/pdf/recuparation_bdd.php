<?php
$eid = $_POST['enseignant'];
$eleven = $_POST['submit'];
$annee = $_POST['annee'];

$sth = $db->prepare("
    SELECT e.id, e.nom, e.classe, c.id as idCompetence, c.nom as nomCompetence, c.image_path as imageCompetence, cat.id as idCategorie, cat.nom as nomCategorie

    FROM eleve e

    LEFT JOIN eleve_competence ec ON ec.eleve_id = e.id

    LEFT JOIN competence c ON c.id = ec.competence_id

    LEFT JOIN categorie_competence cat ON cat.id = c.categorie_competence_id

    WHERE e.nom = '$eleven' AND e.annee = $annee

    ORDER BY cat.id
");
$sth->execute();
// Récupération de toutes les lignes d'un jeu de résultats
$results = $sth->fetchAll();

$competences = [];
foreach ($results as $row) {
	$classe = $row['classe'];
	$id = $row['id'];
	
    if (!array_key_exists($row['nomCategorie'], $competences)) {
        $competences[$row['nomCategorie']] = [];
    }
    $competences[$row['nomCategorie']][] = $row;
}

$competencef ="";

foreach ($competences as $competenceLabel => $competenceRows): 
        $competencef .= "<h3><b><i>$competenceLabel</i></b></h3>";
    foreach ($competenceRows as $competence): 
        if (!empty($competence['imageCompetence'])){
            $catcompetence = $competence['imageCompetence'];
        }
        $competencef .= "<blockquote>".ucfirst($competence['nomCompetence']). "&nbsp; &nbsp; &nbsp; &nbsp; <img src=$catcompetence height='60' width='75'> </blockquote>" ;
    endforeach;
endforeach;

$query = $db->prepare("
SELECT e.id, ea.appreciation as Appreciation

FROM eleve e

LEFT JOIN eleve_appreciation ea ON ea.eleve_id = e.id

WHERE e.id = $id
");
$query->execute();
$nice = $query->fetchAll();

$Appreciation = "";
foreach($nice as $appreciation)
	if(isset($appreciation['Appreciation'])){
		$Appreciation .= "<blockquote>".$appreciation['Appreciation']."<br></blockquote>";
	} 

$enseignants = $db->prepare("
SELECT id, nom, sexe, image

FROM enseignant

WHERE id = $eid
");
$enseignants->execute();

foreach($enseignants as $enseignant)

$nom_enseignant = $enseignant['nom'];
$signature = $enseignant['image'];
$sexe = $enseignant['sexe'];
$anne = (int)$annee;
$annee = $annee . " / " . (string)($anne+1);