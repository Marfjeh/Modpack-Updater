const   http = require('https'),
        fs = require('fs'),
        unzip = require("unzip"),
        sha1File = require('sha1-file'),
        rimraf = require('rimraf'),
        request = require("request");

let modpackjson = JSON.parse(fs.readFileSync('./modpack.json', 'utf8'));
let version = "";
let sha1expected = "";

console.log("SweetNyanCraft Modpack Updater Version 0.2");
console.log("Github: https://github.com/Marfjeh/Modpack-Updater")
console.log("------------------------------------------");
console.log("Modpack: " + modpackjson.Name);
console.log("Modpack version: " + modpackjson.Version);
console.log("------------------------------------------");


function init() {
    console.log("Getting Checksum from Server...");
    request({
        uri: modpackjson.sha1sum,
        method: "GET",
        timeout: 10000,
        followRedirect: true,
        maxRedirects: 10
    }, function(error, response, body) {
        if (error) {
            console.error("Error getting checksum. do you have internet?");
            console.error("Giving up...");
            console.error(error);
            process.exit(1);
        }
        let bodyjson = JSON.parse(body);
        sha1expected = bodyjson.sha1;
        version = bodyjson.version;
        console.log("Version on the server : " + version);
        if (!modpackjson.AutoUpdater) {
            console.log("AutoUpdate is disabled in the modpack.json");
            process.exit(0);
        }
        else if (version === modpackjson.Version) {
            console.log("You're already on the newest version!");
            process.exit(0);
        } else {
            console.log("Update found!");
            console.log("Downloading update...");
            download(modpackjson.URL, "./deployment.zip", null);
        }
    });
}

function clearFolders() {
    rimraf('./mods/*', function () { 
        rimraf('./config/*', function () { 
            unzipjob();
        });
    });
}

function verify() {    
    console.log("Checking SHA-1 sum...");
    let sha1zip = sha1File("./deployment.zip");
    console.log("Expecting SHA-1: " + sha1expected);
    console.log("      Got SHA-1: " + sha1zip);
    if (sha1zip === sha1expected) {
        console.log("SHA-1 Verify success! Unzipping...");
        clearFolders();
    }
    else {
        console.error("ERROR: SHA-1 Verify failed.");
        process.exit(1);
    }
    
}

function unzipjob() {
    fs.createReadStream('./deployment.zip').pipe(unzip.Extract({ path: '.' }));
    fs.unlinkSync("./deployment.zip");
    console.log("Finishing in a bit, still extracting...");
}

var download = function(url, dest, cb) {
  var file = fs.createWriteStream(dest);
  var request = http.get(url, function(response) {
    response.pipe(file);
    file.on('finish', function() {
      file.close(cb);
      verify();
    });
  }).on('error', function(err) { // Handle errors
    fs.unlink(dest);
    if (cb) cb(err.message);
  });
};

init();
