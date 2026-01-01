function convertOptionToLegacy(processor, verifyOption, config) {
  var _a;
  if (processor == null)
    return verifyOption;
  if (typeof processor === "string") {
    return convertOptionToLegacy(
      findProcessor(processor, config),
      verifyOption,
      config
    );
  }
  const filename = (_a = typeof verifyOption === "string" ? verifyOption : verifyOption == null ? void 0 : verifyOption.filename) != null ? _a : "<input>";
  const preprocess = function(code) {
    var _a2;
    const result = (_a2 = processor.preprocess) == null ? void 0 : _a2.call(processor, code, filename);
    return result ? result : [code];
  };
  const postprocess = function(messages) {
    var _a2;
    const result = (_a2 = processor.postprocess) == null ? void 0 : _a2.call(processor, messages, filename);
    return result ? result : messages[0];
  };
  if (verifyOption == null) {
    return { preprocess, postprocess };
  }
  if (typeof verifyOption === "string") {
    return { filename: verifyOption, preprocess, postprocess };
  }
  return { ...verifyOption, preprocess, postprocess };
}
function findProcessor(processor, config) {
  var _a, _b;
  let pluginName, processorName;
  const splitted = processor.split("/")[0];
  if (splitted.length === 2) {
    pluginName = splitted[0];
    processorName = splitted[1];
  } else if (splitted.length === 3 && splitted[0].startsWith("@")) {
    pluginName = `${splitted[0]}/${splitted[1]}`;
    processorName = splitted[2];
  } else {
    throw new Error(`Could not resolve processor: ${processor}`);
  }
  const plugin = (_a = config.plugins) == null ? void 0 : _a[pluginName];
  const resolved = (_b = plugin == null ? void 0 : plugin.processors) == null ? void 0 : _b[processorName];
  if (!resolved) {
    throw new Error(`Could not resolve processor: ${processor}`);
  }
  return resolved;
}

export { convertOptionToLegacy as c };
