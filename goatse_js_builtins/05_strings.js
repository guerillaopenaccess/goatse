// Go functions on unicode strings.
goatse.strings = {};

goatse.strings._entry = function(funcname, mainarg, secondarg) {
    return goatse._wrapgofunc(goatse._strings_entry(funcname, mainarg, secondarg));
};

goatse.strings.urlEscape = function(s) {
    return goatse._wrapgofunc(goatse._strings_urlescape(s));
};

goatse.strings.urlUnescape = function(s) {
    return goatse._wrapgofunc(goatse._strings_urlunescape(s));
};

_.each([
    "contains","containsAny","containsRune","count","equalFold",
    "hasPrefix","hasSuffix","indexRune","lastIndex","lastIndexAny",
    "title","trim","trimSpace","trimLeft","trimRight","trimPrefix",
    "trimSuffix"], function(funcname){
        var func = _.partial(goatse.strings._entry, funcname);
        goatse.strings[funcname] = func;
});
