<?php
require 'menu.php';

$errors = [];
$success = [];
$rawEleves = "";
$eleves = [];

$recups = $db->query("
    SELECT e.annee
    FROM eleve e
    ORDER BY e.annee
")->fetchAll();

if(isset($_GET['annee'])){
    if(!empty($_GET['annee'])){
        $annee = $_GET['annee'];
        $rawEleves = $db->query("
        SELECT e.id, e.nom, e.classe, e.annee, c.id as competenceID, c.nom as competenceNom

        FROM eleve e

        LEFT JOIN eleve_competence ec ON ec.eleve_id = e.id

        LEFT JOIN competence c ON c.id = ec.competence_id

        WHERE e.annee = $annee

        ORDER BY e.classe, e.nom, c.nom;
        ")->fetchAll();

        foreach ($rawEleves as $rawEleve) {
            if (empty($eleves[$rawEleve['id']])) {
                $eleves[$rawEleve['id']] = [
                    'id' => $rawEleve['id'],
                    'nom' => $rawEleve['nom'],
        			'classe' => $rawEleve['classe'],
                    'competences' => [],
                ];
            }

            $eleves[$rawEleve['id']]['competences'][$rawEleve['competenceID']] = $rawEleve['competenceNom'];
        }
    }
}


$competencesData = $db->query('
    SELECT c.id, c.nom, cc.nom as categorie, cc.id as id_cat
    FROM competence c
    INNER JOIN categorie_competence cc
    ON cc.id = c.categorie_competence_id
')->fetchAll();

$competences = [];
foreach ($competencesData as $row) {
    if (!array_key_exists($row['categorie'], $competences)) {
        $competences[$row['categorie']] = [];
    }
    $competences[$row['categorie']][] = $row;
}


if(isset($_POST['submit'])){
    if(!empty($_POST['eleve'])){
        foreach($_POST['eleve'] as $eleveId){

    if (!empty($_POST['competences'] )) {
            //$eleveId = $eleve['id'];
            $query = $db->prepare('INSERT INTO eleve_competence (eleve_id, competence_id) VALUES (:eleveId, :competenceId);');

            $db->beginTransaction();
            foreach ($_POST['competences'] as $competence) {
                $result = $query->execute([
                    'eleveId' => $eleveId,
                    'competenceId' => $competence,
                ]);
                
                if (!$result) {
                    $errors[$query->errorCode()] = $query->errorInfo();
                }	else {
					$success[] = sprintf('La compétence a été ajoutée avec succès !');
				}
            }

            
            $db->commit();
    }

		if(empty($_POST['competences'])){
			$errors[] = sprintf("Aucune compétences n'a été selectionnée !");
		}
        
        
    if (!empty($_POST['appreciation'] )) {
            //$eleveId = $eleve['id'];
            $query = $db->prepare('INSERT INTO eleve_appreciation (eleve_id, appreciation) VALUES (:eleveId, :appreciation);');
        
            $db->beginTransaction();
            $appreciation = $_POST['appreciation'];
			$appreciation = nl2br($appreciation);
                $result = $query->execute([
                    'eleveId' => $eleveId,
                    'appreciation' => $appreciation,
                ]);
                        
                if (!$result) {
                    $errors[$query->errorCode()] = $query->errorInfo();
                }	else {
					$success[] = sprintf("L'appréciation a été ajoutée avec succès !");
				}
            
        
                    
            $db->commit();
    }

	if(empty($_POST['appreciation'])){
		$errors[] = sprintf("Aucune appreciation n'a été rentrée !");
	}
	
    }   }
	if(empty($_POST['eleve'])){
		$errors[] = sprintf("Aucun élève n'a été selectionnée !");
	}
}

if ($recups != []) {
    foreach ($recups as $recup):
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

    location.href="demo.php?page=saisir&menu=modifierEleve&annee="+valeur_cherchee;
}
</script>

<form method="post" name="trouverClient" class="ajouterEleve">
    <center>
        <label for="nom_client" class="center">Année scolaire:</label>
            <select id="annee" name="annee" onchange="redirige()">
                <option value="" disabled selected></option>
                <?php foreach($valeur as $value): ?>
                    <option id="annee" value="<?php echo strtolower($value); ?>"><?php echo $value; ?></option>
                <?php endforeach; ?>
            </select>
    </center>
</form><br>

<div class="container">

	<div class="form-message">
		<?php if (!empty($errors)):?>
			<?php foreach ($errors as $error): ?>
				<div class="alert">
					<span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
					<?= $error ?>
				</div>
			<?php endforeach; ?>
		<?php endif; ?>
		<?php if (!empty($success)): ?>
			<?php foreach($success as $msg): ?>
				<div class="alert sucess">
					<span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
					<?= $msg ?>
				</div>
			<?php endforeach; ?>
		<?php endif; ?>
	</div>

    <form method="post" name="ajouterEleve" class="ajouterEleve_big">
<center>
<div class="wrapper">
<br/><label for="cat_competence" class="center">Catégorie de compétence</label>
	    <select id="parent_select"><option value="" disabled selected></option></select><br>
    <label for="competence" class="center">Compétence</label>
	    <select id="competences" name="competences[]"></select><br>
</div>

<script language="javascript" type="text/javascript">  

    var mList = <?php echo json_encode($competences); ?> 

el_parent = document.getElementById("parent_select");
el_child = document.getElementById("competences");

for (key in mList) {
	el_parent.innerHTML = el_parent.innerHTML + '<option>'+ key +'</option>';
}

el_parent.addEventListener('change', function populate_child(e){
	el_child.innerHTML = '';
	itm = e.target.value;
	if(itm in mList){
			for (i = 0; i < mList[itm].length; i++) {
				el_child.innerHTML = el_child.innerHTML + "<option value="+ mList[itm][i]['id']+ ">"+ mList[itm][i]['nom'] +'</option>';
			}
	}
});

</script>

<?php require 'bouton.html' ?>

<div>
    <input type="button" id="cocherGs" value="Cocher tous les Gs">
	<input type="button" id="cocherMs" value="Cocher tous les Ms">
	<input type="button" id="cocherPs" value="Cocher tous les Ps">
	<input type="button" id="décocher" value="Décocher tous les élèves">
</div><br>
	
    <label for="nom_eleve" class="center">Nom de l'élève:</label><br>
    <?php if ($eleves != ""): ?>
        <?php foreach  ($eleves as $eleve):?>
            <input type="checkbox" name="eleve[]" id="eleve" class="<?= $eleve['classe'] ?>" value="<?= $eleve['id'] ?>" /><?=$eleve['nom']?><br>
        <?php endforeach; ?> <br>
    <?php endif; ?>
    
    <label for="appreciation" class="center">Appréciation:</label>
        <textarea id="appreciation" rows="10" cols="50" name="appreciation" placeholder="appréciation"></textarea><br>
<br>
        <input type="submit" value="Enregistrer" name="submit">
        </center>
    </form>
</div>

<script>
$(document).ready(function(){
	
	$(":button#cocherPs").click(function(){
		$(':checkbox.Ps').prop('checked', true);
	});
	
	$(":button#cocherMs").click(function(){
		$(':checkbox.Ms').prop('checked', true);
	});

    $(":button#cocherGs").click(function(){
		$(':checkbox.Gs').prop('checked', true);
	});
	
	$(":button#décocher").click(function(){
		$(':checkbox.Ms').prop('checked', false);
		$(':checkbox.Ps').prop('checked', false);
        $(':checkbox.Gs').prop('checked', false);
	});
});
</script>

