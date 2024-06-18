<?php
require_once 'menu.php';

$id = $_GET['id'];

$eleves = $db->query("
    SELECT *
    FROM eleve
    WHERE id = $id
")->fetchAll();

$competencesData = $db->query("
    SELECT ec.id as need, ec.competence_id, c.id, c.nom, cc.nom as categorie, cc.id as id_cat
    FROM eleve_competence ec
    INNER JOIN competence c ON c.id = ec.competence_id
    INNER JOIN categorie_competence cc ON cc.id = c.categorie_competence_id
    WHERE ec.eleve_id = $id
")->fetchAll();

$competences = [];
foreach ($competencesData as $row) {
    if (!array_key_exists($row['categorie'], $competences)) {
        $competences[$row['categorie']] = [];
    }
    $competences[$row['categorie']][] = $row;
}

if(isset($_POST['submit'])){
    $id_competence = implode($_POST['competences']);
    
    $delete = $db->query("
    DELETE FROM eleve_competence
    WHERE id = $id_competence
    ")->fetchAll();
    $success[] = sprintf("La compétence a été supprimer avec succès !");
}
?>
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
    <div class="wrapper">
        <center>
            <?php foreach($eleves as $eleve): ?>
                <label for="cat_competence" class="center">Supprimer les compétences de <?php echo $eleve['nom'] ?></label><br>
            <?php endforeach; ?>

            <br/><label for="cat_competence" class="center">Catégorie de compétence</label><br>
	            <select id="parent_select"><option value="" disabled selected></option></select><br>
            <label for="competence" class="center">Compétence</label><br>
	            <select id="competences" name="competences[]"></select><br>

            <input type="submit" value="Supprimer" name="submit">
        </center>
    </div>
</form>

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
				el_child.innerHTML = el_child.innerHTML + "<option value="+ mList[itm][i]['need']+ ">"+ mList[itm][i]['nom'] +'</option>';
			}
	}
});

</script>