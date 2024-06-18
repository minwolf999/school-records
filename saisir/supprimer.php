<?php
require 'menu.php';

$recups = $db->query("
    SELECT e.annee
    FROM eleve e
    ORDER BY e.annee
")->fetchAll();

if ($recups != []) {
    foreach ($recups as $recup):
        $valeurs[] = $recup['annee'];
    endforeach;
    $valeur = array_unique($valeurs);
}

if(isset($_GET['annee'])){
    $annee = $_GET['annee'];
    $eleves = $db->query("
    SELECT *
    FROM eleve
    WHERE annee = $annee
    ")->fetchAll();
}
?>

<script language=javascript>

function redirige()
{
    select = document.getElementById("annee");
    choice = select.selectedIndex  // Récupération de l'index du <option> choisi
 
    valeur_cherchee = select.options[choice].value;

    location.href="demo.php?page=saisir&menu=supprimer&annee="+valeur_cherchee;
}
</script>

<form method="post" name="trouverClient" class="ajouterEleve">
    <center>
        <label for="nom_client" class="center">Année scolaire:</label>
            <select id="annee" name="annee" onchange="redirige()">
                <option value="" disabled selected></option>
                <?php foreach($valeur as $value): ?>
                    <option value="<?php echo strtolower($value); ?>"><?php echo $value; ?></option>
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
            <label for="cat_competence" class="center">Élève :</label>
            <?php if(!empty($eleves)): ?>
                <?php foreach($eleves as $eleve): ?>
                    <?php $id = $eleve['id']; ?>
                    <br><input type="button" value="<?php echo $eleve['nom'] ?>" onclick="lien(<?php echo $id ?>) ">
                <?php endforeach; ?><br>
            <?php endif; ?>
        </div>
    </center>
</form>

<script>
    function lien(id) {
        window.location = "demo.php?page=saisir&menu=supprimer_competence&id="+id;
    }
</script>