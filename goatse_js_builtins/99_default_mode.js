goatse.default_mode = function() {
    while (true) {
        // 0:36 to 5:30
        var between = (goatse.os.randInt(50) * 6000) + 30000;
        goatse.log("Sleep "+ between/1000 +" seconds between list");
        goatse.os.sleep(between);
        try {
            var worklist = goatse.doi.notOpenSample(20);
            var filename = _.map(worklist, goatse.grabone);
            goatse.log("Saved files to "+filename);
        }
        catch(e) {
            goatse.log("ERROR: "+e.message);
        }
    }
};
