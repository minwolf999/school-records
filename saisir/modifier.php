<?php
require_once 'menu.php';
$id = $_GET['id'];

$eleves = $db->query("
SELECT * 
FROM eleve
WHERE id = $id
")->fetchAll();

if(isset($_POST['submit'])){
	if(empty($_POST['classe'])){
		$errors[] = sprintf('La classe a été oubliée !');
	}
	if(empty($_POST['nom'])){
		$errors[] = sprintf("le nom de l'élève a été oubliée !");
	}
    if(empty($_POST['annee'])){
		$errors[] = sprintf("l'année de scolarité a été oubliée !");
	}

    if (!empty($_POST['nom'])) {
        $query = $db->prepare("
        UPDATE `eleve` 
        SET `nom`=:nom,`classe`= :classe, `annee`= :annee
        WHERE id = $id");
        
        foreach($_POST['classe'] as $classe){
			$result = $query->execute([
				'nom' => $_POST['nom'],
				'classe' => $classe,
                'annee' => $_POST['annee']
			]);
		}

		if (!$result) {
			$errors[$query->errorCode()] = $query->errorInfo();
		}
		else if (empty($errors)) {
				$success[] = sprintf('L\'élève `%s` a été créé avec succès!', $_POST['nom']);
		}
	}
}

?>

<div class="container">
    <!-- Gestion messages -->
    <div class="form-message">
        <?php if (!empty($errors)): ?>
            <?php foreach ($errors as $error): ?>
                <div class="alert">
                    <span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
                    <?= $error ?>
                </div>
            <?php endforeach; ?>
        <?php endif; ?>
        <?php if (!empty($success)): ?>
            <?php foreach ($success as $msg): ?>
                <div class="alert sucess">
                    <span class="closebtn" onclick="this.parentElement.style.display='none';">&times;</span>
                    <?= $msg ?>
                </div>
            <?php endforeach; ?>
        <?php endif; ?>
    </div>

<form method="post" class="ajouterEleve" name="ajouterClient">
    <center>
    <?php foreach($eleves as $eleve): ?>
        <label for="nom_eleve" class="center">Nom de l'élève:</label>
            <input id="nom_eleve" type="text" name="nom" value="<?php echo($eleve['nom'])?>" onFocus="this.value='';" placeholder="Nom de l'élève"><br>

        <label for="classe_eleve" class="center">Classe de l'élève:</label>
        <select id="classe" name="classe[]" multiple="multiple">
                <option value="Ps" name="Ps">Petite section</option>
                <option value="Ms" name="Ms">Moyenne section</option>
                <option value="Gs" name="Gs">Grande section</option>
            </select>

        <label for="annee" class="center">Année de scolarité:</label>
            <input id="annee" type="text" name="annee" value="<?php echo($eleve['annee'])?>" onFocus="this.value='';" placeholder="Année de scolarité">

        <input type="submit" value="Créer" name="submit">
    <?php endforeach; ?>
    </center>
    </form>
</div>
