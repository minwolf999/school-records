<?php
require 'observer/pdf/recuparation_bdd.php';


use Spipu\Html2Pdf\Html2Pdf;

    require __DIR__.'/vendor/autoload.php';

    $html2pdf = new Html2Pdf('P','A4','fr');

    $html = "
        <page backtop='25mm' backbottom='10mm' backleft='10mm' backright='10mm'>

		<page_header>
             <blockquote><h4><b>$eleven</b></h4>
             <div align='right'>Année scolaire $annee</div>
             Classe de $classe</blockquote>
		</page_header>

        <br><br><br>
        $competencef

        <h3><i>Appréciation</i></h3>
		
		<p style='line-height: 25px'>
        $Appreciation
        </p>
        
		
		<div style='position:relative'>
		<div style='position:absolute;left:0;'>$nom_enseignant $sexe </div>
		<div style='position:absolute;right:0;'>Signature des parents </div>
		</div>
		
		</page>
        ";
        ob_end_clean();
    $html2pdf->writeHTML($html);
    $html2pdf->output("$eleven.pdf");