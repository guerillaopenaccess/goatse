// Interface to Go http wrapper "httpplus". argobj may be undefined.
goatse.httpMethod = function(method, url, argobj) {
    // make all headers and params strings
    if (argobj !== undefined) {
        if (argobj.Headers !== undefined) {
            _.each(argobj.Headers, function(value, key) {
                argobj.Headers[key] = "" + value;
            });
        }
        if (argobj.PostBodyParams !== undefined) {
            _.each(argobj.PostBodyParams, function(value, key) {
                argobj.PostBodyParams[key] = "" + value;
            });
        }
    }
  return goatse._wrapgofunc(this._httpMethod(method, url, JSON.stringify(argobj)));
};

goatse.head = function(url){
  return goatse.httpMethod("head", url);
};

goatse.get = function(url){
  return goatse.httpMethod("get", url);
};

goatse.post = function(url, body) {
  return goatse.httpMethod("post", url, {"PostBody":body});
};

goatse.postFile = function(url, fileName) {
  return goatse.httpMethod("post", url, {"PostFileByName":fileName});
};
