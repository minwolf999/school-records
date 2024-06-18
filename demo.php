<!doctype html>
<html lang="fr">
<head>
    <meta charset="utf-8">
    <title>Livrets </title>
    <link rel="stylesheet" type="text/css" href="css/main.css">
</head>
<style>
body {
	background:url('images/fond2.png') ;
	background-size: 109.5%;

}
</style>
<body>
<?php
    try {
        $page = 'accueil';
        if (!empty($_GET['page'])) {
            $page = $_GET['page'];
        }

        if (!file_exists(__DIR__ . '/' . $page . '.php')) {
            $page = '404';
        }

        require_once $page . '.php';
    } catch (Exception $ex) {
        require_once 'error.php';
    }
?>

<div id="footer">
	<center>
		Outil créé par Antoine Marvin
	</center>
</div>

</body>
</html>