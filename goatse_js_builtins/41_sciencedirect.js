// ScienceDirect.com
goatse.sciencedirect = {};

goatse.sciencedirect.targetsUrl = _.bind(function(url) {
    return Boolean(url.match(/sciencedirect.com\/science\/article\/pii/));
}, goatse.sciencedirect);

goatse.sciencedirect.getDoi = _.bind(function(doc) {
    return (/^SDM.doi = '([^']+)'/m).exec(doc)[1];
}, goatse.sciencedirect);

goatse.sciencedirect.haveAccess = _.bind(function(doc) {
    var url = this.getUrlPdf(doc);
    return !url.match(/_ob=ShoppingCartURL/);
}, goatse.sciencedirect);

goatse.sciencedirect.getUrlPdf = _.bind(function(doc) {
    return goatse.query(doc).find("#pdfLink").attr("href")[0];
}, goatse.sciencedirect);

goatse.sciencedirect.getPdf = _.bind(function(doc) {
    var pdfurl = this.getUrlPdf(doc);
    var pdf_r = goatse.get(pdfurl);
    if (pdf_r.statusCode === 200) {
        return pdf_r.body;
    } else {
        return "";
    }
}, goatse.sciencedirect);

goatse.addResolver("sciencedirect.com", goatse.sciencedirect);
