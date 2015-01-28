goatse.doi = {};

// doi object has fields from CrossRef. Some are lists.
// `Issued Publisher License Source Url Doi ISSN Score Page Issue
//  Deposited ReferenceCount Author Title Type Volume ContainerTitle Subtitle`

goatse.doi.getMeta = function(doi) {
    return goatse._wrapgofunc(goatse._getDOIMeta(doi));
};

goatse.doi.onLibgen = function(doi) {
    var r = goatse.head("http://libgen.in/scimag/get.php?" + doi);
    if (r.statusCode === 200) {
        return true;
    } else {
        return false;
    }
};

// Returns array of doi objects not oa and not on libgen.
goatse.doi.notOpenSample = function(number) {
    return goatse._wrapgofunc(goatse._notOpenSample(number));
};

// Returns array of doi objects for member by name.
// Only 20 members are supported.
goatse.doi.memberSample = function(number, name) {
    return goatse._wrapgofunc(goatse._memberSample(number, name));
};

_.each(["Elsevier", "Wiley", "Springer", "Informa", "JSTOR", "Ovid", "OUP",
 "Sage", "ACS", "CUP", "Nature", "IEEE", "BMJ", "IOP", "AMA", "APS",
 "AIP", "Thieme", "WalterdeGruyter", "RSC"], function(publisher){
    var function_name = publisher.toLowerCase() + "Sample";
    goatse.doi[function_name] = function(number){
        return goatse.doi.memberSample(number, publisher);
    };
});
