<?php
require 'bdd/connexion.php';
require_once 'observer/menu.php';

$eleves = "";
$valeurs = [];

if(isset($_GET['annee'])){
	if(!empty($_GET['annee'])){
		$annee = $_GET['annee'];
		$eleves = $db->query("
			SELECT nom, classe, annee

			FROM eleve

			WHERE annee = $annee

			ORDER BY classe, nom 
		")->fetchAll();
	}
}

$enseignant = $db->query("
SELECT id, nom, image

FROM enseignant

ORDER BY nom
")->fetchAll();

if(isset($_POST['submit'])){
    if(!empty($_POST['submit'])){
        require 'observer/pdf/index.php';
    }
}

$recups = $db->query("
    SELECT e.annee
    FROM eleve e
    ORDER BY e.annee
")->fetchAll();

foreach ($recups as $recup):
    $valeurs[] = $recup['annee'];
    endforeach;
    if ($valeurs != []) {
        $valeurs = array_unique($valeurs);
    }
?>

<script language=javascript>

function redirige()
{
    select = document.getElementById("annee");
    choice = select.selectedIndex  // Récupération de l'index du <option> choisi
 
    valeur_cherchee = select.options[choice].value;

    location.href="demo.php?page=observer&annee="+valeur_cherchee;
}
</script>

<form method="post" name="trouverEleve">
    <label for="nom_client">Année scolaire:</label>
    <?php if ($valeurs != []): ?>
        <select id="annee" name="annee" onchange="redirige()">
            <option value="" disabled selected></option>
            <?php foreach($valeurs as $value): ?>
                <option value="<?php echo strtolower($value); ?>"><?php echo $value; ?></option>
            <?php endforeach; ?>
        </select>
    <?php else: ?>
        <?php $errors[] = sprintf("Il n'y a pas d'élève existant donc aucune année n'a été entrée dans la base de données"); ?>
    <?php endif; ?>
</form>

<div class="form-message">
	<?php if(!empty($errors)): ?>
		<?php foreach($errors as $error): ?>
			<div class="alert">
				<span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
				<?= $error ?>
			</div>
		<?php endforeach; ?>
	<?php endif; ?>
	<?php if(!empty($success)): ?>
		<?php foreach($success as $msg): ?>
			<div class="alert sucess">
				<span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
				<?= $msg ?>
			</div>
		<?php endforeach; ?>
	<?php endif; ?>
</div>

<div class="container">
	<form method="post" name="ouvrirpdf">
		<br><label for="nom_enseignant">Nom de l'enseignant :</label>
		<select id="enseignant" name="enseignant">
			<option value="" disabled selected></option>
		<?php foreach($enseignant as $prof): ?>
			<option value="<?= $prof['id']?>"><?= $prof['nom']?></option>
		<?php endforeach; ?>
		</select>
</div>
<div class="eleve">
<br><label for="eleve">Élève souhaité: </label><br>
 <form method="post" name="ouvrirpdf">
	<?php if($eleves!=""): ?>
    <?php foreach ($eleves as $eleve): ?>
        <input type="submit" value="<?= $eleve['nom']?>" name="submit" style="width:200px">
		<input type="text" value="<?= $annee ?>" name="annee" style="visibility:hidden;"><br>
    <?php endforeach; ?>
	<?php endif; ?>
 </form>
</div>