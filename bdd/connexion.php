<?php

try{
    $db = new PDO('mysql:host=localhost;dbname=sdz;charset=utf8', 'root');
}catch(PDOException $e){
    die('Erreur connexion : '.$e->getMessage());
}