goatse.wiley = {};

goatse.wiley.targetsUrl = _.bind(function(url) {
    return Boolean(url.match(/onlinelibrary\.wiley\.com\/doi\//));
}, goatse.wiley);

goatse.wiley.getDoi = _.bind(function(doc) {
    return goatse.query(doc).find("meta[name=citation_doi]").attr("content");
}, goatse.wiley);

goatse.wiley.haveAccess = _.bind(function(doc) {
    var pagenav = goatse.query(doc).find("#pageNavAndTools");
    var get_access = _.some(pagenav.text(), function(navtext) {
        return Boolean(!navtext.match(/GET ACCESS/));
    });
    return get_access;
}, goatse.wiley);

goatse.wiley.getUrlPdf = _.bind(function(doc) {
    var relurl = goatse.query(doc).find("#journalToolsPdfLink").attr("href");
    relurl = goatse.strings.trimSpace(relurl);
    if (relurl) {
        return "http://onlinelibrary.wiley.com" + relurl;
    }
    return "";
}, goatse.wiley);

goatse.wiley.getPdf = _.bind(function(doc) {
    var pdfurl = this.getUrlPdf(doc);
    var pdf_r = goatse.get(pdfurl);
    if (pdf_r.statusCode === 200) {
        return pdf_r.body;
    } else {
        return "";
    }    
}, goatse.wiley);

goatse.addResolver("wiley.com", goatse.wiley);
