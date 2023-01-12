# Modpack updater

This is meant as a pre-launch command for multimc, PolyMC, PrismLauncher or any other launcher. And to be included in a multimc instance zip. It will then execute at starting the multimc instance (Example your modpack)
It will check if there is an update on a remote server, if there is it will quietly update the modpack automaticlly and then start the instance.
If there is no update, then it will just exit and then multimc will proceed to start the instance.


This program can also run standalone or with any other minecraft launchers if they support pre-launch commands.
and you might even get this working for different games as well.


It keeps everything simple without configuration hell. 1 zip file and 1 json file is needed really.


This may not make some mod developers happy that i'm rehosting their mods but downloading from curse is just a nightmare.
(And i'm lazy sorry.)


Also a good option to use, if you want to play by the rules is using [packwiz](https://packwiz.infra.link/).

It can download the mods from Modrinth and curseforge instead you rehosting people's mods and making them angry. but it is kind of a configuration hell.


## Features

* Basic SHA-1 Verifying, making sure that the zip file downloaded is correct and not corrupted.
* Downloads the Zip file (Yay)
* extracts the zip file in current directory
* basic version checking, it only downloads if the version is different
* JSON file for the modpack, for name version etc.


## Planned features
* Prevent from deleting user-added mods.


## Building

Compiling this is really simple, as everything is in main.go, you should have go installed on your system.

build it with `go build main.go`

please note this will only compile it for your system. to compile it for different platforms refer to the golang doc.


## using this for your own pack

You're free to use this for your own modpack. you do not need to ask permission only what I ask you to comply to GNU version 3.
and crediting me would be much appreciated. I would love to see people using it or even make it better.

first you need to have a webserver, upload the php file in the `web` directory to your webserver.

This is your version manager, so if you release a new version, change the line where it says `$modpack->version = "0.0.1";` to for example `$modpack->version = "1.0.0";`

now you need to make the downloadable zip, place the `config` folder and `mods` folder in the zip with the `modpack.json` from this project, do not forget to set the right urls etc.

after this, you can test launching it, it should download the zip and then extract it in the same folder where the executable is located.

you can put this as a pre-launch script in multimc. place the binrary in the `.minecraft` folder from the instance from multimc.

`edit intance` > `settings` > `custom commands` > Pre-launch command: `$INST_MC_DIR\modpack-updater64.exe`

Make sure that the file is been unblocked, because windows defender smartscreen likes to lose it mind about it because it's a unknown file and the launcher is unable to launch the file

Done!



## Disclaimer

I'm not making money by rehosting your mods nor do I claim that I made the mods. I'm not those Malvertising websites like 9minecraft that is full with malware and ads.

My only goal is only playing and having fun with your mods, and making it easier for me to distribute updates to my

friends without giving zip files to everyone when I change something.

