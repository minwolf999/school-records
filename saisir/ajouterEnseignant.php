<?php
require_once 'menu.php';

$phpFileUploadErrors = [
	1 => 'La taille du fichier dépasse celle définie dans la configuration du php.ini',
	2 => 'La taille du fichier dépasse celle spécifié dans le formulaire',
	3 => 'Le fichier a été upload partiellement',
	4 => 'Pas de fichier uploadé',
	6 => 'Le dossier temporaire n\'existe pas',
	7 => 'Impossible de sauvegarder le fichier sur le disque. Manque de droits',
	8 => 'Une extension PHP a arrêté l\'upload',
];

$errors = [];
$success = [];

if(isset($_POST['submit'])){
	if(empty($_FILES['image']['name'])){
		$errors[] = sprintf("L'image a été oubliée !");
	}
	if(empty($_POST['sexe'])){
		$errors[] = sprintf("Le genre de l'enseignant a été oubliée !");
	}
	if(empty($_POST['enseignant'])){
		$errors[] = sprintf("Le nom de l'enseignant a été oubliée !");
	}	
}
	if(!empty($_POST['enseignant']) && !empty($_FILES['image']) && !empty($_POST['sexe'])){
		$uploadFile = null;
		if(!empty($_FILES['image'])){
			$extension = pathinfo($_FILES['image']['name'], PATHINFO_EXTENSION);
			if(UPLOAD_ERR_OK === $_FILES['image']['error']){
				$uploadDir = 'images/';
				$uploadFile = $uploadDir . uniqid() . "." .$extension;
				if(!move_uploaded_file($_FILES['image']['tmp_name'], $uploadFile)){
					$errors['move_failed'] = 'Une erreur est survenue lors de l\'upload de l\'image.';
					$uploadFile = null;
				}
			} else {
				$errors[$_FILES['image']['error']] = $phpFileUploadErrors[$_FILES['image']['error']];
			}
		}
		$query = $db->prepare('INSERT INTO `enseignant` (`nom`, `image`, `sexe`) VALUES (:nom, :image_path, :sexe)');
		
		$result = $query->execute([
			'nom' => $_POST['enseignant'],
			'image_path' => $uploadFile,
			'sexe' => $_POST['sexe'],
		]);
		if(!$result){
			$errors[$query->errorCode()] = $query->errorInfo();
		} else{
			$success[] = sprintf("L'enseignant `%s` a été ajouté avec succès !", $_POST['enseignant']);
		}
	}
	
?>

<div class="container">

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

<form method="post" name="ajouterEnseignant" enctype="multipart/form-data" class="ajouterEleve">
<center>
	<label for="nom_enseignant" class="center">Nom de l'enseignant :</label>
	<textarea id="enseignant" name="enseignant" placeholder="Nom de l'enseignant"></textarea><br>
	
	<br><label for="sexe" class="center">Genre de l'enseignant :</label>
	<select id="sexe" name="sexe">
		<option value="" disabled selected></option>
		<option id="sexe" name="sexe" value="enseignant">homme</option>
		<option id="sexe" name="sexe" value="enseignante">femme</option><br>
	</select>
	
	<br><label for="image" class="center">Signature de l'enseignant :</label>
	<input type="file" id="image" name="image">
	
	<br><br><button type="submit" name="submit">Ajouter</button>
</center>
</form>
</div>