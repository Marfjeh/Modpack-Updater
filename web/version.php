<?php
$modpack = new stdClass();
$modpack->version = "0.0.1";
$modpack->sha1 = sha1_file("./modpack.zip");
echo json_encode($modpack);
?>
