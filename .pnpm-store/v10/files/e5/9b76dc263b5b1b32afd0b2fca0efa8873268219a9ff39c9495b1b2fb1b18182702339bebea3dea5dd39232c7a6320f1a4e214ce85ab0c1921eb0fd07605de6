//Abstract stream which allows stash items temporarily
export default class StashingStream {

    constructor() {
        if (this.constructor === StashingStream) {
            throw new TypeError("Can not construct abstract class.");
        }
        this.results = [];
        this.stash = [];
    }

    consumeAll(items) {
        items.forEach(item => this.consume(item));
    }

    consume(item) {
        if (this.shouldStash(item)) {
            if (!this.matchesStash(item)) {
                this.flushStash();
            }
            this.pushOnStash(item);
        } else {
            if (this.stash.length > 0) {
                this.flushStash();
            }
            this.results.push(item);
        }
    }

    pushOnStash(item) {
        this.onPushOnStash(item);
        this.stash.push(item);
    }

    complete() {
        if (this.stash.length > 0) {
            this.flushStash();
        }
        return this.results;
    }

    // return true if the item matches the items of the stack
    matchesStash(item) {
        if (this.stash.length == 0) {
            return true;
        }
        const lastItem = this.stash[this.stash.length - 1];
        return this.doMatchesStash(lastItem, item);
    }

    flushStash() {
        if (this.stash.length > 0) {
            this.doFlushStash(this.stash, this.results);
            this.stash = [];
        }
    }

    onPushOnStash(item) { // eslint-disable-line no-unused-vars
        //sub-classes may override
    }

    shouldStash(item) {
        throw new TypeError("Do not call abstract method foo from child." + item);
    }

    doMatchesStash(lastItem, item) {
        throw new TypeError("Do not call abstract method foo from child." + lastItem + item);
    }

    doFlushStash(stash, results) {
        throw new TypeError("Do not call abstract method foo from child." + stash + results);
    }
}