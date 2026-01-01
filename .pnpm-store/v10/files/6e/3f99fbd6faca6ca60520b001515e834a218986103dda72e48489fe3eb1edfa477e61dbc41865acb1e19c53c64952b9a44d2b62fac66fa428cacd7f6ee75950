"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const groupingUtils_1 = require("./groupingUtils");
class PageFactorFinder {
    find(containers, extractor, config = { sampleCount: 20, minFulfillment: 0.8 }) {
        const containerAnalyzeCount = Math.min(config.sampleCount, containers.length);
        const start = Math.max(containers.length / 2 - containerAnalyzeCount / 2, 0); //start somewhere in the middle
        const pageNumbers = containers
            .slice(start, start + containerAnalyzeCount)
            .map((container) => extractor(container))
            .map((extract) => extract.numbers.map((num) => num - extract.index).filter(groupingUtils_1.onlyUniques));
        const distanceCounts = pageNumbers.reduce((map, indexDistancesPerPage) => {
            indexDistancesPerPage.forEach((indexDistance) => {
                map[indexDistance] = (map[indexDistance] || 0) + 1;
            });
            return map;
        }, {});
        const hits = Object.keys(distanceCounts)
            .filter((distance) => distanceCounts[distance] / containerAnalyzeCount >= config.minFulfillment)
            .sort((d1, d2) => distanceCounts[d1] - distanceCounts[d2]);
        // for all remaining index distance arrays - check y coordinates
        if (hits.length < 1) {
            return undefined;
        }
        return Number.parseInt(hits[0]);
    }
}
exports.default = PageFactorFinder;
