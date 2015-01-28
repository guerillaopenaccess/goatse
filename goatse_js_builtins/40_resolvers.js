// Array of resolvers.
goatse.resolvers = [];

// Resolver must:
// <name>.targetsUrl(url string) -> bool
// <name>.getDoi(doc string) -> string
// <name>.haveAccess(doc string) -> bool
// <name>.getUrlPdf(doc string) -> string
// <name>.getPdf(doc string) -> string //binary
goatse.addResolver = function(name, resolver) {
    var missing_methods = _.difference(['targetsUrl', 'getDoi', 'haveAccess', 'getUrlPdf', 'getPdf'], _.functions(resolver));
    if (!_.isEmpty(missing_methods)) {
        goatse.log("Error: missing on '"+name+"': "+missing_methods);
        throw new Error("Error: missing on '"+name+"': "+missing_methods);
    }
    goatse.resolvers.push(resolver);
};

// Uses response to get resolver by URL.
// First resolver matches URL && access && parse PDF Url returned
goatse.getResolver = function(http_response, require_access) {
    var resolver = _.find(goatse.resolvers, function(resolver) {
        if (!resolver.targetsUrl(http_response.url)) {
            return false;
        }
        if (require_access) {
            if (!resolver.haveAccess(http_response.body)) {
                return false;
            }
            if (!resolver.getUrlPdf(http_response.body)) {
                return false;
            }
        }
        return true;
    });
    if (resolver !== undefined) {
        return resolver;
    } else {
        throw new Error("No have-access resolver with URL: "+http_response.url);
    }
};

goatse.grabPdfByDoi = function(doi) {
    var crossref_meta = goatse.doi.getMeta(doi);
    var filename = "";
    if (crossref_meta.Title.length > 0) {
        filename = goatse.io.fileNameSafe(crossref_meta.Doi + " - " + crossref_meta.Title[0] + ".pdf");
    } else {
        filename = goatse.io.fileNameSafe(crossref_meta.Doi + ".pdf");
    }
    if (goatse.io.fileExists(filename)) {
        throw new Error("Paper file exist's: "+filename);
    }
    var http_response = goatse.get(crossref_meta.Url);
    var pdf = "";
    try {
        var res = goatse.getResolver(http_response, true);
        pdf = res.getPdf(http_response.body);
    }
    catch(e) {
        errfn = goatse.io.fileNameSafe("goatsefail_"+http_response.statusCode+"__"+crossref_meta.Doi+"__"+http_response.url+".html");
        goatse.io.saveToFile(errfn, http_response.body);
        throw e;
    }
    goatse.io.saveToFile(filename, pdf);
    return filename;
};

goatse.grabPdfsFromDoiList = function(doilist) {
    var filenames = {};
    _.each(doilist, function(doi) {
        goatse.log("Try to grab doi: "+doi);
        try {
            var fn = goatse.grabPdfByDoi(doi);
            goatse.log("Grab doi: "+doi);
            filenames[doi] = fn;
        }
        catch(e) {
            goatse.log("Failed doi: "+doi+" -: "+e.message);
            filenames[doi] = "";
        }
    });
    return filenames;
};
