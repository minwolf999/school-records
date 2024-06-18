<?php
require_once 'menu.php';
$competences = "";

if(isset($_GET['annee'])){
    if(!empty($_GET['annee'])){
        if(!empty($_GET['classe'])) {
            $annee = $_GET['annee'];
            $classe = $_GET['classe'];
            if ($classe == 'All') {
                $competences = $db->query("
                    SELECT c.id, c.nom, c.image_path, cc.nom as categorie, GROUP_CONCAT(e.nom ORDER BY e.classe, e.nom SEPARATOR '<br>') as EleveNom, GROUP_CONCAT(e.classe ORDER BY e.classe, e.nom SEPARATOR '<br>') as classe, GROUP_CONCAT(DISTINCT e.annee ORDER BY e.classe, e.nom SEPARATOR '<br>') as annee
        
                    FROM competence c

                    LEFT JOIN categorie_competence cc ON c.categorie_competence_id = cc.id
        
                    LEFT JOIN eleve_competence ec ON ec.competence_id = c.id
        
                    LEFT JOIN eleve e ON e.id = ec.eleve_id
    
                    WHERE e.annee = $annee
    
                    GROUP BY c.nom
        
                    ORDER BY cc.id, c.nom
                ")->fetchAll();
            } else {
                $competences = $db->query("
                    SELECT c.id, c.nom, c.image_path, cc.nom as categorie, GROUP_CONCAT(e.nom ORDER BY e.classe, e.nom SEPARATOR '<br>') as EleveNom, GROUP_CONCAT(e.classe ORDER BY e.classe, e.nom SEPARATOR '<br>') as classe, GROUP_CONCAT(DISTINCT e.annee ORDER BY e.classe, e.nom SEPARATOR '<br>') as annee

                    FROM competence c

                    LEFT JOIN categorie_competence cc ON c.categorie_competence_id = cc.id

                    LEFT JOIN eleve_competence ec ON c.id = ec.competence_id

                    LEFT JOIN eleve e ON ec.eleve_id = e.id

                    WHERE e.classe = '$classe' AND e.annee = $annee

                    GROUP BY c.nom

                    ORDER BY cc.id, c.nom

                ")->fetchAll();
            }
        }
    }
}

$recups = $db->query("
	SELECT e.annee
	FROM eleve e
	ORDER BY e.annee
")->fetchAll();

if ($recups != []) {
    foreach($recups as $recup):
	    $valeurs[] = $recup['annee'];
    endforeach;
    $valeur = array_unique($valeurs);
}

?>

<script language=javascript>

function redirige()
{
    select = document.getElementById("annee");
    choice = select.selectedIndex  // Récupération de l'index du <option> choisi
 
    valeur_cherchee = select.options[choice].value;

    select2 = document.getElementById("classe");
    choice2 = select2.selectedIndex

    valeur_cherchee2 = select2.options[choice2].value;

    if (valeur_cherchee != "" && valeur_cherchee2 != "") {
        location.href="demo.php?page=saisir&menu=listeCompetences&annee="+valeur_cherchee+"&classe="+valeur_cherchee2;
    }
}
</script>

<form method="post" name="trouverClient">
    <label for="nom_client">Année scolaire:</label>
        <select id="annee" name="annee" onchange="redirige()">
			<option value="" disabled selected></option>
            <?php if ($valeur != []):?>
			    <?php foreach($valeur as $value): ?>
				    <option value="<?php echo strtolower($value); ?>"><?php echo $value; ?></option>
			    <?php endforeach; ?>
                <?php endif; ?>
		</select>

        <label>Classe:</label>
        <select id="classe" name="classe" onchange="redirige()">
            <option value="" disabled selected></option>
            <option value="All">Tous les niveaux</option>
            <option value="Ps">Ps</option>
            <option value="Ms">Ms</option>
            <option value="Gs">Gs</option>
        </select>
</form>

<div class="container">
    <table class="styled-table">
        <thead>
            <tr>
                <th class="competence">Categorie</th>
                <th>Compétence</th>
                <th>Image</th>
				<th>Élève</th>
				<th>Classe</th>
                <th>Année</th>
                <th></th>
            </tr>
        </thead>
        <tbody>
        <?php if($competences != ""): ?>
            <?php foreach ($competences as $competence): ?>
            <?php $id = $competence['id'] ?>
            <form method="post" name="modif_competence">
            <tr>
                <td><?= $competence['categorie'] ?></td>
                <td><?= $competence['nom'] ?></td>
                <td>
                    <?php if (!empty($competence['image_path'])): ?>
                    <img width="150px" src="<?= $competence['image_path'] ?>" alt="<?= $competence['nom'] ?>"/>
                    <?php endif; ?>
                </td>
				<td><?= $competence['EleveNom'] ?> </td>
				<td><?= $competence['classe'] ?> </td>
                <td><?= $competence['annee'] ?></td>
                <td><?php echo"<a href='demo.php?page=saisir&menu=modifier_compet&id=$id' name='modifier'>Modifier</a>"?></td>
            </tr>
            </form>
            <?php endforeach; ?>
        <?php endif; ?>
        </tbody>
    </table>
</div>
