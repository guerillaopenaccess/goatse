goatse.doilist_mode = function(doilist) {
    var lines = goatse.io.fileLines(doilist);
    goatse.log("Lines: "+lines);
    goatse.log("Lines length: "+lines.length);
    for (var i=0; i<lines.length; i++) {
        var line = goatse.strings.trimSpace(lines[i]);
        if (goatse.strings.contains(line, "#") == "true") {
            continue;
        }
        if (!goatse.strings.trimSpace(line)) {
            continue;
        }
        goatse.log("Line: "+line);
        var thisdoi = line.split("\t")[0];
        goatse.log("doi: "+thisdoi);
        try {
            goatse.log("Getting meta")
            var doimeta = goatse.doi.getMeta(thisdoi);
            goatse.log("meta for: "+_.first(doimeta.Title));
            var filename = goatse.grabone(doimeta);
            goatse.log("Result saved: "+filename);
        } catch(e) {
            goatse.log("Error on '"+line+"': "+e.message)
            continue
        }
    }
    goatse.log("End of file");
}
