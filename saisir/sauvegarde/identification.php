<?php
$errors  = [];
$success = [];

$DB_HOST = "localhost";
$DB_USER = "root";
$DB_PASS = "";
$DB_NAME = "identification";

$con = new mysqli($DB_HOST, $DB_USER, $DB_PASS, $DB_NAME);
mysqli_set_charset($con, 'utf8mb4');
 
 $tables = array();

$result = mysqli_query($con,"SHOW TABLES");
while ($row = mysqli_fetch_row($result)) {
    $tables[] = $row[0];
}

$return = '';

foreach ($tables as $table) {
    $result = mysqli_query($con, "SELECT * FROM ".$table);
    $num_fields = mysqli_num_fields($result);

    $return .= 'DROP TABLE '.$table.';';
    $row2 = mysqli_fetch_row(mysqli_query($con, 'SHOW CREATE TABLE '.$table));
    $return .= "\n\n".$row2[1].";\n\n";

    for ($i=0; $i < $num_fields; $i++) { 
        while ($row = mysqli_fetch_row($result)) {
            $return .= 'INSERT INTO '.$table.' VALUES(';
            for ($j=0; $j < $num_fields; $j++) { 
                $row[$j] = addslashes($row[$j]);
                if (isset($row[$j])) {
                    $return .= '"'.$row[$j].'"';} else { $return .= '""';}
                    if($j<$num_fields-1){ $return .= ','; }
                }
                $return .= ");\n";
            }
        }
        $return .= "\n\n\n";
    
}

$handle = fopen('C:\wamp64\www\school-records\saisir\sauvegarde\identification.sql', 'w+');
fwrite($handle, $return);
fclose($handle);

if (!$handle) {
    $errors[] = sprintf('Aucun client qui porte le nom où prénom `%s` ne possède de carte complète !', $_POST['nom']);
}
else if (empty($errors)) {
    $success[] = sprintf('La sauvegarde de la base de données d\'indentification a été effectuée !');
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
<!-- Fin gestion messages -->