goatse.springer = {};

goatse.springer.targetsUrl = _.bind(function(url) {
    return Boolean(url.match(/link.springer.com\/article\//));
}, goatse.springer);

goatse.springer.getDoi = _.bind(function(doc) {
    return goatse.query(doc).find("meta[name=citation_doi]").attr("content");
}, goatse.springer);

goatse.springer.haveAccess = _.bind(function(doc) {
    var pagenav = goatse.query(doc).find("#pageNavAndTools");
    var get_access = _.some(pagenav.text(), function(navtext) {
        return Boolean(!navtext.match(/GET ACCESS/));
    });
    return get_access;
}, goatse.springer);

goatse.springer.getUrlPdf = _.bind(function(doc) {
    var relurl = goatse.query(doc).find("#journalToolsPdfLink").attr("href");
    relurl = goatse.strings.trimSpace(relurl);
    if (relurl) {
        return "http://onlinelibrary.wiley.com" + relurl;
    }
    return "";
}, goatse.springer);

goatse.springer.getPdf = _.bind(function(doc) {
    var pdfurl = this.getUrlPdf(doc);
    var pdf_r = goatse.get(pdfurl);
    if (pdf_r.statusCode === 200) {
        return pdf_r.body;
    } else {
        return "";
    }

}, goatse.springer);

goatse.addResolver("springer.com", goatse.springer);
