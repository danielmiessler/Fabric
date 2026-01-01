export class Comment {
    constructor(editor: any);
    renderForToolbar(): any;
    renderForStandalone(): any;
    focusButton(): void;
    onUpdatedColor(): void;
    get commentButtonWidth(): number;
    set commentPopupPositionInLayer(pos: number[] | null);
    get commentPopupPositionInLayer(): number[] | null;
    hasDefaultPopupPosition(): boolean;
    removeStandaloneCommentButton(): void;
    removeToolbarCommentButton(): void;
    setCommentButtonStates({ selected, hasPopup }: {
        selected: any;
        hasPopup: any;
    }): void;
    edit(options: any): void;
    finish(): void;
    isDeleted(): boolean;
    isEmpty(): boolean;
    hasBeenEdited(): boolean;
    serialize(): {
        text: null;
        richText: null;
        date: null;
        deleted: boolean;
    };
    /**
     * Set the comment data.
     */
    set data(text: {
        text: null;
        richText: null;
        date: null;
        deleted: boolean;
    });
    get data(): {
        text: null;
        richText: null;
        date: null;
        deleted: boolean;
    };
    setInitialText(text: any, richText?: null): void;
    shown(): void;
    destroy(): void;
    #private;
}
