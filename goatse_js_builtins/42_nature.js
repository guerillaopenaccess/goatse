// NPG.
goatse.nature = {};

goatse.nature.targetsUrl = _.bind(function(url) {
    return Boolean(url.match(/nature.com\/.+\/journal\/.+\/full/));
}, goatse.nature);

goatse.nature.getDoi = _.bind(function(doc) {
    var doistr = goatse.query(doc).find("dd.doi").text()[0];
    return doistr.slice(doistr.indexOf(":") + 1);
}, goatse.nature);

goatse.nature.haveAccess = _.bind(function(doc) {
    var access = goatse.query(doc).find("meta[name=access]").attr("content");
    return (access.valueOf() != "No");
}, goatse.nature);

goatse.nature.getUrlPdf = _.bind(function(doc) {
    return goatse.query(doc).find(".articlepdf").find("a").attr("href")[0];
}, goatse.nature);

goatse.nature.getPdf = _.bind(function(doc) {
    var pdfurl = this.getUrlPdf(doc);
    var pdf_r = goatse.get(pdfurl);
    if (pdf_r.statusCode === 200) {
        return pdf_r.body;
    } else {
        return "";
    }
}, goatse.nature);

goatse.addResolver("nature.com", goatse.nature);
