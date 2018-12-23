<?php
$modpack->version = "0.0.1";
$modpack->sha1 = sha1_file("./modpack.zip");
$modpackJSON = json_encode($modpack);
echo $modpackJSON;
?>
