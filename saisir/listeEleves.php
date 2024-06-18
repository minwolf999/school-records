<?php
require_once 'menu.php';
$eleves = "";

$db->query("SET @@GROUP_CONCAT_MAX_LEN =2048;")->fetchAll();

if(isset($_GET['annee'])){
    if(!empty($_GET['annee'])){
        $annee = $_GET['annee'];
        $eleves = $db->query("
            SELECT e.id, e.nom, e.classe, e.annee, GROUP_CONCAT(c.nom ORDER BY c.categorie_competence_id, c.nom SEPARATOR '<br>') as comptences
            FROM eleve e
            LEFT JOIN eleve_competence ec ON ec.eleve_id = e.id
            LEFT JOIN competence c ON c.id = ec.competence_id
            WHERE e.annee = $annee
            GROUP BY e.id, e.nom
            ORDER BY e.annee, e.classe, e.nom
        ")->fetchAll();
    }
}

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

    if(isset($_POST['supprimer'])){
        $id = $_POST['id'];
        $supprimer = $db->query("
        DELETE FROM eleve
        WHERE id = $id
        ")->fetchAll();
    }
}
?>

<script language=javascript>

function redirige()
{
    select = document.getElementById("annee");
    choice = select.selectedIndex  // Récupération de l'index du <option> choisi
 
    valeur_cherchee = select.options[choice].value;

    location.href="demo.php?page=saisir&menu=listeEleves&annee="+valeur_cherchee;
}
</script>

<form method="post" name="trouverClient">
    <label for="nom_client">Année scolaire:</label>
            <select id="annee" name="annee" onchange="redirige()">
                <option value="" disabled selected></option>
                <?php foreach($valeur as $value): ?>
                    <option value="<?php echo strtolower($value); ?>"><?php echo $value; ?></option>
                <?php endforeach; ?>
        </select>
</form>

<div class="container">
    <table class="styled-table">
        <thead>
        <tr>
            <th>Nom</th>
			<th>Classe</th>
            <th>Année</th>
            <th>Compétences</th>
            <th></th>
            <th></th>
            <th></th>
        </tr>
        </thead>
        <tbody>
        <?php if($eleves != ""): ?>
            <?php foreach ($eleves as $eleve): ?>
                <?php $id = $eleve['id'] ?>
                <form method="post" name="suppr_eleve">
                <tr>
                    <td><?= $eleve['nom'] ?></td>
    				<td><?= $eleve['classe'] ?></td>
                    <td><?= $eleve['annee']?></td>
                    <td><?= $eleve['comptences'] ?></td>
                    <td><?php echo"<a href='demo.php?page=saisir&menu=modifier&id=$id' name='modifier'>Modifier</a>"?></td>
                    <td><input type="hidden" value="<?php echo($eleve['id']) ?>" name="id"></td>
                    <td><input type="submit" value="Supprimer" name="supprimer"></td>
                </tr> 
            </form>
            <?php endforeach; ?>
        <?php endif;?>
        </tbody>
    </table>
</div>
