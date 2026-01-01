"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.hasVendorPrefix = hasVendorPrefix;
exports.getVendorPrefix = getVendorPrefix;
exports.stripVendorPrefix = stripVendorPrefix;
/**
 * Checks whether given property name has vender prefix
 */
function hasVendorPrefix(prop) {
    return Boolean(getVendorPrefix(prop));
}
/**
 * Get the vender prefix from given property name
 */
function getVendorPrefix(prop) {
    return /^-\w+-/u.exec(prop)?.[0] || '';
}
/**
 * Strip the vender prefix
 */
function stripVendorPrefix(prop) {
    return prop.slice(getVendorPrefix(prop).length);
}
