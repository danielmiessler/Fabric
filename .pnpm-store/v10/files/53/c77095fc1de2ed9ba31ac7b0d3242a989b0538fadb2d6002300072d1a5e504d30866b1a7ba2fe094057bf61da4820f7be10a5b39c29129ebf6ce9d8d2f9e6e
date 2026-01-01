import { CommentDirectives } from './comment-directives';
export declare class Shared {
    readonly commentDirectives: CommentDirectives[];
    newCommentDirectives(options: ConstructorParameters<typeof CommentDirectives>[0]): CommentDirectives;
}
/** Start sharing and make the data available. */
export declare function beginShared(filename: string): void;
/** Get the shared data and end the sharing. */
export declare function terminateShared(filename: string): Shared | null;
/** If sharing has started, get the shared data. */
export declare function getShared(filename: string): Shared | null;
