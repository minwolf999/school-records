<?php
    require_once 'bdd/connexion.php';

    try {
        $page = 'saisir/listeEleves';
        if (!empty($_GET['menu'])) {
            $page = 'saisir/'.$_GET['menu'];
        }

        if (!file_exists(__DIR__ . '/' . $page . '.php')) {
            $page = '../404';
        }

        require_once $page . '.php';
    } catch (Exception $ex) {
        require_once 'error.php';
    }
?>
