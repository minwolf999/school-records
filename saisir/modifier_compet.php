<?php
require_once 'menu.php';

$id = $_GET['id'];

$competences = $db->query("
SELECT * 
FROM competence
WHERE id = $id
")->fetchAll();

$phpFileUploadErrors = [
    1 => 'La taille du fichier dépasse celle définie dans la configuration du php.ini',
    2 => 'La taille du fichier dépasse celle spécifié dans le formulaire',
    3 => 'Le fichier a été upload partiellement',
    4 => 'Pas de fichier uploadé',
    6 => 'Le dossier temporaire n\'existe pas',
    7 => 'Impossible de sauvegarder le fichier sur le disque. Manque de droits',
    8 => 'Une extension PHP a arrêté l\'upload',
];

$errors  = [];
$success = [];

if(isset($_POST['submit'])){
	if(empty($_FILES['image']['name'])){
		$errors[] = sprintf("L'image a été oubliée !");
	}
	
	if(empty($_POST['categorie'])){
		$errors[] = sprintf("La catégorie a été oubliée !");
	}
	
	if(empty($_POST['valeur'])){
		$errors[] = sprintf("Le nom de la compétence a été oubliée !");
}}

if (!empty($_POST['categorie']) && !empty($_POST['valeur']) && !empty($_FILES['image'])) {
    $uploadFile = null;
    if (!empty($_FILES['image'])) {
		$extension = pathinfo($_FILES['image']['name'],PATHINFO_EXTENSION);
        if (UPLOAD_ERR_OK === $_FILES['image']['error']) {
            $uploadDir = 'upload/';
            $uploadFile = $uploadDir . uniqid() . "." . $extension;
            if (!move_uploaded_file($_FILES['image']['tmp_name'], $uploadFile)) {
                $errors['move_failed'] = 'Une erreur est survenue lors de l\'upload de l\'image.';
                $uploadFile = null;
            }
        } else {
            $errors[$_FILES['image']['error']] = $phpFileUploadErrors[$_FILES['image']['error']];
        }
    }

    $query = $db->prepare("UPDATE `competence` SET `nom`=:nom, `categorie_competence_id`=:categorie_competence_id, `image_path`=:image_path WHERE id = $id");

    $result = $query->execute([
        'nom' => $_POST['valeur'],
        'categorie_competence_id' => $_POST['categorie'],
        'image_path' => $uploadFile,
    ]);

    if (!$result) {
        $errors[$query->errorCode()] = $query->errorInfo();
    } else {
        $success[] = sprintf('La compétence `%s` a été créé avec succès!', $_POST['valeur']);
    }
}

$categories = $db->query('SELECT id, nom FROM categorie_competence')->fetchAll();
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
    <!-- Fin gestion messages -->

    <form method="post" name="ajouterCompetence" enctype="multipart/form-data">
        <label for="categorie">Catégorie de compétence</label>
        <select id="categorie" name="categorie">
            <option value="" disabled selected></option>
            <?php foreach ($categories as $categorie): ?>
            <option value="<?= $categorie['id'] ?>"><?= ucfirst($categorie['nom']) ?></option>
            <?php endforeach; ?>
        </select>

        <br><label for="valeur">Nom de la compétence</label>
        <?php foreach ($competences as $competence): ?>
        <input type="text" size="59" id="valeur" name="valeur" placeholder="Nom de la compétence" value="<?php echo($competence['nom'])?>" onFocus="this.value='';"></textarea>
        <?php endforeach; ?>

        <br><label for="image">Image à associer à la compétence</label>
        <input type="file" id="image" name="image">

        <br><br><button type="submit" name="submit">Créer</button>
    </form>
</div>
