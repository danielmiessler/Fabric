"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const EvaluationTracker_1 = require("../debug/EvaluationTracker");
class TransformContext {
    constructor(fontMap, pageViewports, globals, evaluations = new EvaluationTracker_1.default()) {
        this.fontMap = fontMap;
        this.pageViewports = pageViewports;
        this.globals = globals;
        this.evaluations = evaluations;
        this.pageCount = pageViewports.length;
    }
    trackEvaluation(item, score = undefined) {
        this.evaluations.trackEvaluation(item, score);
    }
    globalIsDefined(definition) {
        return this.globals.isDefined(definition);
    }
    getGlobal(definition) {
        return this.globals.get(definition);
    }
    getGlobalOptionally(definition) {
        return this.globals.getOptional(definition);
    }
}
exports.default = TransformContext;
