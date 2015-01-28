goatse.grabone = function(doimeta) {
        var filename = "undefined";
        // 0:01 to 0:05
        var smallsleep = goatse.os.randInt(5) * 1000;
        goatse.os.sleep(smallsleep);
        try {
            filename = goatse.grabPdfByDoi(doimeta.Doi);
        }
        catch(e) {
            goatse.log("Fail with doi (fail dumped to '"+filename+"'): "+JSON.stringify(doimeta));
            return filename;
        }
        if (filename != "undefined") {
            goatse.log("Saved '"+doimeta.Doi+"' to '"+filename+"'.");
            goatse.log("Attempt libgen upload of "+filename);
            try {
                goatse.libgen.urlroot = _.find(['http://libgen.in', 'http://libgen.org'], function(url) {
                    try {
                        if (goatse.head(url).statusCode === 200) {
                            return filename;
                        } else {
                            return filename;
                        }
                    } catch(e) {
                        return filename;
                    }
                });
                if (goatse.libgen.urlroot === undefined) {
                    goatse.libgen.urlroot = goatse.libgen.default_urlroot;
                }
                goatse.libgen.upload(doimeta.Doi, filename);
            }
            catch(e) {
                goatse.log("Failed upload of "+filename+": "+e.message);
                return filename;
            }
        }
        return filename;
};
