goatse.query_find = function(doc, query) {
  return this._wrapgofunc(this._query_find(doc, query, false));
};

goatse.query_findagain = function(doc, query) {
  return this._wrapgofunc(this._query_find(doc, query, true));
};

goatse.query_attr = function(doc, attr) {
  return goatse.strings.trimSpace(this._wrapgofunc(this._query_attr(doc, attr)));
};

goatse.query_children = function(doc) {
  return this._wrapgofunc(this._query_children(doc));
};

goatse.query_text = function(doc) {
  return goatse.strings.trimSpace(this._wrapgofunc(this._query_text(doc)));
};

// Like JQuery
goatse._queryobj = function(doc, subsequent) {
        if (subsequent) {
            this.value = doc;
        } else {
            this.value = [doc];
        }
        this.html = function() {
                return this.value;
            };
        this.text = function() {
                return _.map(this.value, function(currentValue){
                    return goatse.query_text(currentValue);
                });
            };
        this.attr = function(attr_name) {
                return _.map(this.value, function(currentValue){
                    return goatse.query_attr(currentValue, attr_name);
                });
            };
        this.children = function() {
                var children = _.map(this.value, function(currentValue){
                    return goatse.query_children(currentValue);
                });
                children = _.flatten(children);
                return new goatse._queryobj(children, true);
            };
        this.find = function(query_string) {
                var found = _.map(this.value, function(currentValue){
                    if (subsequent) {
                        return goatse.query_findagain(currentValue, query_string);
                    } else {
                        return goatse.query_find(currentValue, query_string);
                    }
                });
                found = _.flatten(found);
                return new goatse._queryobj(found, true);
            };
};

// Like JQuery
goatse.query = function(doc) {
    return new goatse._queryobj(doc);
};
