// Special methods for libgen.
goatse.libgen = {};

goatse.libgen.default_urlroot = "http://libgen.in";
goatse.libgen.urlroot = "http://libgen.in";

goatse.libgen.uploadUrl = function() {
    return this.urlroot + "/scimag/librarian/form.php";
};

goatse.libgen.registerUrl = function() {
    return this.urlroot + "/scimag/librarian/register.php";
};

goatse.libgen.getUrlPdf = function() {
    return this.urlroot + "/scimag/get.php?doi=";
};

goatse.libgen.hasPdf = function(doi) {
    var r = goatse.head(this.getUrlPdf() + doi);
    if (r.statusCode === 200) {
        return true;
    } else {
        return false;
    }
};

goatse.libgen.buildRegister = function(doi, filesize, md5) {
    var doi_crossref = goatse.doi.getMeta(doi);
    var register = {};
    register.doi = doi;
    register.md5 = md5;
    register.filesize = filesize;
    register.year = "";
    register.month = "";
    register.day = "";
    register.firstpage = "";
    register.lastpage = "";
    register.issue = "";
    register.volume = "";
    register.journalid = "";
    register.journal = "";
    register.title = "";
    register.author = "";
    register.isbn = "";
    register.issnp = "";
    register.issne = "";
    register.editmode = 1;
    register.abstracturl = doi_crossref.Url;
    if (doi_crossref.Title !== undefined && doi_crossref.Title.length > 0) {
        register.title = _.first(doi_crossref.Title);
    }
    if (doi_crossref['container-title'] !== undefined && doi_crossref["container-title"].length > 0) {
        register.journal = _.first(doi_crossref["container-title"]);
    }
    if (!!doi_crossref.Author) {
        var authors = _.map(doi_crossref.Author, function(a) {
            return a.Family + ", " + a.Given;
        });
        register.author = authors.join("; ");
    }
    if (doi_crossref.Page !== undefined && doi_crossref.Page.split("-").length===2) {
        var splitpages = doi_crossref.Page.split("-");
        register.firstpage = splitpages[0];
        register.lastpage = splitpages[1];
    }
    if (doi_crossref.Issued['date-parts'] !== undefined && doi_crossref.Issued['date-parts'].length > 0) {
        if (doi_crossref.Issued['date-parts'][0][0] !== 0) {
            register.year = doi_crossref.Issued['date-parts'][0];
        }
        if (doi_crossref.Issued['date-parts'][0][1] !== 0) {
            register.month = doi_crossref.Issued['date-parts'][0];
        }
        if (doi_crossref.Issued['date-parts'][0][2] !== 0) {
            register.day = doi_crossref.Issued['date-parts'][0];
        }
    }
    if (doi_crossref.ISSN !== undefined && doi_crossref.ISSN.length > 0) {
        register.issnp = doi_crossref.ISSN[0];
    }
    if (doi_crossref.ISSN !== undefined && doi_crossref.ISSN.length > 1) {
        register.issne = doi_crossref.ISSN[1];
    }
    return register;
};

goatse.libgen.upload = function(doi, filename) {
    var file_md5 = goatse.io.md5SumFile(filename);
    var file_size = goatse.io.fileSize(filename);
    var upload_args = { "PostFileByName": filename, "PostFileByNameName": "uploadedfile", "PostBodyParams": {"doi": doi}, "BasicAuthUser": "genesis", "BasicAuthPassword": "upload" };
    var first_upload = goatse.httpMethod("post", this.uploadUrl(), upload_args);
    if (first_upload.statusCode !== 200) {
        return false;
    }
    var register_data = this.buildRegister(doi, file_size, file_md5);
    var register_args = { "PostBodyParams": register_data, "BasicAuthUser": "genesis", "BasicAuthPassword": "upload" };
    var register_upload = goatse.httpMethod("post", this.registerUrl(), register_args);
    return register_upload.statusCode === 200;
};

goatse.postFile = function(url, fileName) {
  return goatse.httpMethod("post", url, {"PostFileByName":fileName});
};
